package main

import "math/big"

const (
	BASE_1_DIFFICULTY = "0x00000000ffff0000000000000000000000000000000000000000000000000000"
)


func GetHexTarget(packedTarget *big.Int) *big.Int {
	var mantissa big.Int
	var exponent big.Int
	var t1 big.Int
	var result big.Int
	t1.SetString("0x00ffffff", 0)
	mantissa.And(packedTarget, &t1)
	t1.SetString("0xff000000", 0)
	exponent.And(packedTarget, &t1)
	exponent.Rsh(&exponent, 6*4)
	result.Sub(&exponent, new(big.Int).SetUint64(3))
	result.Mul(&result, new(big.Int).SetInt64(8))
	result.Lsh(&mantissa, uint(result.Uint64()))
	return &result
}

func CalculateDifficulty(t *big.Int) *big.Float {
	var difficulty1Target big.Int
	difficulty1Target.SetString(BASE_1_DIFFICULTY, 0)
	var result big.Float
	result.Quo(new(big.Float).SetInt(&difficulty1Target), new(big.Float).SetInt(t))
	return &result
}

func hashRateApprox(diff uint64, t *big.Int) *big.Int {
	var hashRate big.Int
	diff_big := new(big.Int).SetUint64(diff)
	diff_big.Lsh(diff_big, 32)
	hashRate.Quo(diff_big, t)
	return &hashRate
}

func difficultyToHexTarget(a uint64) *big.Int {
	var difficulty1Target big.Int
	difficulty1Target.SetString(BASE_1_DIFFICULTY, 0)
	diff := new(big.Int).SetUint64(a)
	var target big.Int
	target.Quo(&difficulty1Target, diff)
	return &target
}
