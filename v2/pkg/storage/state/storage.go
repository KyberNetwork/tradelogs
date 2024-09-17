package state

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const (
	ProcessedBlockKey = "processed_block"
)

type Storage interface {
	GetState(key string) (string, error)
	SetState(key string, value string) error
}

type State struct {
	db *sqlx.DB
	l  *zap.SugaredLogger
}

func New(l *zap.SugaredLogger, db *sqlx.DB) Storage {
	return &State{
		db: db,
		l:  l,
	}
}

func (s *State) GetState(key string) (string, error) {
	var data string
	err := s.db.Get(&data, `SELECT value FROM tradelogs_state WHERE key=$1`, key)
	if err != nil {
		return data, fmt.Errorf("failed to get state: %w", err)
	}
	return data, nil
}

func (s *State) SetState(key string, value string) error {
	_, err := s.db.Exec(
		`INSERT INTO tradelogs_state (key,value) VALUES ($1,$2) ON CONFLICT (key) DO UPDATE SET value=excluded.value`,
		key, value)
	if err != nil {
		return fmt.Errorf("failed to set state: %w", err)
	}
	return nil
}
