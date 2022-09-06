package database_test

import (
	"DeployX/database"
	"DeployX/initialize"
	"DeployX/model"
	"testing"
)

func TestGetConfig(t *testing.T) {
	initialize.Database(":memory:")
	testEditor := "TestTextEditor"
	testConfig := model.Config{TextEditor: testEditor}
	database.DBConn.Create(&testConfig)

	configFromDatabase := database.GetConfig()

	if configFromDatabase.TextEditor != testEditor {
		t.Fatal("GetConfig return invalid config")
	}
}
