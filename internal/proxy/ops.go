package proxy

import (
	"errors"
	"fmt"

	"github.com/adi-QTPi/thestral/internal/model"
	"github.com/adi-QTPi/thestral/internal/proxy/route"
)

// creates a cached entry for one route (over writes any previous host - targets mapping)
func (s *service) Create(r *model.Route) error {

	routeHandler, err := route.NewRouteHandler(r)
	if err != nil {
		return fmt.Errorf("error creating registry item : %w", err)
	}

	s.mu.Lock()
	s.registry[r.Host] = routeHandler
	s.mu.Unlock()

	fmt.Printf("new route created for host : %v", r.Host)

	return nil
}

func (s *service) Delete(host string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.registry, host)
}

func (s *service) GetHandler(host string) (*route.Handler, error) {
	s.mu.RLock()
	handler, exists := s.registry[host]
	s.mu.RUnlock()

	if !exists {
		return handler, errors.New("unable to forward your request")
	}

	if !handler.IsActive {
		return nil, errors.New("site under maintainance")
	}

	return handler, nil
}
