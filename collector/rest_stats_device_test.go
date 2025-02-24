package collector

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/prometheus/common/promslog"
)

func TestNewStatsDeviceReport(t *testing.T) {

	jsonResponse, _ := os.ReadFile("fixtures/rest_stats_device_response.json")

	ts := httptest.NewTLSServer(
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

	promslogConfig := &promslog.Config{}
	logger := promslog.New(promslogConfig)

	testToken := "12345"
	expected := `
	# HELP syncthing_rest_stats_device_json_parse_failures Number of errors while parsing JSON.
	# TYPE syncthing_rest_stats_device_json_parse_failures counter
	syncthing_rest_stats_device_json_parse_failures 0
	# HELP syncthing_rest_stats_device_last_connection_duration Duration of last connection with remote device in seconds. If value less 0, means value was missed in syncthing response.
	# TYPE syncthing_rest_stats_device_last_connection_duration gauge
	syncthing_rest_stats_device_last_connection_duration{deviceID="AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA"} 0
	syncthing_rest_stats_device_last_connection_duration{deviceID="AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-BBBBBBB"} 819990.3432191
	syncthing_rest_stats_device_last_connection_duration{deviceID="AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-CCCCCCC"} 0
	# HELP syncthing_rest_stats_device_last_connection_timestamp Timestamp since last connection with remote device expressed in Unix epoch
	# TYPE syncthing_rest_stats_device_last_connection_timestamp gauge
	syncthing_rest_stats_device_last_connection_timestamp{deviceID="AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA"} 3.40704005e+09
	syncthing_rest_stats_device_last_connection_timestamp{deviceID="AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-BBBBBBB"} 3.385540225e+09
	syncthing_rest_stats_device_last_connection_timestamp{deviceID="AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-CCCCCCC"} 4.733856e+08
	# HELP syncthing_rest_stats_device_total_scrapes Current total Syncthings scrapes.
	# TYPE syncthing_rest_stats_device_total_scrapes counter
	syncthing_rest_stats_device_total_scrapes 1
	# HELP syncthing_rest_stats_device_up Was the last scrape of the Syncting system connections endpoint successful.
	# TYPE syncthing_rest_stats_device_up gauge
	syncthing_rest_stats_device_up 1
	`

	err = testutil.CollectAndCompare(
		NewStatsDeviceReport(logger, HTTPClient, u, &testToken),
		strings.NewReader(expected),
	)

	if err != nil {
		t.Errorf("NewStatsDeviceReportError %s", err)
	}
}

func TestStatsDeviceReportWithNoLastConnectionDuration(t *testing.T) {

	jsonResponse, _ := os.ReadFile("fixtures/rest_stats_device_response_missed_lastConnectionDurationS.json")

	ts := httptest.NewTLSServer(
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

	promslogConfig := &promslog.Config{}
	logger := promslog.New(promslogConfig)

	testToken := "12345"
	expected := `
	# HELP syncthing_rest_stats_device_json_parse_failures Number of errors while parsing JSON.
	# TYPE syncthing_rest_stats_device_json_parse_failures counter
	syncthing_rest_stats_device_json_parse_failures 0
	# HELP syncthing_rest_stats_device_last_connection_duration Duration of last connection with remote device in seconds. If value less 0, means value was missed in syncthing response.
	# TYPE syncthing_rest_stats_device_last_connection_duration gauge
	syncthing_rest_stats_device_last_connection_duration{deviceID="AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA"} 0
	syncthing_rest_stats_device_last_connection_duration{deviceID="AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-BBBBBBB"} 819990.3432191
	syncthing_rest_stats_device_last_connection_duration{deviceID="AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-CCCCCCC"} -1
	# HELP syncthing_rest_stats_device_last_connection_timestamp Timestamp since last connection with remote device expressed in Unix epoch
	# TYPE syncthing_rest_stats_device_last_connection_timestamp gauge
	syncthing_rest_stats_device_last_connection_timestamp{deviceID="AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA"} 3.40704005e+09
	syncthing_rest_stats_device_last_connection_timestamp{deviceID="AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-BBBBBBB"} 3.385540225e+09
	syncthing_rest_stats_device_last_connection_timestamp{deviceID="AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-CCCCCCC"} 4.733856e+08
	# HELP syncthing_rest_stats_device_total_scrapes Current total Syncthings scrapes.
	# TYPE syncthing_rest_stats_device_total_scrapes counter
	syncthing_rest_stats_device_total_scrapes 1
	# HELP syncthing_rest_stats_device_up Was the last scrape of the Syncting system connections endpoint successful.
	# TYPE syncthing_rest_stats_device_up gauge
	syncthing_rest_stats_device_up 1
	`

	err = testutil.CollectAndCompare(
		NewStatsDeviceReport(logger, HTTPClient, u, &testToken),
		strings.NewReader(expected),
	)

	if err != nil {
		t.Errorf("NewStatsDeviceReportError %s", err)
	}
}
