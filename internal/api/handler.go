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

// addPath godoc
//
//	@Summary		Add a new path
//	@Description	generate a short path
//	@Tags			pathMeta
//	@Accept			json
//	@Produce		json
//	@Param			input	body		entity.PathMeta	true	"only full_url param is required"
//	@Success		201		{object}	entity.PathMeta
//	@Failure		400		{object}	error
//	@Failure		500		{object}	error
//	@Router			/short/ [post]
func (h *Handlers) addPath(w http.ResponseWriter, r *http.Request) {
	urlPair := &entity.PathMeta{}
	err := json.NewDecoder(r.Body).Decode(urlPair)
	if err != nil {
		responseError(http.StatusBadRequest, "addPath: json decode fail", err, w, r, h.log)
		return
	}

	if err = h.pathMetaUsecase.AddUrlPath(context.Background(), urlPair); err != nil {
		responseError(http.StatusInternalServerError, "failed to add a new path", err, w, r, h.log)
		return
	}
	h.log.Infof("new urlPair with hash '%s' added to redis", urlPair.ShortPath)

	w.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(w).Encode(urlPair); err != nil {
		responseError(http.StatusInternalServerError, "addPath: json encode fail", err, w, r, h.log)
		return
	}
}

// getFullUrl godoc
//
//	@Summary		redirect to original destination
//	@Description	takes a path and returned an original destination
//	@Tags			pathMeta
//	@Param			hash	query		string	true	"a shortened path"
//	@Success		302		{string}	ok
//	@Failure		404		{object}	error
//	@Failure		500		{object}	error
//	@Router			/short/{hash} [get]
func (h *Handlers) getFullUrl(w http.ResponseWriter, r *http.Request) {
	shortUrl := chi.URLParam(r, "hash")

	fullUrl, err := h.pathMetaUsecase.GetFullUrl(context.Background(), shortUrl, utils.GetIPAddress(r))
	if err != nil {
		responseError(http.StatusNotFound, "shortened link doesn't exist", err, w, r, h.log)
		return
	}

	http.Redirect(w, r, fullUrl, http.StatusFound)
	h.log.Infof("client redirected to `%s`", fullUrl)
}

// pathVisits godoc
//
//	@Summary		list of path's visits
//	@Description	takes a path and returned list of visits
//	@Tags			pathMeta
//	@Param			hash	query		string	true	"a shortened path"
//	@Success		200		{object}	entity.PathMeta
//	@Failure		404		{object}	error
//	@Failure		500		{object}	error
//	@Router			/short/{hash}/visits [get]
func (h *Handlers) pathVisits(w http.ResponseWriter, r *http.Request) {
	shortPath := chi.URLParam(r, "hash")
	pathExist, err := h.pathMetaUsecase.IsExists(context.Background(), shortPath)
	if err != nil {
		responseError(http.StatusInternalServerError, "can't check if path exist", err, w, r, h.log)
		return
	}
	if !pathExist {
		responseError(http.StatusNotFound, "shortened link doesn't exist", err, w, r, h.log)
		return
	}

	visits, err := h.pathMetaUsecase.PathVisits(context.Background(), shortPath)
	if err != nil {
		responseError(http.StatusInternalServerError, "can't get list of visits", err, w, r, h.log)
		return
	}
	if err = json.NewEncoder(w).Encode(visits); err != nil {
		responseError(http.StatusInternalServerError, "pathVisits: json encode fail", err, w, r, h.log)
		return
	}
}
