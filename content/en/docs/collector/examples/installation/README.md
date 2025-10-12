# Docker
## -- via -- DockerHub images
* `docker run otel/opentelemetry-collector-contrib`
## custom configuration files
* `docker run -d --name jaeger -p 16686:16686 -p 14250:14250 jaegertracing/all-in-one:latest`
* `docker run -v $(pwd)/otel-collector-config.yaml:/etc/otelcol-contrib/config.yaml otel/opentelemetry-collector-contrib`
  * `docker logs installation-otel-collector-1`
    * check ALL is working fine

# Docker Compose
* `docker compose up -d`
  * `docker logs installation-otel-collector-1`
    * check ALL is working fine

# TODO:
