package proxy

import (
	"log"
	"net/http"

	"github.com/adi-QTPi/thestral/src/config"
	controller "github.com/adi-QTPi/thestral/src/controller/proxy"
	"github.com/adi-QTPi/thestral/src/model"
	"github.com/adi-QTPi/thestral/src/utils"
	"github.com/go-chi/chi/v5"
)

func Router(e *model.Engine) *chi.Mux {
	router := utils.NewStdRouter()

	pc := controller.NewProxyController(e)

	router.NotFound(pc.ProxyLink)

	return router
}

func Serve(router *chi.Mux, cfg *config.Env) {
	// publicIp := "0.0.0.0:80"
	uri := cfg.PROXY_BIND
	log.Printf("Public Proxy listening on %s", uri)
	if err := http.ListenAndServe(uri, router); err != nil {
		log.Fatalf("normal server stopped: %v", err)
	}
}
