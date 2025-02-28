package cpu

import (
	"github.com/ellioht/btc-miner/core"
	"math/big"
)

func IsSolved(header *core.Header) bool {
	hash := header.Hash()
	hashBig := hash.Big()
	target := GetTarget(header.Bits)
	if hashBig.Cmp(target) <= 0 {
		return true
	}
	return false
}

func GetTarget(bits uint32) *big.Int {
	exponent := byte(bits >> 24)
	coefficient := int64(bits & 0x007fffff)
	shift := int(exponent) - 3

	var target big.Int
	if shift >= 0 {
		multiplier := new(big.Int).Exp(big.NewInt(256), big.NewInt(int64(shift)), nil)
		target.Mul(big.NewInt(coefficient), multiplier)
	} else {
		divisor := new(big.Int).Exp(big.NewInt(256), big.NewInt(int64(-shift)), nil)
		target.Div(big.NewInt(coefficient), divisor)
	}

	return &target
}
