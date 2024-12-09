package state

import (
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const (
	ProcessedBlockKey       = "processed_block"
	EnableSystemBackfillKey = "enable_system_backfill"
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
	q, p, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Select("value").
		From("tradelogs_state").
		Where(squirrel.Eq{"key": key}).ToSql()
	if err != nil {
		s.l.Errorw("failed to build query", "err", err)
		return "", fmt.Errorf("failed to build query: %w", err)
	}
	var data string
	err = s.db.Get(&data, q, p...)
	if err != nil {
		s.l.Errorw("failed to get state", "err", err, "key", key)
		return data, fmt.Errorf("failed to get state: %w", err)
	}
	return data, nil
}

func (s *State) SetState(key string, value string) error {
	q, p, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Insert("tradelogs_state").
		Columns("key", "value").
		Values(key, value).
		Suffix("on conflict (key) do update set value = excluded.value").
		ToSql()
	if err != nil {
		s.l.Errorw("failed to build query", "err", err)
		return fmt.Errorf("failed to build query: %w", err)
	}
	_, err = s.db.Exec(q, p...)
	if err != nil {
		s.l.Errorw("failed to set state", "err", err, "key", key, "value", value)
		return fmt.Errorf("failed to set state: %w", err)
	}
	return nil
}
