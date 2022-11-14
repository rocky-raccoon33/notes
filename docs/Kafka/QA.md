# QA

## Kafka 的消息可靠性策略

### 1 `Topic 分区副本` - 副本间的消息状态一致性

> `write-ahead log`

![](./img/fig1.jpeg)
<center>
    <br>
    <div style="color:orange; border-bottom: 1px solid #d9d9d9;
    display: inline-block;
    color: #999;
    padding: 2px;">WAL</div>
</center>

- `Kafka` `topic` 中的每个分区都有一个预写日志（`write-ahead log`），写入 `Kafka` 的消息就存储在这里面。这里面的每条消息都有一个唯一的偏移量，用于标识它在当前分区日志中的位置

> `副本间数据同步`

![](./img/fig22.jpeg)
<center>
    <br>
    <div style="color:orange; border-bottom: 1px solid #d9d9d9;
    display: inline-block;
    color: #999;
    padding: 2px;">leader 同步 follower</div>
</center>

> `ISR in-sync replica`

![](./img/fig51.jpeg)
<center>
    <br>
    <div style="color:orange; border-bottom: 1px solid #d9d9d9;
    display: inline-block;
    color: #999;
    padding: 2px;">ISR：和Leader保持同步的follower副本</div>
</center>

- `min.insync.replicas`：`ISR` 中同步的 `follower` 数量

!!! note

	```markdown
	- leader 允许 ISR 落后的消息数：replica.lag.max.messages
	- follower 在不超过 replica.lag.time.max.ms 时间内向 leader 发送 fetch 请求
	- 不同步的 follower 会从 ISR 中移除
	```

> 如何保证？

- `Leader` **crash** 时，`Kafka`会从`ISR`列表中选择第一个`Follower`作为新的`Leader`，`follower`分区拥有最新的已经 `committed` 的消息。通过这个可以保证已经 `committed` 的消息的数据可靠性

___


### 2 `Producer 消息确认机制`

> **`ack = 0`**：`Producer`通过网络把消息发出去，则认为消息已成功写入Kafka

- 序列化失败，分区离线或整个集群长时间不可用，生产者均不会收到任何错误
- 速度快，但无法保证`server`能收到消息

> **`default:ack = 1`**：`Leader`收到消息并写入分区文件时返回确认或错误响应

- 消息在写入`Leader`，`Follower`写入之前`Leader`奔溃，则消息丢失

> **`可用性保证：ack = -1`**：Leader在所有`Follower`收到消息后，才返回确认或错误响应

```yaml
request.required.acks:-1 # 当leader 同步到所有follower后，才会返回响应
unclean.leader.election.enable:false 
min.insync.replicas:${N/2+2} # 用于保证当前集群中处于正常同步状态的副本数量，当实际值小于配置值时，集群停止服务
```

- 异常情况下，当同步到所有`follower`前`leader` 奔溃，`producer`会重新发送，导致数据重复（需要业务端支持`幂等`）
- 序列化失败，分区离线或整个集群长时间不可用，生产者均不会收到任何错误
- 速度快，但无法保证`server`能收到消息


### 3 `Consumer 可靠性策略`

> `enable.auto.commit:true`：

- consumer 收到消息后即返回给broker，如果消费异常，则内容丢失
___
> `enable.auto.commit:false`：

- consumer 处理流程后手动提交，如果未提交时发生重启，会导致重复消费（需实现幂等）
___

> `Exactly once`：....
___

## 参考


- [1] [Hands-Free Kafka Replication: A Lesson in Operational Simplicity](https://www.confluent.io/blog/hands-free-kafka-replication-a-lesson-in-operational-simplicity/)

- [2] [Kafka 是如何保证数据可靠性和一致性](https://cloud.tencent.com/developer/article/1488458)

- [3] [简单理解 Kafka 的消息可靠性策略](https://cloud.tencent.com/developer/article/1752150)

