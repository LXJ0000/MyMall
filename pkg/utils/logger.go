package util

import (
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"path"
	"time"
)

var LogrusObj *logrus.Logger

func init() {
	src, _ := setOutputFile()

	if LogrusObj != nil {
		LogrusObj.Out = src
		return
	}
	//	实例化
	logger := logrus.New()
	logger.Out = src                   // 输出
	logger.SetLevel(logrus.DebugLevel) // 日志级别
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	LogrusObj = logger
}

func setOutputFile() (*os.File, error) {
	now := time.Now()
	logFilePath := ""
	//获取工作目录
	//dir = /MyMall/
	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + "/logs/"
	}
	_, err := os.Stat(logFilePath)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(logFilePath, 0777); err != nil {
			log.Println(err.Error())
			return nil, err
		}
	}
	logFileName := now.Format("2006-01-02") + ".log"
	//	日志文件
	fileName := path.Join(logFilePath, logFileName)
	_, err = os.Stat(fileName)
	if os.IsNotExist(err); err != nil {
		if _, err = os.Create(fileName); err != nil {
			log.Println(err.Error())
			return nil, err
		}
	}
	//	写入文件
	//os.O_APPEND表示以追加模式打开文件
	//os.O_WRONLY表示以只写模式打开文件
	//os.ModeAppend表示以追加的方式写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return nil, err
	}
	return src, nil
}
