package cmd

import (
	"github.com/spf13/cobra"
)

var checkinCmd = &cobra.Command{
	Use: "checkin",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(checkinCmd)
}
