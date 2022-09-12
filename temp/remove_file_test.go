package temp_test

import (
	"github.com/pawl0wski/DeployX/temp"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveFile(t *testing.T) {
	testFile := temp.FileCreator("")

	temp.RemoveFile(testFile)

	assert.NoFileExists(t, testFile.Name())
}
