package collector

import (
	"crypto/tls"
	"net/http"
)

const (
	namespace = "syncthing"
)

var (
	// Skip insecure connection verify due to using self signed certificate on syncthing side.
	HttpClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
)

func bool2float64(status bool) float64 {
	if status {
		return float64(1)
	}
	return float64(0)
}
