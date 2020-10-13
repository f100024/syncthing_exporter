package collector

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/prometheus/common/promlog"
)

func TestNewSCReport(t *testing.T) {

	jsonResponse, err := ioutil.ReadFile("fixtures/rest_system_connections_response.json")
	if err != nil {
		t.Error("Cant open fixtures file")
	}

	ts := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, string(jsonResponse))
			},
		),
	)
	defer ts.Close()

	u, err := url.Parse(ts.URL)
	if err != nil {
		t.Errorf("url parse error: %s", err)
	}

	promlogConfig := &promlog.Config{}
	logger := promlog.New(*&promlogConfig)

	testToken := "12345"
	expected := `
	# HELP syncthing_rest_system_connections_json_parse_failures Number of errors while parsing JSON.
	# TYPE syncthing_rest_system_connections_json_parse_failures counter
	syncthing_rest_system_connections_json_parse_failures 0
	# HELP syncthing_rest_system_connections_remote_device_connection_info Information of the connection with remote device. 1 - connected; 0 -disconnected.
	# TYPE syncthing_rest_system_connections_remote_device_connection_info gauge
	syncthing_rest_system_connections_remote_device_connection_info{address="",clientVersion="",crypto="",deviceID="AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-CCCCCCC",type=""} 0
	syncthing_rest_system_connections_remote_device_connection_info{address="10.1.0.1:59172",clientVersion="v1.10.0",crypto="TLS1.3-TLS_AES_128_GCM_SHA256",deviceID="AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA",type="tcp-server"} 1
	syncthing_rest_system_connections_remote_device_connection_info{address="10.2.0.2:22000",clientVersion="v1.10.0",crypto="TLS1.3-TLS_CHACHA20_POLY1305_SHA256",deviceID="AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-BBBBBBB",type="tcp-client"} 1
	# HELP syncthing_rest_system_connections_remote_device_in_bytes_total Total bytes received from remote device.
	# TYPE syncthing_rest_system_connections_remote_device_in_bytes_total gauge
	syncthing_rest_system_connections_remote_device_in_bytes_total{deviceID="AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA"} 1.119463078e+09
	syncthing_rest_system_connections_remote_device_in_bytes_total{deviceID="AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-BBBBBBB"} 1.0431439e+07
	syncthing_rest_system_connections_remote_device_in_bytes_total{deviceID="AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-CCCCCCC"} 0
	# HELP syncthing_rest_system_connections_remote_device_is_paused Is sync paused. 1 - paused, 0 - not paused.
	# TYPE syncthing_rest_system_connections_remote_device_is_paused gauge
	syncthing_rest_system_connections_remote_device_is_paused{deviceID="AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA"} 0
	syncthing_rest_system_connections_remote_device_is_paused{deviceID="AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-BBBBBBB"} 0
	syncthing_rest_system_connections_remote_device_is_paused{deviceID="AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-CCCCCCC"} 0
	# HELP syncthing_rest_system_connections_remote_device_out_bytes_total Total bytes transmitted to remote device.
	# TYPE syncthing_rest_system_connections_remote_device_out_bytes_total gauge
	syncthing_rest_system_connections_remote_device_out_bytes_total{deviceID="AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA"} 1.636080926e+10
	syncthing_rest_system_connections_remote_device_out_bytes_total{deviceID="AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-BBBBBBB"} 6.072477885e+09
	syncthing_rest_system_connections_remote_device_out_bytes_total{deviceID="AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-CCCCCCC"} 0
	# HELP syncthing_rest_system_connections_total_in_bytes_total Total bytes received to device.
	# TYPE syncthing_rest_system_connections_total_in_bytes_total gauge
	syncthing_rest_system_connections_total_in_bytes_total 2.458124935e+09
	# HELP syncthing_rest_system_connections_total_out_bytes_total Total bytes transmitted from device.
	# TYPE syncthing_rest_system_connections_total_out_bytes_total gauge
	syncthing_rest_system_connections_total_out_bytes_total 3.3180203458e+10
	# HELP syncthing_rest_system_connections_total_scrapes Current total Syncthings scrapes.
	# TYPE syncthing_rest_system_connections_total_scrapes counter
	syncthing_rest_system_connections_total_scrapes 1
	# HELP syncthing_rest_system_connections_up Was the last scrape of the Syncting system connections endpoint successful.
	# TYPE syncthing_rest_system_connections_up gauge
	syncthing_rest_system_connections_up 1
	`
	err = testutil.CollectAndCompare(
		NewSCReport(logger, &http.Client{}, u, &testToken),
		strings.NewReader(expected),
	)

	if err != nil {
		t.Errorf("NewSCReportError %s", err)
	}

}
