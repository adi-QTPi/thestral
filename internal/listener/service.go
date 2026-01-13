package listener

import (
	"github.com/adi-QTPi/thestral/internal/config"
	"github.com/adi-QTPi/thestral/internal/proxy"
	"github.com/adi-QTPi/thestral/internal/store"
)

type Service interface {
	Run() error
	Load() error
}

type service struct {
	cfg   *config.Env
	proxy proxy.Service
	store store.Service
}

func NewService(cfg *config.Env, p proxy.Service, s store.Service) Service {
	return &service{
		cfg:   cfg,
		proxy: p,
		store: s,
	}
}
