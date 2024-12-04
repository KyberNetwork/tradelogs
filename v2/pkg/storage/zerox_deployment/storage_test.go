package zerox_deployment

import (
	"github.com/KyberNetwork/tradinglib/pkg/testutil"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

func TestState_GetState(t *testing.T) {
	l := zap.S()
	db, tearDown := testutil.MustNewDevelopmentDB(
		"../../../cmd/migrations",
		testutil.DefaultDSN(),
		testutil.RandomString(8),
	)
	defer func() {
		assert.NoError(t, tearDown())
	}()
	s := NewStorage(l, db)

	d := Deployment{BlockNumber: 1, ContractType: 2, Address: "test"}
	err := s.Insert(d)
	assert.NoError(t, err)

	val, err := s.Get()
	assert.NoError(t, err)
	assert.Equal(t, d, val[len(val)-1])
}
