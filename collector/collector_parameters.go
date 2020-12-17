package collector

import (
	"crypto/tls"
	"net/http"
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
