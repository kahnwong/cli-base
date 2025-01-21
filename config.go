package cli_base

import (
	"os"
	"strings"

	"github.com/getsops/sops/v3/decrypt"
	"github.com/rs/zerolog/log"
)

func ReadSops(path string, format string) []byte {
	// check if config exists
	path, err := CheckIfConfigExists(path)
	if err != nil {
		log.Error().Msgf("Config doesn't exist at: %s", path)
		os.Exit(1)
	}

	// decrypt sops
	data, err := decrypt.File(path, format) // format: yaml, txt, etc. Refer to sops docs.
	if err != nil {
		log.Error().Msgf("Failed to decrypt sops config at: %s", path)
		os.Exit(1)
	}

	return data
}

func CheckIfConfigExists(path string) (string, error) {
	// Set config path
	path = expandHome(path)

	// Check if the file exists
	_, err := os.Stat(path)

	return path, err
}

func expandHome(path string) string {
	home, _ := os.UserHomeDir()
	return strings.Replace(path, "~", home, 1)
}
