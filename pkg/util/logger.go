package util

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"time"
)

var LogrusObj *logrus.Logger

// log 初始化
func Init() {
	src, _ := SetOutPutFile()

	if LogrusObj != nil {
		LogrusObj.Out = src //日志输出路径
		return
	}
	logger := logrus.New()

	logger.Out = src
	logger.SetOutput(io.MultiWriter(os.Stdout, src))
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	LogrusObj = logger

}

// 设置日志文件
func SetOutPutFile() (*os.File, error) {
	now := time.Now()

	//查看是否存在日志文件路径
	logFilePath := ""
	if dir, err := os.Getwd(); err == nil {
		fmt.Println("----", dir)
		logFilePath = dir + "\\logs\\"
	}
	_, err := os.Stat(logFilePath)
	if os.IsNotExist(err) {
		if err = os.MkdirAll(logFilePath, 0777); err != nil {
			LogrusObj.Println("MkdirAll err", err.Error())
			return nil, err
		}
	}

	//分成每一天的日志文件
	logFileName := now.Format("2006-01-02") + ".log"
	//日志文件
	fileName := path.Join(logFilePath, logFileName)
	//判断文件是否存在
	if _, err := os.Stat(fileName); err != nil { //若不纯在 则会报错 *PathError
		if _, err := os.Create(fileName); err != nil {
			LogrusObj.Println(err.Error())
			return nil, err
		}
	}
	//写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		return nil, err
	}
	return src, nil
}
