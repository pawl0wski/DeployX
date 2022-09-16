package requests

type ProjectSnapshotRequest struct {
	ProjectID       string `json:"project_id"`
	ProjectPassword string `json:"project_password"`
	Version         string `json:"snapshot_version"`
	Data            string `json:"data"`
}
