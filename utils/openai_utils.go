package utils

import (
	"errors"

	"github.com/blackestwhite/gopenai"
	"github.com/blackestwhite/sam/config"
)

func GetOpenAIResponse(prompt, content, model string) (string, error) {
	instance := gopenai.Setup(config.APIKey)

	res, err := instance.GenerateChatCompletion(gopenai.ChatCompletionRequestBody{
		Stream: true,
		Model:  model,
		Messages: []gopenai.Message{
			{
				Role:    "system",
				Content: prompt,
			},
			{
				Role:    "user",
				Content: content,
			},
		},
	})

	if err != nil {
		return "", err
	}

	if len(res.Choices) <= 0 {
		return "", errors.New("input files are bigger than expected")
	}

	response := res.Choices[0].Message.Content

	return response, nil
}
