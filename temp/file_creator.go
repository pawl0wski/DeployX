package temp

import (
	"os"
)

func FileCreator(content string) *os.File {
	tempFile := createTempFile()
	writeContentToTempFile(tempFile, content)
	return tempFile
}

func createTempFile() *os.File {
	f, err := os.CreateTemp("", "")
	if err != nil {
		panic("Can't create temp file")
	}
	return f
}

func writeContentToTempFile(f *os.File, content string) {
	_, err := f.WriteString(content)
	if err != nil {
		panic("Can't write content to temp file")
	}
}
