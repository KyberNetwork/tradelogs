package server

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/KyberNetwork/tradelogs/internal/storage"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	maxTimeRange uint64 = uint64(24 * time.Hour.Milliseconds())
)

// Server to serve the service.
type Server struct {
	r        *gin.Engine
	bindAddr string
	l        *zap.SugaredLogger
	s        *storage.Storage
}

// New returns a new server.
func New(l *zap.SugaredLogger, s *storage.Storage, bindAddr string) *Server {
	engine := gin.New()
	engine.Use(gin.Recovery())

	server := &Server{
		r:        engine,
		bindAddr: bindAddr,
		l:        l,
		s:        s,
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
	s.r.GET("/tradelogs", s.getTradeLogs)

}
func responseErr(c *gin.Context, status int, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"success": false,
		"error":   err.Error(),
	})
}
func (s *Server) getTradeLogs(c *gin.Context) {
	var (
		query storage.TradeLogsQuery
		err   = c.ShouldBind(&query)
	)
	if err != nil {
		responseErr(c, http.StatusBadRequest, err)
		return
	}
	if query.ToTime-query.FromTime > maxTimeRange {
		responseErr(c, http.StatusBadRequest, errors.New("max time range"))
		return
	}
	data, err := s.s.Get(query)
	if err != nil {
		responseErr(c, http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}
