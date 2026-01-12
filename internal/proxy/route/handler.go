package route

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync/atomic"

	"github.com/adi-QTPi/thestral/internal/model"
)

type Handler struct {
	IsActive bool // to be known by proxy controller
	proxy    *httputil.ReverseProxy
	targets  []*url.URL
	counter  uint64
}

func NewRouteHandler(route *model.Route) (*Handler, error) {
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
	}

	director := func(req *http.Request) {
		// round robin implementation
		idx := atomic.AddUint64(&h.counter, 1) % uint64(len(h.targets))
		target := h.targets[idx]

		req.Header.Set("X-Forwarded-Host", req.Host)
		req.Header.Set("X-Forwarded-Proto", req.URL.Scheme)

		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.Host = target.Host
	}

	h.proxy = &httputil.ReverseProxy{Director: director}

	return h, nil
}

// allows handler to be used directly by the server
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.proxy.ServeHTTP(w, r)
}
