package types

type ContractCallParam struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
	Type  string      `json:"type"`
}
type ContractCall struct {
	ContractType string              `json:"contract_type,omitempty"`
	Name         string              `json:"name"`
	Params       []ContractCallParam `json:"params"`
}
