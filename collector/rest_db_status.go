package collector

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"

	"github.com/prometheus/client_golang/prometheus"
)

// DBStatusResponseBoolMetrics defines struct for boolean metrics
type DBStatusResponseBoolMetrics struct {
	Type  prometheus.ValueType
	Desc  *prometheus.Desc
	Value func(v bool) float64
}

// DBStatusResponseNumericalMetrics defines struct for numeric metrics
type DBStatusResponseNumericalMetrics struct {
	Type  prometheus.ValueType
	Desc  *prometheus.Desc
	Value float64
}

// DBStatusMetrics defines collector struct.
type DBStatusMetrics struct {
	logger    *slog.Logger
	client    *http.Client
	url       *url.URL
	token     *string
	foldersid *[]string

	totalScrapes, jsonParseFailures prometheus.Counter
	boolMetrics                     map[string]*DBStatusResponseBoolMetrics
	numericalMetrics                map[string]*DBStatusResponseNumericalMetrics
}

// NewDBStatusReport returns a new Collector exposing SVCResponse
func NewDBStatusReport(logger *slog.Logger, client *http.Client, url *url.URL, token *string, foldersid *[]string) *DBStatusMetrics {
	subsystem := "rest_db_status"

	return &DBStatusMetrics{
		logger:    logger,
		client:    client,
		url:       url,
		token:     token,
		foldersid: foldersid,

		totalScrapes: prometheus.NewCounter(prometheus.CounterOpts{
			Name: prometheus.BuildFQName(namespace, subsystem, "total_scrapes"),
			Help: "Current total Syncthings scrapes.",
		}),
		jsonParseFailures: prometheus.NewCounter(prometheus.CounterOpts{
			Name: prometheus.BuildFQName(namespace, subsystem, "json_parse_failures"),
			Help: "Number of errors while parsing JSON.",
		}),
		boolMetrics: map[string]*DBStatusResponseBoolMetrics{
			"ignore_patterns": {
				Type: prometheus.GaugeValue,
				Desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "ignore_patterns"),
					"Is using ignore patterns.",
					[]string{"folderID"},
					nil),
				Value: func(v bool) float64 {
					return bool2float64(v)
				},
			},
		},

		numericalMetrics: map[string]*DBStatusResponseNumericalMetrics{
			"errors": {
				Type: prometheus.GaugeValue,
				Desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "errors"),
					"Number of errors for current folder.",
					[]string{"folderID"},
					nil),
			},

			// Global section
			"global_bytes": {
				Type: prometheus.GaugeValue,
				Desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "global_bytes"),
					"Number of bytes globally.",
					[]string{"folderID"},
					nil),
			},
			"global_deleted": {
				Type: prometheus.GaugeValue,
				Desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "global_deleted"),
					"Number of bytes deleted.",
					[]string{"folderID"},
					nil),
			},
			"global_directories": {
				Type: prometheus.GaugeValue,
				Desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "global_directories"),
					"Number of directories globally.",
					[]string{"folderID"},
					nil),
			},
			"global_symlinks": {
				Type: prometheus.GaugeValue,
				Desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "global_symlinks"),
					"Number of symlinks globally.",
					[]string{"folderID"},
					nil),
			},
			"global_total_items": {
				Type: prometheus.GaugeValue,
				Desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "global_total_items"),
					"Number of total items globally.",
					[]string{"folderID"},
					nil),
			},

			// InSync section
			"insync_bytes": {
				Type: prometheus.GaugeValue,
				Desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "insync_bytes"),
					"Number of bytes currently in sync.",
					[]string{"folderID"},
					nil),
			},
			"insync_files": {
				Type: prometheus.GaugeValue,
				Desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "insync_files"),
					"Number of files currently in sync.",
					[]string{"folderID"},
					nil),
			},

			// Local section
			"local_bytes": {
				Type: prometheus.GaugeValue,
				Desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "local_bytes"),
					"Number of bytes locally.",
					[]string{"folderID"},
					nil),
			},
			"local_deleted": {
				Type: prometheus.GaugeValue,
				Desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "local_deleted"),
					"Number of bytes deleted locally.",
					[]string{"folderID"},
					nil),
			},
			"local_directories": {
				Type: prometheus.GaugeValue,
				Desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "local_directories"),
					"Number of local directories.",
					[]string{"folderID"},
					nil),
			},
			"local_symlinks": {
				Type: prometheus.GaugeValue,
				Desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "local_symlinks"),
					"Number of local symlinks",
					[]string{"folderID"},
					nil),
			},
			"local_total_items": {
				Type: prometheus.GaugeValue,
				Desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "local_total_items"),
					"Number of total items locally",
					[]string{"folderID"},
					nil),
			},

			// Need section
			"need_bytes": {
				Type: prometheus.GaugeValue,
				Desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "need_bytes"),
					"Number of bytes need for sync.",
					[]string{"folderID"},
					nil),
			},
			"need_deletes": {
				Type: prometheus.GaugeValue,
				Desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "need_deletes"),
					"Number of bytes need for deletes.",
					[]string{"folderID"},
					nil),
			},
			"need_directories": {
				Type: prometheus.GaugeValue,
				Desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "need_directories"),
					"Number of directories for sync.",
					[]string{"folderID"},
					nil),
			},
			"need_symlinks": {
				Type: prometheus.GaugeValue,
				Desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "need_symlinks"),
					"Number of symlinks need for sync.",
					[]string{"folderID"},
					nil),
			},
			"need_total_items": {
				Type: prometheus.GaugeValue,
				Desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "need_total_items"),
					"Number of total items need to sync.",
					[]string{"folderID"},
					nil),
			},

			// Misc section
			"pull_errors": {
				Type: prometheus.GaugeValue,
				Desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "pull_errors"),
					"Number of pull errors.",
					[]string{"folderID"},
					nil),
			},

			"sequence": {
				Type: prometheus.GaugeValue,
				Desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "sequence"),
					"Total bytes received from remote device.",
					[]string{"folderID", "state", "stateChanged"},
					nil),
			},
		},
	}

}

// Describe set Prometheus metrics descriptions.
func (c *DBStatusMetrics) Describe(ch chan<- *prometheus.Desc) {

	for _, metric := range c.boolMetrics {
		ch <- metric.Desc
	}

	for _, metric := range c.numericalMetrics {
		ch <- metric.Desc
	}

	ch <- c.totalScrapes.Desc()
	ch <- c.jsonParseFailures.Desc()
}

func (c *DBStatusMetrics) fetchDataAndDecode() map[string]DBStatusResponse {

	var chr DBStatusResponse
	data := make(map[string]DBStatusResponse)
	var message string

	for indexfolderid := range *c.foldersid {

		c.totalScrapes.Inc()
		u := *c.url
		folderid := (*c.foldersid)[indexfolderid]
		url, _ := u.Parse(fmt.Sprintf("/rest/db/status?folder=%s", folderid))

		h := make(http.Header)
		h["X-API-Key"] = []string{*c.token}

		request := &http.Request{
			URL:    url,
			Header: h,
		}

		res, err := c.client.Do(request)
		if err != nil {
			message = fmt.Sprintf("Request %s://%s%s%s: failed with code %s",
				request.URL.Scheme, request.URL.Host, request.URL.Path, request.URL.RawQuery, err)
			c.logger.Error(message)
			continue
		}

		defer func() {
			err = res.Body.Close()
			if err != nil {
				c.logger.Info(fmt.Sprintf("%s: %s", "Failed to close http.Client", err))
			}
		}()

		if res.StatusCode != http.StatusOK {
			message = fmt.Sprintf("Request %s://%s%s%s: failed with code %d",
				request.URL.Scheme, request.URL.Host, request.URL.Path, request.URL.RawQuery, res.StatusCode)
			c.logger.Error(message)
			continue
		}

		if err := json.NewDecoder(res.Body).Decode(&chr); err != nil {
			c.logger.Info(fmt.Sprintf("%s: %s", "Failed decode json", err))
			c.jsonParseFailures.Inc()
			continue
		}
		data[folderid] = chr
	}
	return data
}

// Collect collects Syncthing metrics from /rest/db/status.
func (c *DBStatusMetrics) Collect(ch chan<- prometheus.Metric) {

	defer func() {
		ch <- c.totalScrapes
		ch <- c.jsonParseFailures
	}()

	response := c.fetchDataAndDecode()

	for folderid := range response {

		ch <- prometheus.MustNewConstMetric(c.numericalMetrics["errors"].Desc, c.numericalMetrics["errors"].Type, response[folderid].Errors, folderid)

		// Global section
		ch <- prometheus.MustNewConstMetric(c.numericalMetrics["global_bytes"].Desc, c.numericalMetrics["global_bytes"].Type, response[folderid].GlobalBytes, folderid)
		ch <- prometheus.MustNewConstMetric(c.numericalMetrics["global_deleted"].Desc, c.numericalMetrics["global_deleted"].Type, response[folderid].GlobalDeleted, folderid)
		ch <- prometheus.MustNewConstMetric(c.numericalMetrics["global_directories"].Desc, c.numericalMetrics["global_directories"].Type, response[folderid].GlobalDirectories, folderid)
		ch <- prometheus.MustNewConstMetric(c.numericalMetrics["global_symlinks"].Desc, c.numericalMetrics["global_symlinks"].Type, response[folderid].GlobalSymlinks, folderid)
		ch <- prometheus.MustNewConstMetric(c.numericalMetrics["global_total_items"].Desc, c.numericalMetrics["global_total_items"].Type, response[folderid].GlobalTotalItems, folderid)

		// Bool metrics
		ch <- prometheus.MustNewConstMetric(
			c.boolMetrics["ignore_patterns"].Desc,
			c.boolMetrics["ignore_patterns"].Type,
			c.boolMetrics["ignore_patterns"].Value(response[folderid].IgnorePatters),
			folderid,
		)

		// InSync section
		ch <- prometheus.MustNewConstMetric(c.numericalMetrics["insync_bytes"].Desc, c.numericalMetrics["insync_bytes"].Type, response[folderid].InSyncBytes, folderid)
		ch <- prometheus.MustNewConstMetric(c.numericalMetrics["insync_files"].Desc, c.numericalMetrics["insync_files"].Type, response[folderid].InSyncFiles, folderid)

		// Local section
		ch <- prometheus.MustNewConstMetric(c.numericalMetrics["local_bytes"].Desc, c.numericalMetrics["local_bytes"].Type, response[folderid].LocalBytes, folderid)
		ch <- prometheus.MustNewConstMetric(c.numericalMetrics["local_deleted"].Desc, c.numericalMetrics["local_deleted"].Type, response[folderid].LocalDeleted, folderid)
		ch <- prometheus.MustNewConstMetric(c.numericalMetrics["local_directories"].Desc, c.numericalMetrics["local_directories"].Type, response[folderid].LocalDirectories, folderid)
		ch <- prometheus.MustNewConstMetric(c.numericalMetrics["local_symlinks"].Desc, c.numericalMetrics["local_symlinks"].Type, response[folderid].LocalSymlinks, folderid)
		ch <- prometheus.MustNewConstMetric(c.numericalMetrics["local_total_items"].Desc, c.numericalMetrics["local_total_items"].Type, response[folderid].LocalTotalItems, folderid)

		// Need section
		ch <- prometheus.MustNewConstMetric(c.numericalMetrics["need_bytes"].Desc, c.numericalMetrics["need_bytes"].Type, response[folderid].NeedBytes, folderid)
		ch <- prometheus.MustNewConstMetric(c.numericalMetrics["need_deletes"].Desc, c.numericalMetrics["need_deletes"].Type, response[folderid].NeedDeletes, folderid)
		ch <- prometheus.MustNewConstMetric(c.numericalMetrics["need_directories"].Desc, c.numericalMetrics["need_directories"].Type, response[folderid].NeedDirectories, folderid)
		ch <- prometheus.MustNewConstMetric(c.numericalMetrics["need_symlinks"].Desc, c.numericalMetrics["need_symlinks"].Type, response[folderid].NeedSymlinks, folderid)
		ch <- prometheus.MustNewConstMetric(c.numericalMetrics["need_total_items"].Desc, c.numericalMetrics["need_total_items"].Type, response[folderid].NeedTotalItems, folderid)

		// Misc section
		ch <- prometheus.MustNewConstMetric(c.numericalMetrics["pull_errors"].Desc, c.numericalMetrics["pull_errors"].Type, response[folderid].PullErrors, folderid)
		ch <- prometheus.MustNewConstMetric(
			c.numericalMetrics["sequence"].Desc,
			c.numericalMetrics["sequence"].Type,
			response[folderid].Sequence,
			folderid, response[folderid].State, response[folderid].StateChanged,
		)
	}
}
