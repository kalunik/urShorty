package api

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/kalunik/urShorty/internal/entity"
	"github.com/kalunik/urShorty/internal/usecase"
	"github.com/kalunik/urShorty/pkg/logger"
	"github.com/kalunik/urShorty/pkg/utils"
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
	err := json.NewDecoder(r.Body).Decode(urlPair)
	if err != nil {
		h.log.Error("urlPairHandlers: addPair: json decode fail")
		http.Error(w, "Failed to decode JSON data", http.StatusBadRequest)
		return
	}

	urlPair.Short, err = utils.GenerateHash(urlPair.Full)
	if err != nil {
		h.log.Error("urlPairHandlers: addPair: generate hash fail")
		http.Error(w, "Failed to decode JSON data", http.StatusInternalServerError)
		return
	}

	if err := h.urlPairUsecase.AddUrlPair(context.Background(), urlPair); err != nil {
		h.log.Error("urlPairHandlers: fail to add pair")
		http.Error(w, "Failed to add url pair", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(urlPair); err != nil {
		h.log.Error("urlPairHandlers: addPair: encode fail")
		http.Error(w, "Failed to encode json", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *Handlers) addPairHashParam(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *Handlers) getFullUrl(w http.ResponseWriter, r *http.Request) {
	shortUrl := chi.URLParam(r, "hash")
	fullUrl, err := h.urlPairUsecase.FindFullUrl(context.Background(), shortUrl)
	if err != nil {
		h.log.Errorf("urlPairHandlers: getFullUrl: %w", err)
		http.Error(w, "Failed to find full URL", http.StatusBadRequest)
		return
	}
	http.RedirectHandler(fullUrl, http.StatusMovedPermanently)
}
