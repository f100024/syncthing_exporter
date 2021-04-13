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
