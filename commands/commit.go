package commands

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/blackestwhite/sam/utils"
)

func SuggestCommit() {
	diffCmd := exec.Command("git", "diff")
	diffOutput, err := diffCmd.Output()
	if err != nil {
		log.Fatal("Error running git diff:", err)
	}

	gitDiffOutput := string(diffOutput)

	prompt := "You are an AI in a CLI tool. Suggest a commit message based on the output of `git diff` with the following format: 'label: commit message'. The label should be from common labels like: fix, feat, chore, refactor, etc. Don't be verbose. Your response should be a one-line message suitable for a commit message."

	response, err := utils.GetOpenAIResponse(prompt, gitDiffOutput, "gpt-4o-mini")
	if err != nil {
		log.Fatal("Error generating commit message:", err)
	}

	fmt.Println(response)
}
