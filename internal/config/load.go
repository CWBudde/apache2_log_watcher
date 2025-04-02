package config

import (
	"apache2watcher/internal/notifier"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func LoadConfig(path string) (notifier.Config, error) {
	var cfg notifier.Config

	file, err := os.ReadFile(path)
	if err != nil {
		return cfg, fmt.Errorf("could not read config file: %w", err)
	}

	if err := yaml.Unmarshal(file, &cfg); err != nil {
		return cfg, fmt.Errorf("could not parse config: %w", err)
	}

	return cfg, nil
}