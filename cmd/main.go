package main

import (
	"github.com/kalunik/urShorty/config"
	"github.com/kalunik/urShorty/internal/app"
	"github.com/kalunik/urShorty/pkg/db"
	"github.com/kalunik/urShorty/pkg/logger"
)

func main() {
	log := logger.NewLogger()
	log.InitLogger()

	log.Info("launching app")

	configDriver, err := config.LoadNewConfig()
	if err != nil {
		log.Fatal(err)
	}
	appConfig, err := configDriver.ParseConfig()
	if err != nil {
		log.Fatal(err)
	}

	redisClient, err := db.NewRedisConnection(appConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer redisClient.Close()
	log.Info("redis connected")

	app.NewApp(redisClient, log, appConfig).Run()
}
