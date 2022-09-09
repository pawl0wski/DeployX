package prompts

import (
	"github.com/manifoldco/promptui"
)

func SelectDebugMode() bool {
	prompt := promptui.Select{Label: "Enable debug mode", Items: []string{"no", "yes"}}
	_, choice, err := prompt.Run()
	if err != nil {
		panic("Can't get debug mode")
	}
	return choice == "yes"
}
