# Basic Async with Consumer Groups (3)

## Change
add [basic-async](../basic-async/) with 3 consumers (same consumer group)


## Settings
| Name                     	| setting 	  |
|--------------------------	|---------	  |
| topic: replication factor | 1       	  |
| topic: partitions         | 3       	  |
| brokers count            	| 3       	  |
| producer count           	| 1       	  |
| consumer count           	| 3 (changed) |
| -----------           	  | ----      |
| producer send           	| ASNYC  	  |
| -----------           	  | ----      |
| `MSG_COUNT`               | 10000     |

### Pros
- snce we have 3 partitions, using 3 conumser (same group) can enhenche receiver throughput.

## Result
### Output Log
- better performance than [basic-async](../basic-async/) because consumer groups are reading in parallel
```
...
received 10000 messages. 
total time consumption 8 seconds.
```