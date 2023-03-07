package main

import (
	"context"
	"fmt"
	"github.com/otiai10/openaigo-stream"
	"os"
	"strings"
)

func stream() {

	ctx = context.Background()
	client = openaigo.NewClient("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")

	var question string
	fmt.Scanln(&question)

	if strings.EqualFold(question, "重置") {
		fmt.Println("已重置，请继续提问")
		message = defaultMessage()
		//ask("")
	} else if strings.EqualFold(question, "退出") {
		fmt.Println("Bye")
		os.Exit(0)
	} else {
		if len(lastAnswer) > 0 {
			msgAns := openaigo.ChatMessage{Role: "system", Content: lastAnswer}
			message = append(message, msgAns)
		}

		if len(question) > 0 {
			msgNewQuestion := openaigo.ChatMessage{Role: "user", Content: question}
			message = append(message, msgNewQuestion)
		} else {
			//ask(lastAnswer)
		}
	}

	request := openaigo.ChatCompletionRequestBody{
		Model:    "gpt-3.5-turbo-0301",
		Messages: message,
		Stream: true,
	}

	client.ChatStream(ctx, request, func(response openaigo.ChatCompletionStreamResponse, err error) {
		if err == nil {
			fmt.Printf("%s", response.Choices[0].Delta.Content)
		}
	})
}