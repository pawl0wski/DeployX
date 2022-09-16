package responses

import "time"

type DeployResponse struct {
	SnapshotVersion    int       `json:"snapshot_version"`
	WhenWillBeDeployed time.Time `json:"when_will_be_deployed"`
}
