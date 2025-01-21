package cli_base

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog/log"
)

func ExpandHome(path string) string {
	home, _ := os.UserHomeDir()
	return strings.Replace(path, "~", home, 1)
}

func CheckIfConfigExists(path string) (string, error) {
	// Set config path
	path = ExpandHome(path)

	// Check if the file exists
	_, err := os.Stat(path)

	return path, err
}

func CreateConfigIfNotExists(path string) {
	// create path
	dir := filepath.Dir(path)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		log.Fatal().Msgf("Error creating config path: %s", dir)
	}

	// create file
	file, err := os.Create(path)
	if err != nil {
		log.Fatal().Msgf("Error creating config file: %s", path)
	}
	defer file.Close()
}
