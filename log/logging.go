package log

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Logger *logrus.Logger

func Init() {
	fileHooker := NewFileRotateHooker("./log/", 86400)

	Logger = logrus.New()
	Logger.Hooks.Add(fileHooker)
	Logger.Out = os.Stdout
	Logger.Formatter = &logrus.TextFormatter{
		FullTimestamp: true,
	}
	Logger.Level = logrus.InfoLevel
}
