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

func TestNewSVCReport(t *testing.T) {

	jsonResponse, _ := os.ReadFile("fixtures/rest_svc_report_response.json")

	ts := httptest.NewTLSServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, string(jsonResponse))
			},
		),
	)
	defer ts.Close()

	u, _ := url.Parse(ts.URL)

	promslogConfig := &promslog.Config{}
	logger := promslog.New(promslogConfig)

	testToken := "12345"
	expected := `
	# HELP syncthing_rest_svc_report_hash_performance Hash Performance value
	# TYPE syncthing_rest_svc_report_hash_performance gauge
	syncthing_rest_svc_report_hash_performance 187.47
	# HELP syncthing_rest_svc_report_json_parse_failures Number of errors while parsing JSON.
	# TYPE syncthing_rest_svc_report_json_parse_failures counter
	syncthing_rest_svc_report_json_parse_failures 0
	# HELP syncthing_rest_svc_report_memory_size Node memory size in megabytes
	# TYPE syncthing_rest_svc_report_memory_size gauge
	syncthing_rest_svc_report_memory_size 8191
	# HELP syncthing_rest_svc_report_memory_usage_mb Memory usage by syncthc in MB
	# TYPE syncthing_rest_svc_report_memory_usage_mb gauge
	syncthing_rest_svc_report_memory_usage_mb 219
	# HELP syncthing_rest_svc_report_node_info Node's string information
	# TYPE syncthing_rest_svc_report_node_info gauge
	syncthing_rest_svc_report_node_info{longVersion="syncthing v1.9.0 \"Fermium Flea\" (go1.15.1 windows-amd64) teamcity@build.syncthing.net 2020-08-28 05:48:25 UTC",platform="windows-amd64",uniqueID="d9uZew76",version="v1.9.0"} 1
	# HELP syncthing_rest_svc_report_num_cpu Number of node CPU
	# TYPE syncthing_rest_svc_report_num_cpu gauge
	syncthing_rest_svc_report_num_cpu 4
	# HELP syncthing_rest_svc_report_number_devices Number of devices in sync
	# TYPE syncthing_rest_svc_report_number_devices gauge
	syncthing_rest_svc_report_number_devices 3
	# HELP syncthing_rest_svc_report_number_folders Number of folders in sync
	# TYPE syncthing_rest_svc_report_number_folders gauge
	syncthing_rest_svc_report_number_folders 2
	# HELP syncthing_rest_svc_report_sha256_performance SHA256 Performance value
	# TYPE syncthing_rest_svc_report_sha256_performance gauge
	syncthing_rest_svc_report_sha256_performance 216.93
	# HELP syncthing_rest_svc_report_total_data_in_MB Total data in megabytes
	# TYPE syncthing_rest_svc_report_total_data_in_MB gauge
	syncthing_rest_svc_report_total_data_in_MB 3.509532e+06
	# HELP syncthing_rest_svc_report_total_files Total number of files
	# TYPE syncthing_rest_svc_report_total_files gauge
	syncthing_rest_svc_report_total_files 509466
	# HELP syncthing_rest_svc_report_total_scrapes Current total Syncthings scrapes.
	# TYPE syncthing_rest_svc_report_total_scrapes counter
	syncthing_rest_svc_report_total_scrapes 1
	# HELP syncthing_rest_svc_report_up Was the last scrape of the Syncthing endpoint successful.
	# TYPE syncthing_rest_svc_report_up gauge
	syncthing_rest_svc_report_up 1
	# HELP syncthing_rest_svc_report_uptime Syncthing uptime in seconds
	# TYPE syncthing_rest_svc_report_uptime gauge
	syncthing_rest_svc_report_uptime 1.276967e+06
	`

	err := testutil.CollectAndCompare(
		NewSVCReport(logger, HTTPClient, u, &testToken),
		strings.NewReader(expected),
	)

	if err != nil {
		t.Errorf("NewSVCReportError %s", err)
	}
}

func TestFailedNewSVCReport(t *testing.T) {

	u, _ := url.Parse("http://wrong-url")
	promlogConfig := &promslog.Config{}
	logger := promslog.New(promlogConfig)

	testToken := "12345"
	expected := `
	# HELP syncthing_rest_svc_report_json_parse_failures Number of errors while parsing JSON.
	# TYPE syncthing_rest_svc_report_json_parse_failures counter
	syncthing_rest_svc_report_json_parse_failures 0
	# HELP syncthing_rest_svc_report_total_scrapes Current total Syncthings scrapes.
	# TYPE syncthing_rest_svc_report_total_scrapes counter
	syncthing_rest_svc_report_total_scrapes 1
	# HELP syncthing_rest_svc_report_up Was the last scrape of the Syncthing endpoint successful.
	# TYPE syncthing_rest_svc_report_up gauge
	syncthing_rest_svc_report_up 0
	`
	err := testutil.CollectAndCompare(
		NewSVCReport(logger, HTTPClient, u, &testToken),
		strings.NewReader(expected),
	)

	if err != nil {
		t.Errorf("NewSVCReportError %s", err)
	}
}
