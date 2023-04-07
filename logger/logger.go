package logger

import (
	"os"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger
var once sync.Once

func Init() {
	once.Do(func() {
		if log == nil {
			log = logrus.New()
		}
		log.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:    true,
			TimestampFormat:  time.RFC3339Nano,
			QuoteEmptyFields: true,
		})
		log.SetOutput(os.Stdout)
	})
}

func Test() {
	log.Infof("This is an info log. %d", 1)
	log.Error("This is an error log.")
}

func Info(format string, args ...interface{}) {
	log.Infof(format, args...)
}

func Error(format string, args ...interface{}) {
	log.Errorf(format, args...)
}
