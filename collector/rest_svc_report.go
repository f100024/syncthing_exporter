package collector

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/syncthing/syncthing/lib/ur/contract"
)

type svcMetric struct {
	Type  prometheus.ValueType
	Desc  *prometheus.Desc
	Value func(svcMetric *contract.Report) float64
}

type svcData struct {
	Type  prometheus.ValueType
	Desc  *prometheus.Desc
	Value func(svcData *contract.Report) float64
}

// SVCResponse defines collector struct.
type SVCResponse struct {
	logger *log.Logger
	client *http.Client
	url    *url.URL
	token  *string

	up                              prometheus.Gauge
	totalScrapes, jsonParseFailures prometheus.Counter
	metrics                         []*svcMetric
	data                            *svcData
}

// NewSVCReport returns a new Collector exposing SVCResponse
func NewSVCReport(logger log.Logger, client *http.Client, url *url.URL, token *string) *SVCResponse {
	subsystem := "rest_svc_report"

	return &SVCResponse{
		logger: &logger,
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

		metrics: []*svcMetric{
			{
				Type: prometheus.GaugeValue,
				Desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "number_folders"),
					"Number of folders in sync", nil, nil,
				),
				Value: func(svcMetric *contract.Report) float64 {
					return float64(svcMetric.NumFolders)
				},
			},
			{
				Type: prometheus.GaugeValue,
				Desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "number_devices"),
					"Number of devices in sync", nil, nil,
				),
				Value: func(svcMetric *contract.Report) float64 {
					return float64(svcMetric.NumDevices)
				},
			},
			{
				Type: prometheus.GaugeValue,
				Desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "total_files"),
					"Total number of files", nil, nil,
				),
				Value: func(svcMetric *contract.Report) float64 {
					return float64(svcMetric.TotFiles)
				},
			},
			{
				Type: prometheus.GaugeValue,
				Desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "total_data_in_MB"),
					"Total data in megabytes", nil, nil,
				),
				Value: func(svcMetric *contract.Report) float64 {
					return float64(svcMetric.TotMiB)
				},
			},
			{
				Type: prometheus.GaugeValue,
				Desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "memory_usage_mb"),
					"Memory usage by syncthc in MB", nil, nil,
				),
				Value: func(svcMetric *contract.Report) float64 {
					return float64(svcMetric.MemoryUsageMiB)
				},
			},
			{
				Type: prometheus.GaugeValue,
				Desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "sha256_performance"),
					"SHA256 Performance value", nil, nil,
				),
				Value: func(svcMetric *contract.Report) float64 {
					return float64(svcMetric.SHA256Perf)
				},
			},
			{
				Type: prometheus.GaugeValue,
				Desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "hash_performance"),
					"Hash Performance value", nil, nil,
				),
				Value: func(svcMetric *contract.Report) float64 {
					return float64(svcMetric.HashPerf)
				},
			},
			{
				Type: prometheus.GaugeValue,
				Desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "memory_size"),
					"Node memory size in megabytes", nil, nil,
				),
				Value: func(svcMetric *contract.Report) float64 {
					return float64(svcMetric.MemorySize)
				},
			},
			{
				Type: prometheus.GaugeValue,
				Desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "num_cpu"),
					"Number of node CPU", nil, nil,
				),
				Value: func(svcMetric *contract.Report) float64 {
					return float64(svcMetric.NumCPU)
				},
			},
			{
				Type: prometheus.GaugeValue,
				Desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "uptime"),
					"Syncthing uptime in seconds", nil, nil,
				),
				Value: func(svcMetric *contract.Report) float64 {
					return float64(svcMetric.Uptime)
				},
			},
		},

		data: &svcData{
			Type: prometheus.GaugeValue,
			Desc: prometheus.NewDesc(
				prometheus.BuildFQName(namespace, subsystem, "node_info"),
				"Node's string information",
				[]string{"uniqueID", "version", "longVersion", "platform"},
				nil,
			),
			Value: func(svcData *contract.Report) float64 {
				if svcData.UniqueID != "" {
					return float64(1)
				}
				return float64(0)
			},
		},
	}

}

// Describe set Prometheus metrics descriptions.
func (c *SVCResponse) Describe(ch chan<- *prometheus.Desc) {
	for _, metric := range c.metrics {
		ch <- metric.Desc
	}

	ch <- c.data.Desc
	ch <- c.up.Desc()
	ch <- c.totalScrapes.Desc()
	ch <- c.jsonParseFailures.Desc()
}

func (c *SVCResponse) fetchDataAndDecode() (contract.Report, error) {
	var chr contract.Report

	u := *c.url
	url, _ := u.Parse("/rest/svc/report")

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
			_ = level.Warn(*c.logger).Log("msg", "failed to close http.Client", "err", err)
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

// Collect collects Syncthing metrics from /rest/svc/report.
func (c *SVCResponse) Collect(ch chan<- prometheus.Metric) {
	var err error
	c.totalScrapes.Inc()
	defer func() {
		ch <- c.up
		ch <- c.totalScrapes
		ch <- c.jsonParseFailures
	}()

	SVCResp, err := c.fetchDataAndDecode()
	if err != nil {
		c.up.Set(0)
		_ = level.Warn(*c.logger).Log(
			"msg", "failed to fetch and decode data",
			"err", err,
		)
		return
	}
	c.up.Set(1)

	for _, metric := range c.metrics {
		ch <- prometheus.MustNewConstMetric(
			metric.Desc,
			metric.Type,
			metric.Value(&SVCResp),
		)
	}

	ch <- prometheus.MustNewConstMetric(
		c.data.Desc,
		c.data.Type,
		c.data.Value(&SVCResp),
		SVCResp.UniqueID, SVCResp.Version, SVCResp.LongVersion, SVCResp.Platform,
	)
}
