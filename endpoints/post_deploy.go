package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/pawl0wski/DeployX/endpoints/requests"
	"github.com/pawl0wski/DeployX/hasher"
	"github.com/pawl0wski/DeployX/models"
)

func PostDeploy(c *gin.Context) {
	snapshotRequest := requests.DeployRequest{}
	err := c.ShouldBindJSON(&snapshotRequest)
	if err != nil {
		sendError(c, "An error occurred, please check your client version and try again.", 500)
		return
	}
	var passwordCorrect bool
	passwordCorrect, err = checkIfThePasswordInTheRequestIsCorrect(&snapshotRequest)
	if err != nil {
		sendError(c, "There is no such project.", 404)
		return
	}
	if !passwordCorrect {
		sendError(c, "Invalid password.", 401)
		return
	}
	snapshot := createProjectSnapshotFromDeploy(&snapshotRequest)
	snapshot.Save()
}

func sendError(c *gin.Context, errorContent string, errorCode int) {
	c.String(errorCode, errorContent)
}

func checkIfThePasswordInTheRequestIsCorrect(request *requests.DeployRequest) (bool, error) {
	project, err := models.GetProjectByID(request.ProjectID)
	if err != nil {
		return false, err
	}
	hashedRequestPassword := hasher.HashPassword(request.ProjectPassword)
	return hashedRequestPassword == project.Password, nil
}

func createProjectSnapshotFromDeploy(request *requests.DeployRequest) *models.ProjectSnapshot {
	lastVersion := models.GetLastSnapshotVersionByProjectID(request.ProjectID)
	snapshot := models.ProjectSnapshot{
		ProjectID:       request.ProjectID,
		Data:            request.Data,
		Version:         lastVersion + 1,
		CurrentSnapshot: false,
	}
	return &snapshot
}
