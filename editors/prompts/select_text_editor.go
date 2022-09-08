package prompts

import (
	"github.com/manifoldco/promptui"
)

func SelectTextEditor() string {
	prompt := promptui.Select{Label: "Select your favorite text editor", Items: []string{"vi", "nano", "code", "gedit", "kate"}}
	_, result, err := prompt.Run()
	if err != nil {
		panic("Can't get your favorite text editor")
	}
	return result
}
