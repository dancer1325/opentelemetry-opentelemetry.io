---
title: Collector
description: Vendor-agnostic way to receive, process and export telemetry data.
aliases: [./about]
cascade:
  vers: 0.135.0
weight: 270
---

![OpenTelemetry Collector diagram with Jaeger, OTLP and Prometheus integration](img/otel-collector.svg)

## Introduction

* OpenTelemetry Collector
  * == 💡proxy💡
  * provides
    * 💡implementation, about how to receive, process and export telemetry data💡
      * vendor-agnostic
      * avoids run, operate, and maintain MULTIPLE agents/collectors
      * valid | scale environments
      * supports open source observability data formats (Jaeger, Prometheus, Fluent Bit, etc.)
  * goal
    - _Usability_
      - default configuration,
      - supports popular protocols,
      - runs and collects out of the box
    - _Performance_
      - HIGHLY stable and performant | DIFFERENT loads and configurations
    - _Observability_
      - == make a service, observable
    - _Extensibility_
      - WITHOUT touching the core code
    - _Unification_
      - 1! codebase /
        - ways to deploy
          - as an agent
          - as a collector
        - support for traces, metrics, and logs

## When to use a collector vs send DIRECTLY to a backend?

* | MOST language specific instrumentation libraries,
  * there are
    * exporters -- for -- popular backends (Jaeger, Prometheus, ...)
    * OTLP

* approaches
  * Collector
    * service -> collector -> backend
    * use cases
      * production services
        * Reason:🧠offload the service
          * collector can handle
            * retries,
            * batching,
            * encryption
            * sensitive data filtering🧠
  * DIRECT sent
    * service -> backend
    * use cases
      * development OR small-scale environment

## Collector security

* follow
  * best practices
    * [hosting]
    * [configuration]

## Status

* **Collector** status
  * [mixed][]
    * Reason:🧠core Collector components have mixed [stability levels][]🧠
      * [ALL AVAILABLE Collector components][]


[configuration]: /docs/security/config-best-practices/
[hosting]: /docs/security/hosting-best-practices/
[latest release]:
  https://github.com/open-telemetry/opentelemetry-collector-releases/releases/latest
[mixed]: /docs/specs/otel/document-status/#mixed
[ALL AVAILABLE Collector components]: /ecosystem/registry/?language=collector
[stability levels]:
  https://github.com/open-telemetry/opentelemetry-collector#stability-levels
