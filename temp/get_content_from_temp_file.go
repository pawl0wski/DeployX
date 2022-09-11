package temp

import (
	"os"
)

func GetContentFromTempFile(file *os.File) string {
	return getFileContentAsString(file)
}

func getFileContentAsString(file *os.File) string {
	byteContent, err := os.ReadFile(file.Name())
	if err != nil {
		panic("Can't read content from file")
	}
	return string(byteContent)
}
