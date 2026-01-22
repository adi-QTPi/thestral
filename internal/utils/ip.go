package utils

import (
	"net"
	"net/http"
	"strings"
)

// gets the ip address from where the request received
func GetRealIP(r *http.Request) string {
	if cfIP := r.Header.Get("CF-Connecting-IP"); cfIP != "" {
		return cfIP
	}

	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		ips := strings.Split(xff, ",")
		return strings.TrimSpace(ips[0])
	}

	// r.RemoteAddr is ALREADY just an IP (thanks to middleware.RealIP).
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr // Return the whole string as the IP
	}
	return ip
}
