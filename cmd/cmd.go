package cmd

import (
	"github.com/spf13/cobra"
)

// New creates root command for ustimator
func New() *cobra.Command {
	cobra.EnableCommandSorting = false

	c := &cobra.Command{
		Use:           "ustimator",
		Short:         "Estimate output of uniswap v2 token swap",
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	c.AddCommand(NewPair())

	return c
}
