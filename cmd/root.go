package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var verbose bool

var rootCmd = &cobra.Command{
	Use:   "visitor-alert",
	Short: "A tool to watch Apache logs and send notifications on visits",
	Long:  `visitor-alert is a CLI to monitor Apache access logs and trigger alerts like email or Signal.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(watchCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable debug logging")
}
