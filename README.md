# Clash Tracing Dashboard

An example of a clash tracing exporter API.

### Screenshot

![screenshot](./screenshot/screenshot.jpg)

### How to use

1. modify `docker-compose.yaml` and start (`docker-compose up -d`)
2. setup Grafana
3. import `grafana.json` to Grafana

### For influxdb 1.0 compatibility

While using influxdb 2.0, a bucket must be mapped to a database and retention policy (DBRP).

https://docs.influxdata.com/influxdb/v2.0/query-data/influxql/#map-unmapped-buckets

Or manually migrate the `grafana.json` to compatible with `Flux` scripting language