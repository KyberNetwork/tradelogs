package promotionstorage

import (
	"strings"
)

type Promotee struct {
	Promoter    string `db:"promoter" json:"promoter,omitempty"`
	Promotee    string `db:"promotee" json:"promotee,omitempty"`
	Timestamp   uint64 `db:"timestamp" json:"timestamp,omitempty"`
	EventHash   string `db:"event_hash" json:"event_hash,omitempty"`
	ChainId     string `db:"chain_id" json:"chain_id,omitempty"`
	BlockNumber uint64 `db:"block_number" json:"block_number,omitempty"`
}

type PromoteeWithName struct {
	Promoter  string `db:"promoter" json:"promoter,omitempty"`
	Promotee  string `db:"promotee" json:"promotee,omitempty"`
	Timestamp uint64 `db:"timestamp" json:"timestamp,omitempty"`
	EventHash string `db:"event_hash" json:"event_hash,omitempty"`
	ChainId   string `db:"chain_id" json:"chain_id,omitempty"`
	Name      string `json:"name,omitempty"`
}

type PromoteesQuery struct {
	Promoter string `form:"promoter" json:"promoter,omitempty"`
	Promotee string `form:"promotee" json:"promotee,omitempty"`
	ChainId  string `form:"chain_id" json:"chain_id,omitempty"`
}

func (o *Promotee) SerializePromotees() []interface{} {
	return []interface{}{
		strings.ToLower(o.Promoter),
		strings.ToLower(o.Promotee),
		o.Timestamp,
		o.EventHash,
		o.ChainId,
		o.BlockNumber,
	}
}
