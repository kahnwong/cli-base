package cli_base

import (
	"os"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
)

func ReadYaml[T any](path string) *T {
	// check if config exists
	path, err := CheckIfConfigExists(path)
	if err != nil {
		log.Fatal().Msgf("Config doesn't exist at: %s", path)
	}

	// read file
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal().Msgf("Error reading config file: %s", path)
	}

	// unmarshall
	config := new(T)
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Error().Msgf("Error unmarshalling config: %s", path)
	}

	return config
}
