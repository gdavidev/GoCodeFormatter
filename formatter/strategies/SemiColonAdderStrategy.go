package strategies

import (
	"code-formatter/formatter/metadata"
	"strings"
)

type SemiColonAdderStrategy struct{}

func NewSemiColonAdderStrategy(config string) *SemiColonAdderStrategy {
	return new(SemiColonAdderStrategy)
}

func (*SemiColonAdderStrategy) Apply(line *metadata.Line) {
	if line.Length() == 0 {
		return
	}

	var lastChar = string(line.Content[line.Length()-1])
	if strings.ContainsAny(lastChar, "};{") {
		return
	}
	line.Content += ";"
}
