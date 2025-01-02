package runestone

import (
	"github.com/sat20-labs/indexer/indexer/runes/pb"
	"github.com/sat20-labs/indexer/indexer/runes/store"
)

type RuneIdToEntry map[*RuneId]*RuneEntry

type RuneIdToEntryTable struct {
	Table[pb.RuneEntry]
}

func NewRuneIdToEntryTable(store *store.Cache[pb.RuneEntry]) *RuneIdToEntryTable {
	return &RuneIdToEntryTable{Table: Table[pb.RuneEntry]{cache: store}}
}

func (s *RuneIdToEntryTable) Get(key *RuneId) (ret *RuneEntry) {
	tblKey := []byte(store.ID_TO_ENTRY + key.String())
	pbVal := s.cache.Get(tblKey)
	if pbVal != nil {
		ret = &RuneEntry{}
		ret.FromPb(pbVal)
	}
	return
}

func (s *RuneIdToEntryTable) GetFromDB(key *RuneId) (ret *RuneEntry) {
	tblKey := []byte(store.ID_TO_ENTRY + key.String())
	pbVal, _ := s.cache.GetFromDB(tblKey)
	if pbVal != nil {
		ret = &RuneEntry{}
		ret.FromPb(pbVal)
	}
	return
}

func (s *RuneIdToEntryTable) GetListFromDB() (ret map[string]*RuneEntry) {
	prefixKey := []byte(store.ID_TO_ENTRY)
	list := s.cache.GetListFromDB(prefixKey, true)
	if len(list) == 0 {
		return
	}
	ret = make(map[string]*RuneEntry)
	for k, v := range list {
		ret[k] = &RuneEntry{}
		ret[k].FromPb(v)
	}
	return ret
}

func (s *RuneIdToEntryTable) Insert(key *RuneId, value *RuneEntry) (ret *RuneEntry) {
	tblKey := []byte(store.ID_TO_ENTRY + key.String())
	pbVal := s.cache.Set(tblKey, value.ToPb())
	if pbVal != nil {
		ret = &RuneEntry{}
		ret.FromPb(pbVal)
	}
	return
}

func (s *RuneIdToEntryTable) SetToDB(key *RuneId, value *RuneEntry) {
	tblKey := []byte(store.ID_TO_ENTRY + key.String())
	s.cache.SetToDB(tblKey, value.ToPb())
}
