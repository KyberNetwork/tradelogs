package mtm

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/KyberNetwork/tradinglib/pkg/httpclient"
)

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

	httpReq, err := httpclient.NewRequestWithContext(
		ctx,
		http.MethodGet,
		m.baseURL,
		path,
		nil,
		nil)
	if err != nil {
		return nil, fmt.Errorf("new request error: %w", err)
	}

	var resp TokenResponse
	if _, err := httpclient.DoHTTPRequest(
		m.httpClient,
		httpReq,
		&resp,
		httpclient.WithStatusCode(http.StatusOK)); err != nil {
		return nil, fmt.Errorf("do http request error: %w", err)
	}

	return resp.Data, nil
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
	httpReq, err := httpclient.NewRequestWithContext(
		ctx,
		http.MethodGet,
		m.baseURL,
		path,
		httpclient.NewQuery().
			SetString("base", base).
			SetString("quote", quote).
			Int64("chain_id", chainId).
			Int64("time", ts.Unix()),
		nil)
	if err != nil {
		return 0, fmt.Errorf("new request error: %w", err)
	}

	var rate RateV3Response
	if _, err := httpclient.DoHTTPRequest(
		m.httpClient,
		httpReq,
		&rate,
		httpclient.WithStatusCode(http.StatusOK)); err != nil {
		return 0, fmt.Errorf("do http request error: %w", err)
	}

	return rate.Data.Price, nil
}
