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

	RedisDb     string
	RedisAddr   string
	RedisPw     string
	RedisDbName string

	AccessKey        string
	SerectKey        string
	Bucket           string
	QiniuServer      string
	ValidEmail       string
	SmtpHost         string
	SmtpEmail        string
	SmtpPass         string
	Host             string
	ProductPath      string
	AvatarPath       string
	EsHost           string
	EsPort           string
	EsIndex          string
	RabbitMQ         string
	RabbitMQUser     string
	RabbitMQPassWord string
	RabbitMQHost     string
	RabbitMQPort     string
)

func Init() {
	//1、加载配置文件
	file, err := ini.Load("conf/config.ini")
	if err != nil {
		fmt.Println("config file init err:", err)
	}

	//2、读取配置信息
	LoadRedis(file)
	LoadService(file)
	LoadMysql(file)
	LoadQiniu(file)
	LoadPath(file)
	LoadRabbitmq(file)
	LoadEmail(file)
	LoadEs(file)

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
func LoadRedis(file *ini.File) {
	RedisDb = file.Section("redis").Key("RedisDb").String()
	RedisAddr = file.Section("redis").Key("RedisAddr").String()
	RedisPw = file.Section("redis").Key("RedisPw").String()
	RedisDbName = file.Section("redis").Key("RedisDbName").String()
}
func LoadQiniu(file *ini.File) {
	AccessKey = file.Section("qiniu").Key("AccessKey").String()
	SerectKey = file.Section("qiniu").Key("SerectKey").String()
	Bucket = file.Section("qiniu").Key("Bucket").String()
	QiniuServer = file.Section("qiniu").Key("QiniuServer").String()
}
func LoadEmail(file *ini.File) {
	ValidEmail = file.Section("email").Key("ValidEmail").String()
	SmtpHost = file.Section("email").Key("SmtpHost").String()
	SmtpEmail = file.Section("email").Key("SmtpEmail").String()
	SmtpPass = file.Section("email").Key("SmtpPass").String()
}
func LoadPath(file *ini.File) {
	Host = file.Section("path").Key("Host").String()
	ProductPath = file.Section("path").Key("ProductPath").String()
	AvatarPath = file.Section("path").Key("AvatarPath").String()
}
func LoadEs(file *ini.File) {
	EsHost = file.Section("es").Key("EsHost").String()
	EsPort = file.Section("es").Key("EsPort").String()
	EsIndex = file.Section("es").Key("EsIndex").String()
}
func LoadRabbitmq(file *ini.File) {
	RabbitMQ = file.Section("rabbitmq").Key("RabbitMQ").String()
	RabbitMQUser = file.Section("rabbitmq").Key("RabbitMQUser").String()
	RabbitMQPassWord = file.Section("rabbitmq").Key("RabbitMQPassWord").String()
	RabbitMQHost = file.Section("rabbitmq").Key("RabbitMQHost").String()
	RabbitMQPort = file.Section("rabbitmq").Key("RabbitMQPort").String()
}
