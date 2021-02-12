package main

import (
	"forwardcall"
	"forwardcall/pkg/handler"
	"forwardcall/pkg/repository"
	"forwardcall/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		logrus.Fatalf("Error occurred while config initialize: %s", err.Error())
	}

	db, err := repository.NewSqliteDB(repository.Config{
		DatabaseFile: viper.GetString("database"),
	})

	if err != nil {
		logrus.Fatalf("Error occurred while database initialize: %s", err.Error())
	}

	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	srv := new(forwardcall.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes(viper.GetString("mode"), viper.GetString("session_key"))); err != nil {
		logrus.Fatalf("Error occurred while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
