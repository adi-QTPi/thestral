package proxy

import (
	"sync"

	"github.com/adi-QTPi/thestral/internal/admin/dto"
	"github.com/adi-QTPi/thestral/internal/proxy/route"
)

type Service interface {
	Create(r *dto.RouteDisplay) error
	Delete(host string)
	GetHandler(host string) (*route.Handler, error)
	BulkLoad(arr []*dto.RouteDisplay)
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
