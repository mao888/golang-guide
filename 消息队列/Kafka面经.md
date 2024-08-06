## [Apache Kafka是什么？](https://golangguide.top/%E4%B8%AD%E9%97%B4%E4%BB%B6/kafka/%E6%A0%B8%E5%BF%83%E7%9F%A5%E8%AF%86%E7%82%B9/kafka%E6%98%AF%E4%BB%80%E4%B9%88%EF%BC%9F%E6%9E%B6%E6%9E%84%E6%98%AF%E6%80%8E%E4%B9%88%E6%A0%B7%E7%9A%84%EF%BC%9F.html)
### Apache Kafka 介绍
Apache Kafka 是一个分布式流处理平台，最初由 LinkedIn 开发并开源，目前由 Apache 基金会维护。它的核心概念包括：

- **Producer（生产者）**：负责向 Kafka 主题（Topic）发布消息。
- **Consumer（消费者）**：订阅并处理 Kafka 主题中的消息。
- **Broker（代理）**：Kafka 集群中的服务器，每个 broker 处理一定数量的主题分区（Partition）。
- **Topic（主题）**：消息类别，每个主题可以有多个分区，允许并行处理。
- **Partition（分区）**：主题的子单元，消息在分区内有序，但分区之间无序。
- **Offset（偏移量）**：消息在分区中的位置，用于追踪消费进度。
### Kafka 的特点

1. **高吞吐量**：Kafka 能够处理大量的数据流，适用于实时数据管道和事件流处理。
2. **高可用性**：通过分区和复制机制，Kafka 能够保证高可用性和数据冗余。
3. **持久化存储**：Kafka 使用磁盘存储消息，支持持久化和回放。
4. **水平扩展**：Kafka 通过增加 broker 实现水平扩展，处理更大的数据量和并发。
### Kafka 的常见使用场景

1. **实时数据流**：实时日志收集和分析、监控数据流处理。
2. **事件驱动架构**：系统间的事件通知和消息传递。
3. **数据集成**：不同数据源间的数据同步和集成。
4. **日志和指标收集**：集中收集应用程序日志和指标，进行实时分析。
### Kafka 调优建议

1. **Broker 配置**：
   - 调整 `num.partitions` 和 `replication.factor`，根据集群规模和可靠性需求设置分区数和副本数。
   - 配置 `log.retention.hours` 和 `log.segment.bytes`，控制日志保留策略和分段大小。
   - 增加 `socket.send.buffer.bytes` 和 `socket.receive.buffer.bytes`，优化网络传输性能。
2. **生产者配置**：
   - 设置 `acks` 为 `all`，确保消息被所有副本确认后才认为成功，提高数据可靠性。
   - 调整 `batch.size` 和 `linger.ms`，优化批量发送策略，提升吞吐量。
3. **消费者配置**：
   - 使用 `enable.auto.commit` 控制偏移量提交，确保消费进度的可靠性。
   - 调整 `fetch.min.bytes` 和 `fetch.max.wait.ms`，优化拉取消息的策略，提升消费性能。

通过这些配置和优化，可以提升 Kafka 在高并发、大数据量场景下的性能和可靠性。

## 分区策略
Kafka的分区策略分为生产者端的分区策略和消费者端的分区分配策略。以下是详细的分区策略介绍：
### 生产者端的分区策略

1. **轮询策略（Round-Robin Strategy）**：
   - 这是Kafka Java生产者API默认提供的分区策略。如果没有指定分区策略，则会默认使用轮询。
   - 轮询策略按照顺序将消息发送到不同的分区，每个消息被发送到其对应分区，以确保每个分区均匀地接收消息。
   - 这种策略能够实现负载均衡，并且能够最大限度地利用集群资源。
2. **按键分配策略（Key-Based Partitioning）**：
   - 在这种策略中，消息的键被用作决定消息分区的依据。
   - Kafka根据键的哈希值将消息路由到相应的分区，确保具有相同键的消息被发送到同一分区，以提高数据局部性和处理效率。
3. **范围分区策略（Range Partitioning）**：
   - 在这种策略中，Kafka根据消息键的范围将消息分配到不同的分区。
   - 这种策略适用于有序数据的处理，如时间戳或递增的ID，通过将具有相似时间戳或递增ID的消息分配到同一分区，可以提高处理效率并保证数据的顺序性。
4. **自定义分区策略（Custom Partitioning）**：
   - 在某些情况下，可能需要根据特定的业务逻辑或规则来决定消息的分区。
   - 通过实现自定义的分区器类，可以根据应用程序的需求来定义分区的逻辑，如根据地理位置、用户ID等。
5. **粘性分区策略（Sticky Partitioning）**：
   - 在这种策略中，Kafka尽可能将消息分配到与之前消息相同的分区，以减少跨分区的数据移动和复制。
   - 这种策略通过维护一个分区和消费者的映射关系来实现，当消息被发送时，Kafka会尝试将其路由到与之前消息相同的分区。
### 消费者端的分区分配策略

1. **轮询分配（RoundRobinAssignor）**：
   - 将所有可用分区和消费者进行排序，然后按照轮询的方式将分区分配给每个消费者。
2. **范围分配（RangeAssignor）**：
   - 对每个Topic的分区按照序号进行排序，并对消费者按照字母顺序进行排序，然后尽可能均匀地将分区分配给每个消费者。
3. **粘性分配（StickyAssignor）**：
   - 在分配分区时，尽可能保持现有的分区分配方案不变，减少因消费者变动而导致的分区重分配。
4. **合作粘性分配（CooperativeStickyAssignor）**：
   - Kafka 2.4.0版本引入的策略，基于合作协议进行分区分配，将全局分区重平衡分解为多次小规模分区重平衡，以减少重平衡的开销。
### 默认策略

- **生产者端**：默认使用轮询策略（Round-Robin Strategy）。
- **消费者端**：默认策略可能因Kafka版本和配置而异，但常见的默认策略包括范围分配（RangeAssignor）或轮询分配（RoundRobinAssignor），具体取决于Kafka版本和消费者的配置。在Kafka 0.10.1.0及以后的版本中，消费者端默认可能同时使用多种分配策略，如`partition.assignment.strategy=org.apache.kafka.clients.consumer.RangeAssignor,org.apache.kafka.clients.consumer.RoundRobinAssignor`。不过，需要注意的是，Kafka的具体默认行为可能会随着版本的更新而有所变化，因此建议查阅最新版本的官方文档以获取准确信息。
## Kafka 的设计是什么样的？

Kafka 将消息以 topic 为单位进行归纳

将向 Kafka topic 发布消息的程序成为 producers.

将预订 topics 并消费消息的程序成为 consumer.

Kafka 以集群的方式运行，可以由一个或多个服务组成，每个服务叫做一个 broker.

producers 通过网络将消息发送到 Kafka 集群，集群向消费者提供消息

## Kafka 如何保证高可用？

`Kafka` 的基本架构组成是：由多个 `broker` 组成一个集群，每个 `broker` 是一个节点；当创建一个 `topic` 时，这个 `topic` 会被划分为多个 `partition`，每个 `partition` 可以存在于不同的 `broker` 上，每个 `partition` 只存放一部分数据。

这就是**天然的分布式消息队列**，就是说一个 `topic` 的数据，是**分散放在多个机器上的，每个机器就放一部分数据**。

在 `Kafka 0.8` 版本之前，是没有 `HA` 机制的，当任何一个 `broker` 所在节点宕机了，这个 `broker` 上的 `partition` 就无法提供读写服务，所以这个版本之前，`Kafka` 没有什么高可用性可言。

在 `Kafka 0.8` 以后，提供了 `HA` 机制，就是 `replica` 副本机制。每个 `partition` 上的数据都会同步到其它机器，形成自己的多个 `replica` 副本。所有 `replica` 会选举一个 `leader` 出来，消息的生产者和消费者都跟这个 `leader` 打交道，其他 `replica` 作为 `follower`。写的时候，`leader` 会负责把数据同步到所有 `follower` 上去，读的时候就直接读 `leader` 上的数据即可。`Kafka` 负责均匀的将一个 `partition` 的所有 `replica` 分布在不同的机器上，这样才可以提高容错性。

![](http://blog-img.coolsen.cn/img/Solve-MQ-Problem-With-Kafka-01.png#id=rYdVj&originHeight=283&originWidth=553&originalType=binary&ratio=1&rotation=0&showTitle=false&status=done&style=none&title=)

拥有了 `replica` 副本机制，如果某个 `broker` 宕机了，这个 `broker` 上的 `partition` 在其他机器上还存在副本。如果这个宕机的 `broker` 上面有某个 `partition` 的 `leader`，那么此时会从其 `follower` 中重新选举一个新的 `leader` 出来，这个新的 `leader` 会继续提供读写服务，这就有达到了所谓的高可用性。

写数据的时候，生产者只将数据写入 `leader` 节点，`leader` 会将数据写入本地磁盘，接着其他 `follower` 会主动从 `leader` 来拉取数据，`follower` 同步好数据了，就会发送 `ack` 给 `leader`，`leader` 收到所有 `follower` 的 `ack` 之后，就会返回写成功的消息给生产者。

消费数据的时候，消费者只会从 `leader` 节点去读取消息，但是只有当一个消息已经被所有 `follower` 都同步成功返回 `ack` 的时候，这个消息才会被消费者读到。

![](https://gitee.com/dongzl/article-images/raw/master/2020/13-Solve-MQ-Problem-With-Kafka/Solve-MQ-Problem-With-Kafka-02.png#id=It5OH&originHeight=1534&originWidth=1633&originalType=binary&ratio=1&rotation=0&showTitle=false&status=done&style=none&title=)

## Kafka 消息是采用 Pull 模式，还是 Push 模式？

生产者使用push模式将消息发布到Broker，消费者使用pull模式从Broker订阅消息。

push模式很难适应消费速率不同的消费者，如果push的速度太快，容易造成消费者拒绝服务或网络拥塞；如果push的速度太慢，容易造成消费者性能浪费。但是采用pull的方式也有一个缺点，就是当Broker没有消息时，消费者会陷入不断地轮询中，为了避免这点，kafka有个参数可以让消费者阻塞知道是否有新消息到达。

## Kafka 与传统消息系统之间的区别

-  Kafka 持久化日志，这些日志可以被重复读取和无限期保留 
-  Kafka 是一个分布式系统：它以集群的方式运行，可以灵活伸缩，在内部通过复制数据提升容错能力和高可用性 
-  Kafka 支持实时的流式处理 

## 什么是消费者组？

消费者组是Kafka独有的概念，即消费者组是Kafka提供的可扩展且具有容错性的消费者机制。

但实际上，消费者组（Consumer Group）其实包含两个概念，作为队列，消费者组允许你分割数据处理到一组进程集合上（即一个消费者组中可以包含多个消费者进程，他们共同消费该topic的数据），这有助于你的消费能力的动态调整；作为发布-订阅模型（publish-subscribe），Kafka允许你将同一份消息广播到多个消费者组里，以此来丰富多种数据使用场景。

需要注意的是：在消费者组中，多个实例共同订阅若干个主题，实现共同消费。同一个组下的每个实例都配置有相同的组ID，被分配不同的订阅分区。当某个实例挂掉的时候，其他实例会自动地承担起它负责消费的分区。 因此，消费者组在一定程度上也保证了消费者程序的高可用性。

![](http://dockone.io/uploads/article/20201024/7b359b7a1381541fbacf3ecf20dfb347.jpg#id=pdpbH&originHeight=225&originWidth=474&originalType=binary&ratio=1&rotation=0&showTitle=false&status=done&style=none&title=)

## 在Kafka中，ZooKeeper的作用是什么？

目前，Kafka使用ZooKeeper存放集群元数据、成员管理、Controller选举，以及其他一些管理类任务。之后，等KIP-500提案完成后，Kafka将完全不再依赖于ZooKeeper。

- “存放元数据”是指主题分区的所有数据都保存在 ZooKeeper 中，且以它保存的数据为权威，其他 “人” 都要与它保持对齐。
- “成员管理” 是指 Broker 节点的注册、注销以及属性变更，等等。
- “Controller 选举” 是指选举集群 Controller，而其他管理类任务包括但不限于主题删除、参数配置等。

KIP-500 思想，是使用社区自研的基于Raft的共识算法，替代ZooKeeper，实现Controller自选举。

## 解释下Kafka中位移（offset）的作用

在Kafka中，每个主题分区下的每条消息都被赋予了一个唯一的ID数值，用于标识它在分区中的位置。这个ID数值，就被称为位移，或者叫偏移量。一旦消息被写入到分区日志，它的位移值将不能被修改。

## kafka 为什么那么快？

- Cache Filesystem Cache PageCache缓存
- `顺序写`：由于现代的操作系统提供了预读和写技术，磁盘的顺序写大多数情况下比随机写内存还要快。
- `Zero-copy`：零拷技术减少拷贝次数
- `Batching of Messages`：批量量处理。合并小的请求，然后以流的方式进行交互，直顶网络上限。
- `Pull 拉模式`：使用拉模式进行消息的获取消费，与消费端处理能力相符。

## kafka producer发送数据，ack为0，1，-1分别是什么意思？

- `1`（默认） 数据发送到Kafka后，经过leader成功接收消息的的确认，就算是发送成功了。在这种情况下，如果leader宕机了，则会丢失数据。
- `0` 生产者将数据发送出去就不管了，不去等待任何返回。这种情况下数据传输效率最高，但是数据可靠性确是最低的。
- `-1`producer需要等待ISR中的所有follower都确认接收到数据后才算一次发送完成，可靠性最高。当ISR中所有Replica都向Leader发送ACK时，leader才commit，这时候producer才能认为一个请求中的消息都commit了。

## Kafka如何保证消息不丢失?

首先需要弄明白消息为什么会丢失，对于一个消息队列，会有 `生产者`、`MQ`、`消费者` 这三个角色，在这三个角色数据处理和传输过程中，都有可能会出现消息丢失。

![](http://blog-img.coolsen.cn/img/Solve-MQ-Problem-With-Kafka-03.png#id=fyJVA&originHeight=441&originWidth=1429&originalType=binary&ratio=1&rotation=0&showTitle=false&status=done&style=none&title=)

消息丢失的原因以及解决办法：

### 消费者异常导致的消息丢失

**消费者可能导致数据丢失的情况是**：消费者获取到了这条消息后，还未处理，`Kafka` 就自动提交了 `offset`，这时 `Kafka` 就认为消费者已经处理完这条消息，其实消费者才刚准备处理这条消息，这时如果消费者宕机，那这条消息就丢失了。

消费者引起消息丢失的主要原因就是消息还未处理完 `Kafka` 会自动提交了 `offset`，那么只要关闭自动提交 `offset`，消费者在处理完之后手动提交 `offset`，就可以保证消息不会丢失。但是此时需要注意重复消费问题，比如消费者刚处理完，还没提交 `offset`，这时自己宕机了，此时这条消息肯定会被重复消费一次，这就需要消费者根据实际情况保证**幂等性**。

### 生产者数据传输导致的消息丢失

对于生产者数据传输导致的数据丢失主常见情况是生产者发送消息给 `Kafka`，由于网络等原因导致消息丢失，对于这种情况也是通过在 **producer** 端设置 **acks=all** 来处理，这个参数是要求 `leader` 接收到消息后，需要等到所有的 `follower` 都同步到了消息之后，才认为本次写成功了。如果没满足这个条件，生产者会自动不断的重试。

### Kafka 导致的消息丢失

`Kafka` 导致的数据丢失一个常见的场景就是 `Kafka` 某个 `broker` 宕机，而这个节点正好是某个 `partition` 的 `leader` 节点，这时需要重新重新选举该 `partition` 的 `leader`。如果该 `partition` 的 `leader` 在宕机时刚好还有些数据没有同步到 `follower`，此时 `leader` 挂了，在选举某个 `follower` 成 `leader` 之后，就会丢失一部分数据。

对于这个问题，`Kafka` 可以设置如下 4 个参数，来尽量避免消息丢失：

- 给 `topic` 设置 `replication.factor` 参数：这个值必须大于 `1`，要求每个 `partition` 必须有至少 `2` 个副本；
- 在 `Kafka` 服务端设置 `min.insync.replicas` 参数：这个值必须大于 `1`，这个参数的含义是一个 `leader` 至少感知到有至少一个 `follower` 还跟自己保持联系，没掉队，这样才能确保 `leader` 挂了还有一个 `follower` 节点。
- 在 `producer` 端设置 `acks=all`，这个是要求每条数据，必须是写入所有 `replica` 之后，才能认为是写成功了；
- 在 `producer` 端设置 `retries=MAX`（很大很大很大的一个值，无限次重试的意思）：这个参数的含义是一旦写入失败，就无限重试，卡在这里了。

## 怎么确保数据的消费的
Kafka通过偏移量（Offset）来确保数据的消费。每个消费者组（Consumer Group）中的消费者会维护一个偏移量，表示其已经消费到的消息位置。当消费者消费消息时，它会更新自己的偏移量，并将新的偏移量提交给Kafka。如果消费者发生故障并重新启动，它会从上次提交的偏移量开始继续消费消息。
Kafka确保数据的消费主要通过其设计架构和一系列机制来实现，主要包括分区、消费者组、偏移量（Offset）管理、以及消费者策略等。以下是详细解释：
### 分区（Partition）

- Kafka中的每个Topic可以被分为多个分区，每个分区是一个有序的、不可变的消息序列。
- 分区是Kafka实现高吞吐量的关键，因为每个分区都可以独立地进行读写操作，这允许Kafka在多个分区之间并行处理数据。
- 生产者发送消息到指定的分区，消费者从分区中拉取消息进行消费。
### 消费者组（Consumer Group）

- Kafka中的消费者属于一个或多个消费者组，每个消费者组独立地消费Topic中的数据。
- 在同一个消费者组中，每个分区只能被一个消费者实例消费，但不同的消费者组可以消费同一个分区的数据。
- 这种设计允许Kafka实现广播（发布/订阅模式）和点对点（队列模式）两种消息传递模式。
### 偏移量（Offset）管理

- **偏移量**是Kafka中用来标识消费者消费位置的一个整数，它表示消费者已经消费的消息的下一个位置。
- 每个消费者组在每个分区中都有一个唯一的偏移量，用于记录该消费者组在该分区中的消费进度。
- 消费者通过提交偏移量来告诉Kafka它已经消费了哪些消息，以便在发生故障时可以从正确的位置恢复消费。
### 消费者策略
Kafka提供了多种消费者策略来确保数据的均衡消费，包括但不限于：

- **RangeAssignor**：这是Kafka的默认消费者策略，它按照分区的序号范围来分配分区给消费者。
- **RoundRobinAssignor**：使用轮询的方式来分配分区给消费者，确保每个消费者都尽可能均衡地消费数据。
- **StickyAssignor**：在保持现有分区分配的基础上，尽可能地均衡分配新加入的消费者或移除的消费者所负责的分区。
### 消息幂等性和事务

- Kafka还通过消息幂等性和事务机制来确保数据的精确一次消费（Exactly-Once Semantics）。
- 幂等性生产者（Idempotent Producer）可以确保即使在发生故障的情况下，同一条消息也只会被发送到Kafka一次。
- 事务（Transactions）允许生产者将多个消息作为单个原子性操作发送到Kafka，确保这些消息要么全部成功，要么全部失败。
### 数据一致性和可靠性

- Kafka通过多副本机制（Replication）和ISR（In-Sync Replicas）列表来确保数据的一致性和可靠性。
- 即使某个Broker发生故障，Kafka也能从其他副本中恢复数据，从而确保消费者能够继续消费数据而不会丢失。

综上所述，Kafka通过分区、消费者组、偏移量管理、消费者策略、消息幂等性和事务、以及数据一致性和可靠性机制来确保数据的消费。这些机制共同协作，使得Kafka成为一个高性能、高可靠性的消息队列系统。
## kafka如果线上消费端挂了，积压了大量的消息该怎么处理，怎么做，将影响减少

   1. 方法1: 启用备用消费端
      1. **冗余设计**：部署多个消费者实例，确保在一个实例挂掉时，其他实例能继续处理消息。
      2. **自动重启**：配置消费者实例自动重启策略，确保在挂掉后能尽快恢复。
   2. 方法2: 增加消费者实例
      1. **动态扩容**：根据消息积压情况，动态增加消费者实例数量，加速消息处理速度。
      2. **水平扩展**：使用容器化和编排工具（如Docker和Kubernetes）快速扩展消费者实例。
   3. 消费速率控制
      1. **批量处理**：增大批量处理大小，一次性处理更多消息，提高吞吐量。
      2. **异步处理**：使用异步处理方式，提升处理效率，减少处理时间。
   4. 方法5: 优化消费者代码
      1. **优化逻辑**：简化和优化消费者代码逻辑，确保每条消息处理时间最小化。
      2. **提高并发度**：利用多线程或协程并发处理消息，提高消费速度。
   5. 方法6: 持久化和备份
      1. **持久化机制**：确保消息持久化到可靠存储系统，避免数据丢失。
      2. **定期备份**：对消息和偏移量进行定期备份，确保在恢复时数据完整。
   6. 方法7: 设置报警和监控
      1. **实时监控**：使用Kafka监控工具（如Kafka Manager、Prometheus等）实时监控消费者状态和消息积压情况。
      2. **报警机制**：设置报警机制，及时通知运维人员进行处理。
   7. 方法8: 异常处理和降级策略
      1. **异常处理**：完善异常处理逻辑，确保在遇到错误时不会导致消费者挂掉。
      2. **降级策略**：在消费端挂掉时，启用降级策略，如暂时停止非关键业务的消息处理，确保关键业务优先处理。

## Kafka 如何保证消息的顺序性

在某些业务场景下，我们需要保证对于有逻辑关联的多条MQ消息被按顺序处理，比如对于某一条数据，正常处理顺序是`新增-更新-删除`，最终结果是数据被删除；如果消息没有按序消费，处理顺序可能是`删除-新增-更新`，最终数据没有被删掉，可能会产生一些逻辑错误。对于如何保证消息的顺序性，主要需要考虑如下两点：

- 如何保证消息在 `Kafka` 中顺序性；
- 如何保证消费者处理消费的顺序性。

### 如何保证消息在 Kafka 中顺序性

对于 `Kafka`，如果我们创建了一个 `topic`，默认有三个 `partition`。生产者在写数据的时候，可以指定一个 `key`，比如在订单 `topic` 中我们可以指定订单 `id` 作为 `key`，那么相同订单 `id` 的数据，一定会被分发到同一个 `partition` 中去，而且这个 `partition` 中的数据一定是有顺序的。消费者从 `partition` 中取出来数据的时候，也一定是有顺序的。通过制定 `key` 的方式首先可以保证在 `kafka` 内部消息是有序的。

### 如何保证消费者处理消费的顺序性

对于某个 `topic` 的一个 `partition`，只能被同组内部的一个 `consumer` 消费，如果这个 `consumer` 内部还是单线程处理，那么其实只要保证消息在 `MQ` 内部是有顺序的就可以保证消费也是有顺序的。但是单线程吞吐量太低，在处理大量 `MQ` 消息时，我们一般会开启多线程消费机制，那么如何保证消息在多个线程之间是被顺序处理的呢？对于多线程消费我们可以预先设置 `N` 个内存 `Queue`，具有相同 `key` 的数据都放到同一个内存 `Queue` 中；然后开启 `N` 个线程，每个线程分别消费一个内存 `Queue` 的数据即可，这样就能保证顺序性。当然，消息放到内存 `Queue` 中，有可能还未被处理，`consumer` 发生宕机，内存 `Queue` 中的数据会全部丢失，这就转变为上面提到的**如何保证消息的可靠传输**的问题了。

## Kafka中的ISR、AR代表什么？ISR的伸缩指什么？

- `ISR`：In-Sync Replicas 副本同步队列
- `AR`:Assigned Replicas 所有副本

ISR是由leader维护，follower从leader同步数据有一些延迟（包括`延迟时间replica.lag.time.max.ms`和`延迟条数replica.lag.max.messages`两个维度，当前最新的版本0.10.x中只支持`replica.lag.time.max.ms`这个维度），任意一个超过阈值都会把follower剔除出ISR，存入OSR（Outof-Sync Replicas）列表，新加入的follower也会先存放在OSR中。

> AR=ISR+OSR。


## 描述下 Kafka 中的领导者副本（Leader Replica）和追随者副本（Follower Replica）的区别

Kafka副本当前分为领导者副本和追随者副本。只有Leader副本才能对外提供读写服务，响应Clients端的请求。Follower副本只是采用拉（PULL）的方式，被动地同步Leader副本中的数据，并且在Leader副本所在的Broker宕机后，随时准备应聘Leader副本。

加分点：

- 强调Follower副本也能对外提供读服务。自Kafka 2.4版本开始，社区通过引入新的Broker端参数，允许Follower副本有限度地提供读服务。
- 强调Leader和Follower的消息序列在实际场景中不一致。通常情况下，很多因素可能造成Leader和Follower之间的不同步，比如程序问题，网络问题，broker问题等，短暂的不同步我们可以关注（秒级别），但长时间的不同步可能就需要深入排查了，因为一旦Leader所在节点异常，可能直接影响可用性。

注意：之前确保一致性的主要手段是高水位机制（HW），但高水位值无法保证Leader连续变更场景下的数据一致性，因此，社区引入了Leader Epoch机制，来修复高水位值的弊端。

## 分区Leader选举策略有几种？

分区的Leader副本选举对用户是完全透明的，它是由Controller独立完成的。你需要回答的是，在哪些场景下，需要执行分区Leader选举。每一种场景对应于一种选举策略。

- OfflinePartition Leader选举：每当有分区上线时，就需要执行Leader选举。所谓的分区上线，可能是创建了新分区，也可能是之前的下线分区重新上线。这是最常见的分区Leader选举场景。
- ReassignPartition Leader选举：当你手动运行kafka-reassign-partitions命令，或者是调用Admin的alterPartitionReassignments方法执行分区副本重分配时，可能触发此类选举。假设原来的AR是[1，2，3]，Leader是1，当执行副本重分配后，副本集合AR被设置成[4，5，6]，显然，Leader必须要变更，此时会发生Reassign Partition Leader选举。
- PreferredReplicaPartition Leader选举：当你手动运行kafka-preferred-replica-election命令，或自动触发了Preferred Leader选举时，该类策略被激活。所谓的Preferred Leader，指的是AR中的第一个副本。比如AR是[3，2，1]，那么，Preferred Leader就是3。
- ControlledShutdownPartition Leader选举：当Broker正常关闭时，该Broker上的所有Leader副本都会下线，因此，需要为受影响的分区执行相应的Leader选举。

这4类选举策略的大致思想是类似的，即从AR中挑选首个在ISR中的副本，作为新Leader。

## Kafka的哪些场景中使用了零拷贝（Zero Copy）？

在Kafka中，体现Zero Copy使用场景的地方有两处：基于mmap的索引和日志文件读写所用的TransportLayer。

先说第一个。索引都是基于MappedByteBuffer的，也就是让用户态和内核态共享内核态的数据缓冲区，此时，数据不需要复制到用户态空间。不过，mmap虽然避免了不必要的拷贝，但不一定就能保证很高的性能。在不同的操作系统下，mmap的创建和销毁成本可能是不一样的。很高的创建和销毁开销会抵消Zero Copy带来的性能优势。由于这种不确定性，在Kafka中，只有索引应用了mmap，最核心的日志并未使用mmap机制。

再说第二个。TransportLayer是Kafka传输层的接口。它的某个实现类使用了FileChannel的transferTo方法。该方法底层使用sendfile实现了Zero Copy。对Kafka而言，如果I/O通道使用普通的PLAINTEXT，那么，Kafka就可以利用Zero Copy特性，直接将页缓存中的数据发送到网卡的Buffer中，避免中间的多次拷贝。相反，如果I/O通道启用了SSL，那么，Kafka便无法利用Zero Copy特性了。

## 为什么Kafka不支持读写分离？

在 Kafka 中，生产者写入消息、消费者读取消息的操作都是与 leader 副本进行交互的，从 而实现的是一种主写主读的生产消费模型。

Kafka 并不支持主写从读，因为主写从读有 2 个很明 显的缺点:

- **数据一致性问题**。数据从主节点转到从节点必然会有一个延时的时间窗口，这个时间 窗口会导致主从节点之间的数据不一致。某一时刻，在主节点和从节点中 A 数据的值都为 X， 之后将主节点中 A 的值修改为 Y，那么在这个变更通知到从节点之前，应用读取从节点中的 A 数据的值并不为最新的 Y，由此便产生了数据不一致的问题。
- **延时问题**。类似 Redis 这种组件，数据从写入主节点到同步至从节点中的过程需要经历`网络→主节点内存→网络→从节点内存`这几个阶段，整个过程会耗费一定的时间。而在 Kafka 中，主从同步会比 Redis 更加耗时，它需要经历`网络→主节点内存→主节点磁盘→网络→从节点内存→从节点磁盘`这几个阶段。对延时敏感的应用而言，主写从读的功能并不太适用。

## 参考

[http://dockone.io/article/10853](http://dockone.io/article/10853)

[https://segmentfault.com/a/1190000023716306](https://segmentfault.com/a/1190000023716306)

[https://dongzl.github.io/2020/03/16/13-Solve-MQ-Problem-With-Kafka/index.html](https://dongzl.github.io/2020/03/16/13-Solve-MQ-Problem-With-Kafka/index.html)
