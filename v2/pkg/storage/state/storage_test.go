package state

import (
	"testing"

	"github.com/KyberNetwork/tradinglib/pkg/testutil"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
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
	s := New(l, db)
	err := s.SetState("test-key", "test-value")
	assert.NoError(t, err)

	val, err := s.GetState("test-key")
	assert.NoError(t, err)
	assert.Equal(t, "test-value", val)
}
