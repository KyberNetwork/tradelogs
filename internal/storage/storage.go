package storage

import (
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const otcOrderFilledTable = "otc_order_filled"

type Storage struct {
	db *sqlx.DB
	l  *zap.SugaredLogger
}

func New(l *zap.SugaredLogger, db *sqlx.DB) *Storage {
	return &Storage{
		db: db,
		l:  l,
	}
}

func (s *Storage) Insert(orders []OtcOrderFilled) error {
	if len(orders) == 0 {
		return nil
	}
	s.l.Debugw("Request insert", "orders", orders)
	b := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).Insert(otcOrderFilledTable).Columns(
		otcOrderFilledColumns()...,
	)
	for _, order := range orders {
		b = b.Values(
			order.Serialize()...,
		)
	}
	q, p, err := b.Suffix("ON CONFLICT(order_hash) DO NOTHING").ToSql()
	if err != nil {
		s.l.Errorw("Error build insert", "error", err)
		return err
	}
	if _, err := s.db.Exec(q, p...); err != nil {
		s.l.Errorw("Error exec insert", "sql", q, "arg", p, "error", err)
		return err
	}
	return nil
}

func (s *Storage) Get(fromTime uint64, toTime uint64) ([]OtcOrderFilled, error) {
	q, p, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Select(otcOrderFilledColumns()...).
		From(otcOrderFilledTable).
		Where(squirrel.GtOrEq{"timestamp": fromTime}).
		Where(squirrel.LtOrEq{"timestamp": toTime}).
		ToSql()
	if err != nil {
		return nil, err
	}
	var result []OtcOrderFilled
	if err := s.db.Select(&result, q, p...); err != nil {
		return nil, err
	}
	return result, nil
}

func (s *Storage) Delete(orderHash []string) error {
	if len(orderHash) == 0 {
		return nil
	}
	q, p, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Delete(otcOrderFilledTable).
		Where(squirrel.Eq{"order_hash": orderHash}).
		ToSql()
	if err != nil {
		s.l.Errorw("Error while delete", "orderhash", orderHash, "error", err)
		return err
	}
	if _, err := s.db.Exec(q, p...); err != nil {
		return err
	}
	return nil
}

func otcOrderFilledColumns() []string {
	return []string{
		"order_hash",
		"maker",
		"taker",
		"maker_token",
		"taker_token",
		"maker_token_filled_amount",
		"taker_token_filled_amount",
		"address",
		"block_number",
		"tx_hash",
		"timestamp",
	}
}
