package cmd

import (
	"fmt"
	"log"
	"log/slog"

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
			slog.Error("Failed to load config", "error", err)
			log.Fatalf("Failed to load config: %v", err)
		}
	
		err = watcher.WatchLog("/var/log/apache2/access.log", grep, func(line string) {
			slog.Info("Log matched", "line", line)

			msg := "New visitor:\n" + line
			if err := notifier.Send(cfg, msg); err != nil {
				slog.Error("Failed to send notification", "error", err)
			} else {
				slog.Info("Notification sent successfully")
			}
		})
		if err != nil {
			slog.Error("Error watching log", "error", err)
			log.Fatalf("Error watching log: %v", err)
		}
	},
}

func init() {
	watchCmd.Flags().StringVarP(&grep, "grep", "g", "", "Filter log lines that contain this string")
	watchCmd.Flags().StringVar(&configPath, "config", "config.yaml", "Path to config file")
}
