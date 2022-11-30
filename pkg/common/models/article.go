package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model         // adds ID, created_at etc.
	Title       string `json:"title"`
	Content     string `json:"content"`
	Description string `json:"description"`
}
