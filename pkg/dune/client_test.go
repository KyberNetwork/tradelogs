package dune

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

const key = ""

func TestDuneExecute(t *testing.T) {
	t.Skip()
	c := NewClient("https://api.dune.com/api", key, http.DefaultClient)
	res, err := c.ExecuteQuery(3551570, "0x34f57786fb01682fb4eec88d340387ef01a168fe345ea5b76f709d4e560c10eb", 18366958, 19467958)
	require.NoError(t, err)
	t.Log(res.ExecutionID)
}

func TestDuneState(t *testing.T) {
	t.Skip()
	c := NewClient("https://api.dune.com/api", key, http.DefaultClient)
	res, err := c.ExecuteState("01HSB444B5E7WJZDGYS72GFJGY")
	require.NoError(t, err)
	t.Log(res)
}

func TestDuneGetLastestResults(t *testing.T) {
	t.Skip()
	c := NewClient("https://api.dune.com/api", key, http.DefaultClient)
	res, rowcount, err := c.GetLastestExecuteResult(3551570, 10, 1)
	require.NoError(t, err)
	t.Log(res, rowcount)
}
