package watcher

import (
	"fmt"
	"io"
	"strings"

	"github.com/hpcloud/tail"
)

// Callback defines the signature for actions triggered on matching lines.
type Callback func(line string)

// WatchLog tails the log file and triggers a callback if a line contains the grep string.
func WatchLog(logPath string, grep string, callback Callback) error {
	t, err := tail.TailFile(logPath, tail.Config{
		Follow:    true,
		ReOpen:    true,
		MustExist: true,
		Poll:      true,
		Location: &tail.SeekInfo{
			Whence: io.SeekEnd,
		},
	})
	if err != nil {
		return fmt.Errorf("failed to tail file: %w", err)
	}

	for line := range t.Lines {
		if line.Err != nil {
			continue
		}

		if grep == "" || strings.Contains(line.Text, grep) {
			callback(line.Text)
		}
	}

	return nil
}
