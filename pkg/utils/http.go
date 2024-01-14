package utils

import (
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

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
