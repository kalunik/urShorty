package api

import (
	"fmt"
	"github.com/kalunik/urShorty/pkg/logger"
	"github.com/kalunik/urShorty/pkg/utils"
	"net/http"
)

func responseError(status int, message string, err error, w http.ResponseWriter, r *http.Request, log logger.Logger) {
	logResponseError(r, log, fmt.Errorf("%s: %w", message, err))
	http.Error(w, message, status)
}

func logResponseError(r *http.Request, logger logger.Logger, err error) {
	logger.ErrorfCaller(
		1, "RequestID: %s, IPAddress: %s, Error: %s",
		utils.GetRequestID(r),
		utils.GetIPAddress(r),
		err,
	)
}
