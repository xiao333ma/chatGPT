package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/otiai10/openaigo"
	"os"
	"strings"
)

var client *openaigo.Client
var ctx context.Context
var message []openaigo.ChatMessage

func init() {

	ctx = context.Background()
	client = openaigo.NewClient(token)
	fmt.Println("开始提问")
}

func defaultMessage() []openaigo.ChatMessage {

	return []openaigo.ChatMessage{}
}

func ask(lastAnswer string) openaigo.ChatCompletionRequestBody {

	reader := bufio.NewReader(os.Stdin)
	question, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	}

	if strings.EqualFold(question, "重置") {
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
		Model:    "gpt-3.5-turbo",
		Messages: message,
	}

	return request

}
