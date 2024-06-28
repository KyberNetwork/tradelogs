package hashflowv3

type OrderRFQ struct {
	Pool            string  `json:"pool"`
	ExternalAccount string  `json:"externalAccount"`
	Trader          string  `json:"trader"`
	EffectiveTrader string  `json:"effectiveTrader"`
	BaseToken       string  `json:"baseToken"`
	QuoteToken      string  `json:"quoteToken"`
	QuoteExpiry     float64 `json:"quoteExpiry"`
	Nonce           float64 `json:"nonce"`
	Txid            []byte  `json:"txid"`
	Signature       []byte  `json:"signature"`
}
