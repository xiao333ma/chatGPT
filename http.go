package main

import (
	"context"
	"fmt"
	"github.com/otiai10/openaigo-stream"
	"os"
	"strings"
)

var client *openaigo.Client
var ctx context.Context
var message []openaigo.ChatMessage

var lastAnswer string

func http()  {


	ctx = context.Background()
	client = openaigo.NewClient("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")
	fmt.Println("开始提问")
	ask( "")
}

func defaultMessage() []openaigo.ChatMessage {

	return []openaigo.ChatMessage{
	}
}

func ask(lastAnswer string)  {

	var question string
	fmt.Scanln(&question)

	if  strings.EqualFold(question, "重置") {
		fmt.Println("已重置，请继续提问")
		message = defaultMessage()
		ask("")
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
			ask(lastAnswer)
		}
	}

	request := openaigo.ChatCompletionRequestBody{
		Model:  "gpt-3.5-turbo",
		Messages: message,
	}
	response, err := client.Chat(ctx, request)

	fmt.Println("chatGPT:")

	if err != nil {
		fmt.Println("error: ", err)
	}

	if len(response.Choices) > 0 {
		text := response.Choices[0].Message.Content
		text = strings.TrimLeft(text, "\n")
		lastAnswer = text
		fmt.Println(text)
	}
	ask(lastAnswer)
}