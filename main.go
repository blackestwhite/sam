package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	configFilePath := filepath.Join(usr.HomeDir, ".samrc")

	_, err = os.Stat(configFilePath)
	if os.IsNotExist(err) {
		fmt.Print("Enter your OpenAI API key: ")
		var apiKey string
		fmt.Scanln(&apiKey)

		err := os.WriteFile(configFilePath, []byte(apiKey), 0600)
		if err != nil {
			fmt.Println("Error writing to configuration file:", err)
			return
		}

		fmt.Println("API key saved successfully.")
	} else if err != nil {
		fmt.Println("Error:", err)
		return
	}

	apiKeyBytes, err := os.ReadFile(configFilePath)
	if err != nil {
		fmt.Println("Error reading API key from configuration file:", err)
		return
	}

	apiKey := string(apiKeyBytes)
	fmt.Println("API key:", apiKey)
}
