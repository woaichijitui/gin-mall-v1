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
	"gopkg.in/mail.v2"
	"mime/multipart"
	"strings"
	"time"
)

type UserService struct {
	NikeName string `json:"nike_name" form:"nike_name"`
	UserName string `json:"user_name" form:"user_name"`
	Password string `json:"password" form:"password"`
	//Money    string `json:"Money" form:"money"`
	//Avater   string `json:"avater" form:"avater"`
	Key string `json:"key" form:"key"` //前端验证

}

type SendEmailService struct {
	Email         string `json:"email" form:"email"`
	Password      string `json:"password" form:"password"`
	OperationType uint   `json:"operation_type" form:"operation_type"` //邮件类型
	//1、绑定邮箱 2、发送邮件 3、改密码
}
type ValidEmailService struct {
}
type ShowMoneyService struct {
	Key string `json:"key" form:"key"`
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

// 上传头像
func (service *UserService) Post(ctx context.Context, uId uint, file multipart.File, fileSize int64) serializer.Response {
	var code = e.Success
	var user *model.User
	var err error
	userDao := dao.NewUserDao(ctx)

	//查询user
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

// 发送邮件
func (sendEmailService *SendEmailService) Send(ctx context.Context, uId uint) serializer.Response {
	var notice *model.Notice //绑定邮箱，修改密码，模板通知
	var addr string
	code := e.Success

	noticeDao := dao.NewNoticDao(ctx)

	//获取emailtoken
	emailtoken, err := util.GenerateEmailToken(uId, sendEmailService.OperationType, sendEmailService.Email, sendEmailService.Password)
	if err != nil {
		code = e.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	//获取notic
	notice, err = noticeDao.GetNoticeByUId(sendEmailService.OperationType)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	addr = conf.ValidEmail + emailtoken
	mailStr := notice.Text
	mailTex := strings.Replace(mailStr, "email", addr, -1)

	//新邮件
	m := mail.NewMessage()
	m.SetHeader("From", conf.SmtpEmail)
	m.SetHeader("To", sendEmailService.Email)
	m.SetHeader("Subject", "htt")
	m.SetBody("text/html", mailTex)
	d := mail.NewDialer(conf.SmtpHost, 465, conf.SmtpEmail, conf.SmtpPass)
	d.StartTLSPolicy = mail.MandatoryStartTLS

	if err = d.DialAndSend(m); err != nil {
		code = e.ErrorSendEmail
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

// 绑定和解绑邮箱
func (service *ValidEmailService) Valid(ctx context.Context, emailToken string) serializer.Response {

	var userID uint
	var email string
	var password string
	var operationType uint
	var err error

	code := e.Success

	//验证token
	if emailToken == "" {
		code = e.InvalidParams
	} else {
		claims, err := util.ParseEmailToken(emailToken)
		if err != nil {
			code = e.ErrorAuthToken
		} else if claims.ExpiresAt.Unix() < time.Now().Unix() {
			code = e.ErrorAuthCheckTokenTimeOut
		} else { //验证正确 将token解析
			userID = claims.UserID
			email = claims.Email
			password = claims.Password
			operationType = claims.OperationType
		}

	}
	if code != e.Success {
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	//将mail绑定进入user
	var user *model.User
	userDao := dao.NewUserDao(ctx)

	user, err = userDao.GetUserByUId(userID)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	if operationType == 1 {
		//绑定邮箱
		user.Email = email
	} else if operationType == 2 {
		//解绑邮箱
		user.Email = ""
	} else if operationType == 3 {
		//修改密码
		err := user.Setpassword(password)
		if err != nil {
			code = e.Error
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	}

	//修改user参数
	err = userDao.UpdateUserById(userID, user)
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

// 展示用户金额
func (service *ShowMoneyService) Show(ctx context.Context, uId uint) serializer.Response {
	code := e.Success
	userDao := dao.NewUserDao(ctx)

	//查询用户
	user, err := userDao.GetUserByUId(uId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 返回带金额信息
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildMoney(user, service.Key),
	}

}
