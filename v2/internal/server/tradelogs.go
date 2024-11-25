package server

import (
	"fmt"
	"net/http"
	"time"

	storageTypes "github.com/KyberNetwork/tradelogs/v2/pkg/storage/tradelogs/types"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	maxTimeRange = uint64(24 * time.Hour.Milliseconds())
)

type TradeLogs struct {
	r        *gin.Engine
	bindAddr string
	l        *zap.SugaredLogger
	storage  []storageTypes.Storage
}

func NewTradeLogs(l *zap.SugaredLogger, s []storageTypes.Storage, bindAddr string) *TradeLogs {
	engine := gin.New()
	engine.Use(gin.Recovery())

	server := &TradeLogs{
		r:        engine,
		bindAddr: bindAddr,
		l:        l,
		storage:  s,
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
			responseErr(c, http.StatusBadRequest, err)
			return
		}
		data = append(data, tradeLogs...)
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}
