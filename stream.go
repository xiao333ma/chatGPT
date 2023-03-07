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
	client.ChatStream(ctx, request, func(response openaigo.ChatCompletionStreamResponse, err error) {
		if err == nil {
			res := response.Choices[0].Delta.Content
			lastAnswer += res
			fmt.Printf("%s", res)
		} else {
			if err != io.EOF {
				fmt.Println(err)
			}
			fmt.Println("")
			stream(lastAnswer)
		}
	})
}
