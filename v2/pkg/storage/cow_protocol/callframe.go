package cowprotocol

import (
	"strings"

	"github.com/KyberNetwork/tradelogs/v2/pkg/constant"
	"github.com/Masterminds/squirrel"
)

func (s *CowTradeStorage) callFrameTableName() string {
	return constant.TableCowCallFrame
}

func (s *CowTradeStorage) InsertCowCallFrame(txs []CowTradeCallFrame) error {
	if len(txs) == 0 {
		return nil
	}
	b := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).Insert(s.callFrameTableName()).Columns(
		cowCallFrameColumns()...,
	)
	for _, tx := range txs {
		b = b.Values(
			cowCallframeSerialize(&tx)...,
		)
	}
	q, p, err := b.Suffix(`ON CONFLICT (tx_hash) DO UPDATE 
		SET 
			call_frame=excluded.call_frame
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

func (s *CowTradeStorage) GetCowCallFrame(tx_hash string) ([]CowTradeCallFrame, error) {
	builder := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Select("*").
		From(s.callFrameTableName()).
		Where(squirrel.Eq{"tx_hash": strings.ToLower(tx_hash)})
	q, p, err := builder.ToSql()
	if err != nil {
		return nil, err
	}
	var results []CowTradeCallFrame
	if err := s.db.Select(&results, q, p...); err != nil {
		return nil, err
	}
	return results, nil
}

func cowCallframeSerialize(o *CowTradeCallFrame) []interface{} {
	return []interface{}{
		strings.ToLower(o.TxHash),
		o.CallFrame,
	}
}

func cowCallFrameColumns() []string {
	return []string{
		"tx_hash",
		"call_frame",
	}
}
