package config

import "github.com/sirupsen/logrus"

func InitLog() {
	// 初始化logrus
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.DebugLevel)
	
}
