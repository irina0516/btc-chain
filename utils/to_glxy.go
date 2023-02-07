package utils

import "math/big"

// ToGlxy number of GLXY to Wei
func ToGlxy(glxy uint64) *big.Int {
	return new(big.Int).Mul(new(big.Int).SetUint64(glxy), big.NewInt(1e18))
}
