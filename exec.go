package cli_base

import (
	"os/exec"
)

func ExecCommand(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	stdout, err := cmd.Output()

	if err != nil {
		return "", err
	}

	return string(stdout), nil
}
