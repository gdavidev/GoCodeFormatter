package metadata

type Line struct {
	Content  string
	Indent   int
	LineType string
	Scope    []string
	Index    int
}

func (line *Line) Length() int {
	return len(line.Content)
}
