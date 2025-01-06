package runestone

import (
	"strings"

	"github.com/sat20-labs/indexer/common"
	"github.com/sat20-labs/indexer/indexer/runes/pb"
	"github.com/sat20-labs/indexer/indexer/runes/store"
	"lukechampine.com/uint128"
)

type RuneIdOutpointToBalance struct {
	RuneId   *RuneId
	OutPoint *OutPoint
	Balance  *Lot
}

func GenRuneIdOutpointToBalance(str string, balance *Lot) (*RuneIdOutpointToBalance, error) {
	ret := &RuneIdOutpointToBalance{}
	parts := strings.SplitN(str, "-", 4)
	var err error
	ret.RuneId, err = RuneIdFromString(parts[1])
	if err != nil {
		return nil, err
	}
	ret.OutPoint, err = OutPointFromString(parts[2])
	if err != nil {
		return nil, err
	}
	ret.Balance = balance
	return ret, nil
}

func (s *RuneIdOutpointToBalance) String() string {
	return s.RuneId.String() + "-" + s.OutPoint.String()
}

func (s *RuneIdOutpointToBalance) ToPb() *pb.RuneIdToOutpointToBalance {
	pbValue := &pb.RuneIdToOutpointToBalance{
		Balance: &pb.Lot{
			Value: &pb.Uint128{
				Hi: s.Balance.Value.Hi,
				Lo: s.Balance.Value.Lo,
			},
		},
	}
	return pbValue
}

type RuneIdOutpointToBalanceTable struct {
	Table[pb.RuneIdToOutpointToBalance]
}

func NewRuneIdOutpointToBalancesTable(v *store.Cache[pb.RuneIdToOutpointToBalance]) *RuneIdOutpointToBalanceTable {
	return &RuneIdOutpointToBalanceTable{Table: Table[pb.RuneIdToOutpointToBalance]{cache: v}}
}

func (s *RuneIdOutpointToBalanceTable) Get(v *RuneIdOutpointToBalance) (ret *RuneIdOutpointToBalance) {
	tblKey := []byte(store.RUNEID_OUTPOINT_TO_BALANCE + v.String())
	pbVal := s.cache.Get(tblKey)
	if pbVal != nil {
		var err error
		ret, err = GenRuneIdOutpointToBalance(string(tblKey), v.Balance)
		if err != nil {
			common.Log.Panicf("RuneIdOutpointToBalanceTable.Get-> GenRuneIdOutpointToBalance(%s) err:%v", string(tblKey), err)
		}
	}
	return
}

func (s *RuneIdOutpointToBalanceTable) GetBalances(runeId *RuneId) (ret []*RuneIdOutpointToBalance, err error) {
	tblKey := []byte(store.RUNEID_OUTPOINT_TO_BALANCE + runeId.String() + "-")
	pbVal := s.cache.GetList(tblKey, true)
	if pbVal != nil {
		ret = make([]*RuneIdOutpointToBalance, len(pbVal))
		var i = 0
		for k, v := range pbVal {
			var err error
			balance := &Lot{
				Value: &uint128.Uint128{
					Hi: v.Balance.Value.Hi,
					Lo: v.Balance.Value.Lo,
				},
			}
			runeIdOutpointToBalance, err := GenRuneIdOutpointToBalance(string(k), balance)
			if err != nil {
				common.Log.Panicf("RuneIdOutpointToBalanceTable.Get-> GenRuneIdOutpointToBalance(%s) err:%v", string(k), err)
			}
			ret[i] = runeIdOutpointToBalance
			i++
		}
	}
	return
}

func (s *RuneIdOutpointToBalanceTable) Insert(v *RuneIdOutpointToBalance) (ret *RuneIdOutpointToBalance) {
	tblKey := []byte(store.RUNEID_OUTPOINT_TO_BALANCE + v.String())
	pbVal := s.cache.Set(tblKey, v.ToPb())
	if pbVal != nil {
		balance := &Lot{
			Value: &uint128.Uint128{
				Hi: pbVal.Balance.Value.Hi,
				Lo: pbVal.Balance.Value.Lo,
			},
		}
		var err error
		ret, err = GenRuneIdOutpointToBalance(string(tblKey), balance)
		if err != nil {
			common.Log.Panicf("RuneIdOutpointToBalanceTable.Insert-> GenRuneIdOutpointToBalance(%s) err:%v", string(tblKey), err)
		}
	}
	return
}

func (s *RuneIdOutpointToBalanceTable) Remove(key *RuneIdOutpointToBalance) (ret *RuneIdOutpointToBalance) {
	tblKey := []byte(store.RUNEID_OUTPOINT_TO_BALANCE + key.String())
	pbVal := s.cache.Delete(tblKey)
	if pbVal != nil {
		balance := &Lot{
			Value: &uint128.Uint128{
				Hi: pbVal.Balance.Value.Hi,
				Lo: pbVal.Balance.Value.Lo,
			},
		}
		var err error
		ret, err = GenRuneIdOutpointToBalance(string(tblKey), balance)
		if err != nil {
			common.Log.Panicf("RuneIdOutpointToBalanceTable.Insert-> GenRuneIdOutpointToBalance(%s) err:%v", string(tblKey), err)
		}
	}
	return
}