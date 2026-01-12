package http

import (
	"log"
	"net/http"

	"github.com/adi-QTPi/thestral/internal/config"
	"github.com/adi-QTPi/thestral/internal/proxy"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func InitServer(cfg *config.Env, proxy proxy.Service) {
	router := newRouter()

	c := newProxyController(proxy)

	router.NotFound(c.handlePublicRequest)

	uri := cfg.PROXY_BIND
	log.Printf("Public Proxy listening on %s", uri)
	if err := http.ListenAndServe(uri, router); err != nil {
		log.Fatalf("public proxy server stopped: %v", err)
	}
}

func newRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	return r
}
