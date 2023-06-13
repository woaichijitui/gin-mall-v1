package serializer

import (
	"gin-mall/model"
	"gin-mall/pkg/util"
)

type Money struct {
	UserID   uint   `json:"user_id" form:"user_id"'`
	UserName string `json:"user_name" form:"user_name"`
	Money    string `json:"money" form:"money"`
}

func BuildMoney(user *model.User, key string) Money {
	util.Encrypt.SetKey(key)
	return Money{
		UserID:   user.ID,
		UserName: user.UserName,
		Money:    util.Encrypt.AesDecoding(user.Money),
	}
}
