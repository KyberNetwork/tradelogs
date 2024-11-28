package types

import "strings"

type Contract struct {
	Contract     string `db:"contract" json:"contract"`
	ContractName string `db:"contract_name" json:"contract_name"`
}

func (o *Contract) SerializeContract() []interface{} {
	return []interface{}{
		strings.ToLower(o.Contract),
		o.ContractName,
	}
}

func ContractColumns() []string {
	return []string{
		"contract",
		"contract_name",
	}
}
