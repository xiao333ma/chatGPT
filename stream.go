package main

import (
	"fmt"
	"github.com/otiai10/openaigo"
	"io"
)

func stream(question string) {

	request := ask(question)
	request.Stream = true

	lastAnswer := ""

	events := make(chan openaigo.ChatCompletionStreamInfo)
	client.ChatStream(ctx, request, events)

	for event := range events {
		if event.Err == nil {
			if len(event.Rsp.Choices) > 0 {
				res := event.Rsp.Choices[0].Delta.Content
				lastAnswer += res
				fmt.Printf("%s", res)
			}
		} else {
			if event.Err != io.EOF {
				fmt.Println(event.Err)
			}
			fmt.Println('\n')
		}
	}
	stream(lastAnswer)
}
