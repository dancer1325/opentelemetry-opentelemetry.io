* goal
  * OTel collector + telemetrygen
  * OTel collector + prometheus + zipkin + jaeger

# OTel collector + telemetrygen
* follow OWN [quickstart.md](../../quick-start.md)
* `docker cp friendly_gauss:/etc/otelcol-contrib/config.yaml ./otel-collector-default-config.yaml`

# OTel collector + prometheus + zipkin + jaeger
* `docker compose up -d`
* http://localhost:55679/debug/servicez
  * http://localhost:55679/debug/pipelinez
* http://localhost:9090/query
  * `otelcol_exporter_queue_size`
