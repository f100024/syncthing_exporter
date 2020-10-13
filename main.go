package main

import (
	"net/http"
	"net/url"
	"os"

	"github.com/f100024/syncthing_exporter/collector"

	"github.com/go-kit/kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/promlog"
	"github.com/prometheus/common/version"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	var (
		Name          = "syncthing_exporter"
		listenAddress = kingpin.Flag("web.listen-address",
			"Address ot listen on for web interface and telemetry.").
			Default(":9093").
			Envar("WEB_LISTEN_ADDRESS").
			String()

		metricsPath = kingpin.Flag("web.metrics-path",
			"Path under which to expose metrics").
			Default("/metrics").
			Envar("WEB_METRIC_PATH").
			String()

		stURI = kingpin.Flag("st.uri",
			"HTTP API address of Syncthing node").
			Default("http://127.0.0.1:8384").
			Envar("ST_URI").
			String()

		stToken = kingpin.Flag("st.token",
			"Token for authentification Syncthing HTTP API").
			Envar("ST_TOKEN").
			String()

		stTimeout = kingpin.Flag("st.timeout",
			"Timeout for trying to get stats from Syncthing").
			Default("5s").
			Envar("ST_TIMEOUT").
			Duration()
	)

	promlogConfig := &promlog.Config{}

	kingpin.Version(version.Print(Name))
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	logger := promlog.New(*&promlogConfig)

	stURL, err := url.Parse(*stURI)
	if err != nil {
		_ = level.Error(logger).Log(
			"msg", "failed to parse es.uri",
			"err", err,
		)
		os.Exit(1)
	}

	// TODO: Add TLS support
	httpClient := &http.Client{
		Timeout: *stTimeout,
	}

	versionMetric := version.NewCollector(Name)
	prometheus.MustRegister(versionMetric)

	prometheus.MustRegister(collector.NewSVCReport(logger, httpClient, stURL, stToken))
	prometheus.MustRegister(collector.NewSCReport(logger, httpClient, stURL, stToken))

	level.Info(logger).Log("msg", "Starting syncthing_exporter", "version", version.Info())
	level.Info(logger).Log("msg", "Build context", "build_context", version.BuildContext())

	http.Handle(*metricsPath, promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
			<head><title>Syncthing Exporter</title></head>
			<body>
			<h1>Syncthing Exporter</h1>
			<p><a href="` + *metricsPath + `">Metrics</a></p>
			</body>
			</html>`))
	})

	level.Info(logger).Log("msg", "Listening on", "address", *listenAddress)
	http.ListenAndServe(*listenAddress, nil)

}
