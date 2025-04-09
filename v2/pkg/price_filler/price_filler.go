package pricefiller

import (
	"context"
	"errors"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/KyberNetwork/tradelogs/v2/pkg/constant"
	"github.com/KyberNetwork/tradelogs/v2/pkg/mtm"
	dashboardStorage "github.com/KyberNetwork/tradelogs/v2/pkg/storage/dashboard"
	dashboardTypes "github.com/KyberNetwork/tradelogs/v2/pkg/storage/dashboard/types"
	cowProtocolStorage "github.com/KyberNetwork/tradelogs/v2/pkg/storage/tradelogs/cow_protocol"
	storageTypes "github.com/KyberNetwork/tradelogs/v2/pkg/storage/tradelogs/types"
	"go.uber.org/zap"
)

const (
	NetworkETHChainID         = 1
	NetworkETHChainIDString   = "1"
	NetworkETH                = "ETH"
	updateAllCoinInfoInterval = 12 * time.Hour
	backfillTradeLogsLimit    = 60
	addressETH1               = "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"
	addressETH2               = "0x0000000000000000000000000000000000000000"
	coinUSDT                  = "USDT"
	USDTAddress               = "0xdac17f958d2ee523a2206206994597c13d831ec7"
	invalidSymbolErrString    = "<APIError> code=-1121, msg=Invalid symbol."
)

var (
	ErrNoPrice               = errors.New("no price from binance")
	ErrWeirdTokenCatalogResp = errors.New("weird token catalog response")
)

type CoinInfo struct {
	Coin            string
	ChainId         int64
	ContractAddress string
	Decimals        int64
}

type PriceFiller struct {
	l                  *zap.SugaredLogger
	s                  []storageTypes.Storage
	mu                 sync.Mutex
	ksClient           *KsClient
	mappedCoinInfo     map[string]CoinInfo // address - coinInfo
	mtmClient          *mtm.MtmClient
	dashboardStorage   *dashboardStorage.Storage
	cowTransferStorage *cowProtocolStorage.CowTransferStorage
}

func NewPriceFiller(l *zap.SugaredLogger,
	s []storageTypes.Storage,
	mtmClient *mtm.MtmClient,
	dashboardStorage *dashboardStorage.Storage,
	cowTransferStorage *cowProtocolStorage.CowTransferStorage,
) (*PriceFiller, error) {
	p := &PriceFiller{
		l:        l,
		s:        s,
		ksClient: NewKsClient(),
		mappedCoinInfo: map[string]CoinInfo{
			addressETH1: {
				Coin:            "ETH",
				ChainId:         NetworkETHChainID,
				ContractAddress: addressETH1,
				Decimals:        18,
			},
			addressETH2: {
				Coin:            "ETH",
				ChainId:         NetworkETHChainID,
				ContractAddress: addressETH2,
				Decimals:        18,
			},
		},
		mtmClient:          mtmClient,
		dashboardStorage:   dashboardStorage,
		cowTransferStorage: cowTransferStorage,
	}

	if err := p.updateAllCoinInfo(); err != nil {
		return nil, err
	}
	return p, nil
}

func (p *PriceFiller) Run(fillPriceInterval time.Duration) {
	go p.runUpdateAllCoinInfoRoutine()
	p.runBackFillTradelogPriceRoutine(fillPriceInterval)
}

func (p *PriceFiller) getPrice(token string, timestamp int64) (float64, error) {
	price, err := p.mtmClient.GetHistoricalRate(
		context.Background(),
		token,
		USDTAddress,
		NetworkETHChainID,
		time.UnixMilli(timestamp),
	)
	if err != nil {
		return 0, err
	}
	return price, nil
}

func (p *PriceFiller) updateAllCoinInfo() error {
	tokens, err := p.mtmClient.GetListTokens(context.Background())
	if err != nil {
		p.l.Errorw("Failed to get all coins info", "err", err)
		return err
	}

	p.mu.Lock()
	defer p.mu.Unlock()
	for _, info := range tokens {
		if info.ChainId != NetworkETHChainID {
			continue
		}
		p.mappedCoinInfo[info.Address] = CoinInfo{
			Coin:            info.Symbol,
			ChainId:         info.ChainId,
			ContractAddress: info.Address,
			Decimals:        info.Decimals,
		}

	}
	p.l.Infow("Successfully map coin info")
	return nil
}

func (p *PriceFiller) runUpdateAllCoinInfoRoutine() {
	ticker := time.NewTicker(updateAllCoinInfoInterval)
	defer ticker.Stop()

	for range ticker.C {
		if err := p.updateAllCoinInfo(); err != nil {
			p.l.Errorw("Failed to updateAllCoinInfo", "err", err)
		}
	}
}

func (p *PriceFiller) runBackFillTradelogPriceRoutine(fillPriceInterval time.Duration) {
	ticker := time.NewTicker(fillPriceInterval)
	defer ticker.Stop()

	for ; ; <-ticker.C {
		for _, s := range p.s {
			tradeLogs, err := s.GetEmptyPrice(backfillTradeLogsLimit)
			if err != nil {
				p.l.Errorw("Failed to get tradeLogs", "exchange", s.Exchange(), "err", err)
				continue
			}

			p.FullFillTradeLogs(tradeLogs)
			if err = s.Insert(tradeLogs); err != nil {
				p.l.Errorw("Failed to insert tradeLogs", "exchange", s.Exchange(), "err", err)
				continue
			}

			p.l.Infow("backfill tradelog price successfully", "exchange", s.Exchange(), "number", len(tradeLogs))
		}
		// fill price for cow transfer events
		if err := p.insertTokens(); err != nil {
			p.l.Errorw("Failed to insert tokens", "err", err)
		}
		transferEvents, err := p.cowTransferStorage.GetEmptyPrice(backfillTradeLogsLimit)
		if err != nil {
			p.l.Errorw("Failed to get transfer events", "err", err)
			continue
		}
		p.FullFillCowTransferEvents(transferEvents)

		if err = p.cowTransferStorage.Insert(transferEvents); err != nil {
			p.l.Errorw("Failed to insert filled transfer event", "err", err)
		}

	}
}

func (p *PriceFiller) fullFillTradeLog(tradeLog storageTypes.TradeLog) (storageTypes.TradeLog, error) {
	makerPrice, makerUsdAmount, err := p.getPriceAndAmountUsd(strings.ToLower(tradeLog.MakerToken),
		tradeLog.MakerTokenAmount, int64(tradeLog.Timestamp))
	if isConnectionRefusedError(err) {
		p.l.Errorw("Failed to getPriceAndAmountUsd for maker", "err", err)
		return tradeLog, err
	}

	takerPrice, takerUsdAmount, err := p.getPriceAndAmountUsd(strings.ToLower(tradeLog.TakerToken),
		tradeLog.TakerTokenAmount, int64(tradeLog.Timestamp))
	if isConnectionRefusedError(err) {
		p.l.Errorw("Failed to getPriceAndAmountUsd for taker", "err", err)
		return tradeLog, err
	}

	tradeLog.MakerTokenPrice = &makerPrice
	tradeLog.MakerUsdAmount = &makerUsdAmount

	tradeLog.TakerTokenPrice = &takerPrice
	tradeLog.TakerUsdAmount = &takerUsdAmount

	return tradeLog, nil
}

func (p *PriceFiller) fullFillBebopTradeLog(tradeLog storageTypes.TradeLog) (storageTypes.TradeLog, error) {
	makerTokens := strings.Split(strings.ToLower(tradeLog.MakerToken), ",")
	takerTokens := strings.Split(strings.ToLower(tradeLog.TakerToken), ",")
	makerAmounts := strings.Split(tradeLog.MakerTokenAmount, ",")
	takerAmounts := strings.Split(tradeLog.TakerTokenAmount, ",")

	makerPrice, makerSumUsdAmount, err := p.getSumAmountUsd(makerTokens, makerAmounts, int64(tradeLog.Timestamp))
	if isConnectionRefusedError(err) {
		p.l.Errorw("Failed to getSumAndAmountUsd for maker", "err", err)
		return tradeLog, err
	}

	takerPrice, takerUsdAmount, err := p.getSumAmountUsd(takerTokens, takerAmounts, int64(tradeLog.Timestamp))
	if isConnectionRefusedError(err) {
		p.l.Errorw("Failed to getSumAmountUsd for taker", "err", err)
		return tradeLog, err
	}

	tradeLog.MakerTokenPrice = &makerPrice // set zero price for multi-maker trade
	tradeLog.MakerUsdAmount = &makerSumUsdAmount

	tradeLog.TakerTokenPrice = &takerPrice
	tradeLog.TakerUsdAmount = &takerUsdAmount

	return tradeLog, nil
}

func (p *PriceFiller) fullFillCowProtocolTradeLog(tradeLog storageTypes.TradeLog) (storageTypes.TradeLog, error) {
	sellTokenPrice, sellUsdAmount, err := p.getPriceAndAmountUsd(strings.ToLower(tradeLog.SellToken),
		tradeLog.SellAmount, int64(tradeLog.Timestamp))
	if isConnectionRefusedError(err) {
		p.l.Errorw("Failed to getPriceAndAmountUsd for sell token", "err", err)
		return tradeLog, err
	}

	buyTokenPrice, buyUsdAmount, err := p.getPriceAndAmountUsd(strings.ToLower(tradeLog.BuyToken),
		tradeLog.BuyAmount, int64(tradeLog.Timestamp))
	if isConnectionRefusedError(err) {
		p.l.Errorw("Failed to getPriceAndAmountUsd for buy token", "err", err)
		return tradeLog, err
	}

	tradeLog.SellTokenPrice = &sellTokenPrice
	tradeLog.SellUsdAmount = &sellUsdAmount

	tradeLog.BuyTokenPrice = &buyTokenPrice
	tradeLog.BuyUsdAmount = &buyUsdAmount

	return tradeLog, nil
}

func (p *PriceFiller) FullFillCowTransferEvent(cowTransfer storageTypes.CowTransfer) (storageTypes.CowTransfer, error) {
	tokenPrice, tokenUsdAmount, err := p.getPriceAndAmountUsd(strings.ToLower(cowTransfer.Token),
		cowTransfer.Amount, int64(cowTransfer.Timestamp))
	if isConnectionRefusedError(err) {
		p.l.Errorw("Failed to getPriceAndAmountUsd for maker", "err", err)
		return cowTransfer, err
	}

	cowTransfer.TokenPrice = &tokenPrice
	cowTransfer.AmountUsd = &tokenUsdAmount

	return cowTransfer, nil
}

func (p *PriceFiller) getSumAmountUsd(address, rawAmt []string, at int64) (float64, float64, error) {
	var sumAmount, price float64
	for i := range address {
		pr, usdAmount, err := p.getPriceAndAmountUsd(address[i], rawAmt[i], at)
		if isConnectionRefusedError(err) {
			p.l.Errorw("Failed to getPriceAndAmountUsd for address", "err", err)
			return 0, 0, err
		}
		sumAmount += usdAmount
		price = pr
	}
	if len(address) > 1 {
		price = 0
	}
	return price, sumAmount, nil
}

func (p *PriceFiller) getPriceAndAmountUsd(address, rawAmt string, at int64) (float64, float64, error) {
	p.mu.Lock()
	coin, ok := p.mappedCoinInfo[address]
	p.mu.Unlock()
	if ok {
		if coin.Decimals == 0 {
			decimals, symbol, err := p.getDecimalsAndSymbol(address)
			if err != nil {
				return 0, 0, fmt.Errorf("failed to get decimals: %w", err)
			}
			coin.Decimals = decimals
			coin.Coin = symbol
			p.mu.Lock()
			p.mappedCoinInfo[address] = coin
			p.mu.Unlock()
		}

		if coin.Coin == coinUSDT {
			return 1, calculateAmountUsd(rawAmt, coin.Decimals, 1), nil
		}
		price, err := p.getPrice(address, at)
		if err != nil {
			return 0, 0, fmt.Errorf("failed to get price mtm: %w", err)
		}

		return price, calculateAmountUsd(rawAmt, coin.Decimals, price), nil
	}
	return 0, 0, nil
}

func (p *PriceFiller) FullFillTradeLogs(tradeLogs []storageTypes.TradeLog) {
	for idx, tradeLog := range tradeLogs {
		// for the safety, sleep a bit to avoid Binance rate limit
		time.Sleep(10 * time.Millisecond)
		f := p.fullFillTradeLog
		switch tradeLog.Exchange {
		case constant.ExBebop:
			f = p.fullFillBebopTradeLog
		case constant.CowProtocol:
			f = p.fullFillCowProtocolTradeLog
		}
		filledTradeLog, err := f(tradeLog)
		if err != nil {
			p.l.Errorw("Failed to fullFillTradeLog", "err", err, "tradeLog", tradeLog)
			continue
		}
		tradeLogs[idx] = filledTradeLog
	}
}

func (p *PriceFiller) FullFillCowTransferEvents(events []storageTypes.CowTransfer) {
	for idx, event := range events {
		// for the safety, sleep a bit to avoid Binance rate limit
		time.Sleep(10 * time.Millisecond)
		filledEvent, err := p.FullFillCowTransferEvent(event)
		if err != nil {
			p.l.Errorw("Failed to full fill cow transfer event", "err", err, "event", event)
			continue
		}
		events[idx] = filledEvent
	}
}

func (p *PriceFiller) getDecimalsAndSymbol(address string) (int64, string, error) {
	resp, err := p.ksClient.GetTokenCatalog(address)
	if err != nil {
		p.l.Errorw("Failed to GetTokenCatalog", "err", err)
		return 0, "", err
	}
	if len(resp.Data.Tokens) == 1 {
		return resp.Data.Tokens[0].Decimals, resp.Data.Tokens[0].Symbol, nil
	}
	if len(resp.Data.Tokens) > 1 {
		p.l.Warnw("Weird token catalog response", "resp", resp)
		return 0, "", ErrWeirdTokenCatalogResp
	}

	// try to import token if token is not found.
	newResp, err := p.ksClient.ImportToken(NetworkETHChainIDString, address)
	if err != nil {
		p.l.Errorw("Failed to ImportToken", "err", err)
		return 0, "", err
	}
	if len(newResp.Data.Tokens) == 1 {
		return newResp.Data.Tokens[0].Data.Decimals, newResp.Data.Tokens[0].Data.Symbol, nil
	}

	p.l.Warnw("Weird ImportToken response", "resp", newResp)
	return 0, "", ErrWeirdTokenCatalogResp

}

func (p *PriceFiller) insertTokens() error {
	var tokenList []dashboardTypes.Token
	for address, token := range p.mappedCoinInfo {
		tokenList = append(tokenList, dashboardTypes.Token{
			Address:  address,
			Symbol:   token.Coin,
			Decimals: token.Decimals,
			ChainId:  NetworkETHChainID,
		})
	}
	err := p.dashboardStorage.InsertTokens(tokenList)
	if err != nil {
		return err
	}
	return nil
}

func isConnectionRefusedError(err error) bool {
	if err == nil {
		return false
	}
	if errors.Is(err, mtm.ErrRateLimit) {
		return true
	}
	var netErr *net.OpError
	if errors.As(err, &netErr) {
		if strings.Contains(netErr.Err.Error(), "connection refused") {
			return true
		}
	}
	return false
}
