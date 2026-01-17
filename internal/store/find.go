package store

import (
	"github.com/adi-QTPi/thestral/internal/admin/dto"
	"github.com/adi-QTPi/thestral/internal/model"
)

func (s *service) FindOneRoute(filter *model.Route) (*dto.RouteDisplay, error) {
	data, err := findOne[model.Route](s.db, filter)
	if err != nil {
		return nil, err
	}

	out := &dto.RouteDisplay{
		Host:      data.Host,
		Targets:   data.Targets,
		IsActive:  data.IsActive,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}

	return out, nil
}

func (s *service) FindManyRoutes(filter *model.Route) ([]*dto.RouteDisplay, error) {
	data, err := findMany[model.Route](s.db, filter)
	if err != nil {
		return nil, err
	}
	var out []*dto.RouteDisplay
	for _, v := range data {
		out = append(out, &dto.RouteDisplay{
			Host:      v.Host,
			Targets:   v.Targets,
			IsActive:  v.IsActive,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}
	return out, nil
}
