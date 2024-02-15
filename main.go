package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"

	"github.com/blackestwhite/gopenai"
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

	instance := gopenai.Setup(apiKey)

	// Execute git diff command
	diffCmd := exec.Command("git", "diff")
	diffOutput, err := diffCmd.Output()
	if err != nil {
		log.Fatal("Error running git diff:", err)
	}

	// Convert output bytes to string
	gitDiffOutput := string(diffOutput)

	res, err := instance.GenerateChatCompletion(gopenai.ChatCompletionRequestBody{
		Stream: true,
		Model:  "gpt-3.5-turbo",
		Messages: []gopenai.Message{
			{
				Role:    "system",
				Content: "you are a AI in a cli tool. you should suggest commit message based on output of git diff with the following format: 'label: commit message' label should be from common labels like: fix, feat, chore, refactor and ..., and also don't be verbose your response should be a one line response that should be used directly in cli tool.",
			},
			{
				Role:    "user",
				Content: gitDiffOutput,
			},
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	txt := ""
	for comp := range res {
		txt += comp.Choices[0].Delta.Content
	}
	fmt.Println(txt)
}
