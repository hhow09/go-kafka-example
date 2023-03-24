# Replicated
- improve setting from [Basic Batch Timeout Optimized](../basic-batchTimeout-optimized/).
- for better durability.

## Settings
| Name                     	| setting 	    |
|--------------------------	|---------	    |
| topic: replication factor | 3 (updated)   |
| topic: partitions         | 3       	    |
| brokers count            	| 3       	    |
| producer count           	| 1       	    |
| consumer count           	| 1       	    |
| -----------           	| ----          |
| producer send           	| sync  	    |
| producer batch timeout   	| 10 ms         |
| producer Required Acks   	| -1  (all)     |

### Pros
- producer send messages `synchronously`, no message loss
- `replication-factor = 3`, high-availability


## Result
### Output Log
```
received 10000 messages. 
total time consumption 114 seconds.
```

- total time consumption is slower than **the case of replication factor = 1** in [Basic Batch Timeout Optimized](../basic-batchTimeout-optimized/) ( 109 seconds ) since producer needs to wait for every brokers' acks. 


## Ref
- https://github.com/conduktor/kafka-stack-docker-compose
- [Implementing a Kafka Producer and Consumer In Golang (With Full Examples) For Production](https://www.sohamkamani.com/golang/working-with-kafka/)