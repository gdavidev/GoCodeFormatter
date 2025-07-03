package main

import (
	"code-formatter/formatter"
	"fmt"
)

func main() {
	formatter := *new(formatter.Formatter)
	final, err := formatter.Format("./test/samples/index.js")
	if err != nil {
		panic(err)
	}

	for i, line := range final {
		fmt.Printf("Index %d: %s\n", i, line)
	}
}
