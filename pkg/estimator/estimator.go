package estimator

import (
	"math/big"
)

var (
	comissionK = big.NewInt(997)
	grand      = big.NewInt(1000)
)

// Estimate estimates output of uniswap v2 pair swap
func Estimate(inputAmount, inputReserve, outputReserve *big.Int) *big.Int {
	inputAmount.Mul(inputAmount, comissionK)
	// reuse outputReserve as numerator to avoid allocations
	outputReserve.Mul(inputAmount, outputReserve)
	// reuse inputReserve as denominator to avoid extra allocations
	inputReserve.Mul(inputReserve, grand)
	inputReserve.Add(inputReserve, inputAmount)
	// reuse inputAmount as result to avoid extra allocations
	// numerator(outputReserve) / denominator()
	return inputAmount.Div(outputReserve, inputReserve)
}
