package promotees

import (
	"strings"
)

type Promotee struct {
	Promoter    string `db:"promoter" json:"promoter,omitempty"`
	Promotee    string `db:"promotee" json:"promotee,omitempty"`
	Timestamp   uint64 `db:"timestamp" json:"timestamp,omitempty"`
	TxHash      string `db:"tx_hash" json:"tx_hash,omitempty"`
	ChainId     string `db:"chain_id" json:"chain_id,omitempty"`
	BlockNumber uint64 `db:"block_number" json:"block_number,omitempty"`
	Name        string `db:"name" json:"name,omitempty"`
}

type PromoteesQuery struct {
	Promoter string `form:"promoter" json:"promoter,omitempty"`
	Promotee string `form:"promotee" json:"promotee,omitempty"`
	ChainId  string `form:"chain_id" json:"chain_id,omitempty"`
	Name     string `form:"name" json:"name,omitempty"`
}

func (o *Promotee) SerializePromotees() []interface{} {
	return []interface{}{
		strings.ToLower(o.Promoter),
		strings.ToLower(o.Promotee),
		o.Timestamp,
		o.TxHash,
		o.ChainId,
		o.BlockNumber,
	}
}

func (o *Promotee) SerializeNames() []interface{} {
	return []interface{}{
		strings.ToLower(o.Promoter),
		o.Name,
	}
}

func promoteesColumns() []string {
	return []string{
		"promoter",
		"promotee",
		"timestamp",
		"tx_hash",
		"chain_id",
		"block_number",
	}
}

func nameColumns() []string {
	return []string{
		"promoter",
		"name",
	}
}
