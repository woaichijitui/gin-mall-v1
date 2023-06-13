package service

import (
	"gin-mall/conf"
	"io"
	"mime/multipart"
	"os"
	"strconv"
)

// 将文件保存在本地 返回路径
func UploadAvatarToLocalStatic(file multipart.File, uId uint, userName string) (path string, err error) {
	bId := strconv.Itoa(int(uId))

	//项目系统路径
	bassPath := "." + conf.AvatarPath + "user" + bId + "/"

	if !DirExistorNot(bassPath) {
		CreateDir(bassPath)
	}
	//以项目为基础的文件名
	avatarPath := bassPath + userName + ".jpg" // todo: 把file的后缀提取出来
	//将文件转换问byte类型
	bytes, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	//将文件保存在本地
	err = os.WriteFile(avatarPath, bytes, 0666)
	if err != nil {
		return "", err
	}
	return "user" + bId + "/" + userName + ".jpg", nil

}

// 将第一张商品图片保存在本地 返回路径
// 返回的路径是 conf.pruductPath 路径之后的路径
func uploadProductToLocalStatic(file multipart.File, uId uint, productName string) (path string, err error) {
	bId := strconv.Itoa(int(uId))

	//项目系统路径
	bassPath := "." + conf.ProductPath + "boss" + bId + "/"

	if !DirExistorNot(bassPath) {
		CreateDir(bassPath)
	}
	//以项目为基础的文件名
	productPath := bassPath + productName + ".jpg" // todo: 把file的后缀提取出来
	//将文件转换问byte类型
	bytes, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	//将文件保存在本地
	err = os.WriteFile(productPath, bytes, 0666)
	if err != nil {
		return "", err
	}
	return "boss" + bId + "/" + productName + ".jpg", nil

}

// 判断文件夹路径是否存在
func DirExistorNot(fileAddr string) bool {
	stat, err := os.Stat(fileAddr)
	if err != nil {
		return false
	}
	return stat.IsDir()
}

// 创建文件夹
func CreateDir(fileAddr string) bool {
	err := os.MkdirAll(fileAddr, 755)
	if err != nil {
		return false
	}
	return true
}
