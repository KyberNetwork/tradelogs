package server

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/KyberNetwork/tradelogs/pkg/storage"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

var (
	maxTimeRange uint64 = uint64(7 * 24 * time.Hour.Milliseconds())
	wsupgrader          = websocket.Upgrader{
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
	s.r.GET("/1inchws", s.register1InchWS)
}

func responseErr(c *gin.Context, status int, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"success": false,
		"error":   err.Error(),
		"status":  status,
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

func (s *Server) register1InchWS(c *gin.Context) {
	clientID := c.Query("client_id")
	if clientID == "" {
		responseErr(c, http.StatusBadRequest, fmt.Errorf("need client_id param"))
		return
	}
	conn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		s.l.Errorw("Failed to set websocket upgrade", "error", err)
		return
	}
	s.l.Infow("receive ws", "client id", clientID)
	for {
		t, msg, err := conn.ReadMessage()
		s.l.Infow("receive", "msg", msg)
		if err != nil {
			break
		}
		conn.WriteMessage(t, msg[:1])
	}
}
