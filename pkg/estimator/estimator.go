package estimator

import (
	"math/big"
)

// Estimate estimates output of uniswap v2 pair swap
func Estimate(inputAmount, inputReserve, outputReserve *big.Int) *big.Int {
	var inputWithFee, numerator, denominator, result = &big.Int{}, &big.Int{}, &big.Int{}, &big.Int{}
	inputWithFee.Mul(inputAmount, big.NewInt(997))
	numerator.Mul(inputWithFee, outputReserve)
	denominator.Mul(inputReserve, big.NewInt(1000))
	denominator.Add(denominator, inputWithFee)
	result.Div(numerator, denominator)
	return result
}
