package app

import (
	"github.com/kalunik/urShorty/internal/controller"
	"github.com/kalunik/urShorty/pkg/logger"
	"net/http"
)

func Run() {
	log := logger.NewLogger()
	log.InitLogger()

	log.Info("Launching app")
	
	//repository

	//usecase

	r := controller.NewRouter()
	http.ListenAndServe(":3000", r.Mux)
	//shutdown
}
