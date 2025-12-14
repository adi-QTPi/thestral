package proxy

import (
	"log"
	"net/http"

	util "github.com/adi-QTPi/thestral/src/utils"
	"github.com/go-chi/chi/v5"
)

func Router() *chi.Mux {
	router := util.NewStdRouter()

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		log.Println("some request detected on normal.")
	})

	return router
}

func Serve(router *chi.Mux) {
	publicIp := "0.0.0.0:8080"
	log.Printf("Public Proxy listening on %s", publicIp)
	if err := http.ListenAndServe(publicIp, router); err != nil {
		log.Fatalf("normal server stopped: %v", err)
	}
}
