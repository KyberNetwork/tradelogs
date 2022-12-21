package kyberswap

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/KyberNetwork/tradelogs/internal/storage"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	SwappedLog   = "0xd6d4f5681c246c9f42c203e287975af1601f8df8035a9251f79aab5c8f09e2f8"
	SwappedEvent = "Swapped"
	otcOrderABI  = `[{"anonymous":false,"inputs":[{"indexed":false,"internalType":"address","name":"sender","type":"address"},{"indexed":false,"internalType":"contract IERC20","name":"srcToken","type":"address"},{"indexed":false,"internalType":"contract IERC20","name":"dstToken","type":"address"},{"indexed":false,"internalType":"address","name":"dstReceiver","type":"address"},{"indexed":false,"internalType":"uint256","name":"spentAmount","type":"uint256"},{"indexed":false,"internalType":"uint256","name":"returnAmount","type":"uint256"}],"name":"Swapped","type":"event"}]`
)

var ErrInvalidKSSwappedTopic = errors.New("invalid KS Swapped topic")

type swapped struct {
	Sender       common.Address `json:"sender,omitempty"`
	SrcToken     common.Address `json:"src_token,omitempty"`
	DstToken     common.Address `json:"dst_token,omitempty"`
	DstReceiver  common.Address `json:"dst_receiver,omitempty"`
	SpentAmount  *big.Int       `json:"spent_amount,omitempty"`
	ReturnAmount *big.Int       `json:"return_amount,omitempty"`
}

func (s *swapped) toTradeLogs() storage.TradeLogs {
	return storage.TradeLogs{
		Taker:            s.DstReceiver.Hex(),
		MakerToken:       s.SrcToken.Hex(),
		TakerToken:       s.DstToken.Hex(),
		MakerTokenAmount: s.SpentAmount.String(),
		TakerTokenAmount: s.ReturnAmount.String(),
	}
}

type Parser struct {
	abi *abi.ABI
}

func NewParser() *Parser {
	abi, err := abi.JSON(strings.NewReader(otcOrderABI))
	if err != nil {
		panic(err)
	}
	return &Parser{
		abi: &abi,
	}
}

func (p *Parser) Topics() []string {
	return []string{
		SwappedLog,
	}
}

func (p *Parser) Parse(log types.Log, blockNumber, blockTime uint64) (storage.TradeLogs, error) {
	if len(log.Topics) > 0 && log.Topics[0].Hex() != SwappedLog {
		return storage.TradeLogs{}, ErrInvalidKSSwappedTopic
	}
	var event swapped
	err := p.abi.UnpackIntoInterface(&event, SwappedEvent, log.Data)
	if err != nil {
		return storage.TradeLogs{}, err
	}
	result := event.toTradeLogs()
	result.ContractAddress = log.Address.Hex()
	result.TxHash = log.TxHash.Hex()
	result.BlockNumber = blockNumber
	result.Timestamp = blockTime * 1000 // blocktime is second, timestamp is millisecond
	result.LogIndex = uint64(log.Index)
	fmt.Println(hex.EncodeToString(log.Data))
	return result, nil
}
