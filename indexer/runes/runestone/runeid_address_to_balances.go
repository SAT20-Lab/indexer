package runestone

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sat20-labs/indexer/common"
	"github.com/sat20-labs/indexer/indexer/runes/pb"
	"github.com/sat20-labs/indexer/indexer/runes/store"
	"lukechampine.com/uint128"
)

type RuneIdAddressToBalance struct {
	RuneId    *RuneId
	AddressId uint64
	Address   Address
	OutPoint  *OutPoint
	Balance   *Lot
}

func RuneIdAddressToBalanceFromString(str string) (*RuneIdAddressToBalance, error) {
	ret := &RuneIdAddressToBalance{}
	parts := strings.SplitN(str, "-", 3)

	var err error
	ret.RuneId, err = RuneIdFromHex(parts[1])
	if err != nil {
		return nil, err
	}
	ret.AddressId, err = strconv.ParseUint(parts[2], 16, 64)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (s *RuneIdAddressToBalance) Key() string {
	return s.RuneId.Hex() + "-" + fmt.Sprintf("%x", s.AddressId)
}

func (s *RuneIdAddressToBalance) ToPb() *pb.RuneIdAddressToBalance {
	pbValue := &pb.RuneIdAddressToBalance{
		Balance: &pb.Lot{
			Value: &pb.Uint128{
				Hi: s.Balance.Value.Hi,
				Lo: s.Balance.Value.Lo,
			},
		},
		Address: string(s.Address),
		Utxo:    s.OutPoint.Hex(),
	}
	return pbValue
}

type RuneIdAddressToBalanceTable struct {
	Table[pb.RuneIdAddressToBalance]
}

func NewRuneIdAddressToBalanceTable(v *store.Cache[pb.RuneIdAddressToBalance]) *RuneIdAddressToBalanceTable {
	return &RuneIdAddressToBalanceTable{Table: Table[pb.RuneIdAddressToBalance]{cache: v}}
}

func (s *RuneIdAddressToBalanceTable) Get(v *RuneIdAddressToBalance) (ret *RuneIdAddressToBalance) {
	tblKey := []byte(store.RUNEID_ADDRESS_BALANCE + v.Key())
	pbVal := s.cache.Get(tblKey)
	if pbVal != nil {
		var err error
		ret, err = RuneIdAddressToBalanceFromString(string(tblKey))
		if err != nil {
			common.Log.Panicf("RuneIdAddressToBalanceTable.Get-> RuneIdAddressToBalanceFromString(%s) err:%v", string(tblKey), err)
			return nil
		}
		ret.Address = Address(pbVal.Address)
		ret.Balance = &Lot{
			Value: &uint128.Uint128{
				Hi: pbVal.Balance.Value.Hi,
				Lo: pbVal.Balance.Value.Lo,
			},
		}
		ret.OutPoint, err = OutPointFromHex(pbVal.Utxo)
		if err != nil {
			common.Log.Panicf("RuneIdAddressToBalanceTable.Get-> OutPointFromHex(%s) err:%v", pbVal.Utxo, err)
			return nil
		}
	}
	return
}

func (s *RuneIdAddressToBalanceTable) GetBalances(runeId *RuneId) (ret []*RuneIdAddressToBalance, err error) {
	tblKey := []byte(store.RUNEID_ADDRESS_BALANCE + runeId.Hex() + "-")
	pbVal := s.cache.GetList(tblKey, true)
	if pbVal != nil {
		ret = make([]*RuneIdAddressToBalance, len(pbVal))
		var i = 0
		for k, v := range pbVal {
			var err error
			runeIdAddressOutpointToBalance, err := RuneIdAddressToBalanceFromString(string(k))
			if err != nil {
				return nil, err
			}
			ret[i].Address = Address(v.Address)
			ret[i].Balance = &Lot{
				Value: &uint128.Uint128{Hi: v.Balance.Value.Hi, Lo: v.Balance.Value.Lo}}
			ret[i].OutPoint, err = OutPointFromHex(v.Utxo)
			if err != nil {
				return nil, err
			}
			ret[i] = runeIdAddressOutpointToBalance
			i++
		}
	}
	return
}

func (s *RuneIdAddressToBalanceTable) Insert(v *RuneIdAddressToBalance) {
	tblKey := []byte(store.RUNEID_ADDRESS_BALANCE + v.Key())
	s.cache.Set(tblKey, v.ToPb())
}

func (s *RuneIdAddressToBalanceTable) Remove(v *RuneIdAddressToBalance) {
	tblKey := []byte(store.RUNEID_ADDRESS_BALANCE + v.Key())
	s.cache.Delete(tblKey)
}
