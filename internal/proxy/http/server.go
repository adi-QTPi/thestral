package http

import (
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

	c := newProxyController(proxy)

	router.NotFound(c.handlePublicRequest)

	certManager := &autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(cfg.HOST_DOMAIN),
		Cache:      autocert.DirCache("certs"),
	}

	httpsServer := &http.Server{
		Addr:    cfg.PROXY_SSL_BIND,
		Handler: router,
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
			MinVersion:     tls.VersionTLS12,
		},
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
