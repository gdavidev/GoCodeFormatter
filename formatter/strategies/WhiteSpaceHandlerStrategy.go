package strategies

import (
	"code-formatter/formatter/metadata"
	"strings"
)

type WhiteSpaceHandlerStrategy struct{}

func NewWhiteSpaceHandlerStrategy(config string) *WhiteSpaceHandlerStrategy {
	return new(WhiteSpaceHandlerStrategy)
}

func (strategy *WhiteSpaceHandlerStrategy) Apply(line *metadata.Line) {
	if line.Length() == 0 {
		return
	}

	var finalLine = line
	strategy.removeTrailingWhiteSpace(finalLine)
	strategy.removeRedundantWhiteSpace(finalLine)
}

func (*WhiteSpaceHandlerStrategy) removeTrailingWhiteSpace(line *metadata.Line) {
	line.Content = strings.TrimRight(line.Content, " \t")
}

func (*WhiteSpaceHandlerStrategy) removeRedundantWhiteSpace(line *metadata.Line) {
	line.Content = strings.Join(strings.Fields(line.Content), " ")
}
