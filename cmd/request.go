// SPDX-FileCopyrightText: 2026 first-uninteresting-username
//
// SPDX-License-Identifier: GPL-3.0-only
package cmd

import (
	"bytes"
	"context"
	"fmt"
	"os"

	"github.com/sashabaranov/go-openai"
	"github.com/spf13/viper"
)

func makeRequest() (string, error) {
	systemPrompt, err := generateSystemPrompt()
	if err != nil {
		return "", err
	}

	prompt, err := generatePrompt()
	if err != nil {
		return "", err
	}

	apiKey := viper.GetString("api_key")
	if path := viper.GetString("api_key_path"); path != "" {
		data, err := os.ReadFile(path)
		if err != nil {
			return "", fmt.Errorf("reading api key file: %w", err)
		}
		apiKey = string(bytes.TrimSpace(data))
	}

	config := openai.DefaultConfig(apiKey)
	config.BaseURL = viper.GetString("base_url")
	client := openai.NewClientWithConfig(config)

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: viper.GetString("model"),
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: systemPrompt,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
