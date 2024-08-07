package pricefiller

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const ksSettingUrl = "https://ks-setting.kyberswap.com/api/v1"

type KsClient struct {
	client  *http.Client
	baseURL string
}

func NewKsClient() *KsClient {
	return &KsClient{
		client:  &http.Client{},
		baseURL: ksSettingUrl,
	}
}

func (c *KsClient) DoRequest(ctx context.Context, method, path string, jsonData interface{}, out interface{}) error {
	req, err := createRequest(ctx, method, path, jsonData)
	if err != nil {
		return err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bb, err := readResponse(resp.Body, out)
	if err != nil {
		return fmt.Errorf("readResponse error: %w, data: %s", err, string(bb))
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("server return %d - %v", resp.StatusCode, string(bb))
	}

	return nil
}

func createRequest(ctx context.Context, method, url string, jsonData interface{}) (*http.Request, error) {
	var buf io.Reader
	if jsonData != nil {
		body, err := json.Marshal(jsonData)
		if err != nil {
			return nil, err
		}
		buf = bytes.NewBuffer(body)
	}
	req, err := http.NewRequestWithContext(ctx, method, url, buf)
	if err != nil {
		return nil, err
	}
	if jsonData != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	return req, nil
}

func readResponse(data io.Reader, dataField interface{}) ([]byte, error) {
	if dataField == nil {
		return nil, fmt.Errorf("nil data")
	}
	bb, err := io.ReadAll(data)
	if err != nil {
		return nil, err
	}
	return bb, json.Unmarshal(bb, dataField)
}

type TokenCatalogResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Tokens []TokenCatalog `json:"tokens"`
	}
}

type TokenCatalog struct {
	Decimals int64 `json:"decimals"`
}

func (c *KsClient) GetTokenCatalog(address string) (TokenCatalogResp, error) {
	var resp TokenCatalogResp
	err := c.DoRequest(context.Background(), http.MethodGet,
		fmt.Sprintf("%s/tokens?chainIds=%d&query=%s", c.baseURL, NetworkETHChanID, address),
		nil, &resp)
	if err != nil {
		return TokenCatalogResp{}, err
	}

	if resp.Code != 0 {
		return TokenCatalogResp{}, fmt.Errorf("invalid response code: %d", resp.Code)
	}

	return resp, nil
}

type ImportedToken struct {
	ChainID string `json:"chainId"`
	Address string `json:"address"`
}

type ImportTokenParam struct {
	Tokens []ImportedToken `json:"tokens"`
}

type ImportTokenResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Tokens []struct {
			Data TokenCatalog `json:"data"`
		} `json:"tokens"`
	} `json:"data"`
}

func (c *KsClient) ImportToken(chainID, address string) (ImportTokenResp, error) {
	param := ImportTokenParam{
		Tokens: []ImportedToken{
			{
				ChainID: chainID,
				Address: address,
			},
		},
	}
	var resp ImportTokenResp
	err := c.DoRequest(context.Background(), http.MethodPost, c.baseURL+"/tokens/import", param, &resp)
	if err != nil {
		return ImportTokenResp{}, err
	}

	if resp.Code != 0 {
		return ImportTokenResp{}, fmt.Errorf("invalid response code: %d", resp.Code)
	}

	return resp, nil
}
