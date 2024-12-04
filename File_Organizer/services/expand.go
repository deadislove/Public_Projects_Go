package services

import (
	"os"
	"path/filepath"
	"strings"
)

// ExpandPath handles "~" for Git Bash and convert to an absolute path
func ExpandPath(path string) (string, error) {
	// Expand tilde (~) to home directory
	if strings.HasPrefix(path, "~") {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		path = filepath.Join(homeDir, path[1:])
	}

	//Convert to an absolute path
	absPath, err := filepath.Abs(filepath.Clean(path))
	if err != nil {
		return "", err
	}

	return absPath, nil
}
