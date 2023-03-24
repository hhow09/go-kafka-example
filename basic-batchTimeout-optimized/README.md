# Basic Batch Timeout Optimized

## Settings
| Name                     	| setting 	|
|--------------------------	|---------	|
| topic: replication factor | 1       	|
| topic: partitions         | 3       	|
| brokers count            	| 3       	|
| producer count           	| 1       	|
| consumer count           	| 1       	|
| -----------           	  | ----     	|
| producer send           	| sync  	  |
| producer batch timeout   	| 10 ms (default 1 second)  	  |



## [WriterConfig.BatchTimeout](https://github.com/segmentio/kafka-go/blob/main/writer.go#L230)
```go
// Time limit on how often incomplete message batches will be flushed to
// kafka.
// The default is to flush at least every second.
BatchTimeout time.Duration
```
- default `BatchTimeout` in `segmentio/kafka-go` client is 1 second.

### linger.ms 
> linger.ms controls the amount of time to wait for additional messages before send‐ ing the current batch. KafkaProducer sends a batch of messages either when the cur‐ rent batch is full or when the linger.ms limit is reached. By default, the producer will send messages as soon as there is a sender thread available to send them, even if there’s just one message in the batch.

- there is no document but I suppose `BatchTimeout` is same as [linger.ms](https://www.conduktor.io/kafka/kafka-producer-batching/#linger.ms-and-batch.size-0) in Java config

> By setting linger.ms higher than 0, we instruct the producer to wait a few milliseconds to add additional messages to the batch before sending it to the brokers. This increases latency a little and significantly increases throughput—the overhead per message is much lower, and compression, if enabled, is much better.


## Change: BatchTimeout 10 ms
update [basic](../basic/) with producer batch timeout to 10 ms

```yaml
  producer:
    build:
     # ...
    environment: 
     # ...
      - BATCH_TIMEOUT_MS=10
    depends_on:
    # ...
```

### Output Log
```
received 10000 messages. 
total time consumption 109 seconds.
```