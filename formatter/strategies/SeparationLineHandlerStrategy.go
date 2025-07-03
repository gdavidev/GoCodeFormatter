package strategies

import (
	"code-formatter/formatter/metadata"
	"strings"
)

type SeparationLineHandlerStrategy struct{}

func NewSeparationLineHandlerStrategy(config string) *SeparationLineHandlerStrategy {
	return new(SeparationLineHandlerStrategy)
}

func (*SeparationLineHandlerStrategy) Apply(line metadata.Line) metadata.Line {
	var lineLength = len(line.Content)
	if lineLength == 0 {
		return line
	}

	var lastChar = string(line.Content[len(line.Content)-1])
	if strings.ContainsAny(lastChar, "};{") {
		return line
	}
	line.Content += ";"
	return line
}
