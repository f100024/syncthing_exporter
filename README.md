# Syncthing exporter

### Build and run exporter

Clone current repository and 
```bash

cd syncthing_exporter
go build .

./syncthing_exporter --version
```

For pre-built binaries please take a look at the releases.

Basic prometheus configuration:

```yaml
  - job_name: 'syncting_server'
    metrics_path: /metrics
    static_configs:
      - targets: ['127.0.0.1:9093']
        labels:
          service: syncthing_server
```

### Start flags

Name               | Evironment variable | Required | Description
-------------------|---------------------|----------|-------------
web.listen-address | WEB_LISTEN_ADDRESS  |     -    | Address ot listen on for web interface and telemetry  
web.metrics-path   | WEB_METRIC_PATH     |     -    | Path under which to expose metrics  
syncthing.uri      | SYNCTHING_URI       |     +    | HTTP API address of Syncthing node  
syncthing.token    | SYNCTHING_TOKEN     |     +    | Token for authentification Syncthing HTTP API
syncthing.timeout  | SYNCTHING_TIMEOUT   |     -    | Timeout for trying to get stats from Syncthing

### What's and how exported

Example of all metrics related to `syncthing` [here](examples/exposed_parameters.md).

For data obtaining is using two endpoints:

[GET /rest/svc/report](https://docs.syncthing.net/rest/svc-report-get.html)  
[GET /rest/system/connections](https://docs.syncthing.net/rest/system-connections-get.html)

### Grafana dashboard

Example of grafana dashboard:

![screenshot-1.png](./examples/grafana/screenshot-1.png)

## Docker Support

The following environment variables are supported

Env Variable            | Default               | Description
------------------------|-----------------------|--------------------------
ENV WEB_LISTEN_PORT=    | "9112"                | Port to expose telemetry 
ENV WEB_LISTEN_ADDRESS= | ":${PORT}"            | Address & port to expose telemetry
ENV WEB_METRIC_PATH     | "/metrics"            | HTTP Path under which to expose telemetry
ENV ST_URI=             | http://127.0.0.1:8384 | HTTP API address of Syncthing node
ENV ST_TOKEN            | "123"                 | !Change this! Token for authentification Syncthing HTTP API
ENV ST_TIMEOUT          | "5s"                  | Timeout for trying to get stats from Syncthing 

To run, simply

    docker run adenau/syncthing-exporter





## Communal effort
Any ideas and pull requests are appreciated.
