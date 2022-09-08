package editors

import "github.com/fatih/color"

var editorColor *color.Color

func init() {
	editorColor = color.New(color.FgWhite).Add(color.Bold)
}

func printWithEditorColor(text string) {
	_, err := editorColor.Println(text)
	if err != nil {
		panic("Error while using editorColor")
	}
}
