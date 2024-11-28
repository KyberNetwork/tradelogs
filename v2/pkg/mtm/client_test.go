package mtm

import (
	"context"
	"fmt"
	"testing"

	"github.com/test-go/testify/require"
)

func TestNewMtmClient(t *testing.T) {
	// need mtm url
	t.Skip()
	MTM_URL := ""
	client := NewMtmClient(MTM_URL)

	rate, err := client.GetCurrentRate(context.Background(), "0x9be89d2a4cd102d8fecc6bf9da793be995c22541", "0x9702230a8ea53601f5cd2dc00fdbc13d4df4a8c7", "1")
	require.NoError(t, err)
	fmt.Println("rate", rate)

	rate, err = client.GetHistoricalRate(context.Background(), "0x9be89d2a4cd102d8fecc6bf9da793be995c22541", "0x9702230a8ea53601f5cd2dc00fdbc13d4df4a8c7", "1", 1732608687)
	require.NoError(t, err)
	fmt.Println("historical rate", rate)

	tokens, err := client.GetListTokens(context.Background())
	require.NoError(t, err)
	fmt.Println("tokens", tokens)
}
