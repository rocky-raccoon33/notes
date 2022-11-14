# QA

## Kafka 的消息可靠性策略

### Producer 消息确认机制


```java ProducerConfig
/**
 * The number of acknowledgments the producer requires the leader to have received before considering a request complete. 
 * This controls the  durability of records that are sent. The following settings are allowed:  
 * <ul> 
 * <li><code>acks=0</code> If set to zero then the producer will not wait for any acknowledgment from the server at all. The record will be immediately added to the socket buffer and considered sent. No guarantee can be made that the server has received the record in this case, and the <code>retries</code> configuration will not take effect (as the client won't generally know of any failures). The offset given back for each record will always be set to -1. 
 * <li><code>acks=1</code> This will mean the leader will write the record to its local log but will respond without awaiting full acknowledgement from all followers. In this case should the leader fail immediately after acknowledging the record but before the followers have replicated it then the record will be lost. 
 * <li><code>acks=all</code> This means the leader will wait for the full set of in-sync replicas to acknowledge the record. This guarantees that the record will not be lost as long as at least one in-sync replica remains alive. This is the strongest available guarantee. This is equivalent to the acks=-1 setting.
 * 
*/
public static final String ACKS_CONFIG = "acks";

```

> `可用性保证：ack = -1`：Leader在所有`Follower`收到消息后，才返回确认或错误响应

> `ack = 0`：`Producer`通过网络把消息发出去，则认为消息已成功写入Kafka

- 序列化失败，分区离线或整个集群长时间不可用，生产者均不会收到任何错误
- 速度快，但无法保证`server`能收到消息

> `default:ack = 1`：`Leader`收到消息并写入分区文件时返回确认或错误响应

- 消息写入`Leader`，`Follower`写入之前`Leader`奔溃，则消息丢失

___

### Broker

> **副本间的消息状态一致性**

___
### Replia 

> Leader 选举