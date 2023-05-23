package service

import (
	"context"
	"fmt"
	"gin-mall/dao"
	"gin-mall/model"
	"gin-mall/pkg/e"
	"gin-mall/pkg/util"
	"gin-mall/serializer"
)

type UserService struct {
	NikeName string `json:"nike_name" form:"nike_name"`
	UserName string `json:"user_name" form:"user_name"`
	Password string `json:"password" form:"password"`
	//Money    string `json:"Money" form:"money"`
	//Avater   string `json:"avater" form:"avater"`
	Key string `json:"key" form:"key"` //前端验证

}

func (service *UserService) Register(c context.Context) serializer.Response {
	var user model.User
	code := e.Success

	//1、判断key 是否为16字节
	if service.Key == "" || len(service.Key) != 16 {
		code = e.InvalidParams
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "密钥长度不足",
		}
	}

	//??
	util.Encrypt.SetKey(service.Key)

	//2、创建一个连接
	userDao := dao.NewUserDao(c)

	//判断是否存在该用户
	_, exit, err := userDao.ExitOrNotByUserName(service.UserName)

	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	if exit {
		code = e.ErrorExitUser
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	//3、在数据库创建user账户
	user = model.User{
		UserName: service.UserName,
		NickName: service.NikeName,
		Status:   model.Active,
		Avatar:   "avater.JPG",
		Money:    util.Encrypt.AesEncoding("10000"),
	}
	//密码加密
	err = user.Setpassword(service.Password)
	if err != nil {
		code = e.ErrorFailEncry
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	//创建用户
	err = userDao.CreateUser(&user)
	fmt.Println(err)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	//4、创建成功返回resp 并加密密码
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
