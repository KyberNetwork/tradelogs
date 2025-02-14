package worker

import (
	"time"

	dashboardStorage "github.com/KyberNetwork/tradelogs/v2/pkg/storage/dashboard"
	"go.uber.org/zap"
)

type Refresher struct {
	DashStorage *dashboardStorage.Storage
	l           *zap.SugaredLogger
}

func NewRefresher(dashStorage *dashboardStorage.Storage, l *zap.SugaredLogger) *Refresher {
	return &Refresher{
		DashStorage: dashStorage,
		l:           l,
	}
}

func (r *Refresher) ScheduleMViewRefresh() {
	ticker := time.NewTicker(24 * time.Hour)
	defer ticker.Stop()
	for range ticker.C {
		err := r.DashStorage.RefreshTradelogsMView()
		if err != nil {
			r.l.Errorf("Error refreshing materialized view: %v", err)
		} else {
			r.l.Info("Materialized view refreshed successfully.")
		}
	}
}
