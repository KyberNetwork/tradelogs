package server

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/KyberNetwork/tradelogs/internal/bigquery"
	"github.com/KyberNetwork/tradelogs/pkg/parser"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"go.uber.org/zap"
)

// Server to serve the service.
type Server struct {
	l        *zap.SugaredLogger
	r        *gin.Engine
	bq       *bigquery.Worker
	dune     *DuneWorker
	bindAddr string
	parsers  []parser.Parser
}

// New returns a new server.
func New(bindAddr string, bigQueryWorker *bigquery.Worker,
	dune *DuneWorker, parsers []parser.Parser) *Server {
	engine := gin.New()
	engine.Use(gin.Recovery())

	server := &Server{
		l:        zap.S(),
		r:        engine,
		bq:       bigQueryWorker,
		dune:     dune,
		bindAddr: bindAddr,
		parsers:  parsers,
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

func responseOK(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
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

// func (s *Server) backfill(c *gin.Context) {
// 	var params storage.BackfillQuery
// 	if err := c.BindJSON(&params); err != nil {
// 		responseErr(c, err)
// 		return
// 	}

// 	if params.FromBlock > params.ToBlock {
// 		responseErr(c, fmt.Errorf("from block is greater than to block"))
// 		return
// 	}

// 	if params.EventHash == "" && params.Exchange == "" {
// 		responseErr(c, fmt.Errorf("empty event hash or exchange"))
// 		return
// 	}

// 	l := s.l.With("reqID", xid.New().String())
// 	l.Infow("receive backfill params", "params", params)
// 	if params.EventHash != "" {
// 		for _, p := range s.parsers {
// 			for _, t := range p.Topics() {
// 				if strings.EqualFold(params.EventHash, t) {
// 					if err := s.dune.backfill(l, params.QueryID, params.FromBlock, params.ToBlock, p, t); err != nil {
// 						l.Errorw("error when backfill", "error", err)
// 						responseErr(c, err)
// 						return
// 					}
// 					responseOK(c)
// 					return
// 				}
// 			}
// 		}
// 	}

// 	for _, p := range s.parsers {
// 		if strings.EqualFold(params.Exchange, p.Exchange()) {
// 			for _, t := range p.Topics() {
// 				if err := s.dune.backfill(l, params.QueryID, params.FromBlock, params.ToBlock, p, t); err != nil {
// 					l.Errorw("error when backfill", "error", err)
// 					responseErr(c, err)
// 					return
// 				}
// 			}
// 			responseOK(c)
// 			return
// 		}
// 	}

// }

func (s *Server) backfillOneinch(c *gin.Context) {
	var params BackFillOneInchRequest
	if err := c.BindJSON(&params); err != nil {
		responseErr(c, err)
		return
	}

	if err := s.dune.backfillOneInch(s.l.With("reqID", xid.New().String()), params.QueryID, params.EventHash); err != nil {
		responseErr(c, err)
		return
	}
	responseOK(c)
}
