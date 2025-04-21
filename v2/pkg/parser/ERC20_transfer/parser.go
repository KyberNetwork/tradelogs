package erc20transfer

import (
	"fmt"
	"strings"

	cowStorage "github.com/KyberNetwork/tradelogs/v2/pkg/storage/cow_protocol"
	"github.com/KyberNetwork/tradelogs/v2/pkg/util"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	ethereumTypes "github.com/ethereum/go-ethereum/core/types"
)

const (
	TransferEvent = "Transfer"
)

type ERC20TransferParser struct {
	abi               *abi.ABI
	ps                *ERC20TransferFilterer
	transferEventHash string
}

func MustNewParser() (*ERC20TransferParser, error) {
	ps, err := NewERC20TransferFilterer(common.Address{}, nil)
	if err != nil {
		return nil, fmt.Errorf("error when create new cow transfer filterer %w", err)
	}
	ab, err := ERC20TransferMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("error when get abi of cow transfer metadata %w", err)
	}
	transferEvent, ok := ab.Events[TransferEvent]
	if !ok {
		return nil, fmt.Errorf("no such event: Transfer")
	}
	return &ERC20TransferParser{
		ps:                ps,
		abi:               ab,
		transferEventHash: transferEvent.ID.String(),
	}, nil
}

func (p *ERC20TransferParser) Topics() []string {
	return []string{
		p.transferEventHash,
	}
}

func (p *ERC20TransferParser) Parse(log ethereumTypes.Log, blockTime uint64) (cowStorage.CowTransfer, error) {
	if len(log.Topics) > 0 && log.Topics[0].Hex() != p.transferEventHash {
		return cowStorage.CowTransfer{}, util.ErrInvalidTopic
	}
	e, err := p.ps.ParseTransfer(log)
	if err != nil {
		return cowStorage.CowTransfer{}, err
	}
	res := cowStorage.CowTransfer{
		TxHash:       log.TxHash.String(),
		BlockNumber:  log.BlockNumber,
		Timestamp:    blockTime * 1000,
		FromAddress:  e.From.String(),
		ToAddress:    e.To.String(),
		Amount:       e.Value.String(),
		Token:        log.Address.String(),
		TransferType: string(util.TransferTypeERC20),
	}
	return res, nil
}
func (p *ERC20TransferParser) IsContractLog(log ethereumTypes.Log) bool {
	return len(log.Topics) == 3 &&
		strings.EqualFold(log.Topics[0].String(), p.transferEventHash)
}
