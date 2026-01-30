package cli_base

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog/log"
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
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			log.Error().Err(err).Msg("Failed to close config file")
		}
	}(file)

	return nil
}
