package editors

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/pawl0wski/DeployX/models"
)

type ScriptEditor struct {
	Script     *models.Script
	TextEditor string
}

func (e *ScriptEditor) Run() {
	editor := &survey.Editor{Editor: e.TextEditor, FileName: "*.sh", Default: e.Script.Content, HideDefault: false, AppendDefault: true}
	err := survey.AskOne(editor, &e.Script.Content)
	if err != nil {
		panic(err)
	}
}
