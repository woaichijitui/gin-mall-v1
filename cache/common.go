package cache

import (
	"context"
	"fmt"
	"gin-mall/pkg/util"
	"github.com/redis/go-redis/v9"
	"gopkg.in/ini.v1"
	"strconv"
)

var (
	RedisClient *redis.Client
	RedisDb     string
	RedisAddr   string
	RedisPw     string //??
	RedisDbName string

	RedisCtx = context.Background() //需要的环境
)

func Init() {
	//1、加载配置文件
	file, err := ini.Load("conf/config.ini")

	if err != nil {
		fmt.Println("redis config file init err:", err)
	}

	//2、读取配置信息
	LoadRedis(file)

	//2、获取client
	Redis()
}

func Redis() {
	db, _ := strconv.ParseUint(RedisDbName, 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr:     RedisAddr,
		DB:       int(db),
		Password: RedisPw,
	})
	_, err := client.Ping(RedisCtx).Result()

	if err != nil {
		util.LogrusObj.Debug(err)
		panic(err)
	}

	RedisClient = client
}

func LoadRedis(file *ini.File) {
	RedisDb = file.Section("redis").Key("RedisDb").String()
	RedisAddr = file.Section("redis").Key("RedisAddr").String()
	RedisPw = file.Section("redis").Key("RedisPw").String()
	RedisDbName = file.Section("redis").Key("RedisDbName").String()
}
