package dashboard

import (
	"reflect"
	"strings"

	types "github.com/KyberNetwork/tradelogs/v2/pkg/storage/dashboard/types"
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const (
	tokenTable    = "token"
	contractTable = "contract"
	eventTable    = "event"
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

// ----------------------------insert----------------------------
func (s *Storage) InsertTokens(tokens []types.Token) error {
	if len(tokens) == 0 {
		return nil
	}

	b := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).Insert(tokenTable).Columns(
		types.TokenColumns()...,
	)
	for _, token := range tokens {
		b = b.Values(
			token.SerializeToken()...,
		)
	}
	q, p, err := b.Suffix(`ON CONFLICT (address, chain_id) DO UPDATE 
		SET 
			symbol=excluded.symbol, 
			decimals=excluded.decimals
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

func (s *Storage) InsertContract(contracts []types.Contract) error {
	if len(contracts) == 0 {
		return nil
	}

	b := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).Insert(contractTable).Columns(
		types.ContractColumns()...,
	)
	for _, contract := range contracts {
		b = b.Values(
			contract.SerializeContract()...,
		)
	}
	q, p, err := b.Suffix(`ON CONFLICT (contract) DO UPDATE 
		SET 
			contract_name=excluded.contract_name
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

// ----------------------------get----------------------------
func (s *Storage) GetTokens(query types.TokenQuery) ([]types.Token, error) {
	builder := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Select(types.TokenColumns()...).
		From(tokenTable)

	v := reflect.ValueOf(query)
	fields := v.Type()
	for i := 0; i < v.NumField(); i++ {
		tag := string(fields.Field(i).Tag.Get("form"))
		if v.Field(i).IsZero() {
			continue
		}
		if tag == "symbol" {
			builder = builder.Where(squirrel.Eq{tag: strings.ToUpper(v.Field(i).String())})
			continue
		}
		builder = builder.Where(squirrel.Eq{tag: strings.ToLower(v.Field(i).String())})
	}

	q, p, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var tokens []types.Token
	if err := s.db.Select(&tokens, q, p...); err != nil {
		return nil, err
	}

	return tokens, nil
}

func (s *Storage) GetContracts() ([]types.Contract, error) {
	builder := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Select(types.ContractColumns()...).
		From(contractTable)

	q, p, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var contracts []types.Contract
	if err := s.db.Select(&contracts, q, p...); err != nil {
		return nil, err
	}

	return contracts, nil
}
