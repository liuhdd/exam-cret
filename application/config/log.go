package config

import "github.com/sirupsen/logrus"

func init() {
	InitLog()
}
func InitLog() {
	// 初始化logrus
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		ForceColors:     true,
	})
	logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.DebugLevel)

}
