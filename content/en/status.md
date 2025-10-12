---
title: Status
menu: { main: { weight: 30 } }
aliases: [/project-status, /releases]
description: Maturity-level of the main OpenTelemetry components
type: docs
body_class: td-no-left-sidebar
---

* OpenTelemetry
  * == [MULTIPLE components](docs/concepts/components)
    * types
      * language-specific
      * language-agnostic
      * When looking for a
  * [status](docs/specs/status)

## Language APIs & SDKs

* TODO: For the development status, or maturity level, of a
[language API or SDK](/docs/languages/), see the following table:

{{% telemetry-support-table " " %}}

For more details on the specification compliance per implementation, see the
[Spec Compliance Matrix](https://github.com/open-telemetry/opentelemetry-specification/blob/main/spec-compliance-matrix.md).

## Collector

* collector status
  * [mixed](/docs/specs/otel/document-status/#mixed), since
  core collector components currently have mixed
  [stability levels](https://github.com/open-telemetry/opentelemetry-collector#stability-levels).

**Collector components** differ in their maturity levels. Each component has its
stability documented in its `README.md`. You can find a list of all available
collector components in the [registry](/ecosystem/registry/?language=collector).

## Kubernetes Operator

The OpenTelemetry Operator status is
[mixed](/docs/specs/otel/document-status/#mixed), since it deploys components of
differing statuses.

The Operator itself is in a [mixed](/docs/specs/otel/document-status/#mixed)
state with components in `v1alpha1` and `v1beta1` states.

## Specifications

For the development status, or maturity level, of the
[specification](/docs/specs/otel/), see the following:
[Specification Status Summary](/docs/specs/status/).
