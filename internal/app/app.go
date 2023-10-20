package app

import (
	"github.com/kalunik/urShorty/config"
	"github.com/kalunik/urShorty/internal/controller"
	r "github.com/kalunik/urShorty/internal/repository"
	"github.com/kalunik/urShorty/internal/usecase"
	"github.com/kalunik/urShorty/pkg/db"
	"github.com/kalunik/urShorty/pkg/logger"
	"net/http"
)

func Run() {
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

	urlPairUsecase := usecase.NewUrlPairUsecase(r.NewRedisRepository(redisClient), log)

	r := controller.NewRouter()

	log.Infof("server will start on %s port", appConfig.Server.Port)
	http.ListenAndServe(appConfig.Server.Port, r.Mux)
	//shutdown
}
