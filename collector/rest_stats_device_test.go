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

func TestNewStatsDeviceReport(t *testing.T) {

	jsonResponse, _ := ioutil.ReadFile("fixtures/rest_stats_device_response.json")

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

	promlogConfig := &promlog.Config{}
	logger := promlog.New(promlogConfig)

	testToken := "12345"
	expected := `
	# HELP syncthing_rest_stats_device_json_parse_failures Number of errors while parsing JSON.
	# TYPE syncthing_rest_stats_device_json_parse_failures counter
	syncthing_rest_stats_device_json_parse_failures 0
	# HELP syncthing_rest_stats_device_last_connection_duration Duration of last connection with remote device in seconds.
	# TYPE syncthing_rest_stats_device_last_connection_duration gauge
	syncthing_rest_stats_device_last_connection_duration{deviceID="AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA",lastSeen="2077-12-18T00:00:50.3810375-08:00"} 0
	syncthing_rest_stats_device_last_connection_duration{deviceID="AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-BBBBBBB",lastSeen="2077-04-13T04:50:25-07:00"} 819990.3432191
	syncthing_rest_stats_device_last_connection_duration{deviceID="AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-CCCCCCC",lastSeen="1984-12-31T16:00:00-08:00"} 0
	# HELP syncthing_rest_stats_device_total_scrapes Current total Syncthings scrapes.
	# TYPE syncthing_rest_stats_device_total_scrapes counter
	syncthing_rest_stats_device_total_scrapes 1
	# HELP syncthing_rest_stats_device_up Was the last scrape of the Syncting system connections endpoint successful.
	# TYPE syncthing_rest_stats_device_up gauge
	syncthing_rest_stats_device_up 1
	`

	err = testutil.CollectAndCompare(
		NewStatsDeviceReport(logger, HttpClient, u, &testToken),
		strings.NewReader(expected),
	)

	if err != nil {
		t.Errorf("NewStatsDeviceReportError %s", err)
	}
}
