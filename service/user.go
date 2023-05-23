package service

import (
	"context"
	"fmt"
	"gin-mall/conf"
	"gin-mall/dao"
	"gin-mall/model"
	"gin-mall/pkg/e"
	"gin-mall/pkg/util"
	"gin-mall/serializer"
	"mime/multipart"
)

type UserService struct {
	NikeName string `json:"nike_name" form:"nike_name"`
	UserName string `json:"user_name" form:"user_name"`
	Password string `json:"password" form:"password"`
	//Money    string `json:"Money" form:"money"`
	//Avater   string `json:"avater" form:"avater"`
	Key string `json:"key" form:"key"` //前端验证

}

// 用户注册
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

// 用户登录
func (service *UserService) Login(ctx context.Context) serializer.Response {
	var user *model.User
	code := e.Success

	//查看用户是否存在
	//创建一个连接
	userDao := dao.NewUserDao(ctx)

	//查看用户是否存在
	user, exit, err := userDao.ExitOrNotByUserName(service.UserName)
	if err != nil || !exit {
		code = e.ErrorUserNotFount
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "用户不存在，检查用户名",
		}
	}

	// 验证密码

	if user.CheckPassword(service.Password) == false {
		code = e.ErrorNotCompare
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "密码错误,请重新输入",
		}
	}

	//签发token
	token, err := util.GenerateToken(user.ID, user.UserName, 0)
	if err != nil {
		code = e.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return serializer.Response{
		Status: code,
		Data:   serializer.TokenData{User: serializer.BuildUser(user), Token: token},
		Msg:    e.GetMsg(code),
	}

}

// 更改用户信息
func (service *UserService) Update(ctx context.Context, uId uint) serializer.Response {
	var user *model.User
	code := e.Success
	var err error

	userDao := dao.NewUserDao(ctx)

	//根据id查询用户
	user, err = userDao.GetUserByUId(uId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	//更改用户属性
	//更改Nike name
	if service.NikeName != "" {
		user.NickName = service.NikeName
	}
	err = userDao.UpdateUserById(uId, user)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildUser(user),
	}
}

func (service *UserService) Post(ctx context.Context, uId uint, file multipart.File, fileSize int64) serializer.Response {
	var code = e.Success
	var user *model.User
	var err error
	userDao := dao.NewUserDao(ctx)

	//
	user, err = userDao.GetUserByUId(uId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}

	}

	//保存图片到本地函数
	path, err := UploadAvatarToLocalStatic(file, uId, user.UserName)
	if err != nil {
		code = e.ErrorUploadFail
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	//更换数据库User avatar 字段的值
	user.Avatar = conf.Host + conf.HttpPort + conf.AvatarPath + path
	err = userDao.UpdateUserById(uId, user)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildUser(user),
	}
}
