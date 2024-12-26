package main

import (
	"os"

	"github.com/child6yo/forum-sample"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"

	"github.com/child6yo/forum-sample/pkg/handler"
	"github.com/child6yo/forum-sample/pkg/repository"
	"github.com/child6yo/forum-sample/pkg/service"
)

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}


func main() {
	if err := initConfig(); err != nil {
	}

	if err := godotenv.Load(); err != nil {
    }

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
    if err != nil {
	}


    srv := new(forum.Server)
    repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
    srv.Run(viper.GetString("port"), handlers.InitRoutes())
}