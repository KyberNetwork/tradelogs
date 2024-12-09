package types

import (
	"strconv"
	"strings"
)

// -------------table-------------
type Token struct {
	Address  string `db:"address" json:"address,omitempty"`
	ChainId  int64  `db:"chain_id" json:"chain_id,omitempty"`
	Symbol   string `db:"symbol" json:"symbol,omitempty"`
	Decimals int64  `db:"decimals" json:"decimals,omitempty"`
}

type TokenQuery struct {
	Address string `form:"address" json:"address,omitempty"`
	Symbol  string `form:"symbol" json:"symbol,omitempty"`
}

func (o *Token) SerializeToken() []interface{} {
	return []interface{}{
		strings.ToLower(o.Address),
		strconv.FormatInt(o.ChainId, 10),
		strings.ToUpper(o.Symbol),
		o.Decimals,
	}
}

func TokenColumns() []string {
	return []string{
		"address",
		"chain_id",
		"symbol",
		"decimals",
	}
}
