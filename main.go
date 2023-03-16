package main

import (
	"flag"
)

var f = flag.Bool("s", true, "是否要 stream 方式输出")

func main() {

	flag.Parse()
	stream("")
}
