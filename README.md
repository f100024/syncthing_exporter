# Syncthing Exporter (Deprecated)

> **This project is deprecated and archived.** Syncthing has built-in Prometheus metrics since version 1.24.0. Please use the native `/metrics` endpoint instead.

## Why deprecated?

Syncthing added native Prometheus metrics support in [PR #9003](https://github.com/syncthing/syncthing/pull/9003), released around v1.24.0. The native endpoint:

- Requires no separate exporter process
- Supports auth-free scraping
- Includes device and folder labels
- Is actively maintained by the Syncthing team

According to [data.syncthing.net](https://data.syncthing.net/), very few users remain on versions older than 1.24.0. A JSON-scraping middleman exporter no longer provides enough value to justify ongoing maintenance.

## Migration

Configure Prometheus to scrape Syncthing's built-in `/metrics` endpoint directly:

```yaml
- job_name: 'syncthing'
  static_configs:
    - targets: ['127.0.0.1:8384']
```

See the [official metrics documentation](https://docs.syncthing.net/users/metrics) for details.

## What this was

`syncthing_exporter` was a Prometheus exporter that scraped Syncthing's REST API and exposed the data as Prometheus metrics. It supported metrics from endpoints such as `/rest/svc/report`, `/rest/system/connections`, `/rest/stats/device`, `/rest/db/status`, and `/rest/config/devices`.
