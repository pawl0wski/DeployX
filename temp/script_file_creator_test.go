package temp_test

import (
	"DeployX/models"
	"DeployX/temp"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestScriptFileCreator(t *testing.T) {
	testContent := "TestContent"
	testScript := models.Script{Content: testContent}

	tempScriptFile := temp.ScriptFileCreator(&testScript)

	assert.Equal(t, testContent, readContentFromFile(tempScriptFile))
}

func readContentFromFile(file *os.File) string {
	byteContent, err := os.ReadFile(file.Name())
	if err != nil {
		panic("Can't read content from file")
	}
	return string(byteContent)
}
