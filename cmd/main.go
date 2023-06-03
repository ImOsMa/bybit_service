package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/ImOsMa/bybit_service"
	"github.com/ImOsMa/bybit_service/pkg/handler"
	"github.com/ImOsMa/bybit_service/pkg/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// @title ByBit Service API
// @version 1.0
// @description API Server for ByBit Service

func main() {
	if err := initConfig(); err != nil {
		logrus.Fatalf("error init configs: %s", err.Error())
	}
	gin.SetMode(gin.ReleaseMode)

	services := service.NewService()
	handlers := handler.NewHandler(services)

	srv := new(bybit_service.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("ByBit Service Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("ByBit Service Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
