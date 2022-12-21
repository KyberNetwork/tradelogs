package zxotc

import (
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestSimple(t *testing.T) {
	data, _ := hex.DecodeString("848a32ac2a8357254f49b4d4b5f4776f738b9d2c0a7381460908985898492120000000000000000000000000af0b0000f0210d0f421f0009c72406703b50506b00000000000000000000000022f9dcf4647084d6c31b2765f6910cd85c178c180000000000000000000000006b175474e89094c44da98b954eedeac495271d0f000000000000000000000000dac17f958d2ee523a2206206994597c13d831ec7000000000000000000000000000000000000000000000004fba045d36011800000000000000000000000000000000000000000000000000000000000057aa79e")
	p := NewParser()
	var event otcOrderFilled
	err := p.abi.UnpackIntoInterface(&event, OtcOrderFilledEvent, data)
	assert.Nil(t, err)
	makerAmount, _ := big.NewInt(0).SetString("91918545168783998976", 10)
	assert.Equal(t, otcOrderFilled{
		OrderHash:              common.HexToHash("0x848a32ac2a8357254f49b4d4b5f4776f738b9d2c0a7381460908985898492120"),
		Maker:                  common.HexToAddress("0xAF0B0000f0210D0f421F0009C72406703B50506B"),
		Taker:                  common.HexToAddress("0x22F9dCF4647084d6C31b2765F6910cd85C178C18"),
		MakerToken:             common.HexToAddress("0x6B175474E89094C44Da98b954EedeAC495271d0F"),
		TakerToken:             common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7"),
		MakerTokenFilledAmount: makerAmount,
		TakerTokenFilledAmount: big.NewInt(91924382),
	}, event, "otcOrderFilled not match")
}
