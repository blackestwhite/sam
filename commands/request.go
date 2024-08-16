package commands

import (
	"fmt"
	"log"

	"github.com/blackestwhite/sam/utils"
)

func HandleRequest(request string) {
	projectFiles, err := utils.GetProjectFiles()
	if err != nil {
		log.Fatal("Error reading project files:", err)
	}

	prompt := `You are an AI assistant in a CLI tool. Analyze the provided project structure and file contents, then suggest responses for provided request. Format your response as follows(remember to mention which file/path should be worked on):

Responses for your reuqest:
1. [Response]
2. [Response]
3. [Response]
add more steps and responses if needed.


Be concise and focus on the most important suggestions.

The Request is: %s
`
	prompt = fmt.Sprintf(prompt, request)

	response, err := utils.GetOpenAIResponse(prompt, projectFiles, "gpt-4o-mini")
	if err != nil {
		log.Fatal("Error generating suggestions:", err)
	}

	fmt.Println(response)
}
