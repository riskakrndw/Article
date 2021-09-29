package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	CategoryName string `json:"category_name" form:"category_name"`
	CategorySlug string `json:"category_slug" form:"category_slug"`

	//1 to many
	Article []Article `gorm:"foreignKey:CategoryID;references:ID;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;"`
}
