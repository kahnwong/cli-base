package cli_base

import (
	"os"
	"path/filepath"
	"testing"
)

type TestConfig struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Port    int    `yaml:"port"`
}

func TestReadYaml(t *testing.T) {
	tempDir := t.TempDir()

	// Test valid yaml file
	filePath := filepath.Join(tempDir, "valid.yaml")
	content := `name: test-app
version: 1.0.0
port: 8080
`
	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	result, err := ReadYaml[TestConfig](filePath)
	if err != nil {
		t.Errorf("ReadYaml() error = %v", err)
	}
	if result.Name != "test-app" {
		t.Errorf("ReadYaml() Name = %v, want test-app", result.Name)
	}
	if result.Version != "1.0.0" {
		t.Errorf("ReadYaml() Version = %v, want 1.0.0", result.Version)
	}
	if result.Port != 8080 {
		t.Errorf("ReadYaml() Port = %v, want 8080", result.Port)
	}

	// Test non-existent file
	_, err = ReadYaml[TestConfig](filepath.Join(tempDir, "nonexistent.yaml"))
	if err == nil {
		t.Errorf("ReadYaml() expected error for non-existent file")
	}
}
