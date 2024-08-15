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

	prompt := `You are an AI assistant in a CLI tool. Analyze the provided project structure and file contents, then suggest improvements, bug fixes, and new features. Format your response as follows:

Improvements:
1. [Improvement suggestion]
2. [Improvement suggestion]

Bug Fixes:
1. [Bug fix suggestion]
2. [Bug fix suggestion]

New Features:
1. [Feature suggestion]
2. [Feature suggestion]

Be concise and focus on the most important suggestions.`

	response, err := utils.GetOpenAIResponse(prompt, projectFiles, "gpt-4o-mini")
	if err != nil {
		log.Fatal("Error generating suggestions:", err)
	}

	fmt.Println(response)
}
