package utils

import (
	"fmt"
	"net"
	"net/http"
	"strings"
)

func ExtractSubdomain(r *http.Request) string {
	host := r.Host
	fmt.Println(host)
	if h, _, err := net.SplitHostPort(host); err == nil {
		host = h
	}

	parts := strings.Split(host, ".")
	if len(parts) > 1 {
		subdomains := parts[:len(parts)-1]
		return strings.Join(subdomains, ".")
	}

	return ""
}
