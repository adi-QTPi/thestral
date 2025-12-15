package types

import (
	"net/http"
	"time"
)

type ProxyRoute struct {
	ID          string       `json:"id"`
	Source      string       `json:"source"`
	Destination string       `json:"destination"`
	Active      bool         `json:"active"`
	CreatedAt   time.Time    `json:"created_at"`
	Proxy       http.Handler `json:"-"`
}
