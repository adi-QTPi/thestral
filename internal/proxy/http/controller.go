package http

import (
	"errors"
	"fmt"
	"net/http"
	"sync"

	"golang.org/x/time/rate"

	"github.com/adi-QTPi/thestral/internal/proxy"
	"github.com/adi-QTPi/thestral/internal/utils"
)

type IPRateLimiter struct {
	ips map[string]*rate.Limiter
	mu  sync.Mutex
	r   rate.Limit
	b   int
}

type proxyController struct {
	proxy   proxy.Service
	limiter *IPRateLimiter
}

func newProxyController(p proxy.Service, rateLimit int, burst int) *proxyController {
	return &proxyController{
		proxy: p,
		limiter: &IPRateLimiter{
			ips: make(map[string]*rate.Limiter),
			r:   rate.Limit(rateLimit),
			b:   burst,
		},
	}
}

func (pc *proxyController) handlePublicRequest(w http.ResponseWriter, r *http.Request) {
	err := pc.rateLimit(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusTooManyRequests)
		return
	}

	host := r.Host
	data, err := pc.proxy.GetHandler(host)
	if err != nil {
		msg := fmt.Errorf("not found : %v", err)
		http.Error(w, msg.Error(), http.StatusNotFound)
		return
	}

	data.ServeHTTP(w, r)
}

func (i *IPRateLimiter) getLimiter(ip string) *rate.Limiter {
	i.mu.Lock()
	defer i.mu.Unlock()

	limiter, exists := i.ips[ip]
	if !exists {
		limiter = rate.NewLimiter(i.r, i.b)
		i.ips[ip] = limiter
	}
	return limiter
}

// basic header based rate limitter
// please use protect the server using cloud based firewalls for cloudflare ip ranges only
// to protect from header spoofing
func (pc *proxyController) rateLimit(w http.ResponseWriter, r *http.Request) error {

	clientIP := utils.GetRealIP(r)

	limiter := pc.limiter.getLimiter(clientIP)

	if !limiter.Allow() {
		return errors.New("Too Many Requests - Slow Down")
	}
	return nil
}
