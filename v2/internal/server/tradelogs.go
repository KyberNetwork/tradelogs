package server

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/KyberNetwork/tradelogs/v2/pkg/constant"
	cowProtocolStorage "github.com/KyberNetwork/tradelogs/v2/pkg/storage/cow_protocol"
	dashboardStorage "github.com/KyberNetwork/tradelogs/v2/pkg/storage/dashboard"
	dashboardTypes "github.com/KyberNetwork/tradelogs/v2/pkg/storage/dashboard/types"
	storageTypes "github.com/KyberNetwork/tradelogs/v2/pkg/storage/tradelogs/types"
	"github.com/KyberNetwork/tradelogs/v2/pkg/storage/zerox_deployment"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	maxTimeRange        = uint64(24 * time.Hour.Milliseconds())
	ErrExchangeNotFound = errors.New("exchange not found")
)

const (
	maxRangeEmptyAddress    = 7 * 24 * time.Hour
	maxRangeNonEmptyAddress = 30 * 24 * time.Hour
)

type resetTokenPriceParams struct {
	Address  string `form:"address" json:"address"`
	Exchange string `form:"exchange" json:"exchange"`
	From     int64  `form:"from_ts" json:"from_ts"`
	To       int64  `form:"to_ts" json:"to_ts"`
	Rows     int64  `form:"rows" json:"rows"`
}

type TradeLogs struct {
	r               *gin.Engine
	bindAddr        string
	l               *zap.SugaredLogger
	storage         []storageTypes.Storage
	dashStorage     *dashboardStorage.Storage
	deployStorage   zerox_deployment.IStorage
	cowTradeStorage *cowProtocolStorage.CowTradeStorage
}

func NewTradeLogs(
	l *zap.SugaredLogger,
	s []storageTypes.Storage,
	dashStorage *dashboardStorage.Storage,
	deployStorage zerox_deployment.IStorage,
	cowTradeStorage *cowProtocolStorage.CowTradeStorage,
	bindAddr string,
) *TradeLogs {
	engine := gin.New()
	engine.Use(gin.Recovery())

	server := &TradeLogs{
		r:               engine,
		bindAddr:        bindAddr,
		l:               l,
		storage:         s,
		dashStorage:     dashStorage,
		deployStorage:   deployStorage,
		cowTradeStorage: cowTradeStorage,
	}

	gin.SetMode(gin.ReleaseMode)
	server.register()

	return server
}

// Run runs server.
func (s *TradeLogs) Run() error {
	if err := s.r.Run(s.bindAddr); err != nil {
		return fmt.Errorf("run server: %w", err)
	}

	return nil
}

func (s *TradeLogs) register() {
	pprof.Register(s.r, "/debug")
	s.r.GET("/tradelogs", s.getTradeLogs)
	s.r.GET("/tokens", s.getTokens)
	s.r.GET("/makers", s.getMakerName)
	s.r.POST("/makers", s.addMakerName)
	s.r.GET("/txorigin", s.getTxOrigin)
	s.r.POST("/txorigin", s.addTxOrigin)
	s.r.POST("/price_filler/refetch", s.resetTokenPriceToRefetch)
	s.r.GET("/0xv3_deployment", s.get0xv3Deployment)
	s.r.GET("/cow/tx/:tx_hash", s.getInfoCowTx)
}

func (s *TradeLogs) getTradeLogs(c *gin.Context) {
	var (
		query storageTypes.TradeLogsQuery
		err   = c.ShouldBind(&query)
	)

	if err != nil {
		responseErr(c, http.StatusBadRequest, err)
		return
	}

	if query.ToTime < query.FromTime {
		responseErr(c, http.StatusBadRequest, fmt.Errorf("to_time cannot smaller than from_time"))
		return
	}

	if query.ToTime-query.FromTime > maxTimeRange {
		responseErr(c, http.StatusBadRequest, fmt.Errorf("max time range: %v", maxTimeRange))
		return
	}

	s.l.Infow("get trade log query", "query", query)

	var data []storageTypes.TradeLog
	for _, storage := range s.storage {
		tradeLogs, err := storage.Get(query)
		if err != nil {
			responseErr(c, http.StatusInternalServerError, err)
			return
		}
		data = append(data, tradeLogs...)
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}

func (s *TradeLogs) getTokens(c *gin.Context) {
	var queries dashboardTypes.TokenQuery
	if err := c.ShouldBind(&queries); err != nil {
		responseErr(c, http.StatusBadRequest, err)
		return
	}

	data, err := s.dashStorage.GetTokens(queries)
	if err != nil {
		responseErr(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}

func (s *TradeLogs) getMakerName(c *gin.Context) {
	data, err := s.dashStorage.GetMakerName()
	if err != nil {
		responseErr(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}

func (s *TradeLogs) addMakerName(c *gin.Context) {
	var queries []dashboardTypes.MakerName

	if err := c.ShouldBindJSON(&queries); err != nil {
		responseErr(c, http.StatusBadRequest, err)
		return
	}

	if err := s.dashStorage.InsertMakerName(queries); err != nil {
		responseErr(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    queries,
	})
}

func (s *TradeLogs) getTxOrigin(c *gin.Context) {
	var queries dashboardTypes.TxOriginQuery
	if err := c.ShouldBind(&queries); err != nil {
		responseErr(c, http.StatusBadRequest, err)
		return
	}
	data, err := s.dashStorage.GetTxOrigin(queries)
	if err != nil {
		responseErr(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}

func (s *TradeLogs) addTxOrigin(c *gin.Context) {
	var queries []dashboardTypes.TxOrigin

	if err := c.ShouldBindJSON(&queries); err != nil {
		responseErr(c, http.StatusBadRequest, err)
		return
	}

	if err := s.dashStorage.InsertTxOrigin(queries); err != nil {
		responseErr(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    queries,
	})
}

func (s *TradeLogs) resetTokenPriceToRefetch(c *gin.Context) {
	var query resetTokenPriceParams
	if err := c.ShouldBindJSON(&query); err != nil {
		responseErr(c, http.StatusBadRequest, err)
		return
	}

	query, err := validateResetTokenPriceParams(query)
	if err != nil {
		responseErr(c, http.StatusBadRequest, err)
		return
	}
	var rows int64
	switch query.Exchange {
	case constant.CowProtocol:
		rows, err = s.cowTradeStorage.ResetTokenPriceTrades(query.Address, query.From, query.To)
	case constant.CowTransfer:
		rows, err = s.cowTradeStorage.ResetTokenPriceTransfers(query.Address, query.From, query.To)
	default:
		rows, err = s.resetTokenPriceTradelogs(query.Exchange, query.Address, query.From, query.To)
	}
	if errors.Is(err, ErrExchangeNotFound) {
		responseErr(c, http.StatusBadRequest, ErrExchangeNotFound)
		return
	}
	if err != nil {
		responseErr(c, http.StatusInternalServerError, err)
		return
	}
	query.Rows = rows
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    query,
	})
}

func (s *TradeLogs) resetTokenPriceTradelogs(exchange, token string, from int64, to int64) (int64, error) {
	for _, storage := range s.storage {
		if storage.Exchange() != exchange {
			continue
		}
		rows, err := storage.ResetTokenPriceToRefetch(token, from, to)
		return rows, err
	}
	return 0, ErrExchangeNotFound
}

func validateResetTokenPriceParams(query resetTokenPriceParams) (resetTokenPriceParams, error) {
	now := time.Now()
	if query.Address == "" && query.From == 0 && query.To == 0 {
		return query, fmt.Errorf("address is empty and no valid time range provided")
	}

	// Validate `From` timestamp: it cannot be in the future or negative.
	if query.From > now.UnixMilli() || query.From < 0 {
		return query, fmt.Errorf("invalid from_ts")
	}

	// Validate `To` timestamp: it can not be negative,
	// and if provided, it must be greater than or equal to `From`.
	if query.To < 0 || (query.To > 0 && query.To < query.From) {
		return query, fmt.Errorf("invalid to_ts")
	}
	timeRange := maxRangeEmptyAddress
	if len(query.Address) > 0 {
		timeRange = maxRangeNonEmptyAddress
	}

	if query.From == 0 && query.To == 0 {
		query.From = now.Add(-timeRange).UnixMilli()
		query.To = now.UnixMilli()
		return query, nil
	}

	// If `From` is provided but `To` is not, calculate `To` as `From + max range`
	if query.From > 0 && query.To == 0 {
		query.To = min(time.UnixMilli(query.From).Add(timeRange).UnixMilli(), now.UnixMilli())
		return query, nil
	}

	// Adjust `From` as `To - max range
	query.To = min(query.To, now.UnixMilli())
	query.From = max(time.UnixMilli(query.To).Add(-timeRange).UnixMilli(), query.From)
	return query, nil
}

func (s *TradeLogs) get0xv3Deployment(c *gin.Context) {
	result, err := s.deployStorage.Get()
	if err != nil {
		responseErr(c, http.StatusInternalServerError, fmt.Errorf("get 0x deployment failed: %v", err))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    result,
	})
}

func (s *TradeLogs) getInfoCowTx(c *gin.Context) {
	txHash := c.Param("tx_hash")
	if txHash == "" {
		responseErr(c, http.StatusBadRequest, fmt.Errorf("tx_hash is required"))
		return
	}
	type result struct {
		CowTrades   []cowProtocolStorage.CowTrade          `json:"cow_trades"`
		CowTransfer []cowProtocolStorage.CowTransfer       `json:"cow_transfers"`
		CallFrame   []cowProtocolStorage.CowTradeCallFrame `json:"callframe"`
	}

	// get cow trades
	cowTradeQuery := cowProtocolStorage.CowTradeQuery{
		TxHash: txHash,
	}
	cowTrades, err := s.cowTradeStorage.GetCowTrades(cowTradeQuery)
	if err != nil {
		responseErr(c, http.StatusInternalServerError, fmt.Errorf("error when get cow trades: %w", err))
		return
	}

	if len(cowTrades) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data": result{
				CowTrades:   []cowProtocolStorage.CowTrade{},
				CowTransfer: []cowProtocolStorage.CowTransfer{},
				CallFrame:   []cowProtocolStorage.CowTradeCallFrame{},
			},
		})
		return
	}

	// get cow transfer
	cowTransferQuery := cowProtocolStorage.CowTransferQuery{
		TxHash: txHash,
	}
	cowTransfer, err := s.cowTradeStorage.GetCowTransfers(cowTransferQuery)
	if err != nil {
		responseErr(c, http.StatusInternalServerError, fmt.Errorf("error when get cow transfers: %w", err))
		return
	}
	// get cow callframe
	callFrame, err := s.cowTradeStorage.GetCowCallFrame(txHash)
	if err != nil {
		responseErr(c, http.StatusInternalServerError, fmt.Errorf("error when get callFrame: %w", err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": result{
			CowTrades:   cowTrades,
			CowTransfer: cowTransfer,
			CallFrame:   callFrame,
		},
	})
}
