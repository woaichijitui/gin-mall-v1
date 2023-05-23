package api

import (
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
		ctx.JSON(http.StatusBadRequest, err)
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
		ctx.JSON(http.StatusBadRequest, err)
	}

}

// 注册api
func UserUpdate(ctx *gin.Context) {
	var userUpdate service.UserService

	claims, _ := util.ParseToken(ctx.GetHeader("authorization"))

	//绑定UserService
	if err := ctx.ShouldBind(&userUpdate); err == nil {
		res := userUpdate.Update(ctx.Request.Context(), claims.ID)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, err)
	}
}

// 更新头像
// 更新头像上传.jpg 图片
func UpdateAvatar(ctx *gin.Context) {

	//接收图片信息
	file, fileHeader, _ := ctx.Request.FormFile("file")
	fileSize := fileHeader.Size

	var updateAvatar service.UserService
	claims, _ := util.ParseToken(ctx.GetHeader("authorization"))

	//绑定UserService
	if err := ctx.ShouldBind(&updateAvatar); err == nil {
		res := updateAvatar.Post(ctx.Request.Context(), claims.ID, file, fileSize)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, err)
	}
}
