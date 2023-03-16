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

	client.ChatStream(ctx, request, func(info openaigo.ChatCompletionStreamInfo) {
		if info.Err == nil {
			if len(info.Rsp.Choices) > 0 {
				res := info.Rsp.Choices[0].Delta.Content
				lastAnswer += res
				fmt.Printf("%s", res)
			}
		} else {
			if info.Err != io.EOF && info.Err != openaigo.StreamFinish {
				fmt.Println(info.Err)
			}
			fmt.Println("\n")
		}
	})

	stream(lastAnswer)
}
