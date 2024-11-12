package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "acheckin",
	Version: "1.0.0-beta",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// c := cron.New()
	// _, err := c.AddFunc("27 11 * * *", func() {
	// 	now := time.Now()
	// 	log.Printf("Cron job executed at: %v", now.Format(time.RFC3339))
	// })
	// if err != nil {
	// 	log.Fatalf("Error scheduling cron job: %v", err)
	// }

	// log.Println("Starting cron scheduler...")
	// c.Start()
	// select {}
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("config", "c", "", "Config file path")
}
