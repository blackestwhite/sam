package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

var (
	APIKey string
)

func LoadConfig() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	configPath := filepath.Join(home, ".samrc")
	if err := godotenv.Load(configPath); err != nil {
		return err
	}

	APIKey = os.Getenv("OPENAI_API_KEY")
	if APIKey == "" {
		return fmt.Errorf("OPENAI_API_KEY not found in .samrc file")
	}

	return nil
}
