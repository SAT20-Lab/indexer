package runestone

import (
	"github.com/sat20-labs/indexer/common"
	"github.com/sat20-labs/indexer/indexer/runes/pb"
	"github.com/sat20-labs/indexer/indexer/runes/store"
)

type RunesStatus struct {
	Table[pb.RunesStatus]
	Version       string
	Height        uint64
	Number        uint64
	ReservedRunes uint64
}

func (s *RunesStatus) Init() (ret bool) {
	stats := s.store.Get([]byte(store.STATUS_KEY))
	common.Log.Infof("Runes.RunesStatus->Init: stats: %v", stats)
	if stats == nil {
		s.Version = store.DB_VERSION
		ret = false
	} else {
		if stats.Version != store.DB_VERSION {
			common.Log.Panicf("Runes.RunesStatus->Init: db version inconsistent %s", store.DB_VERSION)
		}
		s.Version = stats.Version
		s.Height = stats.Height
		s.Number = stats.Number
		s.ReservedRunes = stats.ReservedRunes
		ret = true
	}
	return
}

func (s *RunesStatus) Update() {
	key := []byte(store.STATUS_KEY)
	value := &pb.RunesStatus{
		Version:       s.Version,
		Height:        s.Height,
		Number:        s.Number,
		ReservedRunes: s.ReservedRunes,
	}
	s.store.Insert(key, value)
}
