package promotionstorage

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/KyberNetwork/tradelogs/internal/testutil"
	"github.com/test-go/testify/assert"
	"go.uber.org/zap"
)

func TestInsert(t *testing.T) {
	db, tearDown := testutil.MustNewDevelopmentDB("../../cmd/promotees/migrations")
	l := zap.S()
	defer func() {
		assert.NoError(t, tearDown())
	}()
	s := New(l, db)
	var promotees []Promotee
	byteValue, err := os.Open("test.json")

	assert.NoError(t, err)
	assert.NoError(t, json.NewDecoder(byteValue).Decode(&promotees), "failed to parse promotees")
	assert.NoError(t, s.Insert(promotees), "failed to insert promotees")
	for _, promotee := range promotees {
		t.Log(promotee)
	}

	// byteValue2, err := os.Open("test2.json")

	// assert.NoError(t, err)
	// assert.NoError(t, json.NewDecoder(byteValue2).Decode(&promotees), "failed to parse promotees")
	// assert.NoError(t, s.Insert(promotees), "failed to insert promotees")
	// t.Log(promotees)
	var query PromoteesQuery
	query.Promoter = ""
	query.Promotee = ""
	query.ChainId = ""
	pwithname, err := s.Get(query)

	assert.NoError(t, err, "Failed to get promotees")

	for _, p := range pwithname {
		t.Log(p)
	}
	blocks := []uint64{3}
	err = s.Delete(blocks)
	assert.NoError(t, err, "Failed to delete")

	pwithname, err = s.Get(query)

	assert.NoError(t, err, "Failed to get promotees")

	for _, p := range pwithname {
		t.Log(p)
	}
}
