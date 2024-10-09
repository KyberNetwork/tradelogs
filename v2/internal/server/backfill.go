package server

import (
	"fmt"
	"net/http"

	"github.com/KyberNetwork/tradelogs/v2/internal/worker"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"go.uber.org/zap"
)

// BackfillServer to serve the service.
type BackfillServer struct {
	l        *zap.SugaredLogger
	r        *gin.Engine
	bindAddr string
	worker   *worker.BackFiller
}

type Query struct {
	FromBlock uint64 `json:"from_block" binding:"required"`
	ToBlock   uint64 `json:"to_block" binding:"required"`
	Exchange  string `json:"exchange" binding:"required"`
}

func NewBackfill(l *zap.SugaredLogger, bindAddr string, w *worker.BackFiller) *BackfillServer {
	engine := gin.New()
	engine.Use(gin.Recovery())

	server := &BackfillServer{
		l:        l,
		r:        engine,
		bindAddr: bindAddr,
		worker:   w,
	}

	gin.SetMode(gin.ReleaseMode)
	server.register()

	return server
}

func (s *BackfillServer) Run() error {
	if err := s.r.Run(s.bindAddr); err != nil {
		return fmt.Errorf("run server: %w", err)
	}

	return nil
}

func (s *BackfillServer) register() {
	pprof.Register(s.r, "/debug")
	s.r.POST("/backfill", s.backfill)
}

func responseErr(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"success": false,
		"error":   err.Error(),
	})
}

func responseOK(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (s *BackfillServer) backfill(c *gin.Context) {
	var params Query
	if err := c.BindJSON(&params); err != nil {
		responseErr(c, err)
		return
	}

	if params.FromBlock > params.ToBlock {
		responseErr(c, fmt.Errorf("from block is greater than to block"))
		return
	}

	l := s.l.With("reqID", xid.New().String())
	l.Infow("receive backfill params", "params", params)

	err := s.worker.BackfillByExchange(params.FromBlock, params.ToBlock, params.Exchange)
	if err != nil {
		l.Errorw("error when backfill", "error", err)
		responseErr(c, err)
		return
	}

	responseOK(c)
}
