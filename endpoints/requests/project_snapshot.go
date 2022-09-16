package requests

import "encoding/base64"

type ProjectSnapshotRequest struct {
	ProjectID       int    `json:"project_id"`
	ProjectPassword string `json:"project_password"`
	Version         string `json:"snapshot_version"`
	Data            string `json:"data"`
}

func (request *ProjectSnapshotRequest) ConvertDataToBinary() []byte {
	encoder := base64.Encoding{}
	binaryData, err := encoder.DecodeString(request.Data)
	if err != nil {
		panic("Can't covert data to binary")
	}
	return binaryData
}
