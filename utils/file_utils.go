package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetProjectFiles() (string, error) {
	ignoredPaths, err := getIgnoredPaths()
	if err != nil {
		return "", err
	}

	var fileContents strings.Builder

	err = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && isIgnored(path, ignoredPaths) {
			return filepath.SkipDir
		}

		if !info.IsDir() && !isIgnored(path, ignoredPaths) {
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			fileContents.WriteString(fmt.Sprintf("File: %s\n", path))
			fileContents.WriteString(string(content))
			fileContents.WriteString("\n\n")
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	return fileContents.String(), nil
}

func getIgnoredPaths() ([]string, error) {
	content, err := os.ReadFile(".samignore")
	if err != nil {
		if os.IsNotExist(err) {
			return []string{}, nil
		}
		return nil, err
	}

	lines := strings.Split(string(content), "\n")
	var ignoredPaths []string
	for _, line := range lines {
		if trimmed := strings.TrimSpace(line); trimmed != "" && !strings.HasPrefix(trimmed, "#") {
			ignoredPaths = append(ignoredPaths, trimmed)
		}
	}

	return ignoredPaths, nil
}

func isIgnored(path string, ignoredPaths []string) bool {
	for _, ignoredPath := range ignoredPaths {
		if strings.HasPrefix(path, ignoredPath) {
			return true
		}
	}
	return false
}
