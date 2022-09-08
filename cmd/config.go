package cmd

import (
	"DeployX/models"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"strconv"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure DeployX",
	Run:   runConfig,
}

func runConfig(cmd *cobra.Command, args []string) {
	color.Blue("Welcome in DeployX configuration! 🛠")
	// Get Config from database
	config := models.Config{}
	config.GetFromDatabaseOrCreate()
	config.TextEditor = selectTextEditor([]string{"vi", "nano", "code", "gedit", "kate"})
	config.ServerPort = getServerPort()
	config.Save()
}

func selectTextEditor(choices []string) string {
	prompt := promptui.Select{Label: "Select your favorite text editor 🗒", Items: choices}
	_, result, err := prompt.Run()
	if err != nil {
		panic("Can't get your favorite text editor")
	}
	return result
}

func getServerPort() uint16 {
	prompt := promptui.Prompt{Label: "Server port", Default: "7777"}
	result, err := prompt.Run()
	if err != nil {
		panic("Can't get server port")
	}
	port, err := strconv.ParseUint(result, 10, 32)
	if err != nil {
		panic("Invalid port")
	}
	return uint16(port)
}

func init() {
	rootCmd.AddCommand(configCmd)
}
