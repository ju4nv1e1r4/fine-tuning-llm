package main

import (
	"finetuning/model"
	"fmt"
	"log"
)

func main() {
	request := model.ChatCompletionRequest{
			Model: "ft:gpt-3.5-turbo-0125:juanvieira::B660pdOh",
			Messages: []model.Message{
					{Role: "developer", Content: "Você é um profissional de carreiras"},
					{Role: "user", Content: "Quem é você?"},
			},
	}

	response, err := model.SendOpenAIRequest(request)
	if err != nil {
			log.Fatalf("Erro ao enviar a requisição: %v", err)
	}

	if len(response.Choices) > 0 {
			fmt.Println("Resposta do Modelo:", response.Choices[0].Message.Content)
	} else {
			fmt.Println("Nenhuma resposta recebida.")
	}
}