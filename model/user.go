package model

import (
	"gorm.io/gorm"
)
import "golang.org/x/crypto/bcrypt"

type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	Email          string
	PasswordDigest string
	NickName       string
	Status         string
	Avatar         string
	Money          string
}

const (
	PasswordCost        = 12       //密码加密难度
	Active       string = "active" //激活用户
)

// 密码加密
func (u *User) Setpassword(password string) (err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PasswordCost)

	if err != nil {
		return err
	}
	u.PasswordDigest = string(bytes)
	return nil

}

// 验证密码
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordDigest), []byte(password))

	return err == nil

}
