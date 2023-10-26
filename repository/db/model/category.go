package model

import "gorm.io/gorm"

// Category 分类
type Category struct {
	gorm.Model
	CategoryName string
}
