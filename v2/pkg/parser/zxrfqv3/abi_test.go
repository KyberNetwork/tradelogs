package zxrfqv3

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/sha3"
	"testing"
)

func TestGetMethodId(t *testing.T) {
	type args struct {
		signature        string
		expectedMethodID string
	}

	testCases := []args{
		{
			signature:        "fillRfqOrderSelfFunded(address,((address,uint256),uint256,uint256),address,bytes,address,uint256)",
			expectedMethodID: "0xfd089b1d",
		},
		{
			signature:        "fillRfqOrderVIP(address,((address,uint256),uint256,uint256),address,bytes,((address,uint256),uint256,uint256),bytes)",
			expectedMethodID: "0xfb55075b",
		},
	}
	for _, tt := range testCases {
		hash := sha3.NewLegacyKeccak256()
		hash.Write([]byte(tt.signature))
		methodID := hash.Sum(nil)[:4]
		assert.Equal(t, tt.expectedMethodID, fmt.Sprintf("0x%x", methodID))
	}
}
