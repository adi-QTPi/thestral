package store

import "github.com/adi-QTPi/thestral/internal/model"

func (s *service) FindOneRoute(filter *model.Route) (*model.Route, error) {
	return findOne[model.Route](s.db, filter)
}

func (s *service) FindManyRoutes(filter *model.Route) ([]*model.Route, error) {
	return findMany[model.Route](s.db, filter)
}
