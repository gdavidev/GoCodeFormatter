package formatter

import (
	"code-formatter/formatter/languages"
	"code-formatter/formatter/metadata"
	"code-formatter/utility"
)

type Formatter struct {
	unformattedLines []string
	formattedLines   []string
}

func (formatter *Formatter) Format(path string) ([]string, error) {
	fileHandler, err := utility.NewFileHandler(path)
	if err != nil {
		panic(err)
	}
	defer fileHandler.Close()

	lines, err := fileHandler.ListLines()
	if err != nil {
		panic(err)
	}
	formatter.unformattedLines = lines
	var lineCollection = metadata.NewLineCollection(formatter.unformattedLines)

	var languageFormatter = languages.NewJavaScriptFormatter()
	for _, line := range lineCollection.Lines {
		languageFormatter.Format(&line)
		formatter.formattedLines = append(formatter.formattedLines, line.Content)
	}

	return formatter.formattedLines, nil
}
