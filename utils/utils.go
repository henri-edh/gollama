package utils

import (
	"os"
	"path/filepath"

	"github.com/sammcj/gollama/logging"
)

func GetHomeDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		logging.ErrorLogger.Printf("Failed to get user home directory: %v\n", err)

		return ""
	}
	return homeDir
}

// getConfigDir returns the directory of the configuration JSON file.
func GetConfigDir() string {
	return filepath.Join(GetHomeDir(), ".config", "gollama")
}

// getConfigPath returns the path to the configuration JSON file.
func GetConfigPath() string {
	return filepath.Join(GetHomeDir(), ".config", "gollama", "config.json")
}