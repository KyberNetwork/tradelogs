package server

import (
	"fmt"
	"net/http"
	"strings"

	promoteeTypes "github.com/KyberNetwork/tradelogs/v2/pkg/storage/promotees"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

type PromoteeServer struct {
	r        *gin.Engine
	bindAddr string
	s        *promoteeTypes.Storage
}

func NewPromoteeServer(s *promoteeTypes.Storage, bindAddr string) *PromoteeServer {
	engine := gin.New()
	engine.Use(gin.Recovery())

	server := &PromoteeServer{
		r:        engine,
		bindAddr: bindAddr,
		s:        s,
	}

	gin.SetMode(gin.DebugMode)
	server.register()

	return server
}

func (s *PromoteeServer) Run() error {
	if err := s.r.Run(s.bindAddr); err != nil {
		return fmt.Errorf("run server: %w", err)
	}

	return nil
}

func (s *PromoteeServer) register() {
	pprof.Register(s.r, "/debug")
	s.r.GET("/promotees", s.getPromotees)
	s.r.POST("/insert_name", s.insertName)
}

func (s *PromoteeServer) getPromotees(c *gin.Context) {
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

func (s *PromoteeServer) insertName(c *gin.Context) {
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
