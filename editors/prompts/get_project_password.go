package prompts

import "github.com/manifoldco/promptui"

func GetProjectPassword() string {
	prompt := promptui.Prompt{Label: "Password", HideEntered: true}
	result, err := prompt.Run()
	if err != nil {
		panic("Can't get password")
	}
	return result
}
