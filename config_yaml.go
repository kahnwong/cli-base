package cli_base

import (
	"os"

	"gopkg.in/yaml.v3"
)

func ReadYaml[T any](path string) (*T, error) {
	// check if config exists
	expandedPath, err := CheckIfConfigExists(path)
	if err != nil {
		return nil, err
	}

	// read file
	data, err := os.ReadFile(expandedPath)
	if err != nil {
		return nil, err
	}

	// unmarshall
	config := new(T)
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
