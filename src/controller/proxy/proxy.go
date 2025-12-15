package controller

import (
	"fmt"
	"net/http"
)

func (pc *ProxyController) ProxyLink(w http.ResponseWriter, r *http.Request) {
	host := r.Host
	proxy, err := pc.Engine.GetProxy(host)
	if err != nil {
		msg := fmt.Errorf("not found : %v", err)
		http.Error(w, msg.Error(), 404)
		return
	}
	proxy.Proxy.ServeHTTP(w, r)
}
