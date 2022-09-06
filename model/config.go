package model

import (
	"gorm.io/gorm"
	"time"
)

type Config struct {
	gorm.Model
	TextEditor         string `json:"string_editor"`
	ServerPort         uint   `json:"server_port"`
	FirstOpenTimestamp int64  `json:"first_open_timestamp"`
}

func (c *Config) SetFirstOpenToNow() {
	currentUnixTimestamp := time.Now().Unix()
	c.FirstOpenTimestamp = currentUnixTimestamp
}
