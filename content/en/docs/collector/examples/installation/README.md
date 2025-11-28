# Docker
## -- via -- DockerHub images
* `docker run otel/opentelemetry-collector-contrib`
* hit [sample.http](sample.http)
  * 's response: `"partialSuccess": {}` == OK
## custom configuration files
* `docker run -d --name jaeger -p 16686:16686 -p 14250:14250 jaegertracing/all-in-one:latest`
* `docker run -v $(pwd)/otel-collector-config.yaml:/etc/otelcol-contrib/config.yaml otel/opentelemetry-collector-contrib`
  * `docker logs installation-otel-collector-1`
    * check ALL is working fine
* hit [sample.http](sample.http)
  * 's response: `"partialSuccess": {}` == OK
* http://localhost:16686/search 
  * \> service: jaeger/all-in-one > Find traces
    * 's output: traces hit

# Docker Compose
* `docker compose up -d`
  * `docker logs installation-otel-collector-1`
    * check ALL is working fine
* hit [sample.http](sample.http)
  * 's response: `"partialSuccess": {}` == OK
* http://localhost:16686/search 
  * \> service: jaeger/all-in-one > Find traces
    * 's output: traces hit

# TODO:
