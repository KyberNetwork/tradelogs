package dune

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"time"
)

type Client struct {
	baseURL string
	c       *http.Client
	apiKey  string
}

func NewClient(url, key string, httpClient *http.Client) *Client {
	return &Client{
		baseURL: url,
		c:       httpClient,
		apiKey:  key,
	}
}

type ExecuteQuery struct {
	QueryParams struct {
		EventHash string `json:"event_hash"`
		BlockFrom uint64 `json:"block_from"`
		BlockTo   uint64 `json:"block_to"`
	} `json:"query_parameters"`
}

type ExecuteResponse struct {
	ExecutionID string `json:"execution_id"`
	State       string `json:"state"`
	Error       string `json:"error"`
}

type StateResponse struct {
	ExecutionID         string    `json:"execution_id"`
	QueryID             int       `json:"query_id"`
	IsExecutionFinished bool      `json:"is_execution_finished"`
	State               string    `json:"state"`
	SubmittedAt         time.Time `json:"submitted_at"`
	ExpiresAt           time.Time `json:"expires_at"`
	ExecutionStartedAt  time.Time `json:"execution_started_at"`
	ExecutionEndedAt    time.Time `json:"execution_ended_at"`
	Error               string    `json:"error"`
}

func (c *Client) ExecuteQuery(queryID int64, evenHash string, blockFrom, blockTo uint64) (*ExecuteResponse, error) {
	params := ExecuteQuery{}
	params.QueryParams.EventHash = evenHash
	params.QueryParams.BlockFrom = blockFrom
	params.QueryParams.BlockTo = blockTo
	data, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/v1/query/%d/execute", c.baseURL, queryID), bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("x-dune-api-key", c.apiKey)
	res, err := c.c.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var executeRes ExecuteResponse
	if err = json.NewDecoder(res.Body).Decode(&executeRes); err != nil {
		return nil, err
	}
	if executeRes.Error != "" {
		return nil, fmt.Errorf(executeRes.Error)
	}
	return &executeRes, nil
}

func (c *Client) ExecuteState(exeID string) (*StateResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1/execution/%s/status", c.baseURL, exeID), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("x-dune-api-key", c.apiKey)
	res, err := c.c.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	var state StateResponse
	if err = json.NewDecoder(res.Body).Decode(&state); err != nil {
		return nil, err
	}
	if state.Error != "" {
		return nil, fmt.Errorf(state.Error)
	}
	return &state, nil
}

type DuneLog struct {
	BlockTime       string `json:"block_time"`
	BlockNumber     uint64 `json:"block_number"`
	BlockHash       string `json:"block_hash"`
	ContractAddress string `json:"contract_address"`
	Topic0          string `json:"topic0"`
	Topic1          string `json:"topic1"`
	Topic2          string `json:"topic2"`
	Topic3          string `json:"topic3"`
	Data            string `json:"data"`
	TxHash          string `json:"tx_hash"`
	Index           uint   `json:"index"`
	TxIndex         uint   `json:"tx_index"`
	BlockDate       string `json:"block_date"`
	TxFrom          string `json:"tx_from"`
	TxTo            string `json:"tx_to"`
}

func (c *Client) GetLastestExecuteResult(queryID int64, limit, offset uint64, out any) (uint64, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1/query/%d/results?allow_partial_results=true&limit=%d&offset=%d", c.baseURL, queryID, limit, offset), nil)
	if err != nil {
		return 0, err
	}
	req.Header.Set("x-dune-api-key", c.apiKey)
	res, err := c.c.Do(req)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	var data map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return 0, fmt.Errorf("failed to decode JSON response: %v", err)
	}
	result, ok := data["result"].(map[string]interface{})
	if !ok {
		return 0, fmt.Errorf("error when parse result")
	}
	metadata, ok := result["metadata"].(map[string]interface{})
	if !ok {
		return 0, fmt.Errorf("error when parse metadata")
	}
	rowCount, ok := metadata["total_row_count"].(float64)
	if !ok {
		return 0, fmt.Errorf("error when parse row count")
	}
	rows, ok := result["rows"].([]interface{})
	if !ok {
		return 0, fmt.Errorf("error when parse rows")
	}
	jsonData, err := json.Marshal(rows)
	if err != nil {
		return 0, err
	}
	if err = json.Unmarshal(jsonData, out); err != nil {
		return 0, err
	}

	return uint64(rowCount), nil
}

type OneInchDuneLog struct {
	ContractAddress string `json:"contract_address"`
	EventIndex      uint64 `json:"evt_index"`
	TxHash          string `json:"call_tx_hash"`
	BlockTime       string `json:"call_block_time"`
	BlockNumber     uint64 `json:"call_block_number"`
	Order           string `json:"order"`
	Output0         BigInt `json:"output_0"`
	Output1         BigInt `json:"output_1"`
	Output2         string `json:"output_2"`
}

type BigInt struct {
	*big.Int
}

func (b *BigInt) UnmarshalJSON(data []byte) error {
	var num json.Number
	if err := json.Unmarshal(data, &num); err != nil {
		return err
	}

	b.Int = new(big.Int)
	b.Int, _ = b.Int.SetString(num.String(), 10)
	return nil
}

func (b *BigInt) MarshalJSON() ([]byte, error) {
	if b.Int == nil {
		return []byte("null"), nil
	}
	return []byte(b.String()), nil
}

func (b *BigInt) Hex() string {
	return "0x" + b.Text(16)
}
