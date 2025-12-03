# Pipelines
## TODO:
## Processors
### BEFORE forwarding data,
#### can transform data
* TODO:
#### TODO:
### uses
#### | MULTIPLE pipelines
* [otel-collector-processor1.yaml](otel-collector-processor1.yaml)

  ```mermaid
  ---
  title: Pipeline "traces"
  ---
  flowchart LR
    R1("`zipkin Receiver`") --> P1["`#quot;batch#quot; Processor`"]
    P1 --> E1[["`#quot;otlp#quot; Exporter`"]]
  ```

  ```mermaid
  ---
  title: Pipeline "traces/2"
  ---
  flowchart LR
    R1("`otlp Receiver`") --> P1["`#quot;batch#quot; Processor`"]
    P1 --> E1[["`#quot;otlp#quot; Exporter`"]]
  ```

##### 1 instance of the processor / EACH pipeline
* `docker compose -f docker-compose-processor.yaml up -d`
* http://localhost:55679/debug/pipelinez
  * DIFFERENT line -> DIFFERENT instance
###### -> OWN state
* TODO:
###### processors are NEVER shared BETWEEN pipelines
* TODO:

# TODO:
