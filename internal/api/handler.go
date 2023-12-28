package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/kalunik/urShorty/internal/entity"
	"github.com/kalunik/urShorty/internal/usecase"
	"github.com/kalunik/urShorty/pkg/logger"
	"github.com/kalunik/urShorty/pkg/utils"
	"net/http"
)

type PathMetaHandlers interface {
	addPath(w http.ResponseWriter, r *http.Request)
	getFullUrl(writer http.ResponseWriter, request *http.Request)
	pathVisits(writer http.ResponseWriter, request *http.Request)
}

type Handlers struct {
	pathMetaUsecase usecase.Usecase
	log             logger.Logger
}

func NewPathMetaHandlers(pathMetaUC usecase.Usecase, log logger.Logger) PathMetaHandlers {
	return &Handlers{
		pathMetaUsecase: pathMetaUC,
		log:             log,
	}
}

// Create a new pair of short and full url
// return a shortUrl (aka key) for getFullUrl handler
func (h *Handlers) addPath(w http.ResponseWriter, r *http.Request) {
	urlPair := &entity.PathMeta{}
	err := json.NewDecoder(r.Body).Decode(urlPair)
	if err != nil {
		utils.LogResponseError(r, h.log, fmt.Errorf("addPath: json decode fail: %w", err))
		http.Error(w, "Failed to decode JSON data", http.StatusBadRequest)
		return
	}

	if err := h.pathMetaUsecase.AddUrlPath(context.Background(), urlPair); err != nil {
		utils.LogResponseError(r, h.log, fmt.Errorf("addPath fail: %w", err))
		http.Error(w, "Failed to add url pair", http.StatusInternalServerError)
		return
	}
	h.log.Infof("new urlPair with hash '%s' added to redis", urlPair.ShortPath)

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(urlPair); err != nil {
		utils.LogResponseError(r, h.log, fmt.Errorf("addPath: encode fail, %w", err))
		http.Error(w, "Failed to encode json", http.StatusInternalServerError)
		return
	}
}

func (h *Handlers) getFullUrl(w http.ResponseWriter, r *http.Request) {
	shortUrl := chi.URLParam(r, "hash")

	fullUrl, err := h.pathMetaUsecase.GetFullUrl(context.Background(), shortUrl, utils.GetIPAddress(r))
	if err != nil {
		utils.LogResponseError(r, h.log, err)
		http.Error(w, "Failed to find full URL", http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, fullUrl, http.StatusFound)
	h.log.Infof("client redirected to `%s`", fullUrl)
}

func (h *Handlers) pathVisits(w http.ResponseWriter, r *http.Request) {
	shortPath := chi.URLParam(r, "hash")
	pathExist, err := h.pathMetaUsecase.IsExists(context.Background(), shortPath)
	if err != nil {
		utils.LogResponseError(r, h.log, err)
		http.Error(w, "Fail to check path existence", http.StatusInternalServerError)
		return
	}
	if !pathExist {
		http.Error(w, "No such URL", http.StatusNotFound)
		return
	}

	visits, err := h.pathMetaUsecase.PathVisits(context.Background(), shortPath)
	if err != nil {
		utils.LogResponseError(r, h.log, fmt.Errorf("pathVisits: usecase fail, %w", err))
		http.Error(w, "Failed to get list of visits", http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(visits); err != nil {
		utils.LogResponseError(r, h.log, fmt.Errorf("pathVisits: encode fail, %w", err))
		http.Error(w, "Failed to encode json", http.StatusInternalServerError)
		return
	}
}
