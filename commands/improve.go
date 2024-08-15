package commands

import (
	"fmt"
	"log"

	"github.com/blackestwhite/sam/utils"
)

func SuggestImprovements() {
	projectFiles, err := utils.GetProjectFiles()
	if err != nil {
		log.Fatal("Error reading project files:", err)
	}

	prompt := `You are an AI assistant in a CLI tool. Analyze the provided project structure and file contents, then suggest improvements, bug fixes, tests, unit tests and new features. Format your response as follows:

Improvements:
1. [Improvement suggestion]
[Code Snippet]
2. [Improvement suggestion]
[Code Snippet]

Bug Fixes:
1. [Bug fix suggestion]
[Code Snippet]
2. [Bug fix suggestion]
[Code Snippet]

New Features:
1. [Feature suggestion]
2. [Feature suggestion]

Tests:
1. [Test suggestion code snippet]
2. [Test suggestion code snippet]

Be concise and focus on the most important suggestions.`

	response, err := utils.GetOpenAIResponse(prompt, projectFiles, "gpt-4o-mini")
	if err != nil {
		log.Fatal("Error generating suggestions:", err)
	}

	fmt.Println(response)
}
