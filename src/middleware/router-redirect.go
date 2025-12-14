package middleware

import (
	"log"
	"net/http"
)

// fallback is the proxy router.
func AdminFilter(fallback http.Handler) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("middleware touched.")
			if 1 > 3 {
				next.ServeHTTP(w, r)
			} else {
				fallback.ServeHTTP(w, r)
			}
		})
	}
}
