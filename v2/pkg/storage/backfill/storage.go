package backfill

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type IStorage interface {
	Get() ([]State, error)
	Update(backfilled uint64) error
}

type Storage struct {
	db *sqlx.DB
	l  *zap.SugaredLogger
}

type State struct {
	Exchange        string `db:"exchange" json:"exchange"`
	DeployBlock     uint64 `db:"deploy_block" json:"deploy_block"`
	BackFilledBlock uint64 `db:"backfilled_block" json:"backfilled_block"`
}

func New(l *zap.SugaredLogger, db *sqlx.DB) *Storage {
	return &Storage{
		db: db,
		l:  l,
	}
}

func (s *Storage) Get() ([]State, error) {
	var infos []State
	// we only need exchanges that not fully filled
	err := s.db.Select(&infos, "SELECT * FROM backfill WHERE backfilled_block > deploy_block OR backfilled_block = 0")
	if err != nil {
		s.l.Errorw("failed to fetch all backfill", "err", err)
		return nil, fmt.Errorf("failed to fetch all backfill state: %w", err)
	}
	return infos, nil
}

func (s *Storage) Update(backfilled uint64) error {
	res, err := s.db.Exec("UPDATE backfill SET backfilled_block=$1 WHERE backfilled_block > $2 OR backfilled_block = 0", backfilled, backfilled)
	if err != nil {
		s.l.Errorw("failed to update backfill", "err", err)
		return fmt.Errorf("failed to update backfill: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		s.l.Errorw("failed to get rows affected", "err", err)
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	s.l.Infow("backfill updated", "rowsAffected", rowsAffected)
	return nil
}
