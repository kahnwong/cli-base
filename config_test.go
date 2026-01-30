package cli_base

import (
	"os"
	"path/filepath"
	"testing"
)

func TestExpandHome(t *testing.T) {
	home, err := os.UserHomeDir()
	if err != nil {
		t.Fatalf("Failed to get user home directory: %v", err)
	}

	result, err := ExpandHome("~/.config/app")
	if err != nil {
		t.Errorf("ExpandHome() error = %v", err)
	}
	expected := filepath.Join(home, ".config/app")
	if result != expected {
		t.Errorf("ExpandHome() = %v, want %v", result, expected)
	}

	result, err = ExpandHome("/etc/config")
	if err != nil {
		t.Errorf("ExpandHome() error = %v", err)
	}
	if result != "/etc/config" {
		t.Errorf("ExpandHome() = %v, want /etc/config", result)
	}
}

func TestCheckIfConfigExists(t *testing.T) {
	tempDir := t.TempDir()
	existingFile := filepath.Join(tempDir, "existing.txt")

	if err := os.WriteFile(existingFile, []byte("test"), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	result, err := CheckIfConfigExists(existingFile)
	if err != nil {
		t.Errorf("CheckIfConfigExists() error = %v", err)
	}
	if result == "" {
		t.Errorf("CheckIfConfigExists() returned empty path")
	}

	_, err = CheckIfConfigExists(filepath.Join(tempDir, "nonexistent.txt"))
	if err == nil {
		t.Errorf("CheckIfConfigExists() expected error for non-existent file")
	}
}

func TestCreateConfigIfNotExists(t *testing.T) {
	tempDir := t.TempDir()

	path := filepath.Join(tempDir, "nested", "dir", "config.txt")
	err := CreateConfigIfNotExists(path)
	if err != nil {
		t.Errorf("CreateConfigIfNotExists() error = %v", err)
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Errorf("CreateConfigIfNotExists() did not create file at %s", path)
	}
}
