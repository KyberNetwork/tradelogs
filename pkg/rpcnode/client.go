package rpcnode

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/KyberNetwork/tradelogs/internal/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"io"
	"net/http"
)

type Client struct {
	rpcUrl    string
	ethClient *ethclient.Client
}

func NewClient(rpcUrl string) (*Client, error) {
	ethClient, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return nil, err
	}
	return &Client{
		rpcUrl:    rpcUrl,
		ethClient: ethClient,
	}, nil
}

func (c *Client) FetchTraceCall(ctx context.Context, txHash string) (types.CallFrame, error) {
	payload := map[string]interface{}{
		"method":  "debug_traceTransaction",
		"id":      1,
		"jsonrpc": "2.0",
		"params": []interface{}{
			txHash,
			map[string]interface{}{
				"tracer": "callTracer",
				"tracerConfig": map[string]interface{}{
					"withLog": true,
				},
			},
		},
	}

	b, err := json.Marshal(payload)
	if err != nil {
		return types.CallFrame{}, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", c.rpcUrl, bytes.NewBuffer(b))

	if err != nil {
		return types.CallFrame{}, err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return types.CallFrame{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return types.CallFrame{}, err
	}

	var rpcResponse types.TraceCallResponse
	if err := json.Unmarshal(body, &rpcResponse); err != nil {
		return types.CallFrame{}, err
	}

	return rpcResponse.Result, err
}
