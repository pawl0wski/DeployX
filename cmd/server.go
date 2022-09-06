package cmd

import (
	"DeployX/database"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run DeployX server",
	Run:   runServer,
}

func runServer(cmd *cobra.Command, args []string) {
	if !database.CheckIfConfigExist() {
		runConfig(cmd, args)
	}
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
