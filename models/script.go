package models

import "gorm.io/gorm"

type Script struct {
	gorm.Model
	Content string `json:"content"`
}
