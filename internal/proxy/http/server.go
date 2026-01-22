package http

import (
	"context"
	"crypto/tls"
	"log"
	"net/http"

	"github.com/adi-QTPi/thestral/internal/config"
	"github.com/adi-QTPi/thestral/internal/proxy"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"golang.org/x/crypto/acme/autocert"
)

func InitServer(cfg *config.Env, proxy proxy.Service) {
	router := newRouter()

	c := newProxyController(proxy, cfg.RATE_LIMIT_REQ_PER_SEC, cfg.RATE_LIMIT_BURST) // 1 req/sec and burst of 5 by default

	router.NotFound(c.handlePublicRequest) // NotFound in chi lies outside the middleware chain, thus rate limitting is implemented inside the controller itself

	if !cfg.ENABLE_TLS {
		log.Println("Disabling TLS")
		log.Printf("Listening on HTTP %s", cfg.PROXY_BIND)

		if err := http.ListenAndServe(cfg.PROXY_BIND, router); err != nil {
			log.Fatalf("Dev server failed: %v", err)
		}
		return
	}

	hostPolicy := func(ctx context.Context, host string) error {
		if _, err := proxy.GetHandler(host); err != nil {
			return err
		}
		return nil
	}

	certManager := &autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: hostPolicy,
		Cache:      autocert.DirCache("certs"),
	}

	tlsConfig := certManager.TLSConfig()
	tlsConfig.MinVersion = tls.VersionTLS12

	httpsServer := &http.Server{
		Addr:      cfg.PROXY_SSL_BIND,
		Handler:   router,
		TLSConfig: tlsConfig,
	}

	// Challenge Handler on HTTP for lets encrypt management
	uri := cfg.PROXY_BIND
	go func() {
		log.Println("ACME Challenge Server listening on ", uri)
		if err := http.ListenAndServe(uri, certManager.HTTPHandler(nil)); err != nil {
			log.Fatalf("HTTP challenge server failed: %v", err)
		}
	}()

	secureUri := cfg.PROXY_SSL_BIND
	log.Println("Public HTTPS Proxy listening on ", secureUri)
	if err := httpsServer.ListenAndServeTLS("", ""); err != nil {
		log.Fatalf("public proxy server stopped: %v", err)
	}
}

func newRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	return r
}
