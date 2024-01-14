package api

import (
	"errors"
	"github.com/golang/mock/gomock"
	mock_logger "github.com/kalunik/urShorty/pkg/logger/mocks"
	"github.com/kalunik/urShorty/pkg/utils"
	"net/http"
	"testing"
)

func TestLogResponseError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockLogger := mock_logger.NewMockLogger(ctrl)

	r := &http.Request{
		RemoteAddr: "127.0.0.1:43647",
	}
	expectedError := errors.New("testError")
	mockLogger.EXPECT().ErrorfCaller(1, "RequestID: %s, IPAddress: %s, Error: %s",
		utils.GetRequestID(r),
		utils.GetIPAddress(r),
		expectedError).Return()

	logResponseError(r, mockLogger, expectedError)
}
