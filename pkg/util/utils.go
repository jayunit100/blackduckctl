package util

import (
	"bytes"
	"fmt"
	"os/exec"
	"time"
)

// runWithTimeout runs a command and times it out at the specified duration
func RunWithTimeout(cmd *exec.Cmd, d time.Duration) (error, string) {
	timeout := time.After(d)

	// Use a bytes.Buffer to get the output
	var buf bytes.Buffer
	cmd.Stdout = &buf

	cmd.Start()

	// Use a channel to signal completion so we can use a select statement
	done := make(chan error)
	go func() { done <- cmd.Wait() }()

	// The select statement allows us to execute based on which channel
	// we get a message from first.
	select {
	case <-timeout:
		// Timeout happened first, kill the process and print a message.
		cmd.Process.Kill()
		return fmt.Errorf("Killed due to timeout"), buf.String()
	case err := <-done:
		if err != nil {
			return nil, buf.String()
		} else {
			return err, buf.String()
		}
	}
}
