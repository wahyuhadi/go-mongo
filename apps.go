package main

import (

	"time"
	"github.com/gin-gonic/gin"
	"gopkg.in/tylerb/graceful.v1"
	"framework/utils"
	"framework/handler"
	"github.com/sirupsen/logrus"
)


func main() {
	var log = logrus.New()
	utils.InitLogFile(log)
	environ := utils.InitEnv(log)
	r := gin.New()

	userHandler := handler.NewUsersHandler(log, environ)
	userHandler.Setup(r)

	// orderHandler := handler.NewOrdersHandler(log, environ)
	// orderHandler.Setup(r)
	graceful.Run(environ.AppHost, 10*time.Second, r)
}