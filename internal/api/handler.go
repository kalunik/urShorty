package api

import (
	"context"
	"encoding/json"
	"github.com/kalunik/urShorty/internal/entity"
	"github.com/kalunik/urShorty/internal/usecase"
	"github.com/kalunik/urShorty/pkg/logger"
	"net/http"
)

type UrlPairHandlers interface {
	addPair(w http.ResponseWriter, r *http.Request)
	addPairHashParam(w http.ResponseWriter, r *http.Request)
	getFullUrl(writer http.ResponseWriter, request *http.Request)
}

type Handlers struct {
	urlPairUsecase usecase.Usecase
	log            logger.Logger
}

func NewUrlPairHandlers(urlPairUC usecase.Usecase, log logger.Logger) UrlPairHandlers {
	return &Handlers{
		urlPairUsecase: urlPairUC,
		log:            log,
	}
}

// Create a new pair of short and full url
// return a shortUrl (aka key) for getFullUrl handler
func (h *Handlers) addPair(w http.ResponseWriter, r *http.Request) {
	urlPair := &entity.UrlPair{}
	if err := json.NewDecoder(r.Body).Decode(urlPair); err != nil {
		h.log.Error("urlPairHandlers: addPair ")
		http.Error(w, "Failed to decode JSON data", http.StatusBadRequest)
		return
	}
	if err := h.urlPairUsecase.AddUrlPair(context.Background(), urlPair); err != nil {
		h.log.Error("urlPairHandlers: fail to add pair")
		http.Error(w, "Failed to decode JSON data", http.StatusBadRequest)
		return
	}

}

func (h *Handlers) addPairHashParam(w http.ResponseWriter, r *http.Request) {
}

func (h *Handlers) getFullUrl(w http.ResponseWriter, r *http.Request) {

}
