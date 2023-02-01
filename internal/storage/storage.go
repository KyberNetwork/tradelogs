package storage

import (
	"reflect"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const tradeLogsTable = "tradelogs"

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

func (s *Storage) Insert(orders []TradeLog) error {
	if len(orders) == 0 {
		return nil
	}
	s.l.Debugw("Request insert", "orders", orders)
	b := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).Insert(tradeLogsTable).Columns(
		tradelogsColumns()...,
	)
	for _, order := range orders {
		b = b.Values(
			order.Serialize()...,
		)
	}
	q, p, err := b.Suffix("ON CONFLICT(block_number, log_index) DO NOTHING").ToSql()
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

func (s *Storage) Get(query TradeLogsQuery) ([]TradeLog, error) {
	builder := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Select(tradelogsColumns()...).
		From(tradeLogsTable)
	if query.FromTime != 0 {
		builder = builder.Where(squirrel.GtOrEq{"timestamp": query.FromTime})
	}
	if query.ToTime != 0 {
		builder = builder.Where(squirrel.LtOrEq{"timestamp": query.ToTime})
	}
	v := reflect.ValueOf(query)
	types := v.Type()
	for i := 0; i < v.NumField(); i++ {
		tag := string(types.Field(i).Tag.Get("form"))
		if tag == "from_time" || tag == "to_time" {
			continue
		}
		if v.Field(i).IsZero() {
			continue
		}
		builder = builder.Where(squirrel.Eq{tag: strings.ToLower(v.Field(i).String())})
	}
	q, p, err := builder.OrderBy("timestamp DESC").ToSql()
	if err != nil {
		return nil, err
	}
	var result []TradeLog
	if err := s.db.Select(&result, q, p...); err != nil {
		return nil, err
	}
	return result, nil
}

func (s *Storage) Delete(blocks []uint64) error {
	if len(blocks) == 0 {
		return nil
	}
	q, p, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Delete(tradeLogsTable).
		Where(squirrel.Eq{"block_number": blocks}).
		ToSql()
	if err != nil {
		s.l.Errorw("Error while delete", "block_number", blocks, "error", err)
		return err
	}
	if _, err := s.db.Exec(q, p...); err != nil {
		return err
	}
	return nil
}

func tradelogsColumns() []string {
	return []string{
		"order_hash",
		"maker",
		"taker",
		"maker_token",
		"taker_token",
		"maker_token_amount",
		"taker_token_amount",
		"contract_address",
		"block_number",
		"tx_hash",
		"log_index",
		"timestamp",
		"event_hash",
	}
}
