package types

import "strings"

type Solver struct {
	Address    string `db:"address" json:"address"`
	SolverName string `db:"solver_name" json:"solver_name"`
}

func (o *Solver) SerializeSolver() []interface{} {
	return []interface{}{
		strings.ToLower(o.Address),
		o.SolverName,
	}
}

func SolverColumns() []string {
	return []string{
		"address",
		"solver_name",
	}
}
