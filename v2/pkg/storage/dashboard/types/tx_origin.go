package types

import "strings"

type TxOrigin struct {
	Address string `db:"address" json:"address"`
	Name    string `db:"name" json:"name"`
	Type    string `db:"type" json:"type"`
}

type TxOriginQuery struct {
	Type    string `form:"type" json:"type,omitempty"`
	Address string `form:"address" json:"address,omitempty"`
}

func (o *TxOrigin) SerializeTxOrigin() []interface{} {
	return []interface{}{
		strings.ToLower(o.Address),
		o.Name,
		o.Type,
	}
}

func TxOriginColumns() []string {
	return []string{
		"address",
		"name",
		"type",
	}
}
