---
title: Quarkus instrumentation
linkTitle: Quarkus
---

* ways to instrument a Quarkus application
  * [Quarkus OpenTelemetry extension](https://quarkus.io/guides/opentelemetry)
    * maintained by Quarkus community
  * [OpenTelemetry Java agent](agent)
    * uses
      * you do NOT run a native image application

## Getting started

* requirements
  * Quarkus v3.16.0+

* steps
  * | your Quarkus application, add `quarkus-opentelemetry` extension dependency
    * | Maven

        ```xml
        <dependency>
            <groupId>io.quarkus</groupId>
            <artifactId>quarkus-opentelemetry</artifactId>
        </dependency>
        ```
    * | Gradle
        ```kotlin
        implementation("io.quarkus:quarkus-opentelemetry")
        ```
  * telemetry data
    * by default,
      * ONLY enabled tracing
    * if you want to enable metrics & logs -> add | "application.properties"

      ```properties
      quarkus.otel.metrics.enabled=true
      quarkus.otel.logs.enabled=true
      ```

* [MORE](https://quarkus.io/guides/opentelemetry#configuration-reference)

## Learn more

- Signal-specific guides for
  - [Tracing](https://quarkus.io/guides/opentelemetry-tracing)
  - [Metrics](https://quarkus.io/guides/opentelemetry-metrics)
  - [Logs](https://quarkus.io/guides/opentelemetry-logging)
