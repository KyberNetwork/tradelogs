package server

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/KyberNetwork/tradelogs/internal/backfill"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Server to serve the service.
type Server struct {
	l              *zap.SugaredLogger
	r              *gin.Engine
	backFillWorker *backfill.Worker
	bindAddr       string
}

// New returns a new server.
func New(bindAddr string, bfWorker *backfill.Worker) *Server {
	engine := gin.New()
	engine.Use(gin.Recovery())

	server := &Server{
		l:              zap.S(),
		r:              engine,
		backFillWorker: bfWorker,
		bindAddr:       bindAddr,
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
	s.r.POST("/backfill-one-inch", s.backFillOneInch)

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
		err = s.backFillWorker.BackFillAllData(exchanges)
	} else {
		err = s.backFillWorker.BackFillPartialData(query.FromTime/1000, query.ToTime/1000, exchanges)
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

func (s *Server) backFillOneInch(c *gin.Context) {
	var (
		query struct {
			FromTime        int64  `form:"from_time" json:"from_time,omitempty"` // milliseconds
			ToTime          int64  `form:"to_time" json:"to_time,omitempty"`     // milliseconds
			ContractAddress string `form:"contract_address" json:"contract_address" `
			FourBytes       string `form:"four_bytes" json:"four_bytes"`
		}
		err = c.ShouldBind(&query)
	)
	if err != nil {
		responseErr(c, err)
		return
	}
	var fourByteMethodIds []string
	if query.FourBytes != "" {
		fourByteMethodIds = strings.Split(query.FourBytes, ",")
	}

	s.l.Infow("Request to backfill 1inch RFQ orders", "query", query)
	err = s.backFillWorker.BackFillOneInch(query.FromTime/1000, query.ToTime/1000, query.ContractAddress, fourByteMethodIds)
	if err != nil {
		responseErr(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"time":    time.Now().UnixMilli(),
	})
}
