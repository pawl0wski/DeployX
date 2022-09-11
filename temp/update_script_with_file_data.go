package temp

import (
	"DeployX/models"
	"os"
)

func UpdateScriptWithFileData(script *models.Script, file *os.File) {
	script.Content = getFileContentAsString(file)
}

func getFileContentAsString(file *os.File) string {
	byteContent, err := os.ReadFile(file.Name())
	if err != nil {
		panic("Can't read content from file")
	}
	return string(byteContent)
}
