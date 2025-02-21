package zerox_deployment

import (
	"fmt"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type IStorage interface {
	Insert(d Deployment) error
	Get() ([]Deployment, error)
}

type Storage struct {
	db *sqlx.DB
	l  *zap.SugaredLogger
}

type Deployment struct {
	BlockNumber  uint64 `db:"block_number" json:"block_number"`
	ContractType int    `db:"contract_type" json:"contract_type"`
	Address      string `db:"contract_address" json:"contract_address"`
}

func NewStorage(l *zap.SugaredLogger, db *sqlx.DB) *Storage {
	return &Storage{db: db, l: l}
}

func (s *Storage) Insert(d Deployment) error {
	q, p, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Insert("zerox_v3_deployment").
		Columns("block_number", "contract_type", "contract_address").
		Values(d.BlockNumber, d.ContractType, strings.ToLower(d.Address)).
		Suffix("on conflict (contract_type, contract_address) do update set block_number = excluded.block_number").
		ToSql()
	if err != nil {
		s.l.Errorw("failed to build query", "err", err)
		return fmt.Errorf("fail to build query: %w", err)
	}
	_, err = s.db.Exec(q, p...)
	if err != nil {
		s.l.Errorw("failed to insert deployment", "err", err)
		return fmt.Errorf("fail to insert deployment: %w", err)
	}
	s.l.Infow("inserted 0xv3 deployment", "block_number", d.BlockNumber, "contract_type", d.ContractType, "contract_address", d.Address)
	return nil
}

func (s *Storage) Get() ([]Deployment, error) {
	q, p, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Select("*").From("zerox_v3_deployment").ToSql()
	if err != nil {
		s.l.Errorw("failed to build query", "err", err)
		return nil, fmt.Errorf("fail to build query: %w", err)
	}
	var d []Deployment
	if err = s.db.Select(&d, q, p...); err != nil {
		return nil, fmt.Errorf("fail to get deployments: %w", err)
	}
	return d, nil
}
