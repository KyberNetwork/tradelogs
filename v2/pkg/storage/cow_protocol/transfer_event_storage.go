package cowprotocol

import (
	"fmt"
	"strings"

	"github.com/KyberNetwork/tradelogs/v2/pkg/constant"
	"github.com/Masterminds/squirrel"
)

func (s *CowTradeStorage) transferTableName() string {
	return constant.TableCowTransfer
}

func (s *CowTradeStorage) InsertCowTransfers(events []CowTransfer) error {
	if len(events) == 0 {
		return nil
	}
	b := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).Insert(s.transferTableName()).Columns(
		cowTransferStorageColumns(false)...,
	)
	for _, event := range events {
		b = b.Values(
			cowTransferStorageSerialize(&event, false)...,
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

func (s *CowTradeStorage) UpdateCowTransfers(events []CowTransfer) error {
	if len(events) == 0 {
		return nil
	}
	b := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).Insert(s.transferTableName()).Columns(
		cowTransferStorageColumns(true)...,
	)
	for _, event := range events {
		b = b.Values(
			cowTransferStorageSerialize(&event, true)...,
		)
	}
	q, p, err := b.Suffix(`ON CONFLICT (transfer_id) DO UPDATE 
		SET 
			token_price=excluded.token_price,
			amount_usd=excluded.amount_usd
	`).ToSql()
	if err != nil {
		s.l.Errorw("Error build insert update", "error", err)
		return err
	}
	if _, err := s.db.Exec(q, p...); err != nil {
		s.l.Errorw("Error exec insert update", "sql", q, "arg", p, "error", err)
		return err
	}
	return nil
}

func (s *CowTradeStorage) GetCowTransfers(query CowTransferQuery) ([]CowTransfer, error) {
	builder := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Select("*").
		From(s.transferTableName())
	if query.FromTime != 0 {
		builder = builder.Where(squirrel.GtOrEq{"timestamp": query.FromTime})
	}
	if query.ToTime != 0 {
		builder = builder.Where(squirrel.LtOrEq{"timestamp": query.ToTime})
	}
	if query.TxHash != "" {
		builder = builder.Where(squirrel.Eq{"tx_hash": strings.ToLower(query.TxHash)})
	}
	q, p, err := builder.OrderBy("timestamp DESC").ToSql()
	if err != nil {
		return nil, err
	}
	var results []CowTransfer
	if err := s.db.Select(&results, q, p...); err != nil {
		return nil, err
	}
	return results, nil
}

func (s *CowTradeStorage) GetEmptyPriceCowTransfers(limit uint64) ([]CowTransfer, error) {
	builder := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Select("*").
		From(s.transferTableName()).Where(squirrel.Eq{"token_price": nil}).Limit(limit)
	q, p, err := builder.ToSql()
	if err != nil {
		return nil, err
	}
	var results []CowTransfer
	if err = s.db.Select(&results, q, p...); err != nil {
		return nil, err
	}
	return results, nil
}

func (s *CowTradeStorage) DeleteCowTransfers(blocks []uint64) error {
	if len(blocks) == 0 {
		return nil
	}
	q, p, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Delete(s.transferTableName()).
		Where(squirrel.Eq{"block_number": blocks}).
		ToSql()
	if err != nil {
		s.l.Errorw("Error while building query for cow transfer", "error", err)
		return fmt.Errorf("failed to build query for cow transfers: %w", err)
	}
	if _, err = s.db.Exec(q, p...); err != nil {
		s.l.Errorw("Error while deleting for cow transfer", "error", err)
		return fmt.Errorf("failed to delete blocks for cow transfers: %w", err)
	}
	return nil
}

func (s *CowTradeStorage) ResetTokenPriceTransfers(token string, from, to int64) (int64, error) {
	builder := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Update(s.transferTableName()).
		Set("token_price", nil).
		Set("amount_usd", nil)
	if token != "" {
		builder = builder.Where(squirrel.Eq{"token": token})
	}
	builder = builder.Where(squirrel.And{
		squirrel.GtOrEq{"timestamp": from},
		squirrel.LtOrEq{"timestamp": to},
	})
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

func cowTransferStorageSerialize(o *CowTransfer, isUpdate bool) []interface{} {
	data := []interface{}{
		strings.ToLower(o.TxHash),
		o.LogIndex,
		strings.ToLower(o.FromAddress),
		strings.ToLower(o.ToAddress),
		strings.ToLower(o.Token),
		o.Timestamp,
		o.BlockNumber,
		o.Amount,
		o.TokenPrice,
		o.AmountUsd,
	}
	if isUpdate {
		return append([]interface{}{o.TransferId}, data...)
	}
	return data
}

func cowTransferStorageColumns(isUpdate bool) []string {
	data := []string{
		"tx_hash",
		"log_index",
		"from_address",
		"to_address",
		"token",
		"timestamp",
		"block_number",
		"amount",
		"token_price",
		"amount_usd",
	}
	if isUpdate {
		return append([]string{"transfer_id"}, data...)
	}
	return data
}
