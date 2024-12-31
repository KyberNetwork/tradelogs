package types

import "strings"

type TxOrigin struct {
	Address string `db:"address" json:"address"`
	Name    string `db:"name" json:"name"`
}

func (o *TxOrigin) SerializeTxOrigin() []interface{} {
	return []interface{}{
		strings.ToLower(o.Address),
		o.Name,
	}
}

func TxOriginColumns() []string {
	return []string{
		"address",
		"name",
	}
}
