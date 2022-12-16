package handler

import (
	"context"

	"github.com/KyberNetwork/tradelogs/internal/evmlistenerclient"
	"github.com/KyberNetwork/tradelogs/internal/worker"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

type Handler struct {
	workers  []worker.Worker
	listener *evmlistenerclient.Client
	l        *zap.SugaredLogger
}

func New(l *zap.SugaredLogger, listener *evmlistenerclient.Client) *Handler {
	return &Handler{
		workers:  make([]worker.Worker, 0),
		listener: listener,
		l:        l,
	}
}

func (h *Handler) AddWoker(w worker.Worker) {
	h.workers = append(h.workers, w)
}

func (h *Handler) Run(ctx context.Context) error {
	for {
		m, err := h.listener.GConsume(ctx)
		if err != nil {
			h.l.Errorw("Error while consume in group")
			return err
		}
		h.l.Infow("Begin process msg")
		var g errgroup.Group
		for i := range h.workers {
			w := h.workers[i]
			g.Go(func() error { return w.Run(m) })
		}
		if err := g.Wait(); err != nil {
			return err
		}
		if err := h.listener.Ack(ctx, m); err != nil {
			return err
		}
	}
}
