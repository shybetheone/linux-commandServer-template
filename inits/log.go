package inits

import (
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

//func LogInit() {
//	// 获取日志文件句柄
//	// 已 只写入文件|没有时创建|文件尾部追加 的形式打开这个文件
//	logFile, err := os.OpenFile(`./logfile.log`, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
//	if err != nil {
//		panic(err)
//	}
//	// 设置存储位置
//	log.SetOutput(logFile)
//}

type CustomFormatter struct {
	logrus.TextFormatter
}

func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	entry.Time = time.Now()
	return f.TextFormatter.Format(entry)
}

func LogInit() {
	log := logrus.New()
	log.SetFormatter(&CustomFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.DebugLevel)
	log.SetFormatter(&logrus.JSONFormatter{})
	log.Debug("Debug message")
	log.Info("Info message")
	log.Warn("Warning message")
	log.Error("Error message")
}
