package hasher_test

import (
	"DeployX/hasher"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashPassword(t *testing.T) {
	passwordToHash := "ABC"
	expectedOutput := "b5d4045c3f466fa91fe2cc6abe79232a1a57cdf104f7a26e716e0a1e2789df78"

	assert.Equal(t, expectedOutput, hasher.HashPassword(passwordToHash))
}
