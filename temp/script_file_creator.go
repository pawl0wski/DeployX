package temp

import (
	"DeployX/models"
	"os"
)

func ScriptFileCreator(script *models.Script) *os.File {
	tempScriptFile := createTempScriptFile()
	writeScriptContentToTempFile(tempScriptFile, script)
	return tempScriptFile
}

func createTempScriptFile() *os.File {
	f, err := os.CreateTemp("", "tempScript")
	if err != nil {
		panic("Can't create temp script file")
	}
	return f
}

func writeScriptContentToTempFile(f *os.File, script *models.Script) {
	_, err := f.WriteString(script.Content)
	if err != nil {
		panic("Can't write script content to temp file")
	}
}
