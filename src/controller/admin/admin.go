package controller

import (
	"fmt"
	"net/http"

	"github.com/adi-QTPi/thestral/src/types"
	"github.com/adi-QTPi/thestral/src/utils"
	"github.com/google/uuid"
)

func (ac *AdminController) Ping(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "pong")
}

func (ac *AdminController) AddRouteHandler(w http.ResponseWriter, r *http.Request) {

	var req types.AddRouteRequest

	if !utils.Validate(w, r, &req) {
		return //http error sent.
	}

	n := &types.ProxyRoute{
		Source:      req.Source,
		Destination: req.Destination,
		Active:      true,
		ID:          uuid.NewString(),
	}

	err := ac.Engine.AddRoute(n)
	if err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, &types.StdResponse{
			Success:     false,
			Description: fmt.Sprintf("error adding : %v", err),
		})
		return
	}

	utils.JSONResponse(w, http.StatusOK, &types.StdResponse{
		Success:     true,
		Description: fmt.Sprintf("[src : %v] <--> [dest : %v]", n.Source, n.Destination),
	})
}

func (ac *AdminController) DeleteRouteHandler(w http.ResponseWriter, r *http.Request) {
	var req types.HostName
	if !utils.Validate(w, r, &req) {
		return //http error sent.
	}
	ac.Engine.DeleteRoute(req.Host)

	utils.JSONResponse(w, http.StatusOK, &types.StdResponse{
		Success:     true,
		Description: fmt.Sprintf("host deleted : %v", req.Host),
	})
}
