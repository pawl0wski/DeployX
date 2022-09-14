package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pawl0wski/DeployX/endpoints"
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
	if !models.DoesConfigExist() {
		runConfig(cmd, args)
	}
	startHttpServer()
}

func startHttpServer() {
	engine := gin.Default()
	endpoints.InitializeEndpoints(engine)
	config := models.Config{}
	config.GetFromDatabaseOrCreate()
	err := engine.Run(fmt.Sprintf(":%d", config.ServerPort))
	if err != nil {
		panic("Can't start HTTP Server")
	}
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
