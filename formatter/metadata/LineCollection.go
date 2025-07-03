package metadata

type LineCollection struct {
	CurrentLine  *Line
	CurrentIndex int
	Lines        []Line
}

func NewLineCollection(lines []string) *LineCollection {
	var collection = new(LineCollection)
	collection.CurrentIndex = 0

	for i, line := range lines {
		var lineObject = new(Line)
		lineObject.Index = i
		lineObject.Content = line

		collection.Lines = append(collection.Lines, *lineObject)
	}

	return collection
}

func (collection *LineCollection) GetNext() *Line {
	return collection.GetLine(collection.CurrentIndex + 1)
}

func (collection *LineCollection) GetPrevious() *Line {
	return collection.GetLine(collection.CurrentIndex - 1)
}

func (collection *LineCollection) GetLine(index int) *Line {
	if collection == nil || collection.Lines == nil {
		return nil
	}
	if index < 0 || index >= len(collection.Lines) {
		return nil
	}
	return &collection.Lines[index]
}
