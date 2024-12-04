package types

import "strings"

type MakerName struct {
	Address string `db:"address" json:"address"`
	Tag     string `db:"tag" json:"tag"`
}

func (o *MakerName) SerializeMakerName() []interface{} {
	return []interface{}{
		strings.ToLower(o.Address),
		o.Tag,
	}
}

func MakerNameColumns() []string {
	return []string{
		"address",
		"tag",
	}
}
