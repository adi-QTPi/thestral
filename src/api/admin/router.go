package admin

import (
	"log"
	"net/http"

	"github.com/adi-QTPi/thestral/src/middleware"
	util "github.com/adi-QTPi/thestral/src/utils"
	"github.com/go-chi/chi/v5"
)

func Router(fallback http.Handler) *chi.Mux {
	router := util.NewStdRouter()
	router.Use(middleware.AdminFilter(fallback)) //redirects non admin requests to proxy router.

	router.Get("/admin",
		func(w http.ResponseWriter, r *http.Request) {
			log.Println("get on admin router detected...")
		})

	return router
}

func Serve(router *chi.Mux) {
	secureIp := "100.114.106.39:8080"
	log.Printf("Admin listening on %s", secureIp)
	if err := http.ListenAndServe(secureIp, router); err != nil {
		log.Fatalf("Admin server stopped: %v", err)
	}
}
