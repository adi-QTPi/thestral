package proxy

import (
	"sync"

	"github.com/adi-QTPi/thestral/internal/model"
	"github.com/adi-QTPi/thestral/internal/proxy/route"
)

type Service interface {
	Create(r *model.Route) error
	GetHandler(host string) (*route.Handler, error)
}

type service struct {
	registry map[string]*route.Handler
	mu       sync.RWMutex
}

func NewService() Service {
	return &service{
		registry: make(map[string]*route.Handler),
	}
}
