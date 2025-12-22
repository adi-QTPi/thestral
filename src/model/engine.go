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

func (e *Engine) createProxy(destination string) (*httputil.ReverseProxy, error) {
	targetURL, err := url.Parse(destination)
	if err != nil {
		return nil, err
	}
	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		req.Header.Set("X-Forwarded-Host", req.Host)
		req.Header.Set("X-Forwarded-Proto", req.URL.Scheme)
		req.Host = targetURL.Host
	}
	return proxy, nil
}

func (e *Engine) LoadRedis() error {
	data, err := e.Store.GetAll()
	if err != nil {
		return fmt.Errorf("error loading from redis : %v", err)
	}

	for _, proxyObj := range data {
		proxy, err := e.createProxy(proxyObj.Destination)
		if err != nil {
			return fmt.Errorf("error creating proxy : %v", err)
		}
		proxyObj.Proxy = proxy
	}

	e.mutex.Lock()
	defer e.mutex.Unlock()

	e.Routes = data
	fmt.Println("Successfully loaded from redis")
	return nil
}

func (e *Engine) AddRoute(route *types.ProxyRoute) error {
	proxy, err := e.createProxy(route.Destination)
	if err != nil {
		return fmt.Errorf("error creating proxy : %v", err)
	}

	e.mutex.Lock()
	_, ok := e.Routes[route.Source]
	if ok {
		e.mutex.Unlock()
		return fmt.Errorf("the source is already bound")
	}
	route.Proxy = proxy
	e.Routes[route.Source] = route
	e.mutex.Unlock()

	if err := e.Store.Add(route); err != nil {
		e.DeleteRoute(route.Source)
		return fmt.Errorf("unable to insert in redis : %v", err)
	}

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
