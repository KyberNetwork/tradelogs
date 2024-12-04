package mtm

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	tokenStorage "github.com/KyberNetwork/tradelogs/v2/pkg/storage/dashboard/types"
)

type MtmClient struct {
	host string
}

func NewMtmClient(host string) *MtmClient {
	return &MtmClient{
		host: host,
	}
}

func (m *MtmClient) GetListTokens(ctx context.Context) ([]tokenStorage.Token, error) {
	const path = "/tokens"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, m.host+path, nil)
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()
	var tokens TokenResponse

	if err := json.NewDecoder(resp.Body).Decode(&tokens); err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %w", err)
	}

	return tokens.Data, nil
}

func (m *MtmClient) GetHistoricalRate(
	ctx context.Context,
	base, quote, chainId string,
	timestamp int64,
) (float64, error) {
	const path = "/v3/historical"
	queryParam := fmt.Sprintf("?base=%s&quote=%s&time=%d&chainid=%s",
		base,
		quote,
		timestamp,
		chainId,
	)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, m.host+path+queryParam, nil)
	if err != nil {
		return 0, fmt.Errorf("unable to create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	var rate RateResponse

	if err := json.NewDecoder(resp.Body).Decode(&rate); err != nil {
		return 0, fmt.Errorf("failed to decode JSON: %w", err)
	}

	return rate.Data.Price, nil
}

func (m *MtmClient) GetCurrentRate(ctx context.Context, base, quote, chainId string) (float64, error) {
	const path = "/v3/rate"
	queryParam := fmt.Sprintf("?base=%s&quote=%s&chainid=%s",
		base,
		quote,
		chainId,
	)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, m.host+path+queryParam, nil)
	if err != nil {
		return 0, fmt.Errorf("unable to create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	var rate RateResponse

	if err := json.NewDecoder(resp.Body).Decode(&rate); err != nil {
		return 0, fmt.Errorf("failed to decode JSON: %w", err)
	}

	return rate.Data.Price, nil
}

type RateResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Price    float64 `json:"price"`
		TimeUnix int64   `json:"timeUnix"`
	} `json:"data"`
}

type TokenResponse struct {
	Success bool                 `json:"success"`
	Data    []tokenStorage.Token `json:"data"`
}
