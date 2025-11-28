---
title: Language APIs & SDKs
description:
  OpenTelemetry code instrumentation is supported for many popular programming
  languages
weight: 250
aliases: [/docs/instrumentation]
redirects:
  - { from: /docs/instrumentation/*, to: ':splat' } # Only for `en`
  - { from: 'net/*', to: 'dotnet/:splat' }
---

* OpenTelemetry 
  * [code instrumentation](../concepts/instrumentation)
    * supported languages -- [registry](../../ecosystem/registry/) --
      * [OFFICIAL ones](#status-and-releases)
      * [Unofficial implementations](other)
  * [zero-code solutions](../zero-code) 
    * ALLOWED |
      * Go,
      * .NET,
      * PHP,
      * Python,
      * Java
      * JavaScript
 to add instrumentation to your
application without code changes.

If you are using Kubernetes, you can use the [OpenTelemetry Operator for
Kubernetes][otel-op] to [inject these zero-code solutions][zero-code] into your
application.

## OpenTelemetry's major functional components' status & releases

* [status definitions](../specs/status.md)
* [semantic conventions](../concepts/semantic-conventions.md)

* [telemetry-support-table](../../../../layouts/_shortcodes/telemetry-support-table.md)

## API references

Special Interest Groups (SIGs) implementing the OpenTelemetry API and SDK in a
specific language also publish API references for developers. The following
references are available:

{{% apidocs %}}

{{% alert title="Note" %}}

The list above is aliased to [`/api`](/api).

{{% /alert %}}

[zero-code]: /docs/platforms/kubernetes/operator/automatic/
[instrumentation]: /docs/concepts/instrumentation/
[otel-op]: /docs/platforms/kubernetes/operator/
