package entity

type ResponseGeolocateAPI struct {
	Country string `json:"country,omitempty"`
	City    string `json:"city,omitempty"`
	Proxy   bool   `json:"proxy,omitempty"`
}
