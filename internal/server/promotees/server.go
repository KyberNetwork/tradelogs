package server

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/KyberNetwork/tradelogs/pkg/promotionstorage"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

type Server struct {
	r        *gin.Engine
	bindAddr string
	s        *promotionstorage.Storage
}

func New(s *promotionstorage.Storage, bindAddr string) *Server {
	engine := gin.New()
	engine.Use(gin.Recovery())

	server := &Server{
		r:        engine,
		bindAddr: bindAddr,
		s:        s,
	}

	gin.SetMode(gin.DebugMode)
	server.register()

	return server
}

func (s *Server) Run() error {
	if err := s.r.Run(s.bindAddr); err != nil {
		return fmt.Errorf("run server: %w", err)
	}

	return nil
}

func (s *Server) register() {
	pprof.Register(s.r, "/debug")
	s.r.GET("/promotees", s.getPromotees)
}

func responseErr(c *gin.Context, status int, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"success": false,
		"error":   err.Error(),
		"status":  status,
	})
}

func (s *Server) getPromotees(c *gin.Context) {
	var (
		query promotionstorage.PromoteesQuery
		err   = c.ShouldBind(&query)
	)
	if query.Promotee != "" {
		query.Promotee = strings.ToLower(query.Promotee)
	}
	if query.Promoter != "" {
		query.Promoter = strings.ToLower(query.Promoter)
	}
	if err != nil {
		responseErr(c, http.StatusBadRequest, err)
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
