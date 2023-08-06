package cmd

import (
	"fmt"
	"math/big"

	"ivanovpetr/ustimator/pkg/estimator"
	"ivanovpetr/ustimator/pkg/uniswapclient"

	"github.com/spf13/cobra"
)

// NewPairEstimate creates pair estimate command
func NewPairEstimate() *cobra.Command {
	c := &cobra.Command{
		Use:           "estimate",
		Short:         "Estimate output of uniswap v2 token swap",
		SilenceUsage:  true,
		SilenceErrors: true,
		Args:          cobra.ExactArgs(4),
		RunE:          pairEstimateHandler,
	}

	c.Flags().String(flagETHRPC, MainnetUrl, "Configure ethereum rpc endpoint")

	return c
}

// pairEstimateHandler handles pair estimate command
func pairEstimateHandler(cmd *cobra.Command, args []string) error {
	var pairAddress, FromToken, ToToken = args[0], args[1], args[2]

	// parse input amount
	flt, _, err := big.ParseFloat(args[3], 10, 0, big.ToNearestEven)
	if err != nil {
		return fmt.Errorf("invalid input amount: %v", err)
	}
	var inputAmount = new(big.Int)
	inputAmount, _ = flt.Int(inputAmount)

	rpcUrl, _ := cmd.Flags().GetString(flagETHRPC)
	contract, err := uniswapclient.NewPoolContract(pairAddress, rpcUrl)
	if err != nil {
		return err
	}

	r0, r1, err := contract.GetReserves()
	if err != nil {
		return err
	}

	t0, t1, err := contract.TokenAddresses()
	if err != nil {
		return err
	}
	if FromToken != t0 && FromToken != t1 {
		return fmt.Errorf("token %s does not belong to pool %s", FromToken, pairAddress)
	}
	if ToToken != t0 && ToToken != t1 {
		return fmt.Errorf("token %s does not belong to pool %s", ToToken, pairAddress)
	}

	var swapResult *big.Int
	if FromToken == t0 {
		swapResult = estimator.Estimate(inputAmount, r0, r1)
	} else {
		swapResult = estimator.Estimate(inputAmount, r1, r0)
	}

	fmt.Printf("Estimated output swappping %d of token %s for token %s is %d \n", inputAmount, FromToken, ToToken, swapResult)
	return nil
}
