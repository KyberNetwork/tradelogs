package kyberswap

import (
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestSimple(t *testing.T) {
	data, _ := hex.DecodeString("00000000000000000000000097a0e5872b47f31321fcaed57d293058bfa9c338000000000000000000000000dac17f958d2ee523a2206206994597c13d831ec7000000000000000000000000108a850856db3f85d0269a2693d896b394c8032500000000000000000000000097a0e5872b47f31321fcaed57d293058bfa9c338000000000000000000000000000000000000000000000000000000001dcd6500000000000000000000000000000000000000000000000794d2be9a29907e56d0")
	p := NewParser()
	var event swapped
	err := p.abi.UnpackIntoInterface(&event, SwappedEvent, data)
	assert.Nil(t, err)
	returnAmount, _ := big.NewInt(0).SetString("35801869247493378561744", 10)
	assert.Equal(t, swapped{
		Sender:       common.HexToAddress("0x97A0E5872B47f31321fcaed57D293058bfA9C338"),
		SrcToken:     common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7"),
		DstToken:     common.HexToAddress("0x108a850856Db3f85d0269a2693D896B394C80325"),
		DstReceiver:  common.HexToAddress("0x97A0E5872B47f31321fcaed57D293058bfA9C338"),
		SpentAmount:  big.NewInt(500000000),
		ReturnAmount: returnAmount,
	}, event, "swap not match")
}
