package webapi

import (
	"bytes"
	"encoding/json"
	"errors"
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
	url := "http://ip-api.com/batch\\?fields\\=country,city,proxy"

	contentType := "application/json"
	bodyJson, err := json.Marshal(userIP)
	if err != nil {
		return entity.ResponseGeolocateAPI{}, errors.New("json marshal fail")
	}
	body := bytes.NewBuffer(bodyJson)

	post, err := http.Post(url, contentType, body)
	if err != nil {
		return entity.ResponseGeolocateAPI{}, fmt.Errorf("post fail: %w", err)
	}
	defer post.Body.Close()

	responseData := entity.ResponseGeolocateAPI{}
	json.NewDecoder(post.Body).Decode(&responseData)
	return responseData, nil
}
