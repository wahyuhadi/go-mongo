package utils

import (
	"io"
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

func InitLogFile(log *logrus.Logger) {
	if dir, exist := os.LookupEnv("LOG_DIR"); !exist {
		f, err := os.OpenFile("logger-apps"+time.Now().Format("_20060102")+".log", os.O_WRONLY|os.O_CREATE, 0755)
		if err == nil {
			mw := io.MultiWriter(os.Stdout, f)
			log.SetOutput(mw)
		} else {
			log.Warn(err.Error())
		}
		log.Warn("No Environment Variables found for LOG_DIR, using default value:" + "logger-apps" + time.Now().Format("_20060102_150405") + ".log")
	} else {
		if strings.HasSuffix(dir, "/") {
			strings.TrimSuffix(dir, "/")
		}
		f, err := os.OpenFile(dir+"/logger-apps"+time.Now().Format("_20060102")+".log", os.O_WRONLY|os.O_CREATE, 0755)
		if err == nil {
			mw := io.MultiWriter(os.Stdout, f)
			log.SetOutput(mw)
		} else {
			log.Warn(err.Error())
		}
	}
}