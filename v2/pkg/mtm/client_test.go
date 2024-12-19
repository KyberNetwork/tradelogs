package mtm

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/test-go/testify/require"
)

func TestNewMtmClient(t *testing.T) {
	// need mtm url
	t.Skip()
	MTM_URL := ""
	httpClient := &http.Client{}
	client := NewMtmClient(MTM_URL, httpClient)
	rate, err := client.GetHistoricalRate(context.Background(), "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2", "0xdac17f958d2ee523a2206206994597c13d831ec7", 1, time.Now().Add(-time.Hour))
	require.NoError(t, err)
	t.Log("historical rate", rate)

	tokens, err := client.GetListTokens(context.Background())
	require.NoError(t, err)
	t.Log("tokens", tokens)
}
