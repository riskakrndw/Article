package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title   string `json:"title" form:"title"`
	Slug    string `json:"slug" form:"slug"`
	Content string `json:"content" form:"content" gorm:"type:text"`

	//FK
	CategoryID uint `json:"category_id" form:"category_id"`
}
