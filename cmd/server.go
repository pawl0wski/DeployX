package cmd

import (
	"github.com/pawl0wski/DeployX/models"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run DeployX server",
	Run:   runServer,
}

func runServer(cmd *cobra.Command, args []string) {
	if models.DoesConfigExist() {
		runConfig(cmd, args)
	}
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
