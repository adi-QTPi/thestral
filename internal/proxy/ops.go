package proxy

import (
	"errors"
	"fmt"

	"github.com/adi-QTPi/thestral/internal/admin/dto"
	"github.com/adi-QTPi/thestral/internal/proxy/route"
)

// creates a cached entry for one route (over writes any previous host - targets mapping)
func (s *service) Create(r *dto.RouteDisplay) error {

	routeHandler, err := route.NewRouteHandler(r)
	if err != nil {
		return fmt.Errorf("error creating registry item : %w", err)
	}

	s.mu.Lock()
	s.registry[r.Host] = routeHandler
	s.mu.Unlock()

	fmt.Println("new route created for host : ", r.Host)

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

// [TODO] error handling if new route handler returns error
func (s *service) BulkLoad(arr []*dto.RouteDisplay) {
	s.mu.Lock()
	defer s.mu.Unlock()

	i := 0
	for _, v := range arr {
		routeHandler, _ := route.NewRouteHandler(v) // [TODO] [CRITICAL] handle error handling
		s.registry[v.Host] = routeHandler
		i++
	}
	fmt.Println("bulk loaded records = ", i)
}
