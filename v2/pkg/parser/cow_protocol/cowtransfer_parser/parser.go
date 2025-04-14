package cowtransferparser

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
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

func MustNewParser() (*CowTransferParser, error) {
	ps, err := NewTransferFilterer(common.Address{}, nil)
	if err != nil {
		return nil, fmt.Errorf("error when create new cow transfer filterer %w", err)
	}
	ab, err := TransferMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("error when get abi of cow transfer metadata %w", err)
	}
	transferEvent, ok := ab.Events[TransferEvent]
	if !ok {
		return nil, fmt.Errorf("no such event: Transfer")
	}
	return &CowTransferParser{
		ps:                ps,
		abi:               ab,
		transferEventHash: transferEvent.ID.String(),
	}, nil
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
	trasferId := generateHash(res)
	res.TransferId = trasferId
	return res, nil
}

func (p *CowTransferParser) Exchange() string {
	return constant.CowProtocol
}

func (p *CowTransferParser) IsMatchLog(log ethereumTypes.Log) bool {
	return len(log.Topics) >= 3 &&
		strings.EqualFold(log.Topics[0].String(), p.transferEventHash)
}

func generateHash(cowTransfer cowStorage.CowTransfer) string {
	data := fmt.Sprintf("%s|%d|%d|%s|%s|%s|%s",
		cowTransfer.TxHash,
		cowTransfer.BlockNumber,
		cowTransfer.Timestamp,
		cowTransfer.FromAddress,
		cowTransfer.ToAddress,
		cowTransfer.Amount,
		cowTransfer.Token,
	)

	hash := sha256.New()
	hash.Write([]byte(data))
	hashBytes := hash.Sum(nil)

	return hex.EncodeToString(hashBytes)
}
