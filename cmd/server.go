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
	fmt.Println("Server is running")
	startHttpServer()
}

func startHttpServer() {
	config := models.Config{}
	config.GetFromDatabaseOrCreate()
	setGinMode(&config)
	engine := gin.Default()
	endpoints.InitializeEndpoints(engine)
	err := engine.Run(fmt.Sprintf(":%d", config.ServerPort))
	if err != nil {
		panic("Can't start HTTP Server")
	}
}

func setGinMode(config *models.Config) {
	if config.DebugMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
