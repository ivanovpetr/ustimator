package cmd

import (
	"github.com/spf13/cobra"
)

const (
	flagETHRPC = "eth-rpc"
)

const (
	MainnetUrl = "https://mainnet.infura.io/v3"
)

// NewPair creates pair command
func NewPair() *cobra.Command {
	c := &cobra.Command{
		Use:           "pair",
		Short:         "Set of utilities for uniswap V2 pools",
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	c.Flags().String(flagETHRPC, MainnetUrl, "Configure ethereum rpc endpoint")

	c.AddCommand(NewPairInfo())
	c.AddCommand(NewPairEstimate())

	return c
}
