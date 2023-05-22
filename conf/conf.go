package conf

import (
	"fmt"
	"gin-mall/dao"
	"gopkg.in/ini.v1"
)

var (
	AppMode    string
	HttpPort   string
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
)

func Init() {
	//1、加载配置文件
	file, err := ini.Load("conf/config.ini")
	if err != nil {
		fmt.Println("config file init err:", err)
	}

	//2、读取配置信息
	LoadService(file)
	LoadMysql(file)

	//3、拼接dsn、mysql初始化
	connRead := DbUser + ":" + DbPassWord + "@tcp(" + DbHost + ":" + DbPort + ")/" + DbName + "?charset=utf8&parseTime=true"
	connWrite := DbUser + ":" + DbPassWord + "@tcp(" + DbHost + ":" + DbPort + ")/" + DbName + "?charset=utf8&parseTime=true"
	dao.Database(connRead, connWrite)
}

func LoadService(file *ini.File) {
	AppMode = file.Section("service").Key("AppMode").String()
	HttpPort = file.Section("service").Key("HttpPort").String()

}
func LoadMysql(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
	DbName = file.Section("mysql").Key("DbName").String()
}
