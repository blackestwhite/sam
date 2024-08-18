package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func GetProjectFiles() (string, error) {
	ignorePatterns, err := getIgnorePatterns()
	if err != nil {
		return "", err
	}

	var fileContents strings.Builder

	err = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if isIgnored(path, ignorePatterns) {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		if !info.IsDir() {
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

type IgnorePattern struct {
	pattern string
	isRegex bool
}

func getIgnorePatterns() ([]IgnorePattern, error) {
	content, err := os.ReadFile(".samignore")
	if err != nil {
		if os.IsNotExist(err) {
			return []IgnorePattern{}, nil
		}
		return nil, err
	}

	lines := strings.Split(string(content), "\n")
	var ignorePatterns []IgnorePattern
	for _, line := range lines {
		if trimmed := strings.TrimSpace(line); trimmed != "" && !strings.HasPrefix(trimmed, "#") {
			if strings.HasPrefix(trimmed, "/") && strings.HasSuffix(trimmed, "/") {
				// It's a regex pattern
				pattern := strings.Trim(trimmed, "/")
				if _, err := regexp.Compile(pattern); err != nil {
					return nil, fmt.Errorf("invalid regex in .samignore: %s", trimmed)
				}
				ignorePatterns = append(ignorePatterns, IgnorePattern{pattern: pattern, isRegex: true})
			} else {
				// It's a glob pattern
				ignorePatterns = append(ignorePatterns, IgnorePattern{pattern: trimmed, isRegex: false})
			}
		}
	}

	return ignorePatterns, nil
}

func isIgnored(path string, ignorePatterns []IgnorePattern) bool {
	for _, pattern := range ignorePatterns {
		if pattern.isRegex {
			if matched, _ := regexp.MatchString(pattern.pattern, path); matched {
				return true
			}
		} else {
			// Check if the pattern matches the full path
			if matched, _ := filepath.Match(pattern.pattern, path); matched {
				return true
			}

			// Check if any part of the path matches the pattern
			pathParts := strings.Split(filepath.ToSlash(path), "/")
			for i := range pathParts {
				if matched, _ := filepath.Match(pattern.pattern, strings.Join(pathParts[i:], "/")); matched {
					return true
				}
			}
		}
	}
	return false
}
