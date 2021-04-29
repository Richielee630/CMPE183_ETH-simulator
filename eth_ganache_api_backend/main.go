package main

import (
	"api_backend/models"
	"api_backend/pkg/logging"
	"api_backend/pkg/setting"
	"api_backend/pkg/util"
	"api_backend/routers"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	setting.Setup()
	logging.Setup()
	util.Setup()
}
func main() {
	models.Init()
	models.InitMysql()

	gin.SetMode(setting.ServerSetting.RunMode)
	gin.ForceConsoleColor()

	router := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        router,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	logging.Info("Start http server listening:", endPoint)
	err := server.ListenAndServe()
	if err != nil {
		logging.Error(err)
	}
}
