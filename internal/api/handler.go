package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/kalunik/urShorty/internal/entity"
	"github.com/kalunik/urShorty/internal/usecase"
	"github.com/kalunik/urShorty/pkg/logger"
	"github.com/kalunik/urShorty/pkg/utils"
	"net/http"
	"net/url"
)

type UrlPairHandlers interface {
	addPair(w http.ResponseWriter, r *http.Request)
	addPairFragment(w http.ResponseWriter, r *http.Request)
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
		utils.LogResponseError(r, h.log, fmt.Errorf("addPair: json decode fail: %w", err))
		http.Error(w, "Failed to decode JSON data", http.StatusBadRequest)
		return
	}

	urlPair.Short, err = utils.GenerateHash(urlPair.Full)
	if err != nil {
		utils.LogResponseError(r, h.log, errors.New("addPair: generate hash fail"))
		http.Error(w, "Failed to generate short url", http.StatusInternalServerError)
		return
	}

	if err := h.urlPairUsecase.AddUrlPair(context.Background(), urlPair); err != nil {
		utils.LogResponseError(r, h.log, fmt.Errorf("addPair fail: %w", err))
		http.Error(w, "Failed to add url pair", http.StatusInternalServerError)
		return
	}
	h.log.Infof("new urlPair with hash '%s' added to redis", urlPair.Short)

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(urlPair); err != nil {
		utils.LogResponseError(r, h.log, fmt.Errorf("addPair: encode fail, %w", err))
		http.Error(w, "Failed to encode json", http.StatusInternalServerError)
		return
	}
}

func (h *Handlers) addPairFragment(w http.ResponseWriter, r *http.Request) {
	u, err := url.ParseRequestURI(r.RequestURI)
	if err != nil {
		utils.LogResponseError(r, h.log, fmt.Errorf("addPairFragment: parse URI fail, %w", err))
		http.Error(w, "Failed to parse URI", http.StatusBadRequest)
		return
	}
	urlPair := &entity.UrlPair{Full: u.EscapedFragment()}

	parsed, _ := url.ParseQuery(u.RequestURI())
	fmt.Println("hey", parsed, "-", u.EscapedFragment(), "-", u.RawFragment, "-", u.EscapedPath(), "-", u.RequestURI(), "-", r.URL, "-", r.RequestURI)
	urlPair.Short, err = utils.GenerateHash(urlPair.Full)
	if err != nil {
		utils.LogResponseError(r, h.log, errors.New("addPair: generate hash fail"))
		http.Error(w, "Failed to generate short url", http.StatusInternalServerError)
		return
	}

	if err := h.urlPairUsecase.AddUrlPair(context.Background(), urlPair); err != nil {
		utils.LogResponseError(r, h.log, fmt.Errorf("addPair fail: %w", err))
		http.Error(w, "Failed to add url pair", http.StatusInternalServerError)
		return
	}
	h.log.Infof("new urlPair with hash '%s' added to redis", urlPair.Short)

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(urlPair); err != nil {
		utils.LogResponseError(r, h.log, fmt.Errorf("addPair: encode fail, %w", err))
		http.Error(w, "Failed to encode json", http.StatusInternalServerError)
		return
	}
}

func (h *Handlers) getFullUrl(w http.ResponseWriter, r *http.Request) {
	shortUrl := chi.URLParam(r, "hash")
	fullUrl, err := h.urlPairUsecase.FindFullUrl(context.Background(), shortUrl)
	if err != nil {
		utils.LogResponseError(r, h.log, err)
		http.Error(w, "Failed to find full URL", http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, fullUrl, http.StatusFound)
	h.log.Infof("client redirected to `%s`", fullUrl)
}
