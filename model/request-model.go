package model

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

type Message struct {
        Role    string `json:"role"`
        Content string `json:"content"`
}

type ChatCompletionRequest struct {
        Model    string    `json:"model"`
        Messages []Message `json:"messages"`
}

type ChatCompletionResponse struct {
        Choices []struct {
                Message Message `json:"message"`
        } `json:"choices"`
}

func LoadAPIkey() (string, error) {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		return "", fmt.Errorf("erro ao ler o arquivo de configuração: %v", err)
	}
	return viper.GetString("OPENAI_API_KEY"), nil
}

func SendOpenAIRequest(request ChatCompletionRequest) (ChatCompletionResponse, error) {
	apiKey, err := LoadAPIkey()
	if err != nil {
		log.Fatal(err)
	}

	url := "https://api.openai.com/v1/chat/completions"

	requestBody, err := json.Marshal(request)
	if err != nil {
			return ChatCompletionResponse{}, fmt.Errorf("erro ao serializar a requisição: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
			return ChatCompletionResponse{}, fmt.Errorf("erro ao criar a requisição: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
			return ChatCompletionResponse{}, fmt.Errorf("erro ao enviar a requisição: %w", err)
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
			return ChatCompletionResponse{}, fmt.Errorf("erro ao ler a resposta: %w", err)
	}

	var chatResponse ChatCompletionResponse
	err = json.Unmarshal(responseBody, &chatResponse)
	if err != nil {
			return ChatCompletionResponse{}, fmt.Errorf("erro ao desserializar a resposta: %w", err)
	}

	return chatResponse, nil
}