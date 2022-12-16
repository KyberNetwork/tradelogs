package zxotc

import (
	"errors"
	"math/big"
	"strings"

	"github.com/KyberNetwork/tradelogs/internal/storage"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	OtcOrderFilledLog   = "0xac75f773e3a92f1a02b12134d65e1f47f8a14eabe4eaf1e24624918e6a8b269f"
	OtcOrderFilledEvent = "OtcOrderFilled"
	otcOrderABI         = `[{"anonymous":false,"inputs":[{"indexed":false,"internalType":"bytes32","name":"orderHash","type":"bytes32"},{"indexed":false,"internalType":"address","name":"maker","type":"address"},{"indexed":false,"internalType":"address","name":"taker","type":"address"},{"indexed":false,"internalType":"address","name":"makerToken","type":"address"},{"indexed":false,"internalType":"address","name":"takerToken","type":"address"},{"indexed":false,"internalType":"uint128","name":"makerTokenFilledAmount","type":"uint128"},{"indexed":false,"internalType":"uint128","name":"takerTokenFilledAmount","type":"uint128"}],"name":"OtcOrderFilled","type":"event"}]`
)

var ErrInvalidOTCTopic = errors.New("invalid OTCFilled topic")

type otcOrderFilled struct {
	OrderHash              common.Hash    `json:"order_hash,omitempty"`
	Maker                  common.Address `json:"maker,omitempty"`
	Taker                  common.Address `json:"taker,omitempty"`
	MakerToken             common.Address `json:"maker_token,omitempty"`
	TakerToken             common.Address `json:"taker_token,omitempty"`
	MakerTokenFilledAmount *big.Int       `json:"maker_token_filled_amount,omitempty"`
	TakerTokenFilledAmount *big.Int       `json:"taker_token_filled_amount,omitempty"`
}

func (o *otcOrderFilled) ToOtcOrderFilled() storage.OtcOrderFilled {
	return storage.OtcOrderFilled{
		OrderHash:              o.OrderHash.Hex(),
		Maker:                  o.Maker.Hex(),
		Taker:                  o.Taker.Hex(),
		MakerToken:             o.MakerToken.Hex(),
		TakerToken:             o.TakerToken.Hex(),
		MakerTokenFilledAmount: o.MakerTokenFilledAmount.String(),
		TakerTokenFilledAmount: o.TakerTokenFilledAmount.String(),
	}
}

type Parser struct {
	abi *abi.ABI
}

func NewParser() (*Parser, error) {
	abi, err := abi.JSON(strings.NewReader(otcOrderABI))
	if err != nil {
		return nil, err
	}
	return &Parser{
		abi: &abi,
	}, nil
}

func (z *Parser) Parse(log types.Log, blockNumber, blockTime uint64) (storage.OtcOrderFilled, error) {
	if log.Topics[0].Hex() != OtcOrderFilledLog {
		return storage.OtcOrderFilled{}, ErrInvalidOTCTopic
	}
	var event otcOrderFilled
	err := z.abi.UnpackIntoInterface(&event, OtcOrderFilledEvent, log.Data)
	if err != nil {
		return storage.OtcOrderFilled{}, err
	}
	result := event.ToOtcOrderFilled()
	result.Address = log.Address.Hex()
	result.TxHash = log.TxHash.Hex()
	result.BlockNumber = blockNumber
	result.Timestamp = blockTime * 1000 // blocktime is second, timestamp is millisecond
	return result, nil
}
