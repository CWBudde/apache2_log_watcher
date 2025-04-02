package notifier

import (
	"fmt"
	"os/exec"
)

// SendSignal sends a Signal message via signal-cli.
func SendSignal(from string, to string, message string) error {
	cmd := exec.Command("signal-cli", "-u", from, "send", "-m", message, to)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to send Signal message: %v\nOutput: %s", err, string(output))
	}

	return nil
}