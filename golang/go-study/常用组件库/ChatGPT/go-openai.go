package main

import (
	"bytes"

	"context"

	"encoding/json"

	"fmt"

	"log"

	"net/http"
)

// github: https://github.com/sashabaranov/go-openai
// https://juejin.cn/post/7210639699583713339?#heading-5

const (
	apiToken = "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"

	apiURL = "https://api.openai.com/v1/chat/completions"
)

type chatCompletionMessage struct {
	Role string `json:"role"`

	Content string `json:"content"`
}

type chatCompletionRequest struct {
	Model string `json:"model"`

	MaxTokens int `json:"max_tokens"`

	Messages []chatCompletionMessage `json:"messages"`
}

func main() {

	message := chatCompletionMessage{

		Role: "user",

		Content: "你好",
	}

	request := chatCompletionRequest{

		Model: "gpt-3.5-turbo",

		MaxTokens: 1024,

		Messages: []chatCompletionMessage{message},
	}

	ctx := context.Background()

	reqBytes, err := json.Marshal(request)

	if err != nil {

		log.Fatalf("error marshaling request: %v", err)

	}

	client := &http.Client{}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, apiURL, bytes.NewBuffer(reqBytes))

	if err != nil {

		log.Fatalf("error creating request: %v", err)

	}

	req.Header.Set("Accept", "application/json; charset=utf-8")

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	res, err := client.Do(req)

	if err != nil {

		log.Fatalf("error sending request: %v", err)

	}

	defer res.Body.Close()

	var v interface{}

	if err = json.NewDecoder(res.Body).Decode(&v); err != nil {

		log.Fatalf("error decoding response: %v", err)

	}

	fmt.Println(v)

}
