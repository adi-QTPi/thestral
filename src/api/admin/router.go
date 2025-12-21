package admin

import (
	"log"
	"net/http"

	"github.com/adi-QTPi/thestral/src/config"
	controller "github.com/adi-QTPi/thestral/src/controller/admin"
	"github.com/adi-QTPi/thestral/src/model"
	"github.com/adi-QTPi/thestral/src/utils"
	"github.com/go-chi/chi/v5"
)

func Router(e *model.Engine) *chi.Mux {
	router := utils.NewStdRouter()

	ac := controller.NewAdminController(e)

	router.Get("/ping", ac.Ping)
	router.Post("/add", ac.AddRouteHandler)
	router.Delete("/delete", ac.DeleteRouteHandler)

	return router
}

func Serve(router *chi.Mux, cfg *config.Env) {
	// secureIp := "100.114.106.39:7008"
	// secureIp := "100.113.160.66:7007" //azkaban
	uri := cfg.ADMIN_BIND
	log.Printf("Admin listening on %s", uri)
	if err := http.ListenAndServe(uri, router); err != nil {
		log.Fatalf("Admin server stopped: %v", err)
	}
}
