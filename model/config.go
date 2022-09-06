package model

import (
	"gorm.io/gorm"
	"time"
)

type Config struct {
	gorm.Model
	TextEditor                  string `json:"string_editor"`
	ServerPort                  uint   `json:"server_port"`
	FirstConfigurationTimestamp int64  `json:"first_configuration_timestamp"`
}

func (c *Config) SetFirstConfigurationToNow() {
	currentUnixTimestamp := time.Now().Unix()
	c.FirstConfigurationTimestamp = currentUnixTimestamp
}
