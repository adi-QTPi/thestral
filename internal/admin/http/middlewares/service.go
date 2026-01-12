package middlewares

import (
	"github.com/adi-QTPi/thestral/internal/admin/http/response"
	"github.com/adi-QTPi/thestral/internal/config"
)

type Service struct {
	config   *config.Env
	response response.Responder
}

func NewService(cfg *config.Env, responder response.Responder) *Service {
	return &Service{
		config:   cfg,
		response: responder,
	}
}
