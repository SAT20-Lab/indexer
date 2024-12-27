package common

import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"strings"
)

const MAX_PRECISION = 18
const DEFAULT_PRECISION = 0

var MAX_PRECISION_STRING = "18"

var precisionFactor [19]*big.Int = [19]*big.Int{
	new(big.Int).Exp(big.NewInt(10), big.NewInt(0), nil),
	new(big.Int).Exp(big.NewInt(10), big.NewInt(1), nil),
	new(big.Int).Exp(big.NewInt(10), big.NewInt(2), nil),
	new(big.Int).Exp(big.NewInt(10), big.NewInt(3), nil),
	new(big.Int).Exp(big.NewInt(10), big.NewInt(4), nil),
	new(big.Int).Exp(big.NewInt(10), big.NewInt(5), nil),
	new(big.Int).Exp(big.NewInt(10), big.NewInt(6), nil),
	new(big.Int).Exp(big.NewInt(10), big.NewInt(7), nil),
	new(big.Int).Exp(big.NewInt(10), big.NewInt(8), nil),
	new(big.Int).Exp(big.NewInt(10), big.NewInt(9), nil),
	new(big.Int).Exp(big.NewInt(10), big.NewInt(10), nil),
	new(big.Int).Exp(big.NewInt(10), big.NewInt(11), nil),
	new(big.Int).Exp(big.NewInt(10), big.NewInt(12), nil),
	new(big.Int).Exp(big.NewInt(10), big.NewInt(13), nil),
	new(big.Int).Exp(big.NewInt(10), big.NewInt(14), nil),
	new(big.Int).Exp(big.NewInt(10), big.NewInt(15), nil),
	new(big.Int).Exp(big.NewInt(10), big.NewInt(16), nil),
	new(big.Int).Exp(big.NewInt(10), big.NewInt(17), nil),
	new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil),
}

// Decimal represents a fixed-point decimal number with 18 decimal places
type Decimal struct {
	Precition int
	Value     *big.Int
}

func NewDefaultDecimal(v int64) *Decimal {
	return &Decimal{Precition: DEFAULT_PRECISION, Value: new(big.Int).SetInt64(v)}
}

func NewDecimal(v int64, p int) *Decimal {
	if p > MAX_PRECISION {
		p = MAX_PRECISION
	}
	return &Decimal{Precition: p, Value: new(big.Int).SetInt64(v)}
}

func NewDecimalCopy(other *Decimal) *Decimal {
	if other == nil {
		return nil
	}
	return &Decimal{Precition: other.Precition, Value: new(big.Int).Set(other.Value)}
}

// NewDecimalFromString creates a Decimal instance from a string
func NewDecimalFromString(s string, maxPrecision int) (*Decimal, error) {
	if s == "" {
		return nil, errors.New("empty string")
	}

	parts := strings.Split(s, ".")
	if len(parts) > 2 {
		return nil, fmt.Errorf("invalid decimal format: %s", s)
	}

	integerPartStr := parts[0]
	if integerPartStr == "" || integerPartStr[0] == '+' {
		return nil, errors.New("empty integer")
	}

	integerPart, ok := new(big.Int).SetString(parts[0], 10)
	if !ok {
		return nil, fmt.Errorf("invalid integer format: %s", parts[0])
	}

	currPrecision := 0
	decimalPart := big.NewInt(0)
	if len(parts) == 2 {
		decimalPartStr := parts[1]
		if decimalPartStr == "" || decimalPartStr[0] == '-' || decimalPartStr[0] == '+' {
			return nil, errors.New("empty decimal")
		}

		currPrecision = len(decimalPartStr)
		if currPrecision > maxPrecision {
			return nil, fmt.Errorf("decimal exceeds maximum precision: %s", s)
		}
		n := maxPrecision - currPrecision
		for i := 0; i < n; i++ {
			decimalPartStr += "0"
		}
		decimalPart, ok = new(big.Int).SetString(decimalPartStr, 10)
		if !ok || decimalPart.Sign() < 0 {
			return nil, fmt.Errorf("invalid decimal format: %s", parts[0])
		}
	}

	value := new(big.Int).Mul(integerPart, precisionFactor[maxPrecision])
	if value.Sign() < 0 {
		value = value.Sub(value, decimalPart)
	} else {
		value = value.Add(value, decimalPart)
	}

	return &Decimal{Precition: int(maxPrecision), Value: value}, nil
}

// String returns the string representation of a Decimal instance
func (d *Decimal) String() string {
	if d == nil {
		return "0"
	}
	value := new(big.Int).Abs(d.Value)
	quotient, remainder := new(big.Int).QuoRem(value, precisionFactor[d.Precition], new(big.Int))
	sign := ""
	if d.Value.Sign() < 0 {
		sign = "-"
	}
	if remainder.Sign() == 0 {
		return fmt.Sprintf("%s%s", sign, quotient.String())
	}
	decimalPart := fmt.Sprintf("%0*d", d.Precition, remainder)
	decimalPart = strings.TrimRight(decimalPart, "0")
	return fmt.Sprintf("%s%s.%s", sign, quotient.String(), decimalPart)
}

// Add adds two Decimal instances and returns a new Decimal instance
func (d *Decimal) Add(other *Decimal) *Decimal {
	if d == nil && other == nil {
		return nil
	}
	if other == nil {
		value := new(big.Int).Set(d.Value)
		return &Decimal{Precition: d.Precition, Value: value}
	}
	if d == nil {
		value := new(big.Int).Set(other.Value)
		return &Decimal{Precition: other.Precition, Value: value}
	}
	if d.Precition != other.Precition {
		Log.Panic("precition not match")
	}
	value := new(big.Int).Add(d.Value, other.Value)
	return &Decimal{Precition: d.Precition, Value: value}
}

// Sub subtracts two Decimal instances and returns a new Decimal instance
func (d *Decimal) Sub(other *Decimal) *Decimal {
	if d == nil && other == nil {
		return nil
	}
	if other == nil {
		value := new(big.Int).Set(d.Value)
		return &Decimal{Precition: d.Precition, Value: value}
	}
	if d == nil {
		value := new(big.Int).Neg(other.Value)
		return &Decimal{Precition: other.Precition, Value: value}
	}
	if d.Precition != other.Precition {
		Log.Panicf("precition not match, (%d != %d)", d.Precition, other.Precition)
	}
	value := new(big.Int).Sub(d.Value, other.Value)
	return &Decimal{Precition: d.Precition, Value: value}
}


// Mul muls two Decimal instances and returns a new Decimal instance
func (d *Decimal) Mul(other *big.Int) *Decimal {
	if d == nil || other == nil {
		return nil
	}
	value := new(big.Int).Mul(d.Value, other)
	// value := new(big.Int).Div(value0, precisionFactor[other.Precition])
	return &Decimal{Precition: d.Precition, Value: value}
}


// Div divs two Decimal instances and returns a new Decimal instance
func (d *Decimal) Div(other *big.Int) *Decimal {
	if d == nil || other == nil {
		return nil
	}
	// value0 := new(big.Int).Mul(d.Value, precisionFactor[other.Precition])
	value := new(big.Int).Div(d.Value, other)
	return &Decimal{Precition: d.Precition, Value: value}
}


func (d *Decimal) Cmp(other *Decimal) int {
	if d == nil && other == nil {
		return 0
	}
	if other == nil {
		return d.Value.Sign()
	}
	if d == nil {
		return -other.Value.Sign()
	}
	if d.Precition != other.Precition {
		Log.Panicf("precition not match, (%d != %d)", d.Precition, other.Precition)
	}
	return d.Value.Cmp(other.Value)
}

func (d *Decimal) CmpAlign(other *Decimal) int {
	if d == nil && other == nil {
		return 0
	}
	if other == nil {
		return d.Value.Sign()
	}
	if d == nil {
		return -other.Value.Sign()
	}
	return d.Value.Cmp(other.Value)
}

func (d *Decimal) Sign() int {
	if d == nil {
		return 0
	}
	return d.Value.Sign()
}

func (d *Decimal) IsOverflowUint64() bool {
	if d == nil {
		return false
	}

	integerPart := new(big.Int).SetInt64(math.MaxInt64)
	value := new(big.Int).Mul(integerPart, precisionFactor[d.Precition])
	return d.Value.Cmp(value) > 0
}

func (d *Decimal) IsOverflowInt64() bool {
	if d == nil {
		return false
	}

	integerPart := new(big.Int).SetInt64(math.MaxInt64)
	value := new(big.Int).Mul(integerPart, precisionFactor[d.Precition])
	return d.Value.Cmp(value) > 0
}

func (d *Decimal) GetMaxInt64() *Decimal {
	if d == nil {
		return nil
	}
	integerPart := new(big.Int).SetInt64(math.MaxInt64)
	value := new(big.Int).Mul(integerPart, precisionFactor[d.Precition])
	return &Decimal{Precition: d.Precition, Value: value}
}

func (d *Decimal) Float64() float64 {
	if d == nil {
		return 0
	}
	value := new(big.Int).Abs(d.Value)
	quotient, remainder := new(big.Int).QuoRem(value, precisionFactor[d.Precition], new(big.Int))
	decimalPart := float64(remainder.Int64()) / float64(precisionFactor[d.Precition].Int64())
	result := float64(quotient.Int64()) + decimalPart
	if d.Value.Sign() < 0 {
		return -result
	}
	return result
}


func (d *Decimal) IntegerPart() int64 {
	if d == nil {
		return 0
	}
	value := new(big.Int).Abs(d.Value)
	quotient, _ := new(big.Int).QuoRem(value, precisionFactor[d.Precition], new(big.Int))
	return quotient.Int64()
}

func (d *Decimal) ToInt64WithMax(max int64) (int64) {
	if d == nil {
		return 0
	}

	if max <= 0 {
		Log.Panicf("invalid max %d", max)
	}

	if d.IntegerPart() > max {
		Log.Panicf("invalid max %d", max)
	}

	scaleIndex := int(math.Log10(float64(math.MaxInt64) / float64(max)))
	if scaleIndex < 0 {
		scaleIndex = 0
	}

	// d.Value 不会
	//value := d.Value.Int64() * (precisionFactor[scaleIndex]).Int64()
	value := new(big.Int).Mul(d.Value, precisionFactor[scaleIndex])
	
	//result := &Decimal{Precition: d.Precition, Value: value}
	
	quotient, _ := new(big.Int).QuoRem(value, precisionFactor[d.Precition], new(big.Int))
	return quotient.Int64()
}

func NewDecimalFromInt64WithMax(value int64, max int64, precision int) (*Decimal, error) {
	if max <= 0 {
		return nil, fmt.Errorf("invalid max %d", max)
	}

	// 根据 max 和 math.MaxInt64 计算放大系数的指数
	scaleIndex := int(math.Log10(float64(math.MaxInt64) / float64(max)))
	if scaleIndex < 0 {
		scaleIndex = 0
	}
	//scaleFactor := math.Pow10(scaleIndex)

	// 计算浮点值
	//floatValue := float64(value) / scaleFactor
	bigValue := new(big.Int).Mul(new(big.Int).SetInt64(value), precisionFactor[precision])
	bigValue = new(big.Int).Div(bigValue, precisionFactor[scaleIndex])
	
	result := &Decimal{Precition: precision, Value: bigValue}
	return result, nil

	// // 转换浮点值为 Decimal 的整数表示
	// integerPart := result.IntegerPart()
	// fractionPart := floatValue - float64(integerPart)

	// valueBigInt := new(big.Int).SetInt64(integerPart)
	// if fractionPart > 0 {
	// 	fractionBigInt := big.NewInt(int64(fractionPart * math.Pow10(int(precision))))
	// 	valueBigInt.Mul(valueBigInt, precisionFactor[precision])
	// 	valueBigInt.Add(valueBigInt, fractionBigInt)
	// } else {
	// 	valueBigInt.Mul(valueBigInt, precisionFactor[precision])
	// }

	// return &Decimal{
	// 	Value:     valueBigInt,
	// 	Precition: precision,
	// }, nil
}
