package main

import (
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/f100024/syncthing_exporter/collector"

	"github.com/go-kit/kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/promlog"
	"github.com/prometheus/common/version"
	kingpin "github.com/alecthomas/kingpin/v2"
)

func main() {
	var (
		Name             = "syncthing_exporter"
		webListenAddress = kingpin.Flag("web.listen-address",
			"Address ot listen on for web interface and telemetry. Environment variable: WEB_LISTEN_ADDRESS").
			Default(":9093").
			Envar("WEB_LISTEN_ADDRESS").
			String()

		webMetricsPath = kingpin.Flag("web.metrics-path",
			"Path under which to expose metrics. Environment variable: WEB_METRIC_PATH").
			Default("/metrics").
			Envar("WEB_METRIC_PATH").
			String()

		syncthingURI = kingpin.Flag("syncthing.uri",
			"HTTP API address of Syncthing node (e.g. http://127.0.0.1:8384). Environment variable: SYNCTHING_URI").
			Required().
			Envar("SYNCTHING_URI").
			URL()

		syncthingToken = kingpin.Flag("syncthing.token",
			"Token for authentification Syncthing API. Environment variable: SYNCTHING_TOKEN").
			Required().
			Envar("SYNCTHING_TOKEN").
			String()

		syncthingTimeout = kingpin.Flag("syncthing.timeout",
			"Timeout for trying to get stats from Syncthing. Environment variable: SYNCTHING_TIMEOUT").
			Default("5s").
			Envar("SYNCTHING_TIMEOUT").
			Duration()

		syncthingFoldersID = kingpin.Flag("syncthing.foldersid",
			"ID of folders for getting db status. Environment variable: SYNCTHING_FOLDERSID").
			Envar("SYNCTHING_FOLDERSID").
			String()
	)

	promlogConfig := &promlog.Config{}

	kingpin.Version(version.Print(Name))
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	logger := promlog.New(promlogConfig)
	stURL := *syncthingURI

	is_valid_url_schema, _ := regexp.MatchString("^(http|https)$", stURL.Scheme)
	if !is_valid_url_schema {
		level.Error(logger).Log("msg", "Syncthing URL schema is not allowed. URL schema should be matched http|https.")
		os.Exit(1)
	}

	collector.HttpClient.Timeout = *syncthingTimeout

	versionMetric := version.NewCollector(Name)
	prometheus.MustRegister(versionMetric)

	prometheus.MustRegister(collector.NewSVCReport(logger, collector.HttpClient, stURL, syncthingToken))
	prometheus.MustRegister(collector.NewSCReport(logger, collector.HttpClient, stURL, syncthingToken))
	prometheus.MustRegister(collector.NewStatsDeviceReport(logger, collector.HttpClient, stURL, syncthingToken))
	if *syncthingFoldersID != "" {
		foldersIDList := func(s *string) *[]string {
			var list []string
			splittedList := strings.Split(*s, ",")
			for folderid := range splittedList {
				list = append(list, strings.TrimSpace(splittedList[folderid]))
			}
			return &list
		}
		prometheus.MustRegister(collector.NewDBStatusReport(logger, collector.HttpClient, stURL, syncthingToken, foldersIDList(syncthingFoldersID)))
	}

	level.Info(logger).Log("msg", "Starting syncthing_exporter", "version", version.Info())
	level.Info(logger).Log("msg", "Build context", "build_context", version.BuildContext())

	http.Handle(*webMetricsPath, promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
			<head><title>Syncthing Exporter</title></head>
			<body>
			<h1>Syncthing Exporter</h1>
			<p><a href="` + *webMetricsPath + `">Metrics</a></p>
			</body>
			</html>`))
	})

	level.Info(logger).Log("msg", "Listening on", "address", *webListenAddress)
	http.ListenAndServe(*webListenAddress, nil)

}
