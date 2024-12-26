package main

import (
	"log"
	"os"

	"github.com/child6yo/forum-sample"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"

	"github.com/child6yo/forum-sample/pkg/handler"
	"github.com/child6yo/forum-sample/pkg/repository"
	"github.com/child6yo/forum-sample/pkg/service"
)

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func main() {
	if err := initConfig(); err != nil {
		log.Fatal("config not initialized")
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal("env not initialized")
    }

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
    if err != nil {
		log.Fatal("not connected to db: ", err)
	}

    repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	
    srv := new(forum.Server)
    srv.Run("8000", handlers.InitRoutes())
}