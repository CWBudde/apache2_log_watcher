package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the version of the application",
	Run: func(cmd *cobra.Command, args []string) {
	  fmt.Println("Apache Visitor Alert CLI v0.1.0")
	},
}

func init() {
}
