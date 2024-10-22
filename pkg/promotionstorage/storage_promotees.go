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
	var query PromoteesQuery
	for i := len(promotees) - 1; i >= 0; i-- {
		promotee := promotees[i]
		query.Promoter = promotee.Promoter
		query.Promotee = promotee.Promotee
		query.ChainId = promotee.ChainId
		pwithname, err := s.Get(query)
		if err != nil {
			s.l.Errorw("Error get info promoter", "error", err)
			return err
		}
		if len(pwithname) > 0 {
			promotees = append(promotees[:i], promotees[i+1:]...)
		}
	}
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
	q, p, err := b.ToSql()
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

func (s *Storage) Get(query PromoteesQuery) ([]PromoteeWithName, error) {
	builder := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Select(promoteesColumns()...).
		From(promoteesTable)
	v := reflect.ValueOf(query)
	types := v.Type()
	for i := 0; i < v.NumField(); i++ {
		tag := string(types.Field(i).Tag.Get("form"))
		if v.Field(i).IsZero() {
			continue
		}
		builder = builder.Where(squirrel.Eq{tag: strings.ToLower(v.Field(i).String())})
	}
	q, p, err := builder.OrderBy("timestamp DESC").ToSql()
	if err != nil {
		return nil, err
	}
	var promotees []Promotee
	if err := s.db.Select(&promotees, q, p...); err != nil {
		return nil, err
	}

	var result []PromoteeWithName
	if len(promotees) == 0 {
		return result, nil
	}

	promoters := make([]string, len(promotees))
	for i, p := range promotees {
		promoters[i] = p.Promoter
	}

	for _, promoter := range promoters {
		fmt.Println(promoter)
	}
	var promoterNames []struct {
		Promoter string `db:"promoter"`
		Name     string `db:"name"`
	}

	q2, p2, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).Select(nameColumns()...).
		From(nameTable).
		Where(squirrel.Eq{"promoter": promoters}).
		ToSql()

	if err != nil {
		return nil, err
	}

	if err := s.db.Select(&promoterNames, q2, p2...); err != nil {
		return nil, err
	}
	fmt.Println(promoterNames)
	nameMap := make(map[string]string)
	for _, pn := range promoterNames {
		nameMap[pn.Promoter] = pn.Name
	}
	fmt.Println(nameMap)

	for i := range promotees {
		if name, ok := nameMap[promotees[i].Promoter]; ok {
			result = append(result, PromoteeWithName{
				Promoter:  promotees[i].Promoter,
				Promotee:  promotees[i].Promotee,
				Timestamp: promotees[i].Timestamp,
				EventHash: promotees[i].EventHash,
				ChainId:   promotees[i].ChainId,
				Name:      name,
			})
		} else {
			result = append(result, PromoteeWithName{
				Promoter:  promotees[i].Promoter,
				Promotee:  promotees[i].Promotee,
				Timestamp: promotees[i].Timestamp,
				EventHash: promotees[i].EventHash,
				ChainId:   promotees[i].ChainId,
				Name:      "",
			})
		}
	}

	return result, nil
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

func nameColumns() []string {
	return []string{
		"promoter",
		"name",
	}
}
