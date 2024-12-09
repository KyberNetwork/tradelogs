package mtm

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/test-go/testify/require"
)

func TestNewMtmClient(t *testing.T) {
	// need mtm url
	t.Skip()
	MTM_URL := "h"
	httpClient := &http.Client{}
	client, err := NewMtmClient(MTM_URL, httpClient)
	require.NoError(t, err)
	rate, err := client.GetHistoricalRate(context.Background(), "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2", "0xdac17f958d2ee523a2206206994597c13d831ec7", 1, time.UnixMilli(1732870268000))
	require.NoError(t, err)
	fmt.Println("historical rate", rate)

	_, err = client.GetListTokens(context.Background())
	require.NoError(t, err)
	// fmt.Println("tokens", tokens)
}
