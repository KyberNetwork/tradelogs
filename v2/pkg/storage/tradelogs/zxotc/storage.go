package zxotc

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/KyberNetwork/tradelogs/v2/pkg/constant"
	storageTypes "github.com/KyberNetwork/tradelogs/v2/pkg/storage/tradelogs/types"
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

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

func (s *Storage) Exchange() string {
	return constant.ExZeroX
}

func (s *Storage) tableName() string {
	return constant.TableZeroX
}

func (s *Storage) Insert(orders []storageTypes.TradeLog) error {
	if len(orders) == 0 {
		return nil
	}
	b := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).Insert(s.tableName()).Columns(
		storageTypes.RFQTradeLogColumns()...,
	)
	for _, order := range orders {
		b = b.Values(
			storageTypes.RFQTradeLogSerialize(&order)...,
		)
	}
	q, p, err := b.Suffix(storageTypes.RFQTradeLogSuffix()).ToSql()
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

func (s *Storage) Get(query storageTypes.TradeLogsQuery) ([]storageTypes.TradeLog, error) {
	builder := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Select(storageTypes.RFQTradeLogColumns()...).
		From(s.tableName())
	if query.FromTime != 0 {
		builder = builder.Where(squirrel.GtOrEq{"timestamp": query.FromTime})
	}
	if query.ToTime != 0 {
		builder = builder.Where(squirrel.LtOrEq{"timestamp": query.ToTime})
	}
	v := reflect.ValueOf(query)
	types := v.Type()
	for i := 0; i < v.NumField(); i++ {
		tag := types.Field(i).Tag.Get("form")
		if tag == "from_time" || tag == "to_time" || tag == "limit" {
			continue
		}
		if v.Field(i).IsZero() {
			continue
		}
		builder = builder.Where(squirrel.Eq{tag: strings.ToLower(v.Field(i).String())})
	}
	if query.Limit != 0 {
		builder = builder.Limit(query.Limit)
	}
	q, p, err := builder.OrderBy("timestamp DESC").ToSql()
	if err != nil {
		return nil, err
	}
	var results []storageTypes.TradeLog
	if err = s.db.Select(&results, q, p...); err != nil {
		return nil, err
	}
	for i := range results {
		results[i].Exchange = s.Exchange()
	}
	return results, nil
}

func (s *Storage) GetEmptyPrice(limit uint64) ([]storageTypes.TradeLog, error) {
	builder := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Select(storageTypes.RFQTradeLogColumns()...).
		From(s.tableName()).Where(squirrel.Eq{"maker_token_price": nil}).Limit(limit)
	q, p, err := builder.ToSql()
	if err != nil {
		return nil, err
	}
	var results []storageTypes.TradeLog
	if err = s.db.Select(&results, q, p...); err != nil {
		return nil, err
	}
	return results, nil
}

func (s *Storage) Delete(blocks []uint64) error {
	if len(blocks) == 0 {
		return nil
	}
	q, p, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Delete(s.tableName()).
		Where(squirrel.Eq{"block_number": blocks}).
		ToSql()
	if err != nil {
		s.l.Errorw("Error while building query", "exchange", s.Exchange(), "error", err)
		return fmt.Errorf("failed to build query for exchange %s: %w", s.Exchange(), err)
	}
	if _, err = s.db.Exec(q, p...); err != nil {
		s.l.Errorw("Error while deleting", "exchange", s.Exchange(), "error", err)
		return fmt.Errorf("failed to delete blocks for exchange %s: %w", s.Exchange(), err)
	}
	return nil
}

func (s *Storage) ResetTokenPriceToRefetch(token string, from, to time.Time) (int64, error) {
	builder := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Update(s.tableName()).
		Set("maker_token_price", nil).
		Set("taker_token_price", nil).
		Set("maker_usd_amount", nil).
		Set("taker_usd_amount", nil).
		Where(squirrel.Or{
			squirrel.Eq{"maker_token": token},
			squirrel.Eq{"taker_token": token},
		})
	if !from.IsZero() {
		builder = builder.Where(squirrel.GtOrEq{"timestamp": from.UnixMilli()})
	}
	if !to.IsZero() {
		builder = builder.Where(squirrel.LtOrEq{"timestamp": to.UnixMilli()})
	}
	q, p, err := builder.ToSql()

	if err != nil {
		return 0, fmt.Errorf("build query error: %w", err)
	}
	result, err := s.db.Exec(q, p...)
	if err != nil {
		return 0, fmt.Errorf("run query error: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("fetch rows affected error: %w", err)
	}
	return rowsAffected, nil
}
