package storage

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/KyberNetwork/tradelogs/internal/testutil"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

type TestCase struct {
	Name    string
	Query   TradeLogsQuery
	Result  []TradeLog
	Success bool
	Err     string
}

func TestSimple(t *testing.T) {
	l := zap.S()
	db, tearDown := testutil.MustNewDevelopmentDB("../../cmd/tradelogs/migrations")
	defer func() {
		assert.NoError(t, tearDown())
	}()
	s := New(l, db)
	var tradelogs []TradeLog
	byteValue, err := os.Open("test.json")
	assert.NoError(t, err)
	assert.NoError(t, json.NewDecoder(byteValue).Decode(&tradelogs), "failed to parse tradelogs")
	assert.NoError(t, s.Insert(tradelogs), "failed to insert tradelogs")
	// test reinsert
	byteValue, err = os.Open("test_2.json")
	assert.NoError(t, err)
	assert.NoError(t, json.NewDecoder(byteValue).Decode(&tradelogs), "failed to parse tradelogs 2")
	assert.NoError(t, s.Insert(tradelogs), "failed to insert tradelogs 2")

	testcase := []TestCase{
		{
			Name: "FromTime-ToTime",
			Query: TradeLogsQuery{
				FromTime: 1671614014900,
				ToTime:   1671614015100,
			},
			Result: []TradeLog{
				{
					Taker:            "0xb94d9c1d39c41fbd61d37cb4235993f4eb4aa160",
					MakerToken:       "0xd4df22556e07148e591b4c7b4f555a17188cf5cf",
					TakerToken:       "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee",
					MakerTokenAmount: "396369187633",
					TakerTokenAmount: "52992604749910750",
					ContractAddress:  "0x617dee16b86534a5d792a4d7a62fb491b544111e",
					BlockNumber:      16232109,
					TxHash:           "0xb61c3c802df945e215d6894a2bc3765e1175ebecdd7148dfc7aa5c9f599b9c13",
					LogIndex:         198,
					Timestamp:        1671614015000,
					Expiry:           1719507125,
				},
			},
			Success: true,
		},
		{
			Name: "Maker",
			Query: TradeLogsQuery{
				Maker: "0xaf0b0000f0210d0f421f0009c72406703b50506b",
			},
			Result: []TradeLog{
				{
					OrderHash:        "0x199cdfca0dad729ec73c03b980964568ee18b7164b1ad42dcfeee05ca789555b",
					Maker:            "0xaf0b0000f0210d0f421f0009c72406703b50506b",
					Taker:            "0x22f9dcf4647084d6c31b2765f6910cd85c178c18",
					MakerToken:       "0x6b175474e89094c44da98b954eedeac495271d0f",
					TakerToken:       "0xdac17f958d2ee523a2206206994597c13d831ec7",
					MakerTokenAmount: "94808075134364237824",
					TakerTokenAmount: "999999999", // should match test_2.json instead of test.json
					ContractAddress:  "0xdef1c0ded9bec7f1a1670819833240f027b25eff",
					BlockNumber:      16232117,
					TxHash:           "0x0e0ec48f90f388a31c637a61ac769b9d0facebff207cb6dc8cf4fc2dacefa55f",
					LogIndex:         202,
					Timestamp:        1671614111000,
					Expiry:           1719507125,
				},
			},
			Success: true,
		},
	}

	for _, test := range testcase {
		result, err := s.Get(test.Query)
		if test.Success {
			assert.NoError(t, err, test.Name)
			assert.Equal(t, test.Result, result, test.Name)
			continue
		}
		assert.EqualError(t, err, test.Err, test.Name)
	}
}
