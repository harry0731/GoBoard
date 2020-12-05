// models/article.go

package models

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	ID      int    `gorm: NOT NULL autoIncrement json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
