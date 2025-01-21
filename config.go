package cli_base

import (
	"os"
	"strings"
)

func expandHome(path string) string {
	home, _ := os.UserHomeDir()
	return strings.Replace(path, "~", home, 1)
}

func CheckIfConfigExists(path string) (string, error) {
	// Set config path
	path = expandHome(path)

	// Check if the file exists
	_, err := os.Stat(path)

	return path, err
}
