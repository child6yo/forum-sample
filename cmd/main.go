package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/child6yo/forum-sample"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"

	"github.com/child6yo/forum-sample/internal/app/handler"
	"github.com/child6yo/forum-sample/internal/app/repository"
	"github.com/child6yo/forum-sample/internal/app/service"
)

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

// @title Forum API
// @version 1.0
// @description API Server for Forum

// @host localhost:8000
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// docs - /swagger/index.html

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
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil && err != http.ErrServerClosed {
			log.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()
	log.Print("Forum Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Print("Forum Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Printf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		log.Printf("error occured on db connection close: %s", err.Error())
	}
}