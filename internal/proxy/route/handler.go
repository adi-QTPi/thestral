package route

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync/atomic"

	"github.com/adi-QTPi/thestral/internal/admin/dto"
	"github.com/adi-QTPi/thestral/internal/utils"
)

type Handler struct {
	IsActive bool // to be known by proxy controller
	proxy    *httputil.ReverseProxy
	targets  []*url.URL
	counter  uint64
}

func NewRouteHandler(route *dto.RouteDisplay) (*Handler, error) {
	var urls []*url.URL
	for _, t := range route.Targets {
		u, err := url.Parse(t)
		if err != nil {
			return nil, err
		}
		urls = append(urls, u)
	}

	h := &Handler{
		IsActive: *route.IsActive,
		targets:  urls,
		counter:  0,
	}

	director := func(req *http.Request) {
		// round robin implementation
		idx := atomic.AddUint64(&h.counter, 1) % uint64(len(h.targets))
		target := h.targets[idx]

		clientIP := utils.GetRealIP(req)
		log.Printf("[proxying] IP: [%s] -> %s to %s", clientIP, req.URL.Path, target.Host)

		req.Header.Set("X-Forwarded-Host", req.Host)
		req.Header.Set("X-Forwarded-Proto", req.URL.Scheme)

		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.Host = target.Host
		targetPath := target.Path
		if targetPath == "" {
			targetPath = "/"
		}
		req.URL.Path = utils.SingleJoiningSlash(targetPath, req.URL.Path)
	}

	h.proxy = &httputil.ReverseProxy{Director: director}

	return h, nil
}

// allows handler to be used directly by the server
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.proxy.ServeHTTP(w, r)
}
