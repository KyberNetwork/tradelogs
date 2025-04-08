package cowprotocol

import (
	"testing"

	"github.com/KyberNetwork/tradelogs/v2/pkg/storage/tradelogs/types"
	"github.com/KyberNetwork/tradinglib/pkg/testutil"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

type TestCase struct {
	Name    string
	Query   types.TradeLogsQuery
	Result  []types.TradeLog
	Success bool
	Err     string
}

func TestSimple(t *testing.T) {
	l := zap.S()
	db, tearDown := testutil.MustNewDevelopmentDB(
		"../../../../cmd/migrations",
		testutil.DefaultDSN(),
		testutil.RandomString(8),
	)
	defer func() {
		assert.NoError(t, tearDown())
	}()
	s := New(l, db)
	var tradeLogs []types.TradeLog
	tradeLogs = append(tradeLogs, types.TradeLog{
		Exchange:        "cow",
		Owner:           "owner",
		SellToken:       "selltoken",
		BuyToken:        "buytoken",
		SellAmount:      "sellamount",
		BuyAmount:       "buyamount",
		FeeAmount:       "feeamount",
		OrderUid:        "orderuid",
		RawTradeData:    "raw",
		ContractAddress: "contract_address",
		BlockNumber:     1,
		TxHash:          "tx_hash",
		LogIndex:        1,
		Timestamp:       1,
		EventHash:       "event_hash",
	})

	assert.NoError(t, s.Insert(tradeLogs), "failed to insert tradelogs")

	query := types.TradeLogsQuery{
		FromTime: 0,
		ToTime:   10,
	}
	trades, err := s.Get(query)
	assert.NoError(t, err)
	t.Log(trades)

}
