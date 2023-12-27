package utils

import (
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kalunik/urShorty/pkg/logger"
	"net/http"
)

func LogResponseError(r *http.Request, logger logger.Logger, err error) {
	logger.ErrorfCaller(
		1, "RequestID: %s, IPAddress: %s, Error: %s",
		GetRequestID(r),
		GetIPAddress(r),
		err,
	)
}

func GetRequestID(r *http.Request) string {
	return middleware.GetReqID(r.Context())
}

func GetIPAddress(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}
