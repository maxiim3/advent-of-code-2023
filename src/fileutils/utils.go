package fileutils

import (
	"os"
	"strings"
)

type File struct {
	Raw   []byte
	Lines []string
}

func ParseFile(path string) File {
	binaryFile := readFile(path)

	lines := readLines(binaryFile)

	return File{Raw: binaryFile, Lines: lines}
}

func readFile(path string) []byte {
	file, err := os.ReadFile(path)

	if err != nil {
		panic("Cannot read the file ")
	}

	return file
}

func readLines(file []byte) []string {
	lines := strings.Split(string(file), "\n")

	return lines
}
