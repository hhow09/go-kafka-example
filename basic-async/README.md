# Basic Async Producer

## Change
update [basic](../basic/) with async producer

```yaml
  producer:
    build:
     # ...
    environment: 
     # ...
      - ASYNC=true 
    depends_on:
    # ...
``` 


## Settings
| Name                     	| setting 	|
|--------------------------	|---------	|
| topic: replication factor | 1       	|
| topic: partitions         | 3       	|
| brokers count            	| 3       	|
| producer count           	| 1       	|
| consumer count           	| 1       	|
| -----------           	  | ----     	|
| producer send           	| ASNYC  	  |

### Pros
- producer send messages `asynchronously`, higher throughput than [basic](../basic/)

### Cons
- `replication-factor = 1`,  this is not good from a high-availability and reliability perspective. 
- producer send messages `asynchronously`, failed retry could lead to messages **out of order**

## Result
### Output Log
- better performance than [basic](../basic/)
```
received 10000 messages. 
total time consumption 14 seconds.
```