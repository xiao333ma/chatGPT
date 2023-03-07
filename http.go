package main

import (
	"fmt"
	"strings"
)

func http(question string) {

	request := ask(question)

	response, err := client.Chat(ctx, request)

	fmt.Println("chatGPT:")

	if err != nil {
		fmt.Println("error: ", err)
	}

	lastAnswer := ""
	if len(response.Choices) > 0 {
		text := response.Choices[0].Message.Content
		text = strings.TrimLeft(text, "\n")
		lastAnswer = text
		fmt.Println(text)
	}
	http(lastAnswer)
}
