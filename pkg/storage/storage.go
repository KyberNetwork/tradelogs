package storage

import (
	"reflect"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const (
	tradeLogsTable = "tradelogs"
	errorLogsTable = "errorlogs"
)

type EVMLog struct {
	Address     string `db:"address"`
	Topics      string `db:"topics"`
	Data        []byte `db:"data"`
	BlockNumber uint64 `db:"block_number"`
	TxHash      string `db:"tx_hash"`
	TxIndex     uint   `db:"tx_index"`
	BlockHash   string `db:"block_hash"`
	Index       uint   `db:"log_index"`
	Removed     bool   `db:"removed"`
	Time        uint64 `db:"time"`
}

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
	b := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).Insert(tradeLogsTable).Columns(
		tradelogsColumns()...,
	)
	for _, order := range orders {
		b = b.Values(
			order.Serialize()...,
		)
	}
	q, p, err := b.Suffix(`ON CONFLICT(block_number, log_index) DO UPDATE 
		SET 
			order_hash=excluded.order_hash,
			maker=excluded.maker,
			taker=excluded.taker,
			maker_token=excluded.maker_token,
			taker_token=excluded.taker_token,
			maker_token_amount=excluded.maker_token_amount,
			taker_token_amount=excluded.taker_token_amount,
			contract_address=excluded.contract_address,
			block_number=excluded.block_number,
			tx_hash=excluded.tx_hash,
			log_index=excluded.log_index,
			timestamp=excluded.timestamp,
			event_hash=excluded.event_hash,
			maker_traits=excluded.maker_traits,
			expiration_date=excluded.expiration_date,
			maker_token_price=excluded.maker_token_price,
			taker_token_price=excluded.taker_token_price,
			maker_usd_amount=excluded.maker_usd_amount,
			taker_usd_amount=excluded.taker_usd_amount,
			state=excluded.state
	`).ToSql()
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
		"maker_traits",
		"expiration_date",
		"maker_token_price",
		"taker_token_price",
		"maker_usd_amount",
		"taker_usd_amount",
		"state",
	}
}

func (s *Storage) InsertErrorLog(log EVMLog) error {
	b := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).Insert(errorLogsTable).Columns(
		[]string{"address", "topics", "data", "block_number", "tx_hash",
			"tx_index", "block_hash", "log_index", "time"}...,
	).Values(log.Address, log.Topics, log.Data, log.BlockNumber, log.TxHash,
		log.TxIndex, log.BlockHash, log.Index, log.Time)
	q, p, err := b.Suffix(`ON CONFLICT(block_number, log_index) DO UPDATE 
		SET 
		address=excluded.address,
		topics=excluded.topics,
		data=excluded.data,
		block_number=excluded.block_number,
		tx_hash=excluded.tx_hash,
		tx_index=excluded.tx_index,
		block_hash=excluded.block_hash,
		log_index=excluded.log_index,
		time=excluded.time
	`).ToSql()
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

func (s *Storage) GetErrorLogs() ([]EVMLog, error) {
	q, p, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Select("*").
		From(errorLogsTable).ToSql()
	if err != nil {
		return nil, err
	}

	var result []EVMLog
	if err := s.db.Select(&result, q, p...); err != nil {
		return nil, err
	}
	return result, nil
}

func (s *Storage) DeleteErrorLogsWithBlock(block_hash string) error {
	q, p, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Delete(errorLogsTable).
		Where(squirrel.Eq{"block_hash": block_hash}).
		ToSql()
	if err != nil {
		s.l.Errorw("Error while delete", "block_number", block_hash, "error", err)
		return err
	}
	if _, err := s.db.Exec(q, p...); err != nil {
		return err
	}
	return nil
}

func (s *Storage) DeleteErrorLogsWithLogIndex(block uint64, logIndex uint) error {
	q, p, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Delete(errorLogsTable).
		Where(squirrel.Eq{"block_number": block, "log_index": logIndex}).
		ToSql()
	if err != nil {
		s.l.Errorw("Error while delete", "block_number", block, "log_index", logIndex, "error", err)
		return err
	}
	if _, err := s.db.Exec(q, p...); err != nil {
		return err
	}
	return nil
}

func (s *Storage) GetErrorLogsSince(t int64) ([]EVMLog, error) {
	q, p, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Select("*").
		From(errorLogsTable).
		Where(squirrel.GtOrEq{"time": t}).
		ToSql()

	if err != nil {
		return nil, err
	}

	var result []EVMLog
	if err := s.db.Select(&result, q, p...); err != nil {
		return nil, err
	}
	return result, nil
}
