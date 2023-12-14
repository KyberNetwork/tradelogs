package worker

import "github.com/KyberNetwork/tradelogs/pkg/storage"

type parseEventLogResult struct {
	tradeLogOrder storage.TradeLog
	err           error
}
