package main

import (
	"apache2watcher/cmd"
	"log/slog"
	"os"
)

func init() {
	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})
	slog.SetDefault(slog.New(handler))
}

func main() {
	cmd.Execute()
}