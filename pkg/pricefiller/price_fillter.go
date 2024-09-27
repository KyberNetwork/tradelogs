package pricefiller

import (
	"context"
	"errors"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/KyberNetwork/go-binance/v2"
	"github.com/KyberNetwork/tradelogs/pkg/storage"
	"go.uber.org/zap"
)

const (
	NetworkETHChanID               = 1
	NetworkETHChanIDString         = "1"
	NetworkETH                     = "ETH"
	updateAllCoinInfoInterval      = time.Hour
	backfillTradeLogsPriceInterval = 10 * time.Minute
	backfillTradeLogsLimit         = 60
	addressETH1                    = "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"
	addressETH2                    = "0x0000000000000000000000000000000000000000"
	coinUSDT                       = "USDT"
	invalidSymbolErrString         = "<APIError> code=-1121, msg=Invalid symbol."
)

var (
	ErrNoPrice               = errors.New(("no price from binance"))
	ErrWeirdTokenCatalogResp = errors.New("weird token catalog response")

	mappedMultiplier = map[string]float64{
		"1MBABYDOGE": 1e-6,
		"1000SATS":   1e-3,
	}
)

type CoinInfo struct {
	Coin            string
	Network         string
	ContractAddress string
	Decimals        int64
}

type PriceFiller struct {
	l              *zap.SugaredLogger
	s              *storage.Storage
	mu             sync.Mutex
	ksClient       *KsClient
	binanceClient  *binance.Client
	mappedCoinInfo map[string]CoinInfo // address - coinInfo
}

func NewPriceFiller(l *zap.SugaredLogger, binanceClient *binance.Client,
	s *storage.Storage) (*PriceFiller, error) {
	p := &PriceFiller{
		l:             l,
		s:             s,
		ksClient:      NewKsClient(),
		binanceClient: binanceClient,
		mappedCoinInfo: map[string]CoinInfo{
			addressETH1: {
				Coin:            "ETH",
				Network:         NetworkETH,
				ContractAddress: addressETH1,
				Decimals:        18,
			},
			addressETH2: {
				Coin:            "ETH",
				Network:         NetworkETH,
				ContractAddress: addressETH2,
				Decimals:        18,
			},
		},
	}

	if err := p.updateAllCoinInfo(); err != nil {
		return nil, err
	}

	go p.runUpdateAllCoinInfoRoutine()
	go p.runBackFillTradelogPriceRoutine()

	return p, nil
}

func (p *PriceFiller) getPrice(token string, timestamp int64) (float64, error) {
	candles, err := p.binanceClient.NewKlinesService().Symbol(withAlias(token) + "USDT").
		Interval("1s").StartTime(timestamp).EndTime(timestamp).Do(context.Background())
	if err != nil {
		return 0, err
	}
	if len(candles) == 0 {
		return 0, ErrNoPrice
	}
	low, err := strconv.ParseFloat(candles[0].Low, 64)
	if err != nil {
		return 0, err
	}
	high, err := strconv.ParseFloat(candles[0].High, 64)
	if err != nil {
		return 0, err
	}
	multiplier := 1.0
	if m, ok := mappedMultiplier[token]; ok {
		multiplier = m
	}
	return multiplier * (low + high) / 2, nil
}

func (p *PriceFiller) updateAllCoinInfo() error {
	resp, err := p.binanceClient.NewAllCoinService().Do(context.Background())
	if err != nil {
		p.l.Errorw("Failed to get all coins info", "err", err)
		return err
	}

	p.mu.Lock()
	defer p.mu.Unlock()
	for _, coinInfo := range resp {
		for _, network := range coinInfo.NetworkList {
			if network.Network == NetworkETH && network.ContractAddress != "" {
				address := strings.ToLower(network.ContractAddress)
				if _, ok := p.mappedCoinInfo[address]; !ok {
					p.mappedCoinInfo[address] = CoinInfo{
						Coin:            network.Coin,
						Network:         network.Network,
						ContractAddress: address,
					}
				}
				break
			}
		}
	}

	p.l.Infow("New mapped coin info", "data", p.mappedCoinInfo)
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

func (p *PriceFiller) runBackFillTradelogPriceRoutine() {
	ticker := time.NewTicker(backfillTradeLogsPriceInterval)
	defer ticker.Stop()

	for range ticker.C {
		tradeLogs, err := p.s.Get(storage.TradeLogsQuery{
			State: string(storage.TradeLogStateNew),
			Limit: backfillTradeLogsLimit,
		})
		if err != nil {
			p.l.Errorw("Failed to get tradeLogs", "err", err)
			continue
		}

		p.FullFillTradeLogs(tradeLogs)
		if err := p.s.Insert(tradeLogs); err != nil {
			p.l.Errorw("Failed to insert tradeLogs", "err", err)
			continue
		}

		p.l.Infow("backfill tradelog price successfully", "trades", tradeLogs)
	}
}

func (p *PriceFiller) fullFillTradeLog(tradeLog storage.TradeLog) (storage.TradeLog, error) {
	makerPrice, makerUsdAmount, err := p.getPriceAndAmountUsd(strings.ToLower(tradeLog.MakerToken),
		tradeLog.MakerTokenAmount, int64(tradeLog.Timestamp))
	if err != nil {
		if err.Error() != invalidSymbolErrString {
			p.l.Errorw("Failed to getPriceAndAmountUsd for maker", "err", err)
			return tradeLog, err
		}
	}

	tradeLog.MakerTokenPrice = makerPrice
	tradeLog.MakerUsdAmount = makerUsdAmount

	takerPrice, takerUsdAmount, err := p.getPriceAndAmountUsd(strings.ToLower(tradeLog.TakerToken),
		tradeLog.TakerTokenAmount, int64(tradeLog.Timestamp))
	if err != nil {
		if err.Error() != invalidSymbolErrString {
			p.l.Errorw("Failed to getPriceAndAmountUsd for taker", "err", err)
			return tradeLog, err
		}
	}

	tradeLog.TakerTokenPrice = takerPrice
	tradeLog.TakerUsdAmount = takerUsdAmount
	tradeLog.State = storage.TradeLogStateProcessed

	return tradeLog, nil
}

func (p *PriceFiller) getPriceAndAmountUsd(address, rawAmt string, at int64) (float64, float64, error) {
	p.mu.Lock()
	coin, ok := p.mappedCoinInfo[address]
	p.mu.Unlock()
	if ok {
		if coin.Decimals == 0 {
			d, err := p.getDecimals(address)
			if err != nil {
				if errors.Is(err, ErrWeirdTokenCatalogResp) {
					return 0, 0, nil
				}
				p.l.Errorw("Failed to getDecimals", "err", err, "address", address)
				return 0, 0, err
			}
			coin.Decimals = d
			p.mu.Lock()
			p.mappedCoinInfo[address] = coin
			p.mu.Unlock()
		}

		if coin.Coin == coinUSDT {
			return 1, calculateAmountUsd(rawAmt, coin.Decimals, 1), nil
		}
		price, err := p.getPrice(coin.Coin, int64(at))
		if err != nil {
			if !errors.Is(err, ErrNoPrice) {
				p.l.Errorw("Failed to getPrice", "err", err, "coin", coin.Coin, "at", at)
				return 0, 0, err
			}
		}

		return price, calculateAmountUsd(rawAmt, coin.Decimals, price), nil
	}
	return 0, 0, nil
}

func (p *PriceFiller) FullFillTradeLogs(tradeLogs []storage.TradeLog) {
	for idx, tradeLog := range tradeLogs {
		// for the safety, sleep a bit to avoid Binance rate limit
		time.Sleep(10 * time.Millisecond)
		filledTradeLog, err := p.fullFillTradeLog(tradeLog)
		if err != nil {
			p.l.Errorw("Failed to fullFillTradeLog", "err", err, "tradeLog", tradeLog)
			continue
		}
		tradeLogs[idx] = filledTradeLog
	}
}

func (p *PriceFiller) getDecimals(address string) (int64, error) {
	resp, err := p.ksClient.GetTokenCatalog(address)
	if err != nil {
		p.l.Errorw("Failed to GetTokenCatalog", "err", err)
		return 0, err
	}

	if len(resp.Data.Tokens) == 1 {
		return resp.Data.Tokens[0].Decimals, nil
	}
	if len(resp.Data.Tokens) > 1 {
		p.l.Warnw("Weird token catalog response", "resp", resp)
		return 0, ErrWeirdTokenCatalogResp
	}

	// try to import token if token is not found.
	newResp, err := p.ksClient.ImportToken(NetworkETHChanIDString, address)
	if err != nil {
		p.l.Errorw("Failed to ImportToken", "err", err)
		return 0, err
	}
	if len(newResp.Data.Tokens) == 1 {
		return newResp.Data.Tokens[0].Data.Decimals, nil
	}

	p.l.Warnw("Weird ImportToken response", "resp", newResp)
	return 0, ErrWeirdTokenCatalogResp
}
