package utils

import (
	"net/http"
	"strings"
)

// GetIP return ip and port from http request
func GetIP(r *http.Request) (string, string) {
	fwd := r.Header.Get("X-Forwarded-For")
	addrStr := ""

	if fwd != "" {
		addrStr = fwd
	} else {
		addrStr = r.RemoteAddr
	}
	addr := strings.Split(addrStr, ":")

	return addr[0], addr[1]
}

// Contains find an element in a slice
func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
