---
title: OpenTelemetry Demo Docs
linkTitle: Demo
cascade:
  repo: https://github.com/open-telemetry/opentelemetry-demo
weight: 180
---

* [OpenTelemetry Demo](/ecosystem/demo/) documentation
  * goal
    * how to install and run the demo
    * view OpenTelemetry in action

## Running the Demo

- [Docker](docker-deployment)
- [Kubernetes](kubernetes-deployment)

## Language Feature Reference


| Language   | Automatic Instrumentation                             | Instrumentation Libraries                                                                    | Manual Instrumentation                                                                       |
| ---------- |-------------------------------------------------------| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| .NET       | [Accounting Service](services/accounting/)            | [Cart Service](services/cart/)                                                               | [Cart Service](services/cart/)                                                               |
| C++        |                                                       |                                                                                              | [Currency Service](services/currency/)                                                       |
| Go         |                                                       | [Checkout Service](services/checkout/), [Product Catalog Service](services/product-catalog/) | [Checkout Service](services/checkout/), [Product Catalog Service](services/product-catalog/) |
| Java       | [Ad Service](services/ad/)                            |                                                                                              | [Ad Service](services/ad/)                                                                   |
| JavaScript |                                                       |                                                                                              | [Payment Service](services/payment/)                                                         |
| TypeScript |                                                       | [Frontend](services/frontend/), [React Native App](services/react-native-app/)               | [Frontend](services/frontend/)                                                               |
| Kotlin     |                                                       | [Fraud Detection Service](services/fraud-detection/)                                         |                                                                                              |
| PHP        |                                                       | [Quote Service](services/quote/)                                                             | [Quote Service](services/quote/)                                                             |
| Python     | [Recommendation Service](services/recommendation/)    |                                                                                              | [Recommendation Service](services/recommendation/)                                           |
| Ruby       |                                                       | [Email Service](services/email/)                                                             | [Email Service](services/email/)                                                             |
| Rust       |                                                       | [Shipping Service](services/shipping/)                                                       | [Shipping Service](services/shipping/)                                                       |

## Service Documentation

* how OpenTelemetry is deployed / EACH service
  - [Accounting Service](services/accounting/)
  - [Ad Service](services/ad/)
  - [Cart Service](services/cart/)
  - [Checkout Service](services/checkout/)
  - [Email Service](services/email/)
  - [Frontend](services/frontend/)
  - [Load Generator](services/load-generator/)
  - [Payment Service](services/payment/)
  - [Product Catalog Service](services/product-catalog/)
  - [Quote Service](services/quote/)
  - [Recommendation Service](services/recommendation/)
  - [Shipping Service](services/shipping/)
  - [Image Provider Service](services/image-provider/)
  - [React Native App](services/react-native-app/)

## Feature Flag Scenarios

* goal
  * use OpenTelemetry -- to -- solve problems

* [feature flag enabled scenarios](feature-flags/)
