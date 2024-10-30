package server

import (
	"fmt"
	"net/http"
	"strings"

	promoteeTypes "github.com/KyberNetwork/tradelogs/v2/pkg/storage/promotees"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

type Server struct {
	r        *gin.Engine
	bindAddr string
	s        *promoteeTypes.Storage
}

func New(s *promoteeTypes.Storage, bindAddr string) *Server {
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
	s.r.POST("/insert_name", s.insertName)
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
		query promoteeTypes.PromoteesQuery
		err   = c.ShouldBind(&query)
	)
	if err != nil {
		responseErr(c, http.StatusBadRequest, err)
		return
	}
	if query.Promotee != "" {
		query.Promotee = strings.ToLower(query.Promotee)
	}
	if query.Promoter != "" {
		query.Promoter = strings.ToLower(query.Promoter)
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

func (s *Server) insertName(c *gin.Context) {
	var queries []promoteeTypes.PromoteesQuery

	if err := c.ShouldBindJSON(&queries); err != nil {
		responseErr(c, http.StatusBadRequest, err)
		return
	}

	promotees := make([]promoteeTypes.Promotee, len(queries))
	for i, query := range queries {
		if query.Name == "" {
			responseErr(c, http.StatusBadRequest, fmt.Errorf("missing field 'Name' in query index %d", i))
			return
		}
		if query.Promoter == "" {
			responseErr(c, http.StatusBadRequest, fmt.Errorf("missing field 'Promoter' in query index %d", i))
			return
		}
		promotees[i] = promoteeTypes.Promotee{
			Promoter: strings.ToLower(query.Promoter),
			Name:     query.Name,
		}
	}

	if err := s.s.InsertPromoterName(promotees); err != nil {
		responseErr(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    promotees,
	})
}
