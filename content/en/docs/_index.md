---
title: Documentation
linkTitle: Docs
menu: { main: { weight: 10 } }
aliases: [/docs/workshop/*]
---

* OpenTelemetry (OTel)
  * == [Observability](concepts/observability-primer.md#what-is-observability) framework
    * vendor-neutral
      * [supported -- by -- > 90 observability vendors](/ecosystem/vendors/)
    * open source
    * ❌NOT by itself
      * observability backend (storage)
      * frontend (visualization)❌
  * uses
    * about telemetry data (_Examples:_ [traces](concepts/signals/traces/), [metrics](concepts/signals/metrics/), [logs](concepts/signals/logs/)),
      * instrumenting,
      * generating,
      * collecting,
      * exporting
  * == CNCF incubating project
  * == OpenTracing + OpenCensus projects
  * integrated -- by -- MANY [libraries, services, and apps](/ecosystem/integrations/)
  * adopted -- by -- [numerous end users](../ecosystem/adopters)
  * supports instrumentation
    * code-based
    * zero-code
  * architecture
    ![OpenTelemetry Reference Architecture](/opentelemetry-opentelemetry.io/static/img/otel-diagram.svg)
