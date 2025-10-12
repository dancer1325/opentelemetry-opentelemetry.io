---
title: Docker deployment
linkTitle: Docker
aliases: [docker_deployment]
cSpell:ignore: otlphttp spanmetrics tracetest tracetesting
---

<!-- markdownlint-disable code-block-style ol-prefix -->

## Prerequisites

- Docker
- [Docker Compose](https://docs.docker.com/compose/install/) v2.0.0+
- Make (optional)
- 6 GB of RAM for the application

## how to run the demo?

* `git clone https://github.com/dancer1325/opentelemetry-demo`
* `cd opentelemetry-demo/`
* ways to start the demo
  * `make start`
  * `docker compose up --force-recreate --remove-orphans --detach`
* enable API observability-driven testing
  * OPTIONAL
  * ways
    * `make run-tracetesting`
    * `docker compose -f docker-compose-tests.yml run traceBasedTests`

* check
  - Web store
    - http://localhost:8080/
      - Problems:
        - Problem1: "{"message":"Cannot GET /","error":"Not Found","statusCode":404}"
          - Solution: kill existing process | 8080
  - Grafana
    - http://localhost:8080/grafana/
      - Problems:
        - Problem1: "no healthy upstream - tls: failed to verify certificate: x509"
          - Attempt1: comment `"GF_INSTALL_PLUGINS=grafana-opensearch-datasource"`
          - Solution: TODO:
  - Load Generator UI
    - http://localhost:8080/loadgen/
  - Jaeger UI
    - http://localhost:8080/jaeger/ui/
  - Tracetest UI
    - requirements
      - you run `make run-tracetesting`
    - http://localhost:11633/
  - Flagd configurator UI
    - http://localhost:8080/feature>

## Changing the demo's primary port number

* demo application
  * 👀start a proxy / ALL browser traffic👀
    * by default, 8080
    * if you want to change the port number -> BEFORE, set `ENVOY_PORT` environment variable
      * run -- via -- make

        ```shell
        ENVOY_PORT=8081 make start
        ```
      * run -- via -- docker compose

        ```shell
        ENVOY_PORT=8081 docker compose up --force-recreate --remove-orphans --detach
        ```

## Bring your own backend

* goal
  * web store -- as a -- demo application for an observability backend

* OpenTelemetry Collector
  * uses
    * export telemetry data -- to -- MULTIPLE backends
      * steps
        * | [src/otel-collector/otelcol-config-extras.yml](https://github.com/open-telemetry/opentelemetry-demo/blob/main/src/otel-collector/otelcol-config-extras.yml)

          ```yaml
            exporters:
              # add a NEW exporter
              otlphttp/example:
                endpoint: <your-endpoint-url>

            # override the exporters
            service:
              pipelines:
              traces:
              exporters: [spanmetrics, otlphttp/example]
          ```

  * by default, configuration files
    - `otelcol-config.yml`
    - `otelcol-config-extras.yml`

* TODO:
{{% alert title="Note" %}} When merging YAML values with the Collector, objects
are merged and arrays are replaced. The `spanmetrics` exporter must be included
in the array of exporters for the `traces` pipeline if overridden. Not including
this exporter will result in an error. {{% /alert %}}

Vendor backends might require you to add additional parameters for
authentication, please check their documentation. Some backends require
different exporters, you may find them and their documentation available at
[opentelemetry-collector-contrib/exporter](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/exporter).

After updating the `otelcol-config-extras.yml`, start the demo by running
`make start`. After a while, you should see the traces flowing into your backend
as well.

[^1]: {{% param notes.docker-compose-v2 %}}
