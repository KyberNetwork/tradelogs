package promotees

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/KyberNetwork/tradelogs/internal/testutil"
	"github.com/test-go/testify/assert"
	"go.uber.org/zap"
)

func TestInsert(t *testing.T) {
	db, tearDown := testutil.MustNewDevelopmentDB("../../../cmd/migrations")
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

	byteValue2, err := os.Open("test2.json")

	assert.NoError(t, err)
	assert.NoError(t, json.NewDecoder(byteValue2).Decode(&promotees), "failed to parse promotees")
	assert.NoError(t, s.Insert(promotees), "failed to insert promotees")
	t.Log(promotees)
	query := PromoteesQuery{}
	pwithname, err := s.Get(query)

	assert.NoError(t, err, "Failed to get promotees")

	for _, p := range pwithname {
		t.Log(p)
	}
	blocks := []uint64{3}
	err = s.Delete(blocks)
	assert.NoError(t, err, "Failed to delete")

	var promotees_name []Promotee
	promotees_name = append(promotees_name, Promotee{
		Promoter: "0xa8be6b2afe6e060985675675615c2108a66135c8",
		Name:     "test",
	})
	assert.NoError(t, s.InsertPromoterName(promotees_name), "failed to insert promotee name")
	pwithname, err = s.Get(query)

	assert.NoError(t, err, "Failed to get promotees")

	for _, p := range pwithname {
		if p.Name == "test" {
			t.Log(p)
		}

	}
	t.Log(s.CheckPromoteeExist("0x0b8a49d816cc709b6eadb09498030ae3416b66dc"))
}
