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

func TestNewConfigDevicesReport(t *testing.T) {
	jsonResponse, _ := os.ReadFile("fixtures/rest_config_devices_response.json")

	ts := httptest.NewTLSServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, _ *http.Request) {
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
	# HELP syncthing_rest_config_devices_is_remote_device_auto_accept_folders Is remote device auto accept folders
	# TYPE syncthing_rest_config_devices_is_remote_device_auto_accept_folders gauge
	syncthing_rest_config_devices_is_remote_device_auto_accept_folders{deviceID="EVOOQ2B-NZSX5XA-MGCOKJ7-M53F43K-HBKQOAS-7S5POY5-DSCKCTH-AI4T2AJ",name="host-1"} 0
	syncthing_rest_config_devices_is_remote_device_auto_accept_folders{deviceID="KXPN5AW-5EAJEA3-F7GXDVS-HPTV57U-NWDMFCS-SCTDFQF-CGGQTQD-FV2HAAZ",name="host-2"} 1
	# HELP syncthing_rest_config_devices_is_remote_device_introducer Is remote device marked as introducer. 1 - introducer, 0 - no.
	# TYPE syncthing_rest_config_devices_is_remote_device_introducer gauge
	syncthing_rest_config_devices_is_remote_device_introducer{deviceID="EVOOQ2B-NZSX5XA-MGCOKJ7-M53F43K-HBKQOAS-7S5POY5-DSCKCTH-AI4T2AJ",name="host-1"} 0
	syncthing_rest_config_devices_is_remote_device_introducer{deviceID="KXPN5AW-5EAJEA3-F7GXDVS-HPTV57U-NWDMFCS-SCTDFQF-CGGQTQD-FV2HAAZ",name="host-2"} 0
	# HELP syncthing_rest_config_devices_is_remote_device_paused Is remote device paused and other device information. 1 - paused, 0 - not paused
	# TYPE syncthing_rest_config_devices_is_remote_device_paused gauge
	syncthing_rest_config_devices_is_remote_device_paused{addresses="dynamic",allowedNetworks="",certName="",compression="COMPRESSION_METADATA",deviceID="EVOOQ2B-NZSX5XA-MGCOKJ7-M53F43K-HBKQOAS-7S5POY5-DSCKCTH-AI4T2AJ",introducedBy="",name="host-1"} 0
	syncthing_rest_config_devices_is_remote_device_paused{addresses="tcp://192.168.1.2:22000",allowedNetworks="",certName="",compression="COMPRESSION_METADATA",deviceID="KXPN5AW-5EAJEA3-F7GXDVS-HPTV57U-NWDMFCS-SCTDFQF-CGGQTQD-FV2HAAZ",introducedBy="",name="host-2"} 0
	# HELP syncthing_rest_config_devices_is_remote_device_skip_introduction_removals Is remote device skip introduction removals
	# TYPE syncthing_rest_config_devices_is_remote_device_skip_introduction_removals gauge
	syncthing_rest_config_devices_is_remote_device_skip_introduction_removals{deviceID="EVOOQ2B-NZSX5XA-MGCOKJ7-M53F43K-HBKQOAS-7S5POY5-DSCKCTH-AI4T2AJ",name="host-1"} 0
	syncthing_rest_config_devices_is_remote_device_skip_introduction_removals{deviceID="KXPN5AW-5EAJEA3-F7GXDVS-HPTV57U-NWDMFCS-SCTDFQF-CGGQTQD-FV2HAAZ",name="host-2"} 0
	# HELP syncthing_rest_config_devices_is_remote_device_untrusted Is remote device auto accept folders
	# TYPE syncthing_rest_config_devices_is_remote_device_untrusted gauge
	syncthing_rest_config_devices_is_remote_device_untrusted{deviceID="EVOOQ2B-NZSX5XA-MGCOKJ7-M53F43K-HBKQOAS-7S5POY5-DSCKCTH-AI4T2AJ",name="host-1"} 0
	syncthing_rest_config_devices_is_remote_device_untrusted{deviceID="KXPN5AW-5EAJEA3-F7GXDVS-HPTV57U-NWDMFCS-SCTDFQF-CGGQTQD-FV2HAAZ",name="host-2"} 0
	# HELP syncthing_rest_config_devices_json_parse_failures Number of errors while parsing JSON.
	# TYPE syncthing_rest_config_devices_json_parse_failures counter
	syncthing_rest_config_devices_json_parse_failures 0
	# HELP syncthing_rest_config_devices_total_scrapes Current total Syncthings scrapes.
	# TYPE syncthing_rest_config_devices_total_scrapes counter
	syncthing_rest_config_devices_total_scrapes 1
	# HELP syncthing_rest_config_devices_up Was the last scrape of the Syncthing endpoint successful.
	# TYPE syncthing_rest_config_devices_up gauge
	syncthing_rest_config_devices_up 1
	`
	err = testutil.CollectAndCompare(
		NewConfigDevicesReport(logger, HTTPClient, u, &testToken),
		strings.NewReader(expected),
	)

	if err != nil {
		t.Errorf("NewConfigDevicesReport %s", err)
	}
}
