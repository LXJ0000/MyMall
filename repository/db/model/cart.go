package model

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserId    uint `gorm:"not null"`
	ProductId uint `gorm:"not null"`
	BossId    uint `gorm:"not null"` // 商家
	Num       uint `gorm:"not null"`
	MaxNum    uint `gorm:"not null"` // 限购
	Check     bool // 是否下单
}
