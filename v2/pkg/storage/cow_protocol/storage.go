package cowprotocol

import (
	"fmt"
	"strings"

	"github.com/KyberNetwork/tradelogs/v2/pkg/constant"
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type CowTradeStorage struct {
	db *sqlx.DB
	l  *zap.SugaredLogger
}

func New(l *zap.SugaredLogger, db *sqlx.DB) *CowTradeStorage {
	return &CowTradeStorage{
		db: db,
		l:  l,
	}
}

func (s *CowTradeStorage) tableName() string {
	return constant.TableCowProtocol
}

func (s *CowTradeStorage) UpsertCowTrades(trades []CowTrade) error {
	if len(trades) == 0 {
		return nil
	}
	b := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).Insert(s.tableName()).Columns(
		cowTradeColumns()...,
	)
	for _, trade := range trades {
		b = b.Values(
			cowTradeSerialize(&trade)...,
		)
	}
	q, p, err := b.Suffix(`ON CONFLICT(block_number, log_index) DO UPDATE 
		SET 
			owner=excluded.owner,
			sell_token=excluded.sell_token,
			buy_token=excluded.buy_token,
			sell_amount=excluded.sell_amount,
			buy_amount=excluded.buy_amount,
			fee_amount=excluded.fee_amount,
			order_uid=excluded.order_uid,
			contract_address=excluded.contract_address,
			tx_hash=excluded.tx_hash,
			timestamp=excluded.timestamp,
			event_hash=excluded.event_hash,
			tx_origin=excluded.tx_origin,
			message_sender=excluded.message_sender,
			interact_contract=excluded.interact_contract,
			buy_token_price=excluded.buy_token_price,
			sell_token_price=excluded.sell_token_price,
			buy_usd_amount=excluded.buy_usd_amount,
			sell_usd_amount=excluded.sell_usd_amount
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

func (s *CowTradeStorage) GetCowTrades(query CowTradeQuery) ([]CowTrade, error) {
	builder := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Select(cowTradeColumns()...).
		From(s.tableName())
	if query.FromTime != 0 {
		builder = builder.Where(squirrel.GtOrEq{"timestamp": query.FromTime})
	}
	if query.ToTime != 0 {
		builder = builder.Where(squirrel.LtOrEq{"timestamp": query.ToTime})
	}
	if query.BlockNumber != 0 {
		builder = builder.Where(squirrel.Eq{"block_number": query.BlockNumber})
	}
	if query.TxHash != "" {
		builder = builder.Where(squirrel.Eq{"tx_hash": query.TxHash})
	}
	if query.Limit != 0 {
		builder = builder.Limit(query.Limit)
	}
	q, p, err := builder.OrderBy("timestamp DESC").ToSql()
	if err != nil {
		return nil, err
	}
	var results []CowTrade
	if err := s.db.Select(&results, q, p...); err != nil {
		return nil, err
	}
	return results, nil
}

func (s *CowTradeStorage) GetEmptyPriceCowTrades(limit uint64) ([]CowTrade, error) {
	builder := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Select(cowTradeColumns()...).
		From(s.tableName()).Where(squirrel.Eq{"sell_token_price": nil}).Limit(limit)
	q, p, err := builder.ToSql()
	if err != nil {
		return nil, err
	}
	var results []CowTrade
	if err = s.db.Select(&results, q, p...); err != nil {
		return nil, err
	}
	return results, nil
}

func (s *CowTradeStorage) DeleteCowTrades(blocks []uint64) error {
	if len(blocks) == 0 {
		return nil
	}
	q, p, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Delete(s.tableName()).
		Where(squirrel.Eq{"block_number": blocks}).
		ToSql()
	if err != nil {
		s.l.Errorw("Error while building query for cow trades", "error", err)
		return fmt.Errorf("failed to build query for cow trades: %w", err)
	}
	if _, err = s.db.Exec(q, p...); err != nil {
		s.l.Errorw("Error while deleting for cow trades", "error", err)
		return fmt.Errorf("failed to delete blocks for cow trades: %w", err)
	}
	return nil
}

func (s *CowTradeStorage) ResetTokenPriceTrades(token string, from, to int64) (int64, error) {
	builder := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Update(s.tableName()).
		Set("sell_token_price", nil).
		Set("buy_token_price", nil).
		Set("sell_usd_amount", nil).
		Set("buy_usd_amount", nil)
	if token != "" {
		builder = builder.Where(squirrel.Or{
			squirrel.Eq{"sell_token": token},
			squirrel.Eq{"buy_token": token},
		})
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

func cowTradeSerialize(o *CowTrade) []interface{} {
	return []interface{}{
		strings.ToLower(o.Owner),
		strings.ToLower(o.SellToken),
		strings.ToLower(o.BuyToken),
		o.SellAmount,
		o.BuyAmount,
		o.FeeAmount,
		strings.ToLower(o.OrderUid),
		strings.ToLower(o.ContractAddress),
		o.BlockNumber,
		o.TxHash,
		o.Timestamp,
		o.EventHash,
		o.TxOrigin,
		o.MessageSender,
		o.InteractContract,
		o.LogIndex,
		o.BuyTokenPrice,
		o.SellTokenPrice,
		o.BuyUsdAmount,
		o.SellUsdAmount,
	}
}

func cowTradeColumns() []string {
	return []string{
		"owner",
		"sell_token",
		"buy_token",
		"sell_amount",
		"buy_amount",
		"fee_amount",
		"order_uid",
		"contract_address",
		"block_number",
		"tx_hash",
		"timestamp",
		"event_hash",
		"tx_origin",
		"message_sender",
		"interact_contract",
		"log_index",
		"buy_token_price",
		"sell_token_price",
		"buy_usd_amount",
		"sell_usd_amount",
	}
}
