package server

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/KyberNetwork/tradelogs/internal/bigquery"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Server to serve the service.
type Server struct {
	l        *zap.SugaredLogger
	r        *gin.Engine
	bq       *bigquery.Worker
	bindAddr string
}

// New returns a new server.
func New(bindAddr string, bq *bigquery.Worker) *Server {
	engine := gin.New()
	engine.Use(gin.Recovery())

	server := &Server{
		l:        zap.S(),
		r:        engine,
		bq:       bq,
		bindAddr: bindAddr,
	}

	gin.SetMode(gin.ReleaseMode)
	server.register()

	return server
}

// Run runs server.
func (s *Server) Run() error {
	if err := s.r.Run(s.bindAddr); err != nil {
		return fmt.Errorf("run server: %w", err)
	}

	return nil
}

func (s *Server) register() {
	pprof.Register(s.r, "/debug")
	s.r.POST("/backfill", s.backfill)
	s.r.POST("/backfill-1inch", s.backfillOneinch)
}

func responseErr(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"success": false,
		"error":   err.Error(),
	})
}

func (s *Server) backfill(c *gin.Context) {
	var (
		query struct {
			FromTime  int64  `form:"from_time" json:"from_time,omitempty"` // milliseconds
			ToTime    int64  `form:"to_time" json:"to_time,omitempty"`     // milliseconds
			Exchanges string `form:"exchanges" json:"exchanges,omitempty"`
		}
		err = c.ShouldBind(&query)
	)
	if err != nil {
		responseErr(c, err)
		return
	}

	var exchanges []string
	if query.Exchanges != "" {
		exchanges = strings.Split(query.Exchanges, ",")
	}
	s.l.Infow("Request backfill", "query", query)
	if query.FromTime == 0 || query.ToTime == 0 {
		err = s.bq.BackFillAllData(exchanges)
	} else {
		err = s.bq.BackFillPartialData(query.FromTime/1000, query.ToTime/1000, exchanges)
	}
	if err != nil {
		responseErr(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"time":    time.Now().UnixMilli(),
	})
}

func (s *Server) backfillOneinch(c *gin.Context) {
	var (
		query BackFillOneInchRequest
	)
	if err := c.ShouldBind(&query); err != nil {
		responseErr(c, err)
		return
	}

	tradeLogs := query.ToTradeLogs()

	err := s.bq.BackfillOneInchRFQ(tradeLogs)
	if err != nil {
		responseErr(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Backfill 1inch rfq orders successfully",
		"success": true,
		"time":    time.Now().UnixMilli(),
	})
}
