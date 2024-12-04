package zxrfqv3

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/KyberNetwork/tradelogs/v2/pkg/constant"
	"github.com/KyberNetwork/tradelogs/v2/pkg/parser/zxrfqv3/deployer"
	"github.com/KyberNetwork/tradelogs/v2/pkg/storage/zerox_deployment"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethereumTypes "github.com/ethereum/go-ethereum/core/types"
	"go.uber.org/zap"
)

type DeployParser struct {
	l               *zap.SugaredLogger
	contractABIs    *ContractABIs
	deployer        *deployer.Deployer
	deployEventHash string
	storage         zerox_deployment.IStorage
}

const (
	deployedEventName = "Deployed"
)

func NewDeployParser(l *zap.SugaredLogger, contractABIs *ContractABIs, d *deployer.Deployer, storage zerox_deployment.IStorage) (*DeployParser, error) {
	deployerABI, err := deployer.DeployerMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get deployer abi: %w", err)
	}
	p := &DeployParser{
		l:               l,
		contractABIs:    contractABIs,
		deployer:        d,
		deployEventHash: deployerABI.Events[deployedEventName].ID.String(),
		storage:         storage,
	}
	return p, p.loadDeployments()
}

func (p *DeployParser) loadDeployments() error {
	deployments, err := p.storage.Get()
	if err != nil {
		p.l.Errorw("failed to get deployments", "err", err)
		return fmt.Errorf("failed to get deployments: %w", err)
	}
	p.l.Infow("get deployments", "deployments", deployments)
	for _, d := range deployments {
		address := common.HexToAddress(d.Address)
		if p.contractABIs.containAddress(address) || isZeroAddress(address) {
			continue
		}

		abi, err := NewABI(ContractABI{Address: address, ContractType: ContractType(d.ContractType)})
		if err != nil {
			p.l.Errorw("failed to create abi for contract", "err", err)
			return fmt.Errorf("create abi error: %w", err)
		}

		p.contractABIs.addContractABI(address, abi)
	}

	p.syncContractAddress()
	return nil
}

func (p *DeployParser) syncContractAddress() {
	contractTypeSupported := []ContractType{SwapContract, GaslessContract}
	ticker := time.NewTicker(time.Hour)
	defer ticker.Stop()
	go func() {
		for ; ; <-ticker.C {
			callOpts := &bind.CallOpts{Context: context.Background()}
			for _, contractType := range contractTypeSupported {
				contractAddress, err := p.deployer.OwnerOf(callOpts, big.NewInt(int64(contractType)))
				if err != nil {
					p.l.Errorw("error to get contract address", "err", err, "contractType", contractType)
					continue
				}

				if p.contractABIs.containAddress(contractAddress) || isZeroAddress(contractAddress) {
					continue
				}

				abi, err := NewABI(ContractABI{Address: contractAddress, ContractType: contractType})
				if err != nil {
					p.l.Errorw("error to create abi", "err", err, "contractAddress", contractAddress, "contractType", contractType)
					continue
				}

				p.l.Infow("add contract abi", "contractAddress", contractAddress)
				p.contractABIs.addContractABI(contractAddress, abi)

				err = p.storage.Insert(zerox_deployment.Deployment{
					Address:      contractAddress.String(),
					ContractType: int(contractType),
				})
				if err != nil {
					p.l.Errorw("failed to insert deployment", "err", err)
				}
			}
		}
	}()
}

func (p *DeployParser) isDeployLog(log ethereumTypes.Log) bool {
	return log.Address.Hex() == constant.Deployer0xV3 && strings.EqualFold(log.Topics[0].String(), p.deployEventHash)
}

func (p *DeployParser) handlerDeployLog(log ethereumTypes.Log) {
	if p.deployer == nil {
		p.l.Errorw("deployer is nil")
		return
	}
	if !p.isDeployLog(log) {
		return
	}
	deployment, err := p.parseDeployLog(log)
	if err != nil {
		p.l.Errorw("failed to parse log", "log", log, "err", err)
		return
	}
	if len(deployment.Address) == 0 {
		return
	}
	err = p.storage.Insert(deployment)
	if err != nil {
		p.l.Errorw("failed to insert deployment", "deployment", deployment, "err", err)
	}
}

func (p *DeployParser) parseDeployLog(log ethereumTypes.Log) (zerox_deployment.Deployment, error) {
	var result zerox_deployment.Deployment
	event, err := p.deployer.ParseDeployed(log)
	if err != nil {
		return result, fmt.Errorf("parse deploy log error %w", err)
	}
	feature, address := event.Arg0, event.Arg2

	contractType := ContractType(feature.Int64())

	if !p.contractABIs.containAddress(address) && !isZeroAddress(address) {
		abi, err := NewABI(ContractABI{Address: address, ContractType: contractType})
		if err != nil {
			return result, fmt.Errorf("create abi error: %w", err)
		}
		p.contractABIs.addContractABI(address, abi)
	}

	result = zerox_deployment.Deployment{
		BlockNumber:  log.BlockNumber,
		ContractType: int(contractType),
		Address:      address.String(),
	}
	return result, nil
}
