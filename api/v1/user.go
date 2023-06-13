package api

import (
	"fmt"
	"gin-mall/pkg/util"
	"gin-mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRegister(ctx *gin.Context) {
	var userRegister service.UserService

	//绑定UserService
	if err := ctx.ShouldBind(&userRegister); err == nil {
		res := userRegister.Register(ctx.Request.Context())
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

// 注册api
func UserLogin(ctx *gin.Context) {
	var userLogin service.UserService

	//绑定UserService
	if err := ctx.ShouldBind(&userLogin); err == nil {
		res := userLogin.Login(ctx.Request.Context())
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}

}

// 更新图片
func UserUpdate(ctx *gin.Context) {
	var userUpdate service.UserService

	claims, _ := util.ParseToken(ctx.GetHeader("authorization"))

	//绑定UserService
	if err := ctx.ShouldBind(&userUpdate); err == nil {
		res := userUpdate.Update(ctx.Request.Context(), claims.ID)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

// 更新头像
// 更新头像上传.jpg 图片
func UpdateAvatar(ctx *gin.Context) {

	//接收图片信息
	file, fileHeader, err := ctx.Request.FormFile("file")
	if err != nil {
		fmt.Println("FormFile err:", err)
	}
	fileSize := fileHeader.Size

	var updateAvatar service.UserService
	claims, _ := util.ParseToken(ctx.GetHeader("authorization"))

	//绑定UserService
	if err := ctx.ShouldBind(&updateAvatar); err == nil {
		res := updateAvatar.Post(ctx.Request.Context(), claims.ID, file, fileSize)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

// 发送邮件
func SendEmail(ctx *gin.Context) {
	var sendEmail service.SendEmailService
	claims, _ := util.ParseToken(ctx.GetHeader("authorization"))

	//绑定UserService
	if err := ctx.ShouldBind(&sendEmail); err == nil {
		res := sendEmail.Send(ctx.Request.Context(), claims.ID)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

// 绑定邮箱
func ValidEmail(ctx *gin.Context) {
	var validEmail service.ValidEmailService

	//绑定UserService
	if err := ctx.ShouldBind(&validEmail); err == nil {
		res := validEmail.Valid(ctx.Request.Context(), ctx.GetHeader("authorization")) //token传入
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

// 查看金额
func ShowMoney(ctx *gin.Context) {
	var showMoney service.ShowMoneyService
	claims, _ := util.ParseToken(ctx.GetHeader("authorization"))
	//绑定UserService
	if err := ctx.ShouldBind(&showMoney); err == nil {
		res := showMoney.Show(ctx.Request.Context(), claims.ID) //token传入
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}
