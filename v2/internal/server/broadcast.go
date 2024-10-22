package server

import (
	"fmt"
	"net/http"

	"github.com/KyberNetwork/tradelogs/v2/internal/worker"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

var (
	wsupgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		// Allow connections from any Origin
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

// Server to serve the service.
type Server struct {
	r        *gin.Engine
	bindAddr string
	l        *zap.SugaredLogger
	bc       *worker.Broadcaster
}

// New returns a new server.
func New(l *zap.SugaredLogger, bindAddr string, bc *worker.Broadcaster) *Server {
	engine := gin.New()
	engine.Use(gin.Recovery())

	server := &Server{
		r:        engine,
		bindAddr: bindAddr,
		l:        l,
		bc:       bc,
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
	s.r.GET("/eventlogws", s.registerEventLogWS)
}

func responseErr(c *gin.Context, status int, err error) {
	c.JSON(status, gin.H{
		"success": false,
		"error":   err.Error(),
	})
}

func (s *Server) registerEventLogWS(c *gin.Context) {
	var param worker.RegisterRequest
	if err := c.BindQuery(&param); err != nil {
		responseErr(c, http.StatusBadRequest, err)
		return
	}

	s.l.Infow("receive ws", "param", param)
	conn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		s.l.Errorw("Failed to set websocket upgrade", "error", err)
		responseErr(c, http.StatusInternalServerError, fmt.Errorf("can't create ws"))
		return
	}
	s.bc.NewConn(param, conn)
}
