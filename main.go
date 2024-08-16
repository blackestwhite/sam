package main

import (
	"fmt"
	"os"

	"github.com/blackestwhite/sam/commands"
	"github.com/blackestwhite/sam/config"
)

func main() {
	if err := config.LoadConfig(); err != nil {
		fmt.Println("Error loading config:", err)
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		printHelp()
		return
	}

	command := os.Args[1]

	switch command {
	case "commit":
		commands.SuggestCommit()
	case "improve":
		commands.SuggestImprovements()
	case "request":
		if len(os.Args) < 3 {
			fmt.Println("Error: No request provided. Usage: sam request <your_request>")
			return
		}
		request := os.Args[2] // Capture the request from command-line arguments
		commands.HandleRequest(request)
	case "help":
		printHelp()
	default:
		fmt.Println("Invalid command. Use 'sam help' for usage information.")
	}
}

func printHelp() {
	fmt.Println("Usage: sam <command>")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  commit   Generate a commit message based on changes in the repository")
	fmt.Println("  improve  Suggest improvements and fixes for the project")
	fmt.Println("  request   Handle specific requests based on project analysis for example: sam request \"how to add feature x\"")
	fmt.Println("  help     Display this help message")
}
