package model

import (
	"time"
)

type Cross struct {
	ID   string `json:"id"`
	Code string `json:"code"` // "EUR/MXN"
}

type Request struct {
	ID      int       `json:"id"`
	Code    string    `json:"code"`
	Updated time.Time `json:"updated"`
	Rate    float64   `json:"rate"`
}
