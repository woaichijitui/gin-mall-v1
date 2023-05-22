package model

import "gorm.io/gorm"

type Favorite struct {
	gorm.Model
	User      User    `gorm:"Foreignkey: UserId"` //外键关联是需要关联的模型
	UserId    uint    `gorm:"not null"`
	Product   Product `gorm:"Foreignkey: ProductId"`
	ProductId uint    `gorm:"not null"`
	Boss      User    `gorm:"Foreignkey: BossId"`
	BossId    uint    `gorm:"not null"`
}
