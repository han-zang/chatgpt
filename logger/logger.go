package logger

import (
	"os"
	"sync"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger
var once sync.Once

func Test() {
	Init()
	log.Infof("This is an info log. %d", 1)
	log.Error("This is an error log.")
}

func Init() {
	once.Do(func() {
		if log == nil {
			log = logrus.New()
		}

		log.SetFormatter(&log_formatter{
			TimestampFormat: "Z07:00 2006-01-02 15:04:05.999999999",
		})
		log.SetOutput(os.Stdout)
	})
}

func Debug(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

func Info(format string, args ...interface{}) {
	log.Infof(format, args...)
}

func Error(format string, args ...interface{}) {
	log.Errorf(format, args...)
}
