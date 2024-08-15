package utils

import (
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

	var response string
	for comp := range res {
		response += comp.Choices[0].Delta.Content
	}

	return response, nil
}
