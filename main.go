package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/coolishbee/go-gin-sample/models"
	"github.com/coolishbee/go-gin-sample/pkg/logging"
	"github.com/coolishbee/go-gin-sample/pkg/setting"
	"github.com/coolishbee/go-gin-sample/pkg/util"
	"github.com/coolishbee/go-gin-sample/routers"
	"github.com/gin-gonic/gin"
)

func init() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	util.Setup()
}

// @title Golang Gin API
// @version 1.0
// @description An example of gin
// @termsOfService https://github.com/coolishbee/go-gin-sample
// @license.name MIT
// @license.url https://github.com/coolishbee/go-gin-sample/blob/main/LICENSE
func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}
	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
}
