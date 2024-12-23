package mtm

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var ErrRateLimit = errors.New("rate limit exceeded")

type MtmClient struct {
	baseURL    string
	httpClient *http.Client
}

func NewMtmClient(baseURL string, httpClient *http.Client) *MtmClient {
	return &MtmClient{
		baseURL:    strings.TrimRight(baseURL, "/"),
		httpClient: httpClient,
	}
}

type Token struct {
	Address  string `json:"address"`
	ChainId  int64  `json:"chain_id"`
	Symbol   string `json:"symbol"`
	Decimals int64  `json:"decimals"`
}
type TokenResponse struct {
	Success bool    `json:"success"`
	Data    []Token `json:"data"`
}

func (m *MtmClient) GetListTokens(ctx context.Context) ([]Token, error) {
	const path = "/tokens"
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, m.baseURL+path, nil)
	if err != nil {
		return nil, fmt.Errorf("new request error: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request error: %w", err)
	}

	var tokens TokenResponse

	if err := json.NewDecoder(resp.Body).Decode(&tokens); err != nil {
		return nil, fmt.Errorf("decode error: %w", err)
	}

	return tokens.Data, nil
}

type RateV3Response struct {
	Success bool `json:"success"`
	Data    struct {
		Price    float64 `json:"price"`
		TimeUnix int64   `json:"timeUnix"`
	} `json:"data"`
}

func (m *MtmClient) GetHistoricalRate(
	ctx context.Context, base, quote string, chainId int64, ts time.Time,
) (float64, error) {
	const path = "/v3/historical"
	params := "?base=" + base +
		"&quote=" + quote +
		"&chain_id=" + strconv.FormatInt(chainId, 10) +
		"&time=" + strconv.FormatInt(ts.Unix(), 10)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, m.baseURL+path+params, nil)
	if err != nil {
		return 0, fmt.Errorf("new request error: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("do request error: %w", err)
	}
	if resp.StatusCode == http.StatusTooManyRequests { // 429
		return 0, ErrRateLimit
	}
	defer resp.Body.Close()

	var rate RateV3Response

	if err := json.NewDecoder(resp.Body).Decode(&rate); err != nil {
		return 0, fmt.Errorf("decode error: %w", err)
	}

	return rate.Data.Price, nil
}
