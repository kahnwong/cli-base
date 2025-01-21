package cli_base

import (
	"os"
	"strings"

	"gopkg.in/yaml.v3"

	"github.com/getsops/sops/v3/decrypt"
	"github.com/rs/zerolog/log"
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

func decryptSops(path string, format string) []byte {
	data, err := decrypt.File(path, format) // format: yaml, txt, etc. Refer to sops docs.
	if err != nil {
		log.Error().Msgf("Failed to decrypt sops config at: %s", path)
		os.Exit(1)
	}

	return data
}

func ReadYamlSops[T any](path string) *T {
	// check if config exists
	path, err := CheckIfConfigExists(path)
	if err != nil {
		log.Error().Msgf("Config doesn't exist at: %s", path)
		os.Exit(1)
	}

	// decrypt sops
	data := decryptSops(path, "yaml")

	// unmarshall
	// ref: <https://stackoverflow.com/a/71955439>
	config := new(T)
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Error().Msgf("Error unmarshalling config: %s", path)
	}

	return config
}
