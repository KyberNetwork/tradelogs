package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/KyberNetwork/tradelogs/v2/internal/service"
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
	service  *service.Backfill
}

type Query struct {
	FromBlock uint64 `json:"from_block" binding:"required"`
	ToBlock   uint64 `json:"to_block" binding:"required"`
	Exchange  string `json:"exchange" binding:"required"`
}

func NewBackfill(l *zap.SugaredLogger, bindAddr string, s *service.Backfill) *BackfillServer {
	engine := gin.New()
	engine.Use(gin.Recovery())

	server := &BackfillServer{
		l:        l,
		r:        engine,
		bindAddr: bindAddr,
		service:  s,
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
	s.r.GET("/backfill", s.getAllTask)
	s.r.GET("/backfill/:id", s.getTask)
	s.r.GET("/backfill/cancel/:id", s.cancelTask)
	s.r.GET("/backfill/restart/:id", s.restartTask)
}

func responseErr(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"success": false,
		"error":   err.Error(),
	})
}

func internalServerError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"success": false,
		"error":   err.Error(),
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

	id, message, err := s.service.NewBackfillTask(params.FromBlock, params.ToBlock, params.Exchange)
	if err != nil {
		l.Errorw("error when backfill", "error", err)
		internalServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"id":      id,
		"message": message,
	})
}

func (s *BackfillServer) getAllTask(c *gin.Context) {
	tasks, err := s.service.ListTask()
	if err != nil {
		internalServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"tasks":   tasks,
	})
}

func (s *BackfillServer) getTask(c *gin.Context) {
	id := c.Param("id")
	taskID, err := strconv.ParseInt(id, 10, 32)
	if err != nil || len(id) == 0 {
		responseErr(c, fmt.Errorf("invalid task id: %s", id))
		return
	}
	task, err := s.service.GetTask(int(taskID))
	if err != nil {
		internalServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"task":    task,
	})
}

func (s *BackfillServer) cancelTask(c *gin.Context) {
	id := c.Param("id")
	taskID, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		responseErr(c, fmt.Errorf("invalid task id: %w", err))
		return
	}
	err = s.service.CancelBackfillTask(int(taskID))
	if err != nil {
		internalServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (s *BackfillServer) restartTask(c *gin.Context) {
	id := c.Param("id")
	taskID, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		responseErr(c, fmt.Errorf("invalid task id: %w", err))
		return
	}
	err = s.service.RestartBackfillTask(int(taskID))
	if err != nil {
		internalServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
