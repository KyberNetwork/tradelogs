package hashflow

type OrderRFQ struct {
	Pool                     string   `json:"pool"`
	ExternalAccount          string   `json:"externalAccount"`
	Trader                   string   `json:"trader"`
	EffectiveTrader          string   `json:"effectiveTrader"`
	BaseToken                string   `json:"baseToken"`
	QuoteToken               string   `json:"quoteToken"`
	EffectiveBaseTokenAmount float64  `json:"effectiveBaseTokenAmount"`
	BaseTokenAmount          float64  `json:"baseTokenAmount"`
	QuoteTokenAmount         float64  `json:"quoteTokenAmount"`
	QuoteExpiry              float64  `json:"quoteExpiry"`
	Nonce                    float64  `json:"nonce"`
	TxId                     [32]byte `json:"txid"`
	Signature                []byte   `json:"signature"`
}
