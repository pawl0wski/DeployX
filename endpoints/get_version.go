package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/pawl0wski/DeployX/endpoints/responses"
)

func GetVersion(c *gin.Context) {
	response := responses.VersionResponse{Version: "1.0.0", Name: "DeployX"}
	c.JSON(200, &response)
}
