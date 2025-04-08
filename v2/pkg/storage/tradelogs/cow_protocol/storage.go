package cowprotocol

import (
	"fmt"
	"reflect"
	"strings"

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
	return constant.CowProtocol
}

func (s *Storage) tableName() string {
	return constant.TableCowProtocol
}

func (s *Storage) Insert(orders []storageTypes.TradeLog) error {
	if len(orders) == 0 {
		return nil
	}
	b := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).Insert(s.tableName()).Columns(
		tradeLogColumns()...,
	)
	for _, order := range orders {
		b = b.Values(
			tradeLogSerialize(&order)...,
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
			raw_trade_data=excluded.raw_trade_data,
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

func (s *Storage) Get(query storageTypes.TradeLogsQuery) ([]storageTypes.TradeLog, error) {
	builder := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Select(tradeLogColumns()...).
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
	if err := s.db.Select(&results, q, p...); err != nil {
		return nil, err
	}
	for i := range results {
		results[i].Exchange = s.Exchange()
	}
	return results, nil
}

func (s *Storage) GetEmptyPrice(limit uint64) ([]storageTypes.TradeLog, error) {
	builder := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Select(tradeLogColumns()...).
		From(s.tableName()).Where(squirrel.Eq{"sell_token_price": nil}).Limit(limit)
	q, p, err := builder.ToSql()
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

func tradeLogSerialize(o *storageTypes.TradeLog) []interface{} {
	return []interface{}{
		strings.ToLower(o.Owner),
		strings.ToLower(o.SellToken),
		strings.ToLower(o.BuyToken),
		o.SellAmount,
		o.BuyAmount,
		o.FeeAmount,
		strings.ToLower(o.OrderUid),
		o.RawTradeData,
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

func tradeLogColumns() []string {
	return []string{
		"owner",
		"sell_token",
		"buy_token",
		"sell_amount",
		"buy_amount",
		"fee_amount",
		"order_uid",
		"raw_trade_data",
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

func (s *Storage) ResetTokenPriceToRefetch(token string, from, to int64) (int64, error) {
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
