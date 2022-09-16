package requests

type DeployRequest struct {
	ProjectID       int    `json:"project_id"`
	ProjectPassword string `json:"project_password"`
	Data            string `json:"data"`
}
