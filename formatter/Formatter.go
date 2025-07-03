package formatter

import (
	"code-formatter/utility"
	"strings"
)

type Formatter struct {
	unformattedLines []string
	formattedLines   []string
}

func (formatter *Formatter) Format(path string) ([]string, error) {
	fileHandler, err := utility.NewFileHandler("./test/samples/index.js")
	if err != nil {
		panic(err)
	}
	defer fileHandler.Close()

	lines, err := fileHandler.ListLines()
	if err != nil {
		panic(err)
	}
	formatter.unformattedLines = lines

	for i, line := range formatter.unformattedLines {
		if line == "" && (i > 0 && formatter.unformattedLines[i-1] == "") {
			continue
		}

		var finalLine = line
		finalLine = strings.TrimRight(line, " \t")

		if formatter.shouldAddSemiColon(finalLine) {
			finalLine += ";"
		}

		formatter.formattedLines = append(formatter.formattedLines, finalLine)
	}

	return formatter.formattedLines, nil
}

func (formatter *Formatter) shouldAddSemiColon(line string) bool {
	var lineLength = len(line)
	if lineLength == 0 {
		return false
	}

	var lastChar = string(line[len(line)-1])
	return !strings.ContainsAny(lastChar, "};{")
}
