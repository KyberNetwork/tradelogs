package server

import (
	"fmt"
	"net/http"
	"time"

	dashboardStorage "github.com/KyberNetwork/tradelogs/v2/pkg/storage/dashboard"
	dashboardTypes "github.com/KyberNetwork/tradelogs/v2/pkg/storage/dashboard/types"
	storageTypes "github.com/KyberNetwork/tradelogs/v2/pkg/storage/tradelogs/types"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	maxTimeRange = uint64(24 * time.Hour.Milliseconds())
)

type TradeLogs struct {
	r           *gin.Engine
	bindAddr    string
	l           *zap.SugaredLogger
	storage     []storageTypes.Storage
	dashStorage *dashboardStorage.Storage
}

func NewTradeLogs(
	l *zap.SugaredLogger,
	s []storageTypes.Storage,
	dashStorage *dashboardStorage.Storage,
	bindAddr string,
) *TradeLogs {
	engine := gin.New()
	engine.Use(gin.Recovery())

	server := &TradeLogs{
		r:           engine,
		bindAddr:    bindAddr,
		l:           l,
		storage:     s,
		dashStorage: dashStorage,
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
	var query struct {
		Address  string `form:"address" json:"address"`
		Exchange string `form:"exchange" json:"exchange"`
		From     int64  `form:"from" json:"from"`
		To       int64  `form:"to" json:"to"`
	}

	if err := c.ShouldBindJSON(&query); err != nil {
		responseErr(c, http.StatusBadRequest, err)
		return
	}

	for _, storage := range s.storage {
		if storage.Exchange() != query.Exchange {
			continue
		}
		from := time.Time{}
		to := time.Time{}
		if query.From > 0 {
			from = time.Unix(query.From, 0)
		}
		if query.To > 0 {
			to = time.Unix(query.To, 0)
		}
		rows, err := storage.ResetTokenPriceToRefetch(query.Address, from, to)
		if err != nil {
			responseErr(c, http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data": struct {
				Token    string `json:"token"`
				Exchange string `json:"exchange"`
				From     int64  `json:"from"`
				To       int64  `json:"to"`
				Rows     int64  `json:"number of row updated"`
			}{
				Token:    query.Address,
				Exchange: query.Exchange,
				From:     query.From,
				To:       query.To,
				Rows:     rows,
			},
		})
		return
	}
	responseErr(c, http.StatusBadRequest, fmt.Errorf("exchange not found"))
}
