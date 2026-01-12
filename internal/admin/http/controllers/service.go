package controllers

import (
	"github.com/adi-QTPi/thestral/internal/admin/http/response"
	"github.com/adi-QTPi/thestral/internal/config"
	"github.com/adi-QTPi/thestral/internal/store"
)

type Service struct {
	config   *config.Env
	response response.Responder
	store    store.Service
}

func NewService(cfg *config.Env, responder response.Responder, store store.Service) *Service {
	return &Service{
		config:   cfg,
		response: responder,
		store:    store,
	}
}
