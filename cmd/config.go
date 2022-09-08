package cmd

import (
	"DeployX/models"
	"errors"
	"fmt"
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
	validator := func(input string) error {
		_, err := strconv.ParseUint(input, 10, 16)
		if err != nil {
			return errors.New(fmt.Sprintf("%s is not valid port", input))
		}
		return nil
	}
	prompt := promptui.Prompt{Label: "Server port", Default: "7777", Validate: validator}
	result, err := prompt.Run()
	if err != nil {
		panic("Can't get server port")
	}
	port, err := strconv.ParseUint(result, 10, 16)
	if err != nil {
		panic("Invalid port")
	}
	return uint16(port)
}

func init() {
	rootCmd.AddCommand(configCmd)
}
