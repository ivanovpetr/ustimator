package estimator

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEstimate(t *testing.T) {
	var (
		inputReserve   = big.NewInt(9000000000000000000)
		outputReserve  = big.NewInt(16461000000)
		inputAmount    = big.NewInt(1000000000000000000)
		expectedOutput = big.NewInt(1641654196)
	)

	result := Estimate(inputAmount, inputReserve, outputReserve)
	require.Equal(t, expectedOutput, result)
}
