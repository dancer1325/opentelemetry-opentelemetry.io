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
          - Solution: TODO:
  - Grafana
    - http://localhost:8080/grafana/
      - Problems:
        - Problem1: "{"message":"Cannot GET /","error":"Not Found","statusCode":404}"
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

By default, the demo application will start a proxy for all browser traffic
bound to port 8080. To change the port number, set the `ENVOY_PORT` environment
variable before starting the demo.

- For example, to use port 8081[^1]:

  {{< tabpane text=true >}} {{% tab Make %}}

```shell
ENVOY_PORT=8081 make start
```

    {{% /tab %}} {{% tab Docker %}}

```shell
ENVOY_PORT=8081 docker compose up --force-recreate --remove-orphans --detach
```

    {{% /tab %}} {{< /tabpane >}}

## Bring your own backend

Likely you want to use the web store as a demo application for an observability
backend you already have (e.g., an existing instance of Jaeger, Zipkin, or one
of the [vendors of your choice](/ecosystem/vendors/)).

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
