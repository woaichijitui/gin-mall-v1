package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	Email          string
	PasswordDigest string
	NickName       string
	Status         string
	Avadar         string
	Money          string
}
