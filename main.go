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
	fmt.Println("  commit   Generate a commit message based on the changes in the repository")
	fmt.Println("  improve  Suggest improvements, fix bugs, and propose new features for the project")
	fmt.Println("  help     Display this help message")
}
