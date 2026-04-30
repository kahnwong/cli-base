package cli_base

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ExpandHome(path string) (string, error) {
	if !strings.HasPrefix(path, "~") {
		return path, nil
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return strings.Replace(path, "~", home, 1), nil
}

func CheckIfConfigExists(path string) (string, error) {
	// Set config path
	expandedPath, err := ExpandHome(path)
	if err != nil {
		return "", err
	}

	// Check if the file exists
	_, err = os.Stat(expandedPath)

	return expandedPath, err
}

func CreateConfigIfNotExists(path string) error {
	// create path
	dir := filepath.Dir(path)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}

	// create file
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	if err = file.Close(); err != nil {
		return fmt.Errorf("failed to close config file: %w", err)
	}

	return nil
}
