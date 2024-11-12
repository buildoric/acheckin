package cmd

import (
	"github.com/spf13/cobra"
)

var checkoutCmd = &cobra.Command{
	Use: "checkout",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(checkoutCmd)
}
