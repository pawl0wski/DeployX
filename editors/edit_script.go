package editors

import (
	"fmt"
	"github.com/pawl0wski/DeployX/models"
	"github.com/pawl0wski/DeployX/temp"
	"os"
	"os/exec"
)

func EditScript(script *models.Script, textEditor string) {
	file := temp.FileCreator(script.Content)
	textEditorCommand := prepareTextEditorToEditScript(textEditor, file)
	err := textEditorCommand.Run()
	if err != nil {
		panic(fmt.Sprintf("Can't execute %s", textEditor))
	}
	script.Content = temp.GetContentFromTempFile(file)
}

func prepareTextEditorToEditScript(textEditor string, file *os.File) *exec.Cmd {
	textEditorCommand := exec.Command(textEditor, file.Name())
	textEditorCommand.Stdin = os.Stdin
	textEditorCommand.Stdout = os.Stdout
	return textEditorCommand
}
