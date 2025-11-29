# == proxy
* `docker compose up -d`
* hit [sample.http](sample.http)
* `docker compose logs otel-collector > collectorLogs.txt`
* | [traces.txt](collectorLogs.txt)
  * look up
    * "otelcol_receiver"
      * == receiver
    * "otelcol_processor"
      * == processor
    * "otelcol_exporter"
      * == exporter
## provides implementation, about how to receive, process and export telemetry data (traces, metrics, logs)
* [otel-collector-config.yaml](otel-collector-config.yaml)
  * == receivers + processors + exporters
  * ALL telemetry data
    * `service.pipelines`
      * `.traces`
      * `.metrics`
      * `.logs`
### vendor-agnostic
* [otel-collector-vendor-agnostic.yaml](otel-collector-vendor-agnostic.yaml)
  * MULTIPLE exporters
* `docker compose -f docker-compose-vendor-agnostic.yml up -d`
# allows
## avoids run, operate, and maintain MULTIPLE agents/collectors
* TODO:
## being used | scaled environments (horizontal OR vertical)
### horizontal
* `docker compose -f docker-compose-horizontal-scaling.yml up -d`
### vertical
* TODO:
## valid | scale environments
* TODO:
# supports open source observability data formats (Jaeger, Prometheus, Fluent Bit, etc.)
* TODO:

# TODO:
* TODO:
