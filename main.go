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
		printHelp()
		return
	}

	command := os.Args[1]

	switch command {
	case "commit":
		suggestCommit()
	case "help":
		printHelp()
	default:
		fmt.Println("Invalid command. Use 'sam help' for usage information.")
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

	diffCmd := exec.Command("git", "diff")
	diffOutput, err := diffCmd.Output()
	if err != nil {
		log.Fatal("Error running git diff:", err)
	}

	gitDiffOutput := string(diffOutput)

	res, err := instance.GenerateChatCompletion(gopenai.ChatCompletionRequestBody{
		Stream: true,
		Model:  "gpt-4",
		Messages: []gopenai.Message{
			{
				Role:    "system",
				Content: "you are an AI in a cli tool. you should suggest a commit message based on the output of `git diff` with the following format: 'label: commit message'. The label should be from common labels like: fix, feat, chore, refactor, etc., and also, don't be verbose. Your response should be a one-line message suitable for a commit message.",
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

func printHelp() {
	fmt.Println("Usage: sam <command>")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  commit   Generate a commit message based on the changes in the repository")
	fmt.Println("  help     Display this help message")
}
