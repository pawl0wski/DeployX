package temp_test

import (
	"DeployX/temp"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestUpdateScriptWithFileData(t *testing.T) {
	testContent := "TestContent"
	testFile := createTestFileAndWriteTestData(testContent)
	defer removeTestFile(testFile)

	contentFromFunction := temp.GetContentFromTempFile(testFile)

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
