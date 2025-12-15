package model

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"

	"github.com/adi-QTPi/thestral/src/types"
)

type Engine struct {
	Routes map[string]*types.ProxyRoute
	mutex  sync.RWMutex
}

func NewEngine() *Engine {
	return &Engine{
		Routes: make(map[string]*types.ProxyRoute),
	}
}

func (e *Engine) AddRoute(route *types.ProxyRoute) error {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	targetURL, err := url.Parse(route.Destination)
	if err != nil {
		return fmt.Errorf("error adding the url : %v", err)
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		req.Header.Set("X-Forwarded-Host", req.Host)
		req.Header.Set("X-Forwarded-Proto", req.URL.Scheme)
		req.Host = targetURL.Host
	}

	route.Proxy = proxy

	e.Routes[route.Source] = route

	return nil
}

func (e *Engine) DeleteRoute(source string) error {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	if _, exists := e.Routes[source]; !exists {
		return fmt.Errorf("no record found for this source host")
	}
	delete(e.Routes, source)
	return nil
}

func (e *Engine) GetProxy(target string) (*types.ProxyRoute, error) {
	e.mutex.RLock()
	proxy, exists := e.Routes[target]
	e.mutex.RUnlock()

	if !exists {
		return nil, fmt.Errorf("unable to forward your request")
	}

	if !proxy.Active {
		return nil, fmt.Errorf("site under maintenance")
	}

	return proxy, nil
}
