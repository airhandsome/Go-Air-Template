package log

import (
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

// 日志文件路径
var logFilePath = "./log/app.log"

// InitLogger 初始化日志配置
func InitLogger() {
	// 创建日志文件
	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatalf("Failed to open log file: %v", err)
	}

	// 设置日志输出到文件
	logrus.SetOutput(file)

	// 设置日志格式为 JSON
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// 设置日志级别
	logrus.SetLevel(logrus.InfoLevel)
}

// GetLogFilePath 获取日志文件绝对路径
func GetLogFilePath() string {
	absPath, _ := filepath.Abs(logFilePath)
	return absPath
}
