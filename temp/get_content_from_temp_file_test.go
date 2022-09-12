package temp_test

import (
	"fmt"
	"github.com/pawl0wski/DeployX/temp"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestUpdateScriptWithFileData(t *testing.T) {
	testContent := "TestContent"
	testFile := createTestFileAndWriteTestData(testContent)
	defer temp.RemoveFile(testFile)

	contentFromFunction := temp.GetContentFromTempFile(testFile)

	assert.FileExists(t, testFile.Name())
	assert.Equal(t, testContent, contentFromFunction)
}

func createTestFileAndWriteTestData(testData string) *os.File {
	tempFile, err := os.CreateTemp("", "testFile")
	if err != nil {
		panic("Can't create test file")
	}
	_, err = tempFile.Write([]byte(testData))
	if err != nil {
		panic(fmt.Sprintf("Can't write test data to %s\n", tempFile.Name()))
	}
	return tempFile
}
