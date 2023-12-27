package webapi

import (
	"encoding/json"
	"fmt"
	"github.com/kalunik/urShorty/internal/entity"
	"net/http"
)

type IPGeolocation interface {
	GetIPLocation(userIP string) (entity.ResponseGeolocateAPI, error)
}

type IPGeoWebAPI struct {
}

func NewIPGeoWebAPI() IPGeolocation {
	return &IPGeoWebAPI{}
}

func (a IPGeoWebAPI) GetIPLocation(userIP string) (entity.ResponseGeolocateAPI, error) {
	apiUrl := fmt.Sprintf("http://ip-api.com/json/%s", userIP)

	resp, err := http.Get(apiUrl)
	if err != nil {
		return entity.ResponseGeolocateAPI{}, fmt.Errorf("post fail: %w", err)
	}
	defer resp.Body.Close()

	respData := entity.ResponseGeolocateAPI{}
	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		return entity.ResponseGeolocateAPI{}, err
	}
	return respData, nil
}
