package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/Andrem19/gpt_trade/variables"
)

type Data struct {
	ID      string                 `json:"id"`
	Object  string                 `json:"object"`
	Created int                    `json:"created"`
	Model   string                 `json:"model"`
	Choices []TextCompletionChoice `json:"choices"`
	Usage   TextCompletionUsage    `json:"usage"`
}

type RequestBody struct {
	Model            string  `json:"model"`
	Prompt           string  `json:"prompt"`
	Temperature      float64 `json:"temperature"`
	MaxTokens        int64   `json:"max_tokens"`
	TopP             float64 `json:"top_p"`
	FrequencyPenalty float64 `json:"frequency_penalty"`
	PresencePenalty  float64 `json:"presence_penalty"`
}

type TextCompletionChoice struct {
	Text         string  `json:"text"`
	Index        int     `json:"index"`
	LogProbs     *string `json:"logprobs"`
	FinishReason string  `json:"finish_reason"`
}

type TextCompletionUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

func AskQuestion(question string, gpt_token string) (string, error) {

	var err error

	client := &http.Client{}
	requestBody := RequestBody{
		Model:            variables.FineTuneModel,
		Prompt:           question,
		Temperature:      0.7,
		MaxTokens:        3000,
		TopP:             1,
		FrequencyPenalty: 0,
		PresencePenalty:  0,
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Send to bot: ", question)
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/completions", bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err, req)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", `Bearer `+gpt_token+``)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("The token is valid?", err)
	}
	// check if the token is valid.
	if resp.StatusCode == 401 {
		fmt.Println("The token is invalid")
		os.Exit(0)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	var response Data
	json.Unmarshal(bodyText, &response)
	if err != nil {
		log.Println(err)
	}
	choice := response.Choices[0]
	text := choice.Text
	fmt.Println("Received from bot: ", text)
	return text, nil
}