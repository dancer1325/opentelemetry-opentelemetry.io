* [OpenTelemetry Collector](../../collector)

{{ if $name }}

## Available exporters

The registry contains a [list of exporters for {{ $name }}][reg].

{{ end }}

{{ if not $name }}

* [AVAILABLE exporters | OpenTelemetry registry](https://opentelemetry.io/ecosystem/registry/?component=exporter&language=)

{{ end }}

* [OpenTelemetry Protocol (OTLP)](../../specs) exporters
  * ' goal
    * emit OTel data WITHOUT loss of information
      * == take in account OpenTelemetry data model

[Jaeger]: /blog/2022/jaeger-native-otlp/
[OTLP]: /docs/specs/otlp/
[Prometheus]:
  https://prometheus.io/docs/prometheus/2.55/feature_flags/#otlp-receiver
[reg]: </ecosystem/registry/?component=exporter&language={{ $lang }}>
[vendors]: /ecosystem/vendors/

{{ if $name }}

* goal
  * main OpenTelemetry {{ $name }} exporters
    * how to set them up

{{ end }}

{{ if $zeroConfigPageExists }}

{{% alert title=Note %}}

* [zero-code instrumentation](</docs/zero-code/{{ $langIdAsPath }}>),
you can learn how to set up exporters by following the
[Configuration Guide](</docs/zero-code/{{ $langIdAsPath }}/configuration/>)

{{% /alert %}}

{{ end }}

{{ if $supportsOTLP }}

## OTLP

### Collector Setup

* way to verify your OTLP exporters
  * `docker run -p 4317:4317 -p 4318:4318 --rm -v $(pwd)/collector-config.yaml:/etc/otelcol/config.yaml otel/opentelemetry-collector`
    * == run the collector | docker container /
      * writes telemetry DIRECTLY | console
      * accept telemetry -- via -- OTLP

* if you want to send your telemetry to your observability backend -> [configure the collector](../../collector/configuration)

{{ end }}
