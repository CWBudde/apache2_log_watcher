package notifier

import "fmt"

func Send(cfg Config, message string) error {
	switch cfg.Channel {
	case Signal:
		return SendSignal(cfg.SignalFrom, cfg.SignalTo, message)
	case Email:
		return SendEmail(cfg, message)
	default:
		return fmt.Errorf("unknown channel: %s", cfg.Channel)
	}
}
