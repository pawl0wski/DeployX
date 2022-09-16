package endpoints

import "github.com/gin-gonic/gin"

func InitializeEndpoints(engine *gin.Engine) {
	engine.GET("/version", GetVersion)
	engine.POST("/project_snapshot", PostProjectSnapshot)
}
