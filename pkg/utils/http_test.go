package utils

import (
	"context"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetIPAddressGetRequestID(t *testing.T) {
	test := make(map[string]string)
	test["reqId_0"] = "req-123456"
	test["reqId_1"] = "278ac745-sdf"

	for _, expectedReqId := range test {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), middleware.RequestIDKey, expectedReqId)
			r = r.WithContext(ctx)

			reqId := GetRequestID(r)

			assert.Equal(t, expectedReqId, reqId, "X-Request-Id should be the same")
		}))
		client := ts.Client()
		_, _ = client.Get(ts.URL)
		ts.Close()
	}
}

func TestGetIPAddress(t *testing.T) {
	test := []struct {
		key        string
		expectedIp string
	}{
		{
			key:        "X-Real-Ip",
			expectedIp: "150.172.238.178",
		},
		{
			key:        "X-Forwarded-For",
			expectedIp: "203.0.113.195",
		},
		{
			key:        "empty",
			expectedIp: "",
		},
	}

	for _, ex := range test {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r.Header.Set(ex.key, ex.expectedIp)
			ipAddr := GetIPAddress(r)

			if ex.key == "empty" {
				assert.Equalf(t, r.RemoteAddr, ipAddr, "ip should be the same: expected %s, got %s\n", ex.expectedIp, ipAddr)
				return
			}
			assert.Equalf(t, ex.expectedIp, ipAddr, "ip should be the same: expected %s, got %s\n", ex.expectedIp, ipAddr)
		}),
		)
		client := ts.Client()
		_, _ = client.Get(ts.URL)
		ts.Close()
	}
}
