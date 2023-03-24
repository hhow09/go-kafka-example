# Go Kafka Example

## Examples
### [basic](./basic/)
### [Basic Async Producer](./basic-async/)
- async producer
### [Basic Async Consumer Group (3)](./basic-async-consumer-group/)
- async producer with 3 consumers (same group)
### [Basic Batch Timeout Optimized](./basic-batchTimeout-optimized/)
- reduce batch timeout
### [replicated](./replicated/)
- adjust replication factor = 3

## How to use
- `make up`: run the services
- `up-rebuild`: run services with dockerfile rebuilt (if [producer](./producer/) or [consumer](./consumer/) need to rebuild)
- `make down`: stop services

### Settings
- change `REMOVE_RAW` to `false` if you want to see the raw logs of conumser

## Packages
- [producer](./producer/): message producer
- [consumer](./consumer/): message consumer (also write logs)
- [log-paresr](./log-paresr/): log paresr / analyzer for performance benchmark

## Ref
- [Kafka: The Definitive Guide v2](https://www.confluent.io/resources/kafka-the-definitive-guide-v2/)
- [conduktor/kafka-stack-docker-compose](https://github.com/conduktor/kafka-stack-docker-compose)
- [Implementing a Kafka Producer and Consumer In Golang (With Full Examples) For Production](https://www.sohamkamani.com/golang/working-with-kafka/)