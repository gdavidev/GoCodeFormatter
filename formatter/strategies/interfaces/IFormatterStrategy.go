package interfaces

import "code-formatter/formatter/metadata"

type IFormatterStrategy interface {
	Apply(*metadata.Line)
}
