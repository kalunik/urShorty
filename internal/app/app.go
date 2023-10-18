package app

import (
	"github.com/kalunik/urShorty/config"
	"github.com/kalunik/urShorty/internal/controller"
	"github.com/kalunik/urShorty/pkg/logger"
	"net/http"
)

func Run() {
	log := logger.NewLogger()
	log.InitLogger()

	log.Info("Launching app")

	configDriver, err := config.LoadNewConfig()
	if err != nil {
		log.Fatal(err)
	}
	appConfig, err := configDriver.ParseConfig()
	if err != nil {
		log.Fatal(err)
	}

	//repository

	//usecase

	r := controller.NewRouter()

	log.Infof("Server will start on %s port", appConfig.Server.Port)
	http.ListenAndServe(appConfig.Server.Port, r.Mux)
	//shutdown
}
