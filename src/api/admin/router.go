package admin

import (
	"log"
	"net/http"

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

func Serve(router *chi.Mux) {
	secureIp := "100.114.106.39:7008"
	log.Printf("Admin listening on %s", secureIp)
	if err := http.ListenAndServe(secureIp, router); err != nil {
		log.Fatalf("Admin server stopped: %v", err)
	}
}
