package zxrfqv3

import (
	"encoding/hex"
	"fmt"
	"github.com/KyberNetwork/tradelogs/pkg/decoder"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/sha3"
	"testing"
)

func TestGetMethodId(t *testing.T) {
	signature := "DecodeParamOfFillRfqOrderVIP(address,((address,uint256),uint256,uint256),address,bytes,((address,uint256),uint256,uint256),bytes)"

	hash := sha3.NewLegacyKeccak256()
	hash.Write([]byte(signature))
	methodID := hash.Sum(nil)[:4]

	assert.Equal(t, "0xbcd804f7", fmt.Sprintf("0x%x", methodID))
}

func TestCustomMethodId(t *testing.T) {
	// Define the custom method ID
	customAbi := loadABI("./abi/custom_abi.json")
	methodsSupported := []decoder.Bytes4{
		methodIdDecodeParamOfFillOrderSelfFunded,
		methodIdDecodeParamOfFillOrderVIP,
	}
	for _, methodID := range methodsSupported {
		_, err := customAbi.MethodById(methodID.Bytes())
		assert.NoError(t, err, "method ID %s is not supported", hex.EncodeToString(methodID.Bytes()))
	}
}
