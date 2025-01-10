// Copyright 2024 The BxELab studyzy Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package runestone

import (
	"errors"
	"math/big"

	"lukechampine.com/uint128"
)

type Address string

var ErrNone = errors.New("none")

func Encode(n *big.Int) []byte {
	var result []byte
	for n.Cmp(big.NewInt(128)) >= 0 {
		temp := new(big.Int).Set(n)
		last := temp.And(n, new(big.Int).SetUint64(0b0111_1111))
		result = append(result, last.Or(last, new(big.Int).SetUint64(0b1000_0000)).Bytes()[0])
		n.Rsh(n, 7)
	}
	if len(n.Bytes()) == 0 {
		result = append(result, 0)
	} else {
		result = append(result, n.Bytes()...)
	}
	return result
}

func Decode(encoded []byte) *big.Int {
	result := new(big.Int)
	for i := len(encoded) - 1; i >= 0; i-- {
		result.Lsh(result, 7)
		byteVal := new(big.Int).SetUint64(uint64(encoded[i] & 0b0111_1111))
		result.Or(result, byteVal)
	}
	return result
}

func EncodeUint64(n uint64) []byte {
	var result []byte
	for n >= 128 {
		result = append(result, byte(n&0x7F|0x80))
		n >>= 7
	}
	result = append(result, byte(n))
	return result
}
func EncodeUint32(n uint32) []byte {
	var result []byte
	for n >= 128 {
		result = append(result, byte(n&0x7F|0x80))
		n >>= 7
	}
	result = append(result, byte(n))
	return result
}
func EncodeUint8(n uint8) []byte {
	var result []byte
	for n >= 128 {
		result = append(result, byte(n&0x7F|0x80))
		n >>= 7
	}
	result = append(result, byte(n))
	return result
}
func EncodeUint128(n uint128.Uint128) []byte {
	return Encode(n.Big())
}
func EncodeChar(r rune) []byte {
	return EncodeUint32(uint32(r))
}
