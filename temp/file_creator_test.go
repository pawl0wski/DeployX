package temp_test

import (
	"DeployX/temp"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScriptFileCreator(t *testing.T) {
	testContent := "TestContent"

	tempFile := temp.FileCreator(testContent)
	defer temp.RemoveFile(tempFile)

	assert.FileExists(t, tempFile.Name())
	assert.Equal(t, testContent, temp.GetContentFromTempFile(tempFile))
}
