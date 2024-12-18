package oneinchv6

import (
	promoteesStorage "github.com/KyberNetwork/tradelogs/v2/pkg/storage/promotees"
	"go.uber.org/zap"
)

type MarkFusion struct {
	promoteeStorage *promoteesStorage.Storage
	promoteeMap     map[string]bool
	l               *zap.SugaredLogger
}

func NewMarkFusion(promoteeStorage *promoteesStorage.Storage, l *zap.SugaredLogger) (*MarkFusion, error) {
	m := &MarkFusion{
		promoteeStorage: promoteeStorage,
		promoteeMap:     make(map[string]bool),
		l:               l,
	}
	err := m.CreatePromoteeMap()
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (m *MarkFusion) CreatePromoteeMap() error {
	query := promoteesStorage.PromoteesQuery{}
	promoteeList, err := m.promoteeStorage.Get(query)
	if err != nil {
		m.l.Errorw("Error get promoteeList", "error", err)
		return err
	}
	for _, promotee := range promoteeList {
		if promotee.ChainId == "1" {
			m.promoteeMap[promotee.Promotee] = true
		}
	}
	return nil
}

func (m *MarkFusion) CheckPromoteeExist(promotee string) bool {
	exists := m.promoteeMap[promotee]
	return exists
}
