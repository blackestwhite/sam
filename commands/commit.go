package commands

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/blackestwhite/sam/utils"
)

func SuggestCommit() {
	projectFiles, err := utils.GetProjectFiles()
	if err != nil {
		log.Fatal("Error reading project files:", err)
	}

	diffCmd := exec.Command("git", "diff")
	diffOutput, err := diffCmd.Output()
	if err != nil {
		log.Fatal("Error running git diff:", err)
	}

	statusCmd := exec.Command("git", "status")
	statusOutput, err := statusCmd.Output()
	if err != nil {
		log.Fatal("Error running git status:", err)
	}

	logCmd := exec.Command("git", "log", "-1", "--pretty=%B")
	logOutput, err := logCmd.Output()
	if err != nil {
		log.Fatal("Error running git log:", err)
	}

	gitDiffOutput := string(diffOutput)
	gitStatusOutput := string(statusOutput)
	gitLogOutput := string(logOutput)
	projectFilesContent := string(projectFiles)

	prompt := "You are an AI in a CLI tool. Suggest a commit message based on the output of `git diff`, `git status`, the last commit message, and the contents of the project files with the following format: 'label: commit message'. The label should be from common labels like: fix, feat, chore, refactor, etc. Don't be verbose. Your response should be a one-line message suitable for a commit message."

	fullContext := fmt.Sprintf("Diff:\n%s\nStatus:\n%s\nLast Commit Message:\n%s\nProject Files:\n%s", gitDiffOutput, gitStatusOutput, gitLogOutput, projectFilesContent)

	response, err := utils.GetOpenAIResponse(prompt, fullContext, "gpt-4o-mini")
	if err != nil {
		log.Fatal("Error generating commit message:", err)
	}

	fmt.Println(response)
}
