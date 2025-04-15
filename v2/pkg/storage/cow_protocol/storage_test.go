package cowprotocol

import (
	"testing"

	"github.com/KyberNetwork/tradinglib/pkg/testutil"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

type TestCase struct {
	Name    string
	Query   CowTradeQuery
	Result  []CowTrade
	Success bool
	Err     string
}

func TestSimple(t *testing.T) {
	l := zap.S()
	db, tearDown := testutil.MustNewDevelopmentDB(
		"../../../cmd/migrations",
		testutil.DefaultDSN(),
		testutil.RandomString(8),
	)
	defer func() {
		assert.NoError(t, tearDown())
	}()

	s := New(l, db)

	expectedTrade := CowTrade{
		Owner:           "owner",
		SellToken:       "selltoken",
		BuyToken:        "buytoken",
		SellAmount:      "sellamount",
		BuyAmount:       "buyamount",
		FeeAmount:       "feeamount",
		OrderUid:        "orderuid",
		ContractAddress: "contract_address",
		BlockNumber:     1,
		TxHash:          "tx_hash",
		LogIndex:        1,
		Timestamp:       1,
		EventHash:       "event_hash",
	}

	tradeLogs := []CowTrade{expectedTrade}

	assert.NoError(t, s.UpsertCowTrades(tradeLogs), "failed to insert tradelogs")

	query := CowTradeQuery{
		FromTime: 0,
		ToTime:   10,
	}
	trades, err := s.GetCowTrades(query)
	assert.NoError(t, err)

	assert.Equal(t, []CowTrade{expectedTrade}, trades, "queried trades should match inserted trades")
}
