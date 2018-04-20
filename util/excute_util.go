package cloudimageeditor

/*
 * cloudimageeditor package
 *
 * Authors:
 *     zhenwei pi <cloudimageeditor@126.com>
 *
 * This project is under Apache v2 License.
 *
 */

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

// Execute : excute @binary with @args
func Execute(binary string, args []string) (string, error) {
	cmd := exec.Command(binary, args...)
	result, err := cmd.CombinedOutput()

	return string(result), err
}

// ExecuteTimeout : excute @binary with @seconds at most @timeout seconds
func ExecuteTimeout(binary string, args []string, seconds int) (string, error) {
	var result []byte
	var err error
	cmd := exec.Command(binary, args...)
	done := make(chan struct{})

	go func() {
		result, err = cmd.CombinedOutput()
		done <- struct{}{}
	}()

	select {
	case <-done:
	case <-time.After(time.Duration(seconds) * time.Second):
		if cmd.Process != nil {
			cmd.Process.Kill()
		}
		return "", fmt.Errorf("Timeout executing : %v %v, result %+v, error %v", binary, args, string(result), err)
	}

	if err != nil {
		if !strings.Contains(err.Error(), "no child processes") {
			return "", fmt.Errorf("Failed to execute : %v %v, result %+v, error %v", binary, args, string(result), err)
		}
	}

	return string(result), nil
}
