package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/adi-QTPi/thestral/src/types"
)

func (ac *AdminController) Ping(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "pong")
}

func (ac *AdminController) AddRouteHandler(w http.ResponseWriter, r *http.Request) {

	n := &types.ProxyRoute{
		Source:      "google.localhost:7007",
		Destination: "http://google.com",
		Active:      true,
	}

	ac.Engine.AddRoute(n)

	fmt.Fprintf(w, "added, src %v, dest %v", n.Source, n.Destination)
}

func (ac *AdminController) DeleteRouteHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("delete service controller called")

	fmt.Fprint(w, "delete service called")
}
