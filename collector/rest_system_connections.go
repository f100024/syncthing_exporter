package collector

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
)

// SCResponseBoolMetrics defines struct for boolean metrics
type SCResponseBoolMetrics struct {
	Type  prometheus.ValueType
	Desc  *prometheus.Desc
	Value func(v bool) float64
}

// SCResponseNumericalMetrics defines struct for numeric metrics
type SCResponseNumericalMetrics struct {
	Type  prometheus.ValueType
	Desc  *prometheus.Desc
	Value func(v int) float64
}

// SCResponseTotalNumericalMetrics defines struct for total metrics
type SCResponseTotalNumericalMetrics struct {
	Type  prometheus.ValueType
	Desc  *prometheus.Desc
	Value func(v int) float64
}

// SCResponse defines collector struct.
type SCResponse struct {
	logger *log.Logger
	client *http.Client
	url    *url.URL
	token  *string

	up                              prometheus.Gauge
	totalScrapes, jsonParseFailures prometheus.Counter
	boolMetrics                     map[string]*SCResponseBoolMetrics
	numericalMetrics                map[string]*SCResponseNumericalMetrics
	totalNumericalMetrics           map[string]*SCResponseTotalNumericalMetrics
}

// NewSCReport returns a new Collector exposing SVCResponse
func NewSCReport(logger log.Logger, client *http.Client, url *url.URL, token *string) *SCResponse {
	subsystem := "rest_system_connections"

	return &SCResponse{
		logger: &logger,
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
		boolMetrics: map[string]*SCResponseBoolMetrics{
			"remote_device_connection_info": {
				Type: prometheus.GaugeValue,
				Desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "remote_device_connection_info"),
					"Information of the connection with remote device. 1 - connected; 0 -disconnected.",
					[]string{"deviceID", "clientVersion", "crypto", "type", "address"},
					nil),
				Value: func(v bool) float64 {
					return bool2float64(v)
				},
			},
			"remote_device_is_paused": {
				Type: prometheus.GaugeValue,
				Desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "remote_device_is_paused"),
					"Is sync paused. 1 - paused, 0 - not paused.",
					[]string{"deviceID"},
					nil),
				Value: func(v bool) float64 {
					return bool2float64(v)
				},
			},
		},

		numericalMetrics: map[string]*SCResponseNumericalMetrics{
			"remote_device_in_bytes_total": {
				Type: prometheus.GaugeValue,
				Desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "remote_device_in_bytes_total"),
					"Total bytes received from remote device.",
					[]string{"deviceID"},
					nil),
				Value: func(v int) float64 {
					return float64(v)
				},
			},

			"remote_device_out_bytes_total": {
				Type: prometheus.GaugeValue,
				Desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "remote_device_out_bytes_total"),
					"Total bytes transmitted to remote device.",
					[]string{"deviceID"},
					nil),
				Value: func(v int) float64 {
					return float64(v)
				},
			},
		},

		totalNumericalMetrics: map[string]*SCResponseTotalNumericalMetrics{
			"total_in_bytes_total": {
				Type: prometheus.GaugeValue,
				Desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "total_in_bytes_total"),
					"Total bytes received to device.", nil, nil),
				Value: func(v int) float64 {
					return float64(v)
				},
			},
			"total_out_bytes_total": {
				Type: prometheus.GaugeValue,
				Desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "total_out_bytes_total"),
					"Total bytes transmitted from device.", nil, nil),
				Value: func(v int) float64 {
					return float64(v)
				},
			},
		},
	}

}

// Describe set Prometheus metrics descriptions.
func (c *SCResponse) Describe(ch chan<- *prometheus.Desc) {

	for _, metric := range c.boolMetrics {
		ch <- metric.Desc
	}

	for _, metric := range c.numericalMetrics {
		ch <- metric.Desc
	}

	for _, metric := range c.totalNumericalMetrics {
		ch <- metric.Desc
	}

	ch <- c.up.Desc()
	ch <- c.totalScrapes.Desc()
	ch <- c.jsonParseFailures.Desc()
}

func (c *SCResponse) fetchDataAndDecode() (SystemConnectionsResponse, error) {
	var chr SystemConnectionsResponse

	u := *c.url
	url, _ := u.Parse("/rest/system/connections")

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
			level.Warn(*c.logger).Log("msg", "failed to close http.Client", "err", err)
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

// Collect collects Syncthing metrics from /rest/system/connections.
func (c *SCResponse) Collect(ch chan<- prometheus.Metric) {
	var err error

	c.totalScrapes.Inc()
	defer func() {
		ch <- c.up
		ch <- c.totalScrapes
		ch <- c.jsonParseFailures
	}()

	SCResponse, err := c.fetchDataAndDecode()
	if err != nil {
		c.up.Set(0)
		level.Warn(*c.logger).Log(
			"msg", "failed to fetch and decode data",
			"err", err,
		)
		return
	}
	c.up.Set(1)

	for deviceID, deviceData := range SCResponse.CS {

		ch <- prometheus.MustNewConstMetric(
			c.boolMetrics["remote_device_connection_info"].Desc,
			c.boolMetrics["remote_device_connection_info"].Type,
			c.boolMetrics["remote_device_connection_info"].Value(deviceData.Connected),
			deviceID, deviceData.ClientVersion, deviceData.CryptoConfig, deviceData.Type, deviceData.IPAddress,
		)
		ch <- prometheus.MustNewConstMetric(
			c.boolMetrics["remote_device_is_paused"].Desc,
			c.boolMetrics["remote_device_is_paused"].Type,
			c.boolMetrics["remote_device_is_paused"].Value(deviceData.Paused),
			deviceID,
		)

		ch <- prometheus.MustNewConstMetric(
			c.numericalMetrics["remote_device_in_bytes_total"].Desc,
			c.numericalMetrics["remote_device_in_bytes_total"].Type,
			c.numericalMetrics["remote_device_in_bytes_total"].Value(deviceData.InBytesTotal),
			deviceID,
		)
		ch <- prometheus.MustNewConstMetric(
			c.numericalMetrics["remote_device_out_bytes_total"].Desc,
			c.numericalMetrics["remote_device_out_bytes_total"].Type,
			c.numericalMetrics["remote_device_out_bytes_total"].Value(deviceData.OutBytesTotal),
			deviceID,
		)

	}

	ch <- prometheus.MustNewConstMetric(
		c.totalNumericalMetrics["total_in_bytes_total"].Desc,
		c.totalNumericalMetrics["total_in_bytes_total"].Type,
		c.totalNumericalMetrics["total_in_bytes_total"].Value(SCResponse.Total.InBytesTotal),
	)
	ch <- prometheus.MustNewConstMetric(
		c.totalNumericalMetrics["total_out_bytes_total"].Desc,
		c.totalNumericalMetrics["total_out_bytes_total"].Type,
		c.totalNumericalMetrics["total_out_bytes_total"].Value(SCResponse.Total.OutBytesTotal),
	)
}
