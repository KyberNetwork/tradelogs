package promotionstorage

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const (
	promoteesTable = "promotees"
	nameTable      = "promoteesName"
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

func (s *Storage) Insert(promotees []Promotee) error {
	if len(promotees) == 0 {
		return nil
	}

	b := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).Insert(promoteesTable).Columns(
		promoteesColumns()...,
	)
	for _, promotee := range promotees {
		b = b.Values(
			promotee.SerializePromotees()...,
		)
	}
	q, p, err := b.Suffix("ON CONFLICT (promotee, promoter, chain_id) DO NOTHING").ToSql()
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

func (s *Storage) Get(query PromoteesQuery) ([]Promotee, error) {
	builder := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Select(fmt.Sprintf("%s.promoter, promotee, chain_id, event_hash, timestamp, name", promoteesTable)).
		From(promoteesTable).
		Join(fmt.Sprintf("%s ON %s.promoter = %s.promoter", nameTable, promoteesTable, nameTable))

	v := reflect.ValueOf(query)
	types := v.Type()
	for i := 0; i < v.NumField(); i++ {
		tag := string(types.Field(i).Tag.Get("form"))
		if v.Field(i).IsZero() {
			continue
		}
		builder = builder.Where(squirrel.Eq{tag: strings.ToLower(v.Field(i).String())})
	}

	q, p, err := builder.OrderBy(fmt.Sprintf("%s.timestamp DESC", promoteesTable)).ToSql()
	if err != nil {
		return nil, err
	}

	var promotees []Promotee
	if err := s.db.Select(&promotees, q, p...); err != nil {
		return nil, err
	}

	return promotees, nil
}

func (s *Storage) Delete(blocks []uint64) error {
	if len(blocks) == 0 {
		return nil
	}
	q, p, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Delete(promoteesTable).
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

func promoteesColumns() []string {
	return []string{
		"promoter",
		"promotee",
		"timestamp",
		"event_hash",
		"chain_id",
		"block_number",
	}
}
