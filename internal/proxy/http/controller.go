package http

import (
	"fmt"
	"net/http"

	"github.com/adi-QTPi/thestral/internal/proxy"
)

type proxyController struct {
	proxy proxy.Service
}

func newProxyController(p proxy.Service) *proxyController {
	return &proxyController{
		proxy: p,
	}
}

func (pc *proxyController) handlePublicRequest(w http.ResponseWriter, r *http.Request) {
	host := r.Host
	data, err := pc.proxy.GetHandler(host)
	if err != nil {
		msg := fmt.Errorf("not found : %v", err)
		http.Error(w, msg.Error(), 404)
		return
	}

	data.ServeHTTP(w, r)
}
