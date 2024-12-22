package runestone

import (
	"lukechampine.com/uint128"
)

type Lot struct {
	Value *uint128.Uint128
}

func NewLot(value *uint128.Uint128) *Lot {
	return &Lot{Value: value}
}

func (l Lot) N() *uint128.Uint128 {
	return l.Value
}

func (l Lot) Add(rhs *Lot) Lot {
	result := l.Value.Add(*rhs.Value)
	return Lot{Value: &result}
}

func (l *Lot) AddAssign(rhs *Lot) {
	*l = l.Add(rhs)
}

func (l Lot) AddUint128(rhs *uint128.Uint128) Lot {
	return l.Add(&Lot{Value: rhs})
}

func (l *Lot) AddAssignUint128(rhs *uint128.Uint128) {
	l.AddAssign(&Lot{Value: rhs})
}

func (l Lot) Sub(rhs Lot) Lot {
	result := l.Value.Sub(*rhs.Value)
	return Lot{Value: &result}
}

func (l *Lot) SubAssign(rhs Lot) {
	*l = l.Sub(rhs)
}

func (l Lot) Div(rhs *uint128.Uint128) Lot {
	value := l.Value.Div(*rhs)
	return Lot{Value: &value}
}

func (l Lot) Rem(rhs *uint128.Uint128) Lot {
	value := l.Value.Mod(*rhs)
	return Lot{Value: &value}
}

func (l Lot) Eq(rhs *uint128.Uint128) bool {
	return l.Value.Equals(*rhs)
}

func (l Lot) Cmp(rhs *uint128.Uint128) int {
	return l.Value.Cmp(*rhs)
}