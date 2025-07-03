package languages

import (
	"code-formatter/formatter/metadata"
	strategies "code-formatter/formatter/strategies"
	strategies_interfaces "code-formatter/formatter/strategies/interfaces"
)

type JavaScriptFormatter struct {
	strategiesList []strategies_interfaces.IFormatterStrategy
}

func NewJavaScriptFormatter() *JavaScriptFormatter {
	var formatter = new(JavaScriptFormatter)
	formatter.strategiesList = []strategies_interfaces.IFormatterStrategy{
		strategies.NewWhiteSpaceHandlerStrategy(""),
		strategies.NewSemiColonAdderStrategy(""),
	}

	return formatter
}

func (formatter *JavaScriptFormatter) Format(line *metadata.Line) {
	for _, strategy := range formatter.strategiesList {
		strategy.Apply(line)
	}
}
