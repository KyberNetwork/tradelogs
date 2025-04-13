package cowtransferparser

import (
	"errors"
	"strings"

	"github.com/KyberNetwork/tradelogs/v2/pkg/constant"
	cowStorage "github.com/KyberNetwork/tradelogs/v2/pkg/storage/cow_protocol"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	ethereumTypes "github.com/ethereum/go-ethereum/core/types"
)

const (
	TransferEvent = "Transfer"
)

type CowTransferParser struct {
	abi               *abi.ABI
	ps                *TransferFilterer
	transferEventHash string
}

func MustNewParser() *CowTransferParser {
	ps, err := NewTransferFilterer(common.Address{}, nil)
	if err != nil {
		panic(err)
	}
	ab, err := TransferMetaData.GetAbi()
	if err != nil {
		panic(err)
	}
	transferEvent, ok := ab.Events[TransferEvent]
	if !ok {
		panic("no such event: Trade")
	}
	return &CowTransferParser{
		ps:                ps,
		abi:               ab,
		transferEventHash: transferEvent.ID.String(),
	}
}

func (p *CowTransferParser) Topics() []string {
	return []string{
		p.transferEventHash,
	}
}

func (p *CowTransferParser) Parse(log ethereumTypes.Log, blockTime uint64) (cowStorage.CowTransfer, error) {
	if len(log.Topics) > 0 && log.Topics[0].Hex() != p.transferEventHash {
		return cowStorage.CowTransfer{}, errors.New("invalid order topic")
	}
	e, err := p.ps.ParseTransfer(log)
	if err != nil {
		return cowStorage.CowTransfer{}, err
	}
	res := cowStorage.CowTransfer{
		TxHash:      log.TxHash.String(),
		BlockNumber: log.BlockNumber,
		Timestamp:   blockTime * 1000,
		FromAddress: e.From.String(),
		ToAddress:   e.To.String(),
		Amount:      e.Value.String(),
		Token:       log.Address.String(),
	}
	return res, nil
}

func (p *CowTransferParser) Exchange() string {
	return constant.CowProtocol
}

func (p *CowTransferParser) IsMatchLog(log ethereumTypes.Log) bool {
	return len(log.Topics) >= 3 &&
		strings.EqualFold(log.Topics[0].String(), p.transferEventHash)
}
