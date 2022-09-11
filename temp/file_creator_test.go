package temp_test

import (
	"DeployX/temp"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestScriptFileCreator(t *testing.T) {
	testContent := "TestContent"

	tempScriptFile := temp.FileCreator(testContent)
	defer removeTestFile(tempScriptFile)

	assert.Equal(t, testContent, readContentFromFile(tempScriptFile))
}

func readContentFromFile(file *os.File) string {
	byteContent, err := os.ReadFile(file.Name())
	if err != nil {
		panic("Can't read content from file")
	}
	return string(byteContent)
}

func removeTestFile(file *os.File) {
	fileName := file.Name()
	err := file.Close()
	if err != nil {
		panic(fmt.Sprintf("Can't close %s file", fileName))
	}
	err = os.Remove(fileName)
	if err != nil {
		panic(fmt.Sprintf("Can't remove %s file", fileName))
	}
}
