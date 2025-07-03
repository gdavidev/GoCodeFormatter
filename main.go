package main

import (
	"code-formatter/formatter"
	"fmt"
	"strings"
)

func main() {
	formatter := new(formatter.Formatter)
	final, err := formatter.Format("./test/samples/badCodeExample.js")
	if err != nil {
		panic(err)
	}

	for i, line := range final {
		var lineToPrint = line
		lineToPrint = strings.ReplaceAll(lineToPrint, " ", "•")
		lineToPrint = strings.ReplaceAll(lineToPrint, "\n", "↳")
		lineToPrint = strings.ReplaceAll(lineToPrint, "\t", "→")
		fmt.Printf("Index %d: %s\n", i, lineToPrint)
	}
}
