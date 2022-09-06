package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure DeployX",
	Run:   runCommand,
}

func runCommand(cmd *cobra.Command, args []string) {
	fmt.Println("Welcome in DeployX configuration! 🛠")
}

func init() {
	rootCmd.AddCommand(configCmd)
}
