package cmd

import (
	"fmt"

	"ivanovpetr/ustimator/pkg/uniswapclient"

	"github.com/spf13/cobra"
)

// NewPairInfo creates pair info command
func NewPairInfo() *cobra.Command {
	c := &cobra.Command{
		Use:           "info",
		Short:         "Get info about pair reserves",
		SilenceUsage:  true,
		SilenceErrors: true,
		Args:          cobra.ExactArgs(1),
		RunE:          pairInfoHandler,
	}

	c.Flags().String(flagETHRPC, MainnetUrl, "Configure ethereum rpc endpoint")

	return c
}

// pairInfoHandler handles pair info command
func pairInfoHandler(cmd *cobra.Command, args []string) error {

	rpcUrl, _ := cmd.Flags().GetString(flagETHRPC)
	contract, err := uniswapclient.NewPoolContract(args[0], rpcUrl)
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
	fmt.Printf("Reserve for %s: %d\n", t0, r0)
	fmt.Printf("Reserve for %s: %d\n", t1, r1)
	return nil
}
