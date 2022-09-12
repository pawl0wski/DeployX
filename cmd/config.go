package cmd

import (
	"github.com/pawl0wski/DeployX/editors"
	"github.com/pawl0wski/DeployX/models"
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure DeployX",
	Run:   runConfig,
}

func runConfig(cmd *cobra.Command, args []string) {
	// Get Config from database
	config := models.Config{}
	config.GetFromDatabaseOrCreate()
	editors.EditConfig(&config)
	config.Save()
}

func init() {
	rootCmd.AddCommand(configCmd)
}
