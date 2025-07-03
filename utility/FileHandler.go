package utility

import (
	"bufio"
	"io"
	"os"
)

type FileHandler struct {
	fileName string
	file     *os.File
	scanner  bufio.Scanner
}

func NewFileHandler(targetFilePath string) (*FileHandler, error) {
	file, err := os.Open(targetFilePath)
	if err != nil {
		return new(FileHandler), err
	}

	var newFileReader = new(FileHandler)
	newFileReader.fileName = targetFilePath
	newFileReader.file = file
	newFileReader.scanner = *bufio.NewScanner(file)

	return newFileReader, nil
}

func (f *FileHandler) ListLines() ([]string, error) {
	var lines []string
	for {
		line, err := f.nextLine()

		if err != nil && err != io.EOF {
			return []string{}, err
		} else if err == io.EOF {
			break
		}

		lines = append(lines, line)
	}
	return lines, nil
}

func (f *FileHandler) nextLine() (string, error) {
	if !f.scanner.Scan() {
		if err := f.scanner.Err(); err != nil {
			return "", err
		}
		return "", io.EOF
	}

	line := f.scanner.Text() // Get the current line as a string
	return line, nil
}

func (f *FileHandler) Close() {
	if err := f.file.Close(); err != nil {
		panic(err)
	}
}
