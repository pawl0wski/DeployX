package editors

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/pawl0wski/DeployX/models"
)

func EditScript(script *models.Script, textEditor string) {
	editor := &survey.Editor{Editor: textEditor, FileName: "*.sh", Default: script.Content, HideDefault: false, AppendDefault: true}
	err := survey.AskOne(editor, &script.Content)
	if err != nil {
		panic(err)
	}
}
