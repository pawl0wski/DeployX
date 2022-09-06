package cmd

import (
	"DeployX/model"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure DeployX",
	Run:   runCommand,
}

func runCommand(cmd *cobra.Command, args []string) {
	color.Blue("Welcome in DeployX configuration! 🛠")
	// Create blank Config
	config := model.Config{}
	config.TextEditor = selectTextEditor([]string{"vi", "nano", "code"})
}

func selectTextEditor(choices []string) string {
	prompt := promptui.Select{Label: "Select your favorite text editor 🗒", Items: choices}
	_, result, err := prompt.Run()
	if err != nil {
		panic("Can't get your favorite text editor")
	}
	return result
}

func init() {
	rootCmd.AddCommand(configCmd)
}
