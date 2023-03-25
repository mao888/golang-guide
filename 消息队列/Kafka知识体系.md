# 一、消息队列

## 1. 消息队列的介绍

消息（Message）是指在应用之间传送的数据，消息可以非常简单，比如只包含文本字符串，也可以更复杂，可能包含嵌入对象。消息队列（Message Queue）是一种应用间的通信方式，消息发送后可以立即返回，有消息系统来确保信息的可靠专递，消息发布者只管把消息发布到MQ中而不管谁来取，消息使用者只管从MQ中取消息而不管谁发布的，这样发布者和使用者都不用知道对方的存在。

## 2. 消息队列的应用场景

消息队列在实际应用中包括如下四个场景：

- **应用耦合**：多应用间通过消息队列对同一消息进行处理，避免调用接口失败导致整个过程失败；
- **异步处理**：多应用对消息队列中同一消息进行处理，应用间并发处理消息，相比串行处理，减少处理时间；
- **限流削峰**：广泛应用于秒杀或抢购活动中，避免流量过大导致应用系统挂掉的情况；
- **消息驱动的系统**：系统分为消息队列、消息生产者、消息消费者，生产者负责产生消息，消费者(可能有多个)负责对消息进行处理；

下面详细介绍上述四个场景以及消息队列如何在上述四个场景中使用：

### 异步处理

具体场景：用户为了使用某个应用，进行注册，系统需要发送注册邮件并验证短信。对这两个操作的处理方式有两种：**串行**及**并行**。

**串行方式**：新注册信息生成后，先发送注册邮件，再发送验证短信； 在这种方式下，需要最终发送验证短信后再返回给客户端。

![img](https://cdn.nlark.com/yuque/0/2023/png/22219483/1678885094809-3fbc8d86-6af6-44d7-957a-cb17a6c24bc5.png)

**并行处理**：新注册信息写入后，由发短信和发邮件并行处理； 

![img](https://cdn.nlark.com/yuque/0/2023/png/22219483/1678885115071-c16b7dd5-4b13-4f32-af5d-65b926cc48e8.png)

在这种方式下，发短信和发邮件 需处理完成后再返回给客户端。假设以上三个子系统处理的时间均为50ms，且不考虑网络延迟，则总的处理时间：

串行：50+50+50=150ms

并行：50+50 = 100ms

若使用消息队列：

![img](https://cdn.nlark.com/yuque/0/2023/png/22219483/1678885143645-e4d82be1-9788-4d9e-be67-a9b687575449.png)

在写入消息队列后立即返回成功给客户端，则总的响应时间依赖于写入消息队列的时间，而写入消息队列的时间本身是可以很快的，基本可以忽略不计，因此总的处理时间相比串行提高了2倍，相比并行提高了一倍；

### 应用耦合

具体场景：用户使用QQ相册上传一张图片，人脸识别系统会对该图片进行人脸识别，一般的做法是，服务器接收到图片后，图片上传系统立即调用人脸识别系统，调用完成后再返回成功，如下图所示：

![img](https://cdn.nlark.com/yuque/0/2023/png/22219483/1678885318587-b8bf48f1-56d4-44f6-97db-fae2fadb5bde.png)

该方法有如下缺点： 

- 人脸识别系统被调失败，导致图片上传失败；
- 延迟高，需要人脸识别系统处理完成后，再返回给客户端，即使用户并不需要立即知道结果；
- 图片上传系统与人脸识别系统之间互相调用，需要做耦合；

若使用消息队列：

![img](https://cdn.nlark.com/yuque/0/2023/png/22219483/1678931938551-b383174c-c906-41f6-b6b8-5a7ce875e43c.png)

客户端上传图片后，图片上传系统将图片信息如uin、批次写入消息队列，直接返回成功；而人脸识别系统则定时从消息队列中取数据，完成对新增图片的识别。

此时图片上传系统并不需要关心人脸识别系统是否对这些图片信息的处理、以及何时对这些图片信息进行处理。事实上，由于用户并不需要立即知道人脸识别结果，人脸识别系统可以选择不同的调度策略，按照闲时、忙时、正常时间，对队列中的图片信息进行处理。

### 限流削峰

具体场景：购物网站开展秒杀活动，一般由于瞬时访问量过大，服务器接收过大，会导致流量暴增，相关系统无法处理请求甚至崩溃。而加入消息队列后，系统可以从消息队列中取数据，相当于消息队列做了一次缓冲。

![img](https://cdn.nlark.com/yuque/0/2023/png/22219483/1678933058622-fcc41fd2-debc-4279-a863-9afd5e20260f.png)

该方法有如下优点：

- 请求先入消息队列，而不是由业务处理系统直接处理，做了一次缓冲,极大地减少了业务处理系统的压力；
- 队列长度可以做限制，事实上，秒杀时，后入队列的用户无法秒杀到商品，这些请求可以直接被抛弃，返回活动已结束或商品已售完信息；

### 消息驱动的系统

具体场景：用户新上传了一批照片，人脸识别系统需要对这个用户的所有照片进行聚类，聚类完成后由对账系统重新生成用户的人脸索引(加快查询)。这三个子系统间由消息队列连接起来，前一个阶段的处理结果放入队列中，后一个阶段从队列中获取消息继续处理。  

![img](https://cdn.nlark.com/yuque/0/2023/png/22219483/1678933331342-3d9e0eca-5a4d-4d33-8ed2-3b7555978ea1.png)

该方法有如下优点：

- 避免了直接调用下一个系统导致当前系统失败；
- 每个子系统对于消息的处理方式可以更为灵活，可以选择收到消息时就处理，可以选择定时处理，也可以划分时间段按不同处理速度处理；

## 3. 消息队列的两种模式

消息队列包括两种模式，点对点模式（point to point， queue）和发布/订阅模式（publish/subscribe，topic）

### 1) 点对点模式

点对点模式下包括三个角色：

- 消息队列
- 发送者 (生产者)
- 接收者（消费者）  

![img](https://cdn.nlark.com/yuque/0/2023/png/22219483/1678933556594-4eb8f012-53f3-4a88-bd5a-8f63452dac23.png)

消息发送者生产消息发送到queue中，然后消息接收者从queue中取出并且消费消息。消息被消费以后，queue中不再有存储，所以消息接收者不可能消费到已经被消费的消息。

点对点模式特点：

- 每个消息只有一个接收者（Consumer）(即一旦被消费，消息就不再在消息队列中)；
- 发送者和接发收者间没有依赖性，发送者发送消息之后，不管有没有接收者在运行，都不会影响到发送者下次发送消息；
- 接收者在成功接收消息之后需向队列应答成功，以便消息队列删除当前接收的消息；

### 2) 发布/订阅模式

发布/订阅模式下包括三个角色：

- 角色主题（Topic）
- 发布者(Publisher)
- 订阅者(Subscriber)  

![img](https://cdn.nlark.com/yuque/0/2023/png/22219483/1678935912021-cd7e68ce-7453-4332-8730-7573a49f1274.png)

发布者将消息发送到Topic，系统将这些消息传递给多个订阅者。

发布/订阅模式特点：

- 每个消息可以有多个订阅者；
- 发布者和订阅者之间有时间上的依赖性。针对某个主题（Topic）的订阅者，它必须创建一个订阅者之后，才能消费发布者的消息。
- 为了消费消息，订阅者需要提前订阅该角色主题，并保持在线运行；

## 4. 常用的消息队列介绍

### 1) RabbitMQ

RabbitMQ 2007年发布，是一个在AMQP(高级消息队列协议)基础上完成的，可复用的企业消息系统，是当前最主流的消息中间件之一。

### 2) ActiveMQ

ActiveMQ是由Apache出品，ActiveMQ 是一个完全支持JMS1.1和J2EE 1.4规范的 JMS Provider实现。它非常快速，支持多种语言的客户端和协议，而且可以非常容易的嵌入到企业的应用环境中，并有许多高级功能。

### 3) RocketMQ

RocketMQ出自 阿里公司的开源产品，用 Java 语言实现，在设计时参考了 Kafka，并做出了自己的一些改进，消息可靠性上比 Kafka 更好。RocketMQ在阿里集团被广泛应用在订单，交易，充值，流计算，消息推送，日志流式处理等。

### 4) Kafka

Apache Kafka是一个分布式消息发布订阅系统。它最初由LinkedIn公司基于独特的设计实现为一个分布式的提交日志系统( a distributed commit log)，，之后成为Apache项目的一部分。Kafka系统快速、可扩展并且可持久化。它的分区特性，可复制和可容错都是其不错的特性。

### 5. Pulsar

Apahce Pulasr是一个企业级的发布-订阅消息系统，最初是由雅虎开发，是下一代云原生分布式消息流平台，集消息、存储、轻量化函数式计算为一体，采用计算与存储分离架构设计，支持多租户、持久化存储、多机房跨区域数据复制，具有强一致性、高吞吐、低延时及高可扩展性等流数据存储特性。

Pulsar 非常灵活：它既可以应用于像 Kafka 这样的分布式日志应用场景，也可以应用于像 RabbitMQ 这样的纯消息传递系统场景。它支持多种类型的订阅、多种交付保证、保留策略以及处理模式演变的方法，以及其他诸多特性。

### 7. 其他消息队列与Kafka对比

![img](https://cdn.nlark.com/yuque/0/2023/png/22219483/1678936443185-6cc113e5-088f-4bdd-94a9-1ff00b29a87c.png)

# 二、Kafka基础

## 1. kafka的基本介绍

官网：http://kafka.apache.org/

kafka是最初由linkedin公司开发的，使用scala语言编写，kafka是一个分布式，分区的，多副本的，多订阅者的日志系统（分布式MQ系统），可以用于搜索日志，监控日志，访问日志等。

Kafka is a distributed,partitioned,replicated commit logservice。它提供了类似于JMS的特性，但是在设计实现上完全不同，此外它并不是JMS规范的实现。kafka对消息保存时根据Topic进行归类，发送消息者成为Producer,消息接受者成为Consumer,此外kafka集群有多个kafka实例组成，每个实例(server)成为broker。无论是kafka集群，还是producer和consumer都依赖于zookeeper来保证系统可用性集群保存一些meta信息。

## 2. kafka的好处

- 可靠性：分布式的，分区，复本和容错的。
- 可扩展性：kafka消息传递系统轻松缩放，无需停机。
- 耐用性：kafka使用分布式提交日志，这意味着消息会尽可能快速的保存在磁盘上，因此它是持久的。
- 性能：kafka对于发布和定于消息都具有高吞吐量。即使存储了许多TB的消息，他也爆出稳定的性能。
- kafka非常快：保证零停机和零数据丢失。

## 3. 分布式的发布与订阅系统

apache kafka是一个分布式发布-订阅消息系统和一个强大的队列，可以处理大量的数据，并使能够将消息从一个端点传递到另一个端点，kafka适合离线和在线消息消费。kafka消息保留在磁盘上，并在集群内复制以防止数据丢失。kafka构建在zookeeper同步服务之上。它与apache和spark非常好的集成，应用于实时流式数据分析。

## 4. kafka的主要应用场景

**1. 指标分析**

kafka 通常用于操作监控数据。这设计聚合来自分布式应用程序的统计信息， 以产生操作的数据集中反馈

**2. 日志聚合解决方法**

kafka可用于跨组织从多个服务器收集日志，并使他们以标准的格式提供给多个服务器。

**3. 流式处理**

流式处理框架（spark，storm，ﬂink）重主题中读取数据，对齐进行处理，并将处理后的数据写入新的主题，供 用户和应用程序使用，kafka的强耐久性在流处理的上下文中也非常的有用。

# 三、Kafka架构及组件

## 1. kafka架构

1. ![img](https://cdn.nlark.com/yuque/0/2023/png/22219483/1678937134718-929b6f4d-eadf-42e7-b2c2-908a75dc0dc5.png)
2. **生产者API**

1. 1. 允许应用程序发布记录流至一个或者多个kafka的主题（topics）。

1. **消费者API**

1. 1. 允许应用程序订阅一个或者多个主题，并处理这些主题接收到的记录流。

1. **StreamsAPI**

1. 1. 允许应用程序充当流处理器（stream processor），从一个或者多个主题获取输入流，并生产一个输出流到一个或者多个主题，能够有效的变化输入流为输出流。

1. **ConnectAPI**

1. 1. 允许构建和运行可重用的生产者或者消费者，能够把kafka主题连接到现有的应用程序或数据系统。例如：一个连接到关系数据库的连接器可能会获取每个表的变化。

![img](https://cdn.nlark.com/yuque/0/2023/png/22219483/1678953033506-45455ca0-ee85-4152-a501-60f9e0b55623.png)

注：**在Kafka 2.8.0 版本，移除了对Zookeeper的依赖，通过KRaft进行自己的集群管理**，使用Kafka内部的Quorum控制器来取代ZooKeeper，因此用户第一次可在完全不需要ZooKeeper的情况下执行Kafka，这不只节省运算资源，并且也使得Kafka效能更好，还可支持规模更大的集群。

过去Apache ZooKeeper是Kafka这类分布式系统的关键，ZooKeeper扮演协调代理的角色，所有代理服务器启动时，都会连接到Zookeeper进行注册，当代理状态发生变化时，Zookeeper也会储存这些数据，在过去，ZooKeeper是一个强大的工具，但是毕竟ZooKeeper是一个独立的软件，使得Kafka整个系统变得复杂，因此官方决定使用内部Quorum控制器来取代ZooKeeper。

这项工作从去年4月开始，而现在这项工作取得部分成果，用户将可以在2.8版本，在没有ZooKeeper的情况下执行Kafka，官方称这项功能为**Kafka Raft元数据模式（KRaft）**。在KRaft模式，过去由Kafka控制器和ZooKeeper所操作的元数据，将合并到这个新的Quorum控制器，并且在Kafka集群内部执行，当然，如果使用者有特殊使用情境，Quorum控制器也可以在专用的硬件上执行。

好，说完在新版本中移除zookeeper这个事，咱们在接着聊kafka的其他功能：

kafka支持消息持久化，消费端是主动拉取数据，消费状态和订阅关系由客户端负责维护，**消息消费完后，不会立即删除，会保留历史消息**。因此支持多订阅时，消息只会存储一份就可以。

- **broker**：kafka集群中包含一个或者多个服务实例（节点），这种服务实例被称为broker（一个broker就是一个节点/一个服务器）；
- **topic**：每条发布到kafka集群的消息都属于某个类别，这个类别就叫做topic；
- **partition**：partition是一个物理上的概念，每个topic包含一个或者多个partition；
- **segment**：一个partition当中存在多个segment文件段，每个segment分为两部分，.log文件和 .index 文件，其中 .index 文件是索引文件，主要用于快速查询， .log 文件当中数据的偏移量位置；
- **producer**：消息的生产者，负责发布消息到 kafka 的 broker 中；
- **consumer**：消息的消费者，向 kafka 的 broker 中读取消息的客户端；
- **consumer group**：消费者组，每一个 consumer 属于一个特定的 consumer group（可以为每个consumer指定 groupName）；
- **.log**：存放数据文件；
- **.index**：存放.log文件的索引数据。

## 2. Kafka 主要组件

### 1. producer（生产者）

producer主要是用于生产消息，是kafka当中的消息生产者，生产的消息通过topic进行归类，保存到kafka的broker里面去。

### 2. topic（主题）

- kafka将消息以topic为单位进行归类；
- topic特指kafka处理的消息源（feeds of messages）的不同分类；
- topic是一种分类或者发布的一些列记录的名义上的名字。kafka主题始终是支持多用户订阅的；也就是说，一 个主题可以有零个，一个或者多个消费者订阅写入的数据；
- 在kafka集群中，可以有无数的主题；
- 生产者和消费者消费数据一般以主题为单位。更细粒度可以到分区级别。

### 3. partition（分区）

- kafka当中，topic是消息的归类，一个topic可以有多个分区（partition），每个分区保存部分topic的数据，所有的partition当中的数据全部合并起来，就是一个topic当中的所有的数据。
- 一个broker服务下，可以创建多个分区，broker数与分区数没有关系；
- 在kafka中，每一个分区会有一个编号：编号从0开始。
- **每一个分区内的数据是有序的，但全局的数据不能保证是有序的。**（有序是指生产什么样顺序，消费时也是什么样的顺序）

### 4. consumer（消费者）

consumer是kafka当中的消费者，主要用于消费kafka当中的数据，消费者一定是归属于某个消费组中的。

### 5. consumer group（消费者组）

消费者组由一个或者多个消费者组成，**同一个组中的消费者对于同一条消息只消费一次**。

每个消费者都属于某个消费者组，如果不指定，那么所有的消费者都属于默认的组。

每个消费者组都有一个ID，即group ID。组内的所有消费者协调在一起来消费一个订阅主题( topic)的所有分区(partition)。当然，**每个分区只能由同一个消费组内的一个消费者(consumer)来消费，可以由不同的消费组来消费**。

**partition数量决定了每个consumer group中并发消费者的最大数量。**如下图：

![img](https://cdn.nlark.com/yuque/0/2023/png/22219483/1678954737008-dbc8fb5a-5971-4c10-9d34-e4c9b7fe282d.png)

如上面左图所示，如果只有两个分区，即使一个组内的消费者有4个，也会有两个空闲的。

如上面右图所示，有4个分区，每个消费者消费一个分区，并发量达到最大4。

在来看如下一幅图：

![img](https://cdn.nlark.com/yuque/0/2023/png/22219483/1678955044673-5e1e5a56-f934-42f0-99f4-85183bfa7629.png)

如上图所示，不同的消费者组消费同一个topic，这个topic有4个分区，分布在两个节点上。左边的 消费组1有两个消费者，每个消费者就要消费两个分区才能把消息完整的消费完，右边的 消费组2有四个消费者，每个消费者消费一个分区即可。

**总结下kafka中分区与消费组的关系：**

**消费组**： 由一个或者多个消费者组成，同一个组中的消费者对于同一条消息只消费一次。 某一个主题下的分区数，对于消费该主题的同一个消费组下的消费者数量，应该小于等于该主题下的分区数。

如：某一个主题有4个分区，那么消费组中的消费者应该小于等于4，而且最好与分区数成整数倍 1 2 4 这样。同一个分区下的数据，在同一时刻，不能同一个消费组的不同消费者消费。

**总结**：分区数越多，同一时间可以有越多的消费者来进行消费，消费数据的速度就会越快，提高消费的性能。

### 6. partition replicas（分区副本）

kafka 中的分区副本如下图所示：

![img](https://cdn.nlark.com/yuque/0/2023/png/22219483/1678955580026-45695d51-ef29-4981-9abb-c4928730cd3d.png)

- **副本数**（replication-factor）：控制消息保存在几个broker（服务器）上，一般情况下副本数等于broker的个数。
- 一个broker服务下，不可以创建多个副本因子。**创建主题时，副本因子应该小于等于可用的broker数。**
- 副本因子操作以分区为单位的。每个分区都有各自的主副本和从副本；
- 主副本叫做leader，从副本叫做 follower（在有多个副本的情况下，kafka会为同一个分区下的所有分区，设定角色关系：一个leader和N个 follower），处于同步状态的副本叫做in-sync-replicas(ISR);
- follower通过拉的方式从leader同步数据。 消费者和生产者都是从leader读写数据，不与follower交互。
- **副本因子的作用**：让kafka读取数据和写入数据时的可靠性。
- 副本因子是包含本身，同一个副本因子不能放在同一个broker中。
- 如果某一个分区有三个副本因子，就算其中一个挂掉，那么只会剩下的两个中，选择一个leader，但不会在其他的broker中，另启动一个副本（因为在另一台启动的话，存在数据传递，只要在机器之间有数据传递，就会长时间占用网络IO，kafka是一个高吞吐量的消息系统，这个情况不允许发生）所以不会在另一个broker中启动。
- 如果所有的副本都挂了，生产者如果生产数据到指定分区的话，将写入不成功。
- lsr表示：当前可用的副本。

### 7. segment文件

一个partition当中由多个segment文件组成，每个segment文件，包含两部分，一个是 .log 文件，另外一个是 .index 文件，其中 .log 文件包含了我们发送的数据存储，.index 文件，记录的是我们.log文件的数据索引值，以便于我们加快数据的查询速度。

**索引文件与数据文件的关系**

既然它们是一一对应成对出现，必然有关系。索引文件中元数据指向对应数据文件中message的物理偏移地址。

比如索引文件中 3,497 代表：数据文件中的第三个message，它的偏移地址为497。

再来看数据文件中，Message 368772表示：在全局partiton中是第368772个message。

注：segment index file 采取稀疏索引存储方式，减少索引文件大小，通过mmap（内存映射）可以直接内存操作，稀疏索引为数据文件的每个对应message设置一个元数据指针，它比稠密索引节省了更多的存储空间，但查找起来需要消耗更多的时间。

.index 与 .log 对应关系如下：

![img](https://cdn.nlark.com/yuque/0/2023/png/22219483/1678956071469-d5f329ae-fb74-43a5-bf17-0e26db97a63f.png)

上图左半部分是索引文件，里面存储的是一对一对的key-value，其中key是消息在数据文件（对应的log文件）中的编号，比如“1,3,6,8……”，分别表示在log文件中的第1条消息、第3条消息、第6条消息、第8条消息……

那么为什么在index文件中这些编号不是连续的呢？这是因为index文件中并没有为数据文件中的每条消息都建立索引，而是采用了稀疏存储的方式，每隔一定字节的数据建立一条索引。 这样避免了索引文件占用过多的空间，从而可以将索引文件保留在内存中。 但缺点是没有建立索引的Message也不能一次定位到其在数据文件的位置，从而需要做一次顺序扫描，但是这次顺序扫描的范围就很小了。

value 代表的是在全局partiton中的第几个消息。

以索引文件中元数据 3,497 为例，其中3代表在右边log数据文件中从上到下第3个消息，497表示该消息的物理偏移地址（位置）为497(也表示在全局partiton表示第497个消息-顺序写入特性)。

log日志目录及组成 kafka在我们指定的log.dir目录下，会创建一些文件夹；名字是 （主题名字-分区名） 所组成的文件夹。 在（主题名字-分区名）的目录下，会有两个文件存在，如下所示：

\#索引文件

00000000000000000000.index

\#日志内容

00000000000000000000.log

在目录下的文件，会根据log日志的大小进行切分，.log文件的大小为1G的时候，就会进行切分文件；如下：

-rw-r--r--. 1 root root 389k  1月  17  18:03   00000000000000000000.index

-rw-r--r--. 1 root root 1.0G  1月  17  18:03   00000000000000000000.log

-rw-r--r--. 1 root root  10M  1月  17  18:03   00000000000000077894.index

-rw-r--r--. 1 root root 127M  1月  17  18:03   00000000000000077894.log

在kafka的设计中，将offset值作为了文件名的一部分。

segment文件命名规则：partion全局的第一个segment从0开始，后续每个segment文件名为上一个全局 partion的最大offset（偏移message数）。数值最大为64位long大小，20位数字字符长度，没有数字就用 0 填充。

通过索引信息可以快速定位到message。通过index元数据全部映射到内存，可以避免segment File的IO磁盘操作；

通过索引文件稀疏存储，可以大幅降低index文件元数据占用空间大小。

**稀疏索引**：为了数据创建索引，但范围并不是为每一条创建，而是为某一个区间创建；好处：就是可以减少索引值的数量。 不好的地方：找到索引区间之后，要得进行第二次处理。

### 8. message的物理结构

生产者发送到kafka的每条消息，都被kafka包装成了一个message

message 的物理结构如下图所示：

![img](https://cdn.nlark.com/yuque/0/2023/png/22219483/1678957877666-54e2dbed-d0ba-4364-ba51-082957f62786.png)

所以生产者发送给kafka的消息并不是直接存储起来，而是经过kafka的包装，每条消息都是上图这个结构，只有最后一个字段才是真正生产者发送的消息数据。

# 四、Kafka集群操作

## 1. 创建topic

创建一个名字为test的主题， 有三个分区，有两个副本：

```shell
bin/kafka-topics.sh --create --zookeeper node01:2181 --replication-factor 2 --partitions 3 --topic test
```

## 2. 查看主题命令

查看kafka当中存在的主题：

```shell
bin/kafka-topics.sh  --list --zookeeper node01:2181,node02:2181,node03:2181
```

## 3. 生产者生产数据

模拟生产者来生产数据:

```shell
bin/kafka-console-producer.sh --broker-list node01:9092,node02:9092,node03:9092 --topic test
```

## 4. 消费者消费数据

执行以下命令来模拟消费者进行消费数据:

```shell
bin/kafka-console-consumer.sh --from-beginning --topic test  --zookeeper node01:2181,node02:2181,node03:2181
```

## 5. 运行describe topics命令

执行以下命令运行describe查看topic的相关信息:

```shell
bin/kafka-topics.sh --describe --zookeeper node01:2181 --topic test
```

结果说明：

- 这是输出的解释。第一行给出了所有分区的摘要，每个附加行提供有关一个分区的信息。由于我们只有一个分 区用于此主题，因此只有一行。
- “leader”是负责给定分区的所有读取和写入的节点。每个节点将成为随机选择的分区部分的领导者。（因为在kafka中 如果有多个副本的话，就会存在leader和follower的关系，表示当前这个副本为leader所在的broker是哪一个）
- “replicas”是复制此分区日志的节点列表，无论它们是否为领导者，或者即使它们当前处于活动状态。（所有副本列表0,1,2）
- “isr”是“同步”复制品的集合。这是副本列表的子集，该列表当前处于活跃状态并且已经被领导者捕获。（可用的列表数）

## 6. 增加topic分区数

执行以下命令可以增加topic分区数:

```shell
bin/kafka-topics.sh --zookeeper zkhost:port --alter --topic topicName --partitions 8
```

## 7. 增加配置

动态修改kakfa的配置:

```shell
bin/kafka-topics.sh --zookeeper node01:2181 --alter --topic test --config flush.messages=1
```

## 8. 删除配置

动态删除kafka集群配置:

```shell
bin/kafka-topics.sh --zookeeper node01:2181 --alter --topic test --delete-config flush.messages
```

## 9. 删除topic

目前删除topic在默认情况下知识打上一个删除的标记，在重新启动kafka后才删除。

如果需要立即删除，则需要在server.properties中配置：

delete.topic.enable=true

然后执行以下命令进行删除topic:

```shell
kafka-topics.sh --zookeeper zkhost:port --delete --topic topicName
```

# 五、Kafka的GoAPI操作

使用Go语言操作Kafka可以使用Sarama库，Sarama是一个专门为Kafka设计的Go语言库，可以帮助您轻松地与Kafka进行交互。

以下是一些基本的Kafka Go API操作示例：

## 1. 生产者操作

首先，您需要创建一个生产者并配置它，然后可以使用它来将消息发送到Kafka集群中的主题：

```go
import (
    "fmt"
    "github.com/Shopify/sarama"
)

func main() {
    config := sarama.NewConfig()
    producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
    if err != nil {
        fmt.Println("Error creating producer:", err)
        return
    }
    defer producer.Close()

    message := &sarama.ProducerMessage{
        Topic: "test",
        Value: sarama.StringEncoder("hello world"),
    }

    partition, offset, err := producer.SendMessage(message)
    if err != nil {
        fmt.Println("Error sending message:", err)
        return
    }

    fmt.Printf("Message sent to partition %d at offset %d\n", partition, offset)
}
```

## 2. 消费者操作

您可以使用Sarama创建一个消费者并订阅一个主题，然后在一个无限循环中读取消息：

```go
import (
    "fmt"
    "github.com/Shopify/sarama"
)

func main() {
    config := sarama.NewConfig()
    consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
    if err != nil {
        fmt.Println("Error creating consumer:", err)
        return
    }
    defer consumer.Close()

    partitionConsumer, err := consumer.ConsumePartition("test", 0, sarama.OffsetNewest)
    if err != nil {
        fmt.Println("Error consuming partition:", err)
        return
    }
    defer partitionConsumer.Close()

    for {
        select {
        case message := <-partitionConsumer.Messages():
            fmt.Printf("Received message with value %s\n", string(message.Value))
        case err := <-partitionConsumer.Errors():
            fmt.Println("Error consuming message:", err)
        }
    }
}
```

这些示例可以帮助您开始使用Kafka的Go API。您还可以参考Sarama文档以了解更多操作。

Sarama: https://github.com/Shopify/sarama

# 六、Kafka中的数据不丢失机制

## 1. 生产者生产数据不丢失

### 发送消息方式

生产者发送给kafka数据，可以采用同步方式或异步方式

**同步方式：**

发送一批数据给kafka后，等待kafka返回结果：

- 生产者等待10s，如果broker没有给出ack响应，就认为失败。
- 生产者重试3次，如果还没有响应，就报错.

异步方式：

- 发送一批数据给kafka，只是提供一个回调函数：
- 先将数据保存在生产者端的buffer中。buffer大小是2万条 。
- 满足数据阈值或者数量阈值其中的一个条件就可以发送数据。
- 发送一批数据的大小是500条。

注：如果broker迟迟不给ack，而buffer又满了，开发者可以设置是否直接清空buffer中的数据。

### ack机制（确认机制）

生产者数据发送出去，需要服务端返回一个确认码，即ack响应码；ack的响应有三个状态值0,1，-1

- 0：生产者只负责发送数据，不关心数据是否丢失，丢失的数据，需要再次发送
- 1：partition的leader收到数据，不管follow是否同步完数据，响应的状态码为1
- -1：所有的从节点都收到数据，响应的状态码为-1

如果broker端一直不返回ack状态，producer永远不知道是否成功；producer可以设置一个超时时间10s，超过时间认为失败。

## 2. broker中数据不丢失

在broker中，保证数据不丢失主要是通过副本因子（冗余），防止数据丢失。

## 3. 消费者消费数据不丢失

在消费者消费数据的时候，只要每个消费者记录好offset值即可，就能保证数据不丢失。也就是需要我们自己维护偏移量(offset)，可保存在 Redis 中。

# 七、Kafka配置文件说明

## Server.properties配置文件说明：

```shell
#broker的全局唯一编号，不能重复
broker.id=0
#用来监听链接的端口，producer或consumer将在此端口建立连接
port=9092
#处理网络请求的线程数量
num.network.threads=3
#用来处理磁盘IO的线程数量
num.io.threads=8
#发送套接字的缓冲区大小
socket.send.buffer.bytes=102400
#接受套接字的缓冲区大小
socket.receive.buffer.bytes=102400
#请求套接字的缓冲区大小
socket.request.max.bytes=104857600
#kafka运行日志存放的路径
log.dirs=/export/data/kafka/
#topic在当前broker上的分片个数
num.partitions=2
#用来恢复和清理data下数据的线程数量
num.recovery.threads.per.data.dir=1
#segment文件保留的最长时间，超时将被删除
log.retention.hours=168
#滚动生成新的segment文件的最大时间
log.roll.hours=1
#日志文件中每个segment的大小，默认为1G
log.segment.bytes=1073741824
#周期性检查文件大小的时间
log.retention.check.interval.ms=300000
#日志清理是否打开
log.cleaner.enable=true
#broker需要使用zookeeper保存meta数据
zookeeper.connect=zk01:2181,zk02:2181,zk03:2181
#zookeeper链接超时时间
zookeeper.connection.timeout.ms=6000
#partion buffer中，消息的条数达到阈值，将触发flush到磁盘
log.flush.interval.messages=10000
#消息buffer的时间，达到阈值，将触发flush到磁盘
log.flush.interval.ms=3000
#删除topic需要server.properties中设置delete.topic.enable=true否则只是标记删除
delete.topic.enable=true
#此处的host.name为本机IP(重要),如果不改,则客户端会抛出:Producer connection to localhost:9092 unsuccessful 错误!
host.name=kafka01

advertised.host.name=192.168.140.128

producer生产者配置文件说明
#指定kafka节点列表，用于获取metadata，不必全部指定
metadata.broker.list=node01:9092,node02:9092,node03:9092
# 指定分区处理类。默认kafka.producer.DefaultPartitioner，表通过key哈希到对应分区
#partitioner.class=kafka.producer.DefaultPartitioner
# 是否压缩，默认0表示不压缩，1表示用gzip压缩，2表示用snappy压缩。压缩后消息中会有头来指明消息压缩类型，故在消费者端消息解压是透明的无需指定。
compression.codec=none
# 指定序列化处理类
serializer.class=kafka.serializer.DefaultEncoder
# 如果要压缩消息，这里指定哪些topic要压缩消息，默认empty，表示不压缩。
#compressed.topics=
# 设置发送数据是否需要服务端的反馈,有三个值0,1,-1
# 0: producer不会等待broker发送ack 
# 1: 当leader接收到消息之后发送ack 
# -1: 当所有的follower都同步消息成功后发送ack. 
request.required.acks=0 
# 在向producer发送ack之前,broker允许等待的最大时间 ，如果超时,broker将会向producer发送一个error ACK.意味着上一次消息因为某种原因未能成功(比如follower未能同步成功) 
request.timeout.ms=10000
# 同步还是异步发送消息，默认“sync”表同步，"async"表异步。异步可以提高发送吞吐量,
也意味着消息将会在本地buffer中,并适时批量发送，但是也可能导致丢失未发送过去的消息
producer.type=sync
# 在async模式下,当message被缓存的时间超过此值后,将会批量发送给broker,默认为5000ms
# 此值和batch.num.messages协同工作.
queue.buffering.max.ms = 5000
# 在async模式下,producer端允许buffer的最大消息量
# 无论如何,producer都无法尽快的将消息发送给broker,从而导致消息在producer端大量沉积
# 此时,如果消息的条数达到阀值,将会导致producer端阻塞或者消息被抛弃，默认为10000
queue.buffering.max.messages=20000
# 如果是异步，指定每次批量发送数据量，默认为200
batch.num.messages=500
# 当消息在producer端沉积的条数达到"queue.buffering.max.meesages"后 
# 阻塞一定时间后,队列仍然没有enqueue(producer仍然没有发送出任何消息) 
# 此时producer可以继续阻塞或者将消息抛弃,此timeout值用于控制"阻塞"的时间 
# -1: 无阻塞超时限制,消息不会被抛弃 
# 0:立即清空队列,消息被抛弃 
queue.enqueue.timeout.ms=-1

# 当producer接收到error ACK,或者没有接收到ACK时,允许消息重发的次数 
# 因为broker并没有完整的机制来避免消息重复,所以当网络异常时(比如ACK丢失) 
# 有可能导致broker接收到重复的消息,默认值为3.
message.send.max.retries=3
# producer刷新topic metada的时间间隔,producer需要知道partition leader的位置,以及当前topic的情况 
# 因此producer需要一个机制来获取最新的metadata,当producer遇到特定错误时,将会立即刷新 
# (比如topic失效,partition丢失,leader失效等),此外也可以通过此参数来配置额外的刷新机制，默认值600000 
topic.metadata.refresh.interval.ms=60000
```

## consumer消费者配置详细说明:

```shell
# zookeeper连接服务器地址
zookeeper.connect=zk01:2181,zk02:2181,zk03:2181
# zookeeper的session过期时间，默认5000ms，用于检测消费者是否挂掉
zookeeper.session.timeout.ms=5000
#当消费者挂掉，其他消费者要等该指定时间才能检查到并且触发重新负载均衡
zookeeper.connection.timeout.ms=10000
# 指定多久消费者更新offset到zookeeper中。注意offset更新时基于time而不是每次获得的消息。一旦在更新zookeeper发生异常并重启，将可能拿到已拿到过的消息
zookeeper.sync.time.ms=2000
#指定消费 
group.id=itcast
# 当consumer消费一定量的消息之后,将会自动向zookeeper提交offset信息 
# 注意offset信息并不是每消费一次消息就向zk提交一次,而是现在本地保存(内存),并定期提交,默认为true
auto.commit.enable=true
# 自动更新时间。默认60 * 1000
auto.commit.interval.ms=1000
# 当前consumer的标识,可以设定,也可以有系统生成,主要用来跟踪消息消费情况,便于观察
conusmer.id=xxx 
# 消费者客户端编号，用于区分不同客户端，默认客户端程序自动产生
client.id=xxxx
# 最大取多少块缓存到消费者(默认10)
queued.max.message.chunks=50
# 当有新的consumer加入到group时,将会reblance,此后将会有partitions的消费端迁移到新  的consumer上,如果一个consumer获得了某个partition的消费权限,那么它将会向zk注册 "Partition Owner registry"节点信息,但是有可能此时旧的consumer尚没有释放此节点, 此值用于控制,注册节点的重试次数. 
rebalance.max.retries=5
# 获取消息的最大尺寸,broker不会像consumer输出大于此值的消息chunk 每次feth将得到多条消息,此值为总大小,提升此值,将会消耗更多的consumer端内存
fetch.min.bytes=6553600
# 当消息的尺寸不足时,server阻塞的时间,如果超时,消息将立即发送给consumer
fetch.wait.max.ms=5000
socket.receive.buffer.bytes=655360
# 如果zookeeper没有offset值或offset值超出范围。那么就给个初始的offset。有smallest、largest、anything可选，分别表示给当前最小的offset、当前最大的offset、抛异常。默认largest
auto.offset.reset=smallest
# 指定序列化处理类
derializer.class=kafka.serializer.DefaultDecoder
```

# 八、CAP理论

## 1. 分布式系统当中的CAP理论

分布式系统（distributed system）正变得越来越重要，大型网站几乎都是分布式的。

分布式系统的最大难点，就是各个节点的状态如何同步。

为了解决各个节点之间的状态同步问题，在1998年，由加州大学的计算机科学家 Eric Brewer 提出分布式系统的三个指标，分别是:

- Consistency：一致性
- Availability：可用性
- Partition tolerance：分区容错性

Eric Brewer 说，这三个指标不可能同时做到。最多只能同时满足其中两个条件，这个结论就叫做 CAP 定理。

**CAP理论是指：分布式系统中，一致性、可用性和分区容忍性最多只能同时满足两个。**

### 一致性：Consistency

- 通过某个节点的写操作结果对后面通过其它节点的读操作可见
- 如果更新数据后，并发访问情况下后续读操作可立即感知该更新，称为强一致性
- 如果允许之后部分或者全部感知不到该更新，称为弱一致性
- 若在之后的一段时间（通常该时间不固定）后，一定可以感知到该更新，称为最终一致性

### 可用性：Availability

- 任何一个没有发生故障的节点必须在有限的时间内返回合理的结果

### 分区容错性：Partition tolerance

- 部分节点宕机或者无法与其它节点通信时，各分区间还可保持分布式系统的功能

一般而言，都要求保证分区容忍性。所以在CAP理论下，更多的是需要在可用性和一致性之间做权衡。

![img](https://cdn.nlark.com/yuque/0/2023/png/22219483/1678959165472-de5a638c-221d-44e3-84d4-2ef9c68ae996.png)

## 2. 分区容错性：Partition tolerance

先看 Partition tolerance，中文叫做"分区容错"。

大多数分布式系统都分布在多个子网络。每个子网络就叫做一个区（partition）。分区容错的意思是，区间通信可能失败。比如，一台服务器放在中国，另一台服务器放在美国，这就是两个区，它们之间可能无法通信。

![img](https://cdn.nlark.com/yuque/0/2023/png/22219483/1678959312317-745a4a64-5f05-4446-a44f-4b7e16860b10.png)

上图中，G1 和 G2 是两台跨区的服务器。G1 向 G2 发送一条消息，G2 可能无法收到。系统设计的时候，必须考虑到这种情况。

一般来说，分区容错无法避免，因此可以认为 CAP 的 P 总是存在的。即永远可能存在分区容错这个问题

## 3. 一致性：Consistency

Consistency 中文叫做"一致性"。意思是，写操作之后的读操作，必须返回该值。举例来说，某条记录是 v0，用户向 G1 发起一个写操作，将其改为 v1。  

![img](https://cdn.nlark.com/yuque/0/2023/png/22219483/1678959402478-25467153-a749-4ba0-a94f-5dbe349d4c0b.png)接下来，用户的读操作就会得到 v1。这就叫一致性。![img](https://cdn.nlark.com/yuque/0/2023/png/22219483/1678959479995-37b3fe9b-678d-4182-9709-3f1b0766fa82.png)  问题是，用户有可能向 G2 发起读操作，由于 G2 的值没有发生变化，因此返回的是 v0。G1 和 G2 读操作的结果不一致，这就不满足一致性了。

![img](https://cdn.nlark.com/yuque/0/2023/png/22219483/1678959874110-c67c8570-f5d2-412e-913e-676ab0fa4f16.png)

为了让 G2 也能变为 v1，就要在 G1 写操作的时候，让 G1 向 G2 发送一条消息，要求 G2 也改成 v1。

![img](https://cdn.nlark.com/yuque/0/2023/png/22219483/1678959891726-2b802610-9d7f-4ab7-9dee-74358da28763.png)

这样的话，用户向 G2 发起读操作，也能得到 v1。

![img](https://cdn.nlark.com/yuque/0/2023/png/22219483/1678959899186-4c565218-40a5-4ff0-843b-2b5c601e52be.png)

## 4. 可用性：Availability

Availability 中文叫做"可用性"，意思是只要收到用户的请求，服务器就必须给出回应。用户可以选择向 G1 或 G2 发起读操作。不管是哪台服务器，只要收到请求，就必须告诉用户，到底是 v0 还是 v1，否则就不满足可用性。

# 九、Kafka中的CAP机制

kafka是一个分布式的消息队列系统，既然是一个分布式的系统，那么就一定满足CAP定律，那么在kafka当中是如何遵循CAP定律的呢？kafka满足CAP定律当中的哪两个呢？

**kafka满足的是CAP定律当中的CA，其中Partition tolerance通过的是一定的机制尽量的保证分区容错性。**

**其中C表示的是数据一致性。A表示数据可用性。**

kafka首先将数据写入到不同的分区里面去，每个分区又可能有好多个副本，数据首先写入到leader分区里面去，**读写的操作都是与leader分区进行通信，保证了数据的****一致性****原则，也就是满足了Consistency原则**。然后**kafka通过分区副本机制，来保证了kafka当中数据的可用性**。但是也存在另外一个问题，就是副本分区当中的数据与leader当中的数据存在差别的问题如何解决，这个就是Partition tolerance的问题。

**kafka为了解决Partition tolerance的问题，使用了ISR的同步策略，来尽最大可能减少Partition tolerance的问题。**

每个leader会维护一个ISR（a set of in-sync replicas，基本同步）列表。

ISR列表主要的作用就是决定哪些副本分区是可用的，也就是说可以将leader分区里面的数据同步到副本分区里面去，决定一个副本分区是否可用的条件有两个：

- replica.lag.time.max.ms=10000 副本分区与主分区心跳时间延迟
- replica.lag.max.messages=4000 副本分区与主分区消息同步最大差

produce 请求被认为完成时的确认值：request.required.acks=0。

- ack=0：producer不等待broker同步完成的确认，继续发送下一条(批)信息。
- ack=1（默认）：producer要等待leader成功收到数据并得到确认，才发送下一条message。
- ack=-1：producer得到follwer确认，才发送下一条数据。



# 十一、Kafka面经

## 1. 为什么要使用 kafka？

- **缓冲和削峰**：上游数据时有突发流量，下游可能扛不住，或者下游没有足够多的机器来保证冗余，kafka在中间可以起到一个缓冲的作用，把消息暂存在kafka中，下游服务就可以按照自己的节奏进行慢慢处理。
- **解耦和扩展性**：项目开始的时候，并不能确定具体需求。消息队列可以作为一个接口层，解耦重要的业务流程。只需要遵守约定，针对数据编程即可获取扩展能力。
- **冗余**：可以采用一对多的方式，一个生产者发布消息，可以被多个订阅topic的服务消费到，供多个毫无关联的业务使用。
- **健壮性**：消息队列可以堆积请求，所以消费端业务即使短时间死掉，也不会影响主要业务的正常进行。
- **异步通信**：很多时候，用户不想也不需要立即处理消息。消息队列提供了异步处理机制，允许用户把一个消息放入队列，但并不立即处理它。想向队列中放入多少消息就放多少，然后在需要的时候再去处理它们。

## 2. Kafka消费过的消息如何再消费？

kafka消费消息的offset是定义在zookeeper中的， 如果想重复消费kafka的消息，可以在redis中自己记录offset的checkpoint点（n个），当想重复消费消息时，通过读取redis中的checkpoint点进行zookeeper的offset重设，这样就可以达到重复消费消息的目的了

## 3. kafka的数据是放在磁盘上还是内存上，为什么速度会快？

kafka使用的是磁盘存储。

速度快是因为：

- 顺序写入：因为硬盘是机械结构，每次读写都会寻址->写入，其中寻址是一个“机械动作”，它是耗时的。所以硬盘 “讨厌”随机I/O， 喜欢顺序I/O。为了提高读写硬盘的速度，Kafka就是使用顺序I/O。
- Memory Mapped Files（内存映射文件）：64位操作系统中一般可以表示20G的数据文件，它的工作原理是直接利用操作系统的Page来实现文件到物理内存的直接映射。完成映射之后你对物理内存的操作会被同步到硬盘上。
- Kafka高效文件存储设计： 

- - **Kafka把topic中一个parition大文件分成多个小文件段**，通过多个小文件段，就容易定期清除或删除已经消费完文件，减少磁盘占用。
  - 通过索引信息可以快速定位message和确定response的 大 小。
  - **通过index元数据全部映射到memory**（内存映射文件）， 可以避免segment file的IO磁盘操作。
  - **通过索引文件稀疏存储**，可以大幅降低index文件元数据占用空间大小。

注：

- Kafka解决查询效率的手段之一是将数据文件分段，比如有100条Message，它们的offset是从0到99。假设将数据文件分成5段，第一段为0-19，第二段为20-39，以此类推，每段放在一个单独的数据文件里面，数据文件以该段中 小的offset命名。这样在查找指定offset的Message的时候，用二分查找就可以定位到该Message在哪个段中。
- 为数据文件建 索引数据文件分段 使得可以在一个较小的数据文件中查找对应offset的Message 了，但是这依然需要顺序扫描才能找到对应offset的Message。
- 为了进一步提高查找的效率，Kafka为每个分段后的数据文件建立了索引文件，文件名与数据文件的名字是一样的，只是文件扩展名为.index。

## 4. Kafka数据怎么保障不丢失？

分三个点说，一个是**生产者端**，一个**消费者端**，一个**broker端**。

### 生产者数据的不丢失

kafka的ack机制：在kafka发送数据的时候，每次发送消息都会有一个确认反馈机制，确保消息正常的能够被收到，其中状态有0，1，-1。

**如果是同步模式：**

ack设置为0，风险很大，一般不建议设置为0。即使设置为1，也会随着leader宕机丢失数据。所以如果要严格保证生产端数据不丢失，可设置为-1。

**如果是异步模式：**

也会考虑ack的状态，除此之外，异步模式下的有个buffer，通过buffer来进行控制数据的发送，有两个值来进行控制，时间阈值与消息的数量阈值，如果buffer满了数据还没有发送出去，有个选项是配置是否立即清空buffer。可以设置为-1，永久阻塞，也就数据不再生产。异步模式下，即使设置为-1。也可能因为程序员的不科学操作，操作数据丢失，比如kill -9，但这是特别的例外情况。

注：

ack=0：producer不等待broker同步完成的确认，继续发送下一条(批)信息。

ack=1（默认）：producer要等待leader成功收到数据并得到确认，才发送下一条message。

ack=-1：producer得到follwer确认，才发送下一条数据。

### 消费者数据的不丢失

通过offset commit 来保证数据的不丢失，kafka自己记录了每次消费的offset数值，下次继续消费的时候，会接着上次的offset进行消费。

而offset的信息在kafka0.8版本之前保存在zookeeper中，在0.8版本之后保存到topic中，即使消费者在运行过程中挂掉了，再次启动的时候会找到offset的值，找到之前消费消息的位置，接着消费，由于 offset 的信息写入的时候并不是每条消息消费完成后都写入的，所以这种情况有可能会造成重复消费，但是不会丢失消息。

唯一例外的情况是，我们在程序中给原本做不同功能的两个consumer组设置KafkaSpoutConfig.bulider.setGroupid的时候设置成了一样的groupid，这种情况会导致这两个组共享同一份数据，就会产生组A消费partition1，partition2中的消息，组B消费partition3的消息，这样每个组消费的消息都会丢失，都是不完整的。 为了保证每个组都独享一份消息数据，groupid一定不要重复才行。

### kafka集群中的broker的数据不丢失

每个broker中的partition我们一般都会设置有replication（副本）的个数，生产者写入的时候首先根据分发策略（有partition按partition，有key按key，都没有轮询）写入到leader中，follower（副本）再跟leader同步数据，这样有了备份，也可以保证消息数据的不丢失。

## 5. 采集数据为什么选择kafka？

采集层 主要可以使用Flume, Kafka等技术。

Flume：Flume 是管道流方式，提供了很多的默认实现，让用户通过参数部署，及扩展API.

Kafka：Kafka是一个可持久化的分布式的消息队列。 Kafka 是一个非常通用的系统。你可以有许多生产者和很多的消费者共享多个主题Topics。

相比之下,Flume是一个专用工具被设计为旨在往HDFS，HBase发送数据。它对HDFS有特殊的优化，并且集成了Hadoop的安全特性。

所以，Cloudera 建议如果数据被多个系统消费的话，使用kafka；如果数据被设计给Hadoop使用，使用Flume。

## 6. kafka 重启是否会导致数据丢失？

1. kafka是将数据写到磁盘的，一般数据不会丢失。
2. 但是在重启kafka过程中，如果有消费者消费消息，那么kafka如果来不及提交offset，可能会造成数据的不准确（丢失或者重复消费）。

## 7. kafka 宕机了如何解决？

1. 先考虑业务是否受到影响
2. kafka 宕机了，首先我们考虑的问题应该是所提供的服务是否因为宕机的机器而受到影响，如果服务提供没问题，如果实现做好了集群的容灾机制，那么这块就不用担心了。
3. 节点排错与恢复  想要恢复集群的节点，主要的步骤就是通过日志分析来查看节点宕机的原因，从而解决，重新恢复节点。

## 8. 为什么Kafka不支持读写分离？

在 Kafka 中，生产者写入消息、消费者读取消息的操作都是与 leader 副本进行交互的，从 而实现的是一种主写主读的生产消费模型。Kafka 并不支持 主写从读，因为主写从读有 2 个很明显的缺点:

1. 数据一致性问题：数据从主节点转到从节点必然会有一个延时的时间窗口，这个时间 窗口会导致主从节点之间的数据不一致。某一时刻，在主节点和从节点中 A 数据的值都为 X， 之后将主节点中 A 的值修改为 Y，那么在这个变更通知到从节点之前，应用读取从节点中的 A 数据的值并不为最新的 Y，由此便产生了数据不一致的问题。
2. 延时问题：类似 Redis 这种组件，数据从写入主节点到同步至从节点中的过程需要经历 网络→主节点内存→网络→从节点内存 这几个阶段，整个过程会耗费一定的时间。而在 Kafka 中，主从同步会比 Redis 更加耗时，它需要经历 网络→主节点内存→主节点磁盘→网络→从节 点内存→从节点磁盘 这几个阶段。对延时敏感的应用而言，主写从读的功能并不太适用。

而kafka的**主写主读**的优点就很多了：

1. 可以简化代码的实现逻辑，减少出错的可能;
2. 将负载粒度细化均摊，与主写从读相比，不仅负载效能更好，而且对用户可控;
3. 没有延时的影响;
4. 在副本稳定的情况下，不会出现数据不一致的情况。

## 9. kafka数据分区和消费者的关系？

每个分区只能由同一个消费组内的一个消费者(consumer)来消费，可以由不同的消费组的消费者来消费，同组的消费者则起到并发的效果。

## 10. kafka的数据offset读取流程

1. 连接ZK集群，从ZK中拿到对应topic的partition信息和partition的Leader的相关信息
2. 连接到对应Leader对应的broker
3. consumer将⾃自⼰己保存的offset发送给Leader
4. Leader根据offset等信息定位到segment（索引⽂文件和⽇日志⽂文件）
5. 根据索引⽂文件中的内容，定位到⽇日志⽂文件中该偏移量量对应的开始位置读取相应⻓长度的数据并返回给consumer

## 11. kafka内部如何保证顺序，结合外部组件如何保证消费者的顺序？

kafka只能保证partition内是有序的，但是partition间的有序是没办法的。爱奇艺的搜索架构，是从业务上把需要有序的打到同⼀个partition。

## 12. Kafka消息数据积压，Kafka消费能力不足怎么处理？

1. 如果是Kafka消费能力不足，则可以考虑增加Topic的分区数，并且同时提升消费组的消费者数量，消费者数=分区数。（两者缺一不可）
2. 如果是下游的数据处理不及时：提高每批次拉取的数量。批次拉取数据过少（拉取数据/处理时间<生产速度），使处理的数据小于生产的数据，也会造成数据积压。

## 13. Kafka单条日志传输大小

kafka对于消息体的大小默认为单条最大值是1M但是在我们应用场景中, 常常会出现一条消息大于1M，如果不对kafka进行配置。则会出现生产者无法将消息推送到kafka或消费者无法去消费kafka里面的数据, 这时我们就要对kafka进行以下配置：server.properties

```shell
replica.fetch.max.bytes: 1048576  broker可复制的消息的最大字节数, 默认为1M
message.max.bytes: 1000012   kafka 会接收单个消息size的最大限制， 默认为1M左右
```

**注意：message.max.bytes必须小于等于replica.fetch.max.bytes，否则就会导致replica之间数据同步失败。**

参考文章：

- https://blog.51cto.com/u_14932245/4340967

- https://zhuanlan.zhihu.com/p/37405836