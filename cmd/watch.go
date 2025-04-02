package cmd

import (
	"fmt"
	"log"

	"apache2watcher/internal/config"
	"apache2watcher/internal/notifier"
	"apache2watcher/internal/watcher"

	"github.com/spf13/cobra"
)

var (
	grep string
  configPath string
)

var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Watch the access log and trigger alerts on visits",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Watching logs...")
		fmt.Printf("Grep filter: %s\n", grep)
	
		cfg, err := config.LoadConfig(configPath)
		if err != nil {
			log.Fatalf("Failed to load config: %v", err)
		}
	
		err = watcher.WatchLog("/var/log/apache2/access.log", grep, func(line string) {
			msg := "New visitor:\n" + line
			if err := notifier.Send(cfg, msg); err != nil {
				log.Println("Notify error:", err)
			}
		})
		if err != nil {
			log.Fatalf("Error watching log: %v", err)
		}
	},
}

func init() {
	watchCmd.Flags().StringVarP(&grep, "grep", "g", "", "Filter log lines that contain this string")
	watchCmd.Flags().StringVar(&configPath, "config", "config.yaml", "Path to config file")
}
