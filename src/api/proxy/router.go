package proxy

import (
	"log"
	"net/http"

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

func Serve(router *chi.Mux) {
	publicIp := "0.0.0.0:7007"
	log.Printf("Public Proxy listening on %s", publicIp)
	if err := http.ListenAndServe(publicIp, router); err != nil {
		log.Fatalf("normal server stopped: %v", err)
	}
}
