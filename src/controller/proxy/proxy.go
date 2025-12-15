package controller

import (
	"fmt"
	"net/http"
)

func (pc *ProxyController) ProxyLink(w http.ResponseWriter, r *http.Request) {
	host := r.Host
	data, err := pc.Engine.GetProxy(host)
	if err != nil {
		msg := fmt.Errorf("not found : %v", err)
		http.Error(w, msg.Error(), 404)
		return
	}
	data.Proxy.ServeHTTP(w, r)
}
