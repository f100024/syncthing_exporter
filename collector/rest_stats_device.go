package collector

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

// StatsDeviceResponseNumericalMetrics defines struct for numeric metrics
type StatsDeviceResponseNumericalMetrics struct {
	Type  prometheus.ValueType
	Desc  *prometheus.Desc
	Value func(v float64) float64
}

// StatsDeviceResponse defines collector struct.
type StatsDeviceResponse struct {
	logger *slog.Logger
	client *http.Client
	url    *url.URL
	token  *string

	up                              prometheus.Gauge
	totalScrapes, jsonParseFailures prometheus.Counter
	numericalMetrics                map[string]*StatsDeviceResponseNumericalMetrics
}

// NewStatsDeviceReport generate report from device
func NewStatsDeviceReport(logger *slog.Logger, client *http.Client, url *url.URL, token *string) *StatsDeviceResponse {
	subsystem := "rest_stats_device"

	return &StatsDeviceResponse{
		logger: logger,
		client: client,
		url:    url,
		token:  token,

		up: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: prometheus.BuildFQName(namespace, subsystem, "up"),
			Help: "Was the last scrape of the Syncting system connections endpoint successful.",
		}),
		totalScrapes: prometheus.NewCounter(prometheus.CounterOpts{
			Name: prometheus.BuildFQName(namespace, subsystem, "total_scrapes"),
			Help: "Current total Syncthings scrapes.",
		}),
		jsonParseFailures: prometheus.NewCounter(prometheus.CounterOpts{
			Name: prometheus.BuildFQName(namespace, subsystem, "json_parse_failures"),
			Help: "Number of errors while parsing JSON.",
		}),
		numericalMetrics: map[string]*StatsDeviceResponseNumericalMetrics{
			"last_connection_duration": {
				Type: prometheus.GaugeValue,
				Desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "last_connection_duration"),
					"Duration of last connection with remote device in seconds. If value less 0, means value was missed in syncthing response.",
					[]string{"deviceID"}, nil),
				Value: func(v float64) float64 {
					return v
				},
			},
			"last_connection_timestamp": {
				Type: prometheus.GaugeValue,
				Desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "last_connection_timestamp"),
					"Timestamp since last connection with remote device expressed in Unix epoch",
					[]string{"deviceID"}, nil),
				Value: func(v float64) float64 {
					return v
				},
			},
		},
	}

}

// Describe set Prometheus metrics descriptions.
func (c *StatsDeviceResponse) Describe(ch chan<- *prometheus.Desc) {

	for _, metric := range c.numericalMetrics {
		ch <- metric.Desc
	}

	ch <- c.up.Desc()
	ch <- c.totalScrapes.Desc()
	ch <- c.jsonParseFailures.Desc()
}

func (c *StatsDeviceResponse) fetchDataAndDecode() (map[string]interface{}, error) {
	chr := make(map[string]interface{})

	u := *c.url
	url, _ := u.Parse("/rest/stats/device")

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

// Collect collects Syncthing metrics from /rest/stats/device.
func (c *StatsDeviceResponse) Collect(ch chan<- prometheus.Metric) {
	var err error

	c.totalScrapes.Inc()
	defer func() {
		ch <- c.up
		ch <- c.totalScrapes
		ch <- c.jsonParseFailures
	}()
	statsDeviceResponse, err := c.fetchDataAndDecode()
	if err != nil {
		c.up.Set(0)
		c.logger.Info(fmt.Sprintf("%s: %s", "failed to fetch and decode data", err))
		return
	}
	c.up.Set(1)

	for deviceID, deviceData := range statsDeviceResponse {
		deviceDataAssertion := deviceData.(map[string]interface{})

		if deviceDataAssertion["lastConnectionDurationS"] == nil {
			deviceDataAssertion["lastConnectionDurationS"] = -1.0
		}

		ch <- prometheus.MustNewConstMetric(
			c.numericalMetrics["last_connection_duration"].Desc,
			c.numericalMetrics["last_connection_duration"].Type,
			c.numericalMetrics["last_connection_duration"].Value(deviceDataAssertion["lastConnectionDurationS"].(float64)),
			deviceID,
		)
		thetime, err := time.Parse(time.RFC3339, deviceDataAssertion["lastSeen"].(string))
		if err != nil {
			c.logger.Warn(fmt.Sprintf("%s: %s", "failed to parse timestamp", err))
			return
		}
		ch <- prometheus.MustNewConstMetric(
			c.numericalMetrics["last_connection_timestamp"].Desc,
			c.numericalMetrics["last_connection_timestamp"].Type,
			c.numericalMetrics["last_connection_timestamp"].Value(float64(thetime.Unix())),
			deviceID,
		)
	}

}
