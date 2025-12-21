package model

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"

	"github.com/adi-QTPi/thestral/src/config"
	"github.com/adi-QTPi/thestral/src/types"
)

type Engine struct {
	Routes map[string]*types.ProxyRoute
	Store  types.Storage
	mutex  sync.RWMutex
}

func NewEngine(cfg *config.Env) *Engine {

	client, err := NewRedisClient(cfg.REDIS_HOST, cfg.REDIS_PORT, cfg.REDIS_PASSWORD)
	if err != nil {
		log.Fatal("redis connection issue : ", err)
	}

	return &Engine{
		Routes: make(map[string]*types.ProxyRoute),
		Store:  client,
	}
}

func (e *Engine) AddRoute(route *types.ProxyRoute) error {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	targetURL, err := url.Parse(route.Destination)
	if err != nil {
		return fmt.Errorf("error adding the url : %v", err)
	}

	_, ok := e.Routes[route.Source]
	if ok {
		return fmt.Errorf("the source is already bound")
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
