package collector

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/syncthing/syncthing/lib/config"
)

// CDResponseBoolMetrics defines struct for boolean metrics
type CDResponseBoolMetrics struct {
	Type  prometheus.ValueType
	Desc  *prometheus.Desc
	Value func(v bool) float64
}

// ConfigDevicesResponse defines collector struct.
type ConfigDevicesResponse struct {
	logger *slog.Logger
	client *http.Client
	url    *url.URL
	token  *string

	up                              prometheus.Gauge
	totalScrapes, jsonParseFailures prometheus.Counter
	boolMetrics                     map[string]*CDResponseBoolMetrics
}

// NewConfigDevicesReport returns a new Collector exposing ConfigDevicesResponse
func NewConfigDevicesReport(logger *slog.Logger, client *http.Client, url *url.URL, token *string) *ConfigDevicesResponse {
	subsystem := "rest_config_devices"
	return &ConfigDevicesResponse{
		logger: logger,
		client: client,
		url:    url,
		token:  token,
		up: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: prometheus.BuildFQName(namespace, subsystem, "up"),
			Help: "Was the last scrape of the Syncthing endpoint successful.",
		}),
		totalScrapes: prometheus.NewCounter(prometheus.CounterOpts{
			Name: prometheus.BuildFQName(namespace, subsystem, "total_scrapes"),
			Help: "Current total Syncthings scrapes.",
		}),
		jsonParseFailures: prometheus.NewCounter(prometheus.CounterOpts{
			Name: prometheus.BuildFQName(namespace, subsystem, "json_parse_failures"),
			Help: "Number of errors while parsing JSON.",
		}),
		boolMetrics: map[string]*CDResponseBoolMetrics{
			"is_remote_device_paused": {
				Type: prometheus.GaugeValue,
				Desc: prometheus.NewDesc(prometheus.BuildFQName(namespace, subsystem, "is_remote_device_paused"),
					"Is remote device paused and other device information. 1 - paused, 0 - not paused",
					[]string{"deviceID", "name", "addresses", "compression", "certName", "introducedBy", "allowedNetworks"}, nil),
				Value: func(v bool) float64 {
					return bool2float64(v)
				},
			},
			"is_remote_device_introducer": {
				Type: prometheus.GaugeValue,
				Desc: prometheus.NewDesc(prometheus.BuildFQName(namespace, subsystem, "is_remote_device_introducer"),
					"Is remote device marked as introducer. 1 - introducer, 0 - no.",
					[]string{"deviceID", "name"}, nil),
				Value: func(v bool) float64 {
					return bool2float64(v)
				},
			},
			"is_remote_device_skip_introduction_removals": {
				Type: prometheus.GaugeValue,
				Desc: prometheus.NewDesc(prometheus.BuildFQName(namespace, subsystem, "is_remote_device_skip_introduction_removals"),
					"Is remote device skip introduction removals",
					[]string{"deviceID", "name"}, nil),
				Value: func(v bool) float64 {
					return bool2float64(v)
				},
			},
			"is_remote_device_auto_accept_folders": {
				Type: prometheus.GaugeValue,
				Desc: prometheus.NewDesc(prometheus.BuildFQName(namespace, subsystem, "is_remote_device_auto_accept_folders"),
					"Is remote device auto accept folders",
					[]string{"deviceID", "name"}, nil),
				Value: func(v bool) float64 {
					return bool2float64(v)
				},
			},
			"is_remote_device_untrusted": {
				Type: prometheus.GaugeValue,
				Desc: prometheus.NewDesc(prometheus.BuildFQName(namespace, subsystem, "is_remote_device_untrusted"),
					"Is remote device auto accept folders",
					[]string{"deviceID", "name"}, nil),
				Value: func(v bool) float64 {
					return bool2float64(v)
				},
			},
		},
	}
}

// Describe set Prometheus metrics descriptions.
func (c *ConfigDevicesResponse) Describe(ch chan<- *prometheus.Desc) {
	for _, metric := range c.boolMetrics {
		ch <- metric.Desc
	}

	ch <- c.up.Desc()
	ch <- c.totalScrapes.Desc()
	ch <- c.jsonParseFailures.Desc()
}

func (c *ConfigDevicesResponse) fetchDataAndDecode() ([]config.DeviceConfiguration, error) {
	var chr []config.DeviceConfiguration

	u := *c.url
	url, _ := u.Parse("/rest/config/devices")

	h := make(http.Header)
	h["X-API-Key"] = []string{*c.token}

	request := &http.Request{
		URL:    url,
		Header: h,
	}

	res, err := c.client.Do(request)
	if err != nil {
		return chr, fmt.Errorf("failed to get data from %s://%s:%s%s: %s",
			u.Scheme, u.Hostname(), u.Port(), u.Path, err)
	}

	defer func() {
		err = res.Body.Close()
		if err != nil {
			c.logger.Warn(fmt.Sprintf("%s: %s", "failed to close http.Client", err))
		}
	}()

	if res.StatusCode != http.StatusOK {
		return chr, fmt.Errorf("HTTP Request failed with code %d", res.StatusCode)
	}

	if err := json.NewDecoder(res.Body).Decode(&chr); err != nil {
		c.jsonParseFailures.Inc()
		return chr, err
	}

	return chr, nil
}

// Collect collects Syncthing metrics from /rest/config/devices.
func (c *ConfigDevicesResponse) Collect(ch chan<- prometheus.Metric) {
	var err error
	c.totalScrapes.Inc()
	defer func() {
		ch <- c.up
		ch <- c.totalScrapes
		ch <- c.jsonParseFailures
	}()

	CDResponse, err := c.fetchDataAndDecode()

	if err != nil {
		c.up.Set(0)
		c.logger.Warn(fmt.Sprintf("%s: %s", "failed to fetch and decode data", err))
		return
	}
	c.up.Set(1)

	for _, deviceProperty := range CDResponse {
		ch <- prometheus.MustNewConstMetric(
			c.boolMetrics["is_remote_device_paused"].Desc,
			c.boolMetrics["is_remote_device_paused"].Type,
			c.boolMetrics["is_remote_device_paused"].Value(deviceProperty.Paused),
			deviceProperty.DeviceID.String(), deviceProperty.Name, strings.Join(deviceProperty.Addresses, ","),
			deviceProperty.Compression.ToProtocol().Enum().String(), deviceProperty.CertName, deviceProperty.IntroducedBy.GoString(),
			strings.Join(deviceProperty.AllowedNetworks, ","),
		)
		ch <- prometheus.MustNewConstMetric(
			c.boolMetrics["is_remote_device_introducer"].Desc,
			c.boolMetrics["is_remote_device_introducer"].Type,
			c.boolMetrics["is_remote_device_introducer"].Value(deviceProperty.Introducer),
			deviceProperty.DeviceID.String(), deviceProperty.Name,
		)
		ch <- prometheus.MustNewConstMetric(
			c.boolMetrics["is_remote_device_skip_introduction_removals"].Desc,
			c.boolMetrics["is_remote_device_skip_introduction_removals"].Type,
			c.boolMetrics["is_remote_device_skip_introduction_removals"].Value(deviceProperty.SkipIntroductionRemovals),
			deviceProperty.DeviceID.String(), deviceProperty.Name,
		)
		ch <- prometheus.MustNewConstMetric(
			c.boolMetrics["is_remote_device_auto_accept_folders"].Desc,
			c.boolMetrics["is_remote_device_auto_accept_folders"].Type,
			c.boolMetrics["is_remote_device_auto_accept_folders"].Value(deviceProperty.AutoAcceptFolders),
			deviceProperty.DeviceID.String(), deviceProperty.Name,
		)
		ch <- prometheus.MustNewConstMetric(
			c.boolMetrics["is_remote_device_untrusted"].Desc,
			c.boolMetrics["is_remote_device_untrusted"].Type,
			c.boolMetrics["is_remote_device_untrusted"].Value(deviceProperty.Untrusted),
			deviceProperty.DeviceID.String(), deviceProperty.Name,
		)
	}
}
