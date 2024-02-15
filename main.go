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
	if len(os.Args) < 2 {
		fmt.Println("Usage: sam <command>")
		return
	}

	command := os.Args[1]

	switch command {
	case "commit":
		suggestCommit()
	default:
		fmt.Println("Invalid command. Available commands: commit")
	}
}

func suggestCommit() {
	usr, err := user.Current()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	configFilePath := filepath.Join(usr.HomeDir, ".samrc")

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
			gopenai.Message{
				Role:    "system",
				Content: "you are a AI in a cli tool. you should suggest commit message based on output of git diff with the following format: 'label: commit message' label should be from common labels like: fix, feat, chore, refactor and ..., and also don't be verbose your response should be a one line response that should be used directly in cli tool.",
			},
			gopenai.Message{
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
