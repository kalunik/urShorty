package entity

import "time"

type PathMeta struct {
	FullUrl   string    `json:"full_url,omitempty"`
	ShortPath string    `json:"short_path,omitempty"`
	Domain    string    `json:"domain,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	VisitedAt time.Time `json:"visited_at,omitempty"`
	Latitude  float32   `json:"latitude,omitempty"`
	Longitude float32   `json:"longitude,omitempty"`
	Country   string    `json:"country,omitempty"`
	City      string    `json:"city,omitempty"`
	Proxy     bool      `json:"proxy,omitempty"`
}

// all visits of path response
type PathVisitsList struct {
	TotalCount int          `ch:"count" json:"total_count"`
	Visits     []*VisitMeta `json:"visits"`
}

// base for response list
type VisitMeta struct {
	VisitedAt time.Time `json:"visited_at,omitempty"`
	Latitude  float32   `json:"latitude,omitempty"`
	Longitude float32   `json:"longitude,omitempty"`
	Country   string    `json:"country,omitempty"`
	City      string    `json:"city,omitempty"`
	Proxy     bool      `json:"proxy,omitempty"`
}
