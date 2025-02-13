package uniswapx

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
	return constant.ExUniswapX
}

func (s *Storage) tableName() string {
	return constant.TableUniswapx
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
			tx_origin=excluded.tx_origin,
			message_sender=excluded.message_sender,
			interact_contract=excluded.interact_contract,
			maker_token_price=excluded.maker_token_price,
			taker_token_price=excluded.taker_token_price,
			maker_usd_amount=excluded.maker_usd_amount,
			taker_usd_amount=excluded.taker_usd_amount,
			expiration=excluded.expiration,
			decay_start_time=excluded.decay_start_time,
			decay_end_time=excluded.decay_end_time,
			exclusive_filler=excluded.exclusive_filler
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

func (s *Storage) ResetTokenPriceToRefetch(token string, from, to int64) (int64, error) {
	builder := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Update(s.tableName()).
		Set("maker_token_price", nil).
		Set("taker_token_price", nil).
		Set("maker_usd_amount", nil).
		Set("taker_usd_amount", nil)
	if token != "" {
		builder = builder.Where(squirrel.Or{
			squirrel.Eq{"maker_token": token},
			squirrel.Eq{"taker_token": token},
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

func tradeLogSerialize(o *storageTypes.TradeLog) []interface{} {
	return []interface{}{
		o.OrderHash,
		strings.ToLower(o.Maker),
		strings.ToLower(o.Taker),
		strings.ToLower(o.MakerToken),
		strings.ToLower(o.TakerToken),
		o.MakerTokenAmount,
		o.TakerTokenAmount,
		strings.ToLower(o.ContractAddress),
		o.BlockNumber,
		o.TxHash,
		o.LogIndex,
		o.Timestamp,
		o.EventHash,
		o.TxOrigin,
		o.MessageSender,
		o.InteractContract,
		o.MakerTokenPrice,
		o.TakerTokenPrice,
		o.MakerUsdAmount,
		o.TakerUsdAmount,
		o.Expiry,
		o.DecayStartTime,
		o.DecayEndTime,
		strings.ToLower(o.ExclusiveFiller),
	}
}

func tradeLogColumns() []string {
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
		"tx_origin",
		"message_sender",
		"interact_contract",
		"maker_token_price",
		"taker_token_price",
		"maker_usd_amount",
		"taker_usd_amount",
		"expiration",
		"decay_start_time",
		"decay_end_time",
		"exclusive_filler",
	}
}
