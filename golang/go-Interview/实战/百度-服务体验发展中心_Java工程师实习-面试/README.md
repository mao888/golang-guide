![img](https://cdn.nlark.com/yuque/0/2022/png/22219483/1661163101342-3d9a6843-74ad-40c4-9e93-60945e9709bf.png)

------

## 王雅晴面经

- https://leetcode.cn/problems/shortest-unsorted-continuous-subarray/

1. **自我介绍**
2. **实习经历**

1. 1. **新乡市卡口项目**

1. 1. 1. **redis缓存减少数据库压力**

1. 1. 1. 1. hash的长度，数不需要设置长度，但是数据量大的话，单key会比较大，影响redis读写性能

1. 1. 1. **rabbitMq削峰，**

1. 1. 1. 1. 调入外部接口需要实时数据的话，会对响应的速度有要求，把参数和结果在数据库做一个记录，为了不影响调用人家接口的速度，放入消息队列里，请求问之后插入到库里边不影响速度，提高了接口的性能
         2. 对接口的影响是什么，跟之前接口的对比？

1. 1. 1. 1. 1. 面试官：节省的时间，是相当于把请求的一些信息存在mysql里面，就几十毫秒，貌似并没有多少提升

1. 1. **怎么学习到的redis、mq、mysql这些组件？**

1. **求平衡二叉树的高度**

1. 1. ![img](https://cdn.nlark.com/yuque/0/2022/png/22219483/1661182065490-1de1a0b1-2b97-4517-a851-c40e3c4849ad.png)
   2. **时间复杂度**

1. **计网**

1. 1. **七层网络模型、四层网络模型、分别有什么协议（io、ip、tcp/udp、https）**

1. 1. 1. ISO 七层模型

- - - - 应用层：网络服务与最终用户的一个接口，常见的协议有：**HTTP FTP  SMTP SNMP DNS**.
      - 表示层：数据的表示、安全、压缩。，确保一个系统的应用层所发送的信息可以被另一个系统的应用层读取。
      - 会话层：建立、管理、终止会话, 对应主机进程，指本地主机与远程主机正在进行的会话.
      - 传输层：定义传输数据的协议端口号，以及流控和差错校验, 协议有**TCP UDP**.
      - 网络层：进行逻辑地址寻址，实现不同网络之间的路径选择, 协议有**ICMP IGMP IP 等**.
      - 数据链路层：在物理层提供比特流服务的基础上，建立相邻结点之间的数据链路。
      - 物理层：建立、维护、断开物理连接。

#### TCP/IP 四层模型

★

- 应用层：对应于 OSI 参考模型的（应用层、表示层、会话层）。
- 传输层: 对应 OSI 的传输层，为应用层实体提供端到端的通信功能，保证了数据包的顺序传送及数据的完整性。
- 网际层：对应于 OSI 参考模型的网络层，主要解决主机到主机的通信问题。
- 网络接口层：与 OSI 参考模型的数据链路层、物理层对应。

#### 五层体系结构



★

- 应用层：对应于 OSI 参考模型的（应用层、表示层、会话层）。
- 传输层：对应 OSI 参考模型的的传输层
- 网络层：对应 OSI 参考模型的的网络层
- 数据链路层：对应 OSI 参考模型的的数据链路层
- 物理层：对应 OSI 参考模型的的物理层。

1. 1. **http的错误码**

- - 200：服务器已成功处理了请求。 通常，这表示服务器提供了请求的网页。
  - 301 ： (永久移动) 请求的网页已永久移动到新位置。 服务器返回此响应(对 GET 或 HEAD 请求的响应)时，会自动将请求者转到新位置。
  - 302：(临时移动) 服务器目前从不同位置的网页响应请求，但请求者应继续使用原有位置来进行以后的请求。
  - 400 ：客户端请求有语法错误，不能被服务器所理解。
  - 403 ：服务器收到请求，但是拒绝提供服务。
  - 404 ：(未找到) 服务器找不到请求的网页。
  - 500： (服务器内部错误) 服务器遇到错误，无法完成请求。

1. 1. **错误码可以自定义吗？**

1. 1. 1. 常见自定义错误码

- - - 暂定都是5位数字
    - 0标识成功,其他都表示错误
    - 错误码按模块按功能场景分级分段

- - - - 前三位表示模块
      - 第四位表示模块下的功能

- - - 示例:

- - - - 商城系统里有交易模块和商品模块

- - - - - 401开头的表示交易模块

- - - - - - 4011开头的表示交易模块里的下单场景需要用到的错误码

- - - - - 402开头的表示商品模块

- - - - - - 4021表示商品模块下的添加商品场景里需要用到的错误码

- - - - - 如果某个场景功能下需要的比较多的错误码，则可以使用其他未被使用的码段，即该场景功能可以拥有多个码段，然后通过添加注释等方式让人理解即可。

- - - 数字 1 开头的错误码表示

- - - - 系统级别的错误

- - - - - 缺少某种字符集，连不上数据库之类的，系统级的错误码不需要分模块，可以按照自增方式进行添加

- - - 数字 4 开头的错误码表示

- - - - API参数校验失败

- - - - - “交易模块下单场景中，订单金额参数不能为空” 可以用 40111 错误码来表示

- - - 数字 5 开头的错误码表示

- - - - 后台业务校验失败

- - - - - “交易模块下单场景中，该用户没有下单权限” 可以用 50111 错误码来表示

- - - 注意

- - - - 数字 4 开头的错误码与数字 5 开头的错误码对应的模块分类需要保持一致，即 4011 表示交易模块下单场景的API错误，5011 表示交易模块下单场景的业务错误错误码按需分配，逐步增加，灵活扩展

1. 1. **http的请求头和响应头里的信息**

1. 1. 1. [https://blog.csdn.net/kxkltey/article/details/106683790?spm=1001.2101.3001.6650.1&utm_medium=distribute.pc_relevant.none-task-blog-2%7Edefault%7ECTRLIST%7ERate-1-106683790-blog-89843937.t0_layer_eslanding_s&depth_1-utm_source=distribute.pc_relevant.none-task-blog-2%7Edefault%7ECTRLIST%7ERate-1-106683790-blog-89843937.t0_layer_eslanding_s&utm_relevant_index=2](https://blog.csdn.net/kxkltey/article/details/106683790?spm=1001.2101.3001.6650.1&utm_medium=distribute.pc_relevant.none-task-blog-2~default~CTRLIST~Rate-1-106683790-blog-89843937.t0_layer_eslanding_s&depth_1-utm_source=distribute.pc_relevant.none-task-blog-2~default~CTRLIST~Rate-1-106683790-blog-89843937.t0_layer_eslanding_s&utm_relevant_index=2)

1. **mysql**

1. 1. **聚簇索引**

1. 1. 1. 在 InnoDB 里，索引B+ Tree的叶子节点存储了整行数据的是主键索引，也被称之为聚簇索引，即将数据存储与索引放到了一块，找到索引也就找到了数据。而索引B+ Tree的叶子节点存储了主键的值的是非主键索引，也被称之为非聚簇索引、二级索引。
      2. 聚簇索引与非聚簇索引的区别：

1. 1. 1. 1. 非聚集索引与聚集索引的区别在于非聚集索引的叶子节点不存储表中的数据，而是存储该列对应的主键（行号） 
         2. 对于InnoDB来说，想要查找数据我们还需要根据主键再去聚集索引中进行查找，这个再根据聚集索引查找数据的过程，我们称为**回表**。第一次索引一般是顺序IO，回表的操作属于随机IO。需要回表的次数越多，即随机IO次数越多，我们就越倾向于使用全表扫描 。通常情况下， **主键索引（聚簇索引）****查询只会查一次**，而**非主键索引（非聚簇索引）****需要回表查询多次**。当然，如果是覆盖索引的话，查一次即可 
         3. 注意：MyISAM无论主键索引还是二级索引都是非聚簇索引，而InnoDB的主键索引是聚簇索引，二级索引是非聚簇索引。我们自己建的索引基本都是非聚簇索引。

1. 1. **使用mysql建表需要考虑哪些呢？**

1. 1. 1. [https://shuen.blog.csdn.net/article/details/107319061?spm=1001.2101.3001.6650.1&utm_medium=distribute.pc_relevant.none-task-blog-2%7Edefault%7ECTRLIST%7ERate-1-107319061-blog-106685471.pc_relevant_multi_platform_whitelistv4eslandingctr&depth_1-utm_source=distribute.pc_relevant.none-task-blog-2%7Edefault%7ECTRLIST%7ERate-1-107319061-blog-106685471.pc_relevant_multi_platform_whitelistv4eslandingctr&utm_relevant_index=2](https://shuen.blog.csdn.net/article/details/107319061?spm=1001.2101.3001.6650.1&utm_medium=distribute.pc_relevant.none-task-blog-2~default~CTRLIST~Rate-1-107319061-blog-106685471.pc_relevant_multi_platform_whitelistv4eslandingctr&depth_1-utm_source=distribute.pc_relevant.none-task-blog-2~default~CTRLIST~Rate-1-107319061-blog-106685471.pc_relevant_multi_platform_whitelistv4eslandingctr&utm_relevant_index=2)

1. 1. **建了一个索引，想看一下有没有生效，用的什么工具？**

1. **linux**

1. 1. **查看某个端口被哪个进程占用了？**

1. 1. 1. netstat -ano |findstr XXX端口号

1. 1. **怎么kill到一个进程**
   2. **如果有一个文本，求一下某个接口的平均耗时？（一个文件里面有很多行，但只有一列就一个数字，求一下这个文件所有列的平均值）**

1. **多个系统串联的时候需要logId概念，有类似的吗？**
2. **这个项目的开发流程是什么？上线呢？**
3. **有没有遇到线上问题？**
4. **操作系统**

1. 1. **进程和线程的区别？**

1. 1. 1. 调度：进程是资源管理的基本单位，线程是程序执行的基本单位。

并发性：不仅进程之间可以并发执行，同一个进程的多个线程之间也可并发执行
切换：线程上下文切换比进程上下文切换要快得多。
拥有资源： 进程是拥有资源的一个独立单位，线程不拥有系统资源，但是可以访问隶属于进程的资源。
系统开销： 创建或撤销进程时，系统都要为之分配或回收系统资源，如内存空间，I/O设备等，OS所付出的开销显著大于在创建或撤销线程时的开销，进程切换的开销也远大于线程切换的开销。

1. 1. **go里边的协程？**

1. 1. 1. 进程：是应用程序的启动实例，每个进程都有独立的内存空间，不同的进程通过进程间的通信方式来通信。

线程：从属于进程，每个进程至少包含一个线程，线程是 CPU 调度的基本单位，多个线程之间可以共享进程的资源并通过共享内存等线程间的通信方式来通信。

协程：为轻量级线程，与线程相比，协程不受操作系统的调度，协程的调度器由用户应用程序提供，协程调度器按照调度策略把协程调度到线程中运行

1. **除了go、学过c++吗？**
2. **编译原理学过吗？**

1. 1. **编译的几个环节？**

1. 1. 1. https://blog.csdn.net/weixin_43999496/article/details/102802765
      2. （1）先要识别出句子中的一个个单词；

（2）分析句子的语法结构；

（3）根据句子的含义进行初步翻译；

（4）对译文进行修饰；

（5）写出最后的译文。

1. **go里面的多个channel**
2. **有没有遇到个内存泄露和pannic的情况？**

1. 1. **开启了一个管道，开启了一个协程从管道里读数据，但不给管道写数据，然后这个协程就会一直堵塞着，造成内存泄露**
   2. https://zhuanlan.zhihu.com/p/469817707

1. 1. 1. 答：go 中的内存泄漏一般都是 goroutine 泄漏，就是 goroutine 没有被关闭，或者没有添加超时控制，让 goroutine 一只处于阻塞状态，不能被 GC。

**内存泄露有下面一些情况**

1）如果 goroutine 在执行时被阻塞而无法退出，就会导致 goroutine 的内存泄漏，一个 goroutine 的最低栈大小为 2KB，在高并发的场景下，对内存的消耗也是非常恐怖的。

2）互斥锁未释放或者造成死锁会造成内存泄漏

3）time.Ticker 是每隔指定的时间就会向通道内写数据。作为循环触发器，必须调用 stop 方法才会停止，从而被 GC 掉，否则会一直占用内存空间。

4）字符串的截取引发临时性的内存泄漏

func main() { 	var str0 = "12345678901234567890" 	str1 := str0[:10] }

5）切片截取引起子切片内存泄漏

func main() { 	var s0 = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} 	s1 := s0[:3] }

6）函数数组传参引发内存泄漏【如果我们在函数传参的时候用到了数组传参，且这个数组够大（我们假设数组大小为 100 万，64 位机上消耗的内存约为 800w 字节，即 8MB 内存），或者该函数短时间内被调用 N 次，那么可想而知，会消耗大量内存，对性能产生极大的影响，如果短时间内分配大量内存，而又来不及 GC，那么就会产生临时性的内存泄漏，对于高并发场景相当可怕。】

**排查方式：**

一般通过 pprof 是 Go 的性能分析工具，在程序运行过程中，可以记录程序的运行信息，可以是 CPU 使用情况、内存使用情况、goroutine 运行情况等，当需要性能调优或者定位 Bug 时候，这些记录的信息是相当重要。

**当然你能说说具体的分析指标更加分咯，有的面试官就喜欢他问什么，你简洁的回答什么，不喜欢巴拉巴拉详细解释一通，比如虾P面试官，不过他考察的内容特别多，可能是为了节约时间。**

1. 1. **pannic：**

1. 1. 1. **调用空指针会引发panic**
      2. **关闭一个已经关闭的管道会引发pannic** 
      3. **从一个nil的管道读会引发panic**
      4. **切片的小标越界会引发panic**

- - 给一个 nil channel 发送数据，造成永远阻塞
  - 从一个 nil channel 接收数据，造成永远阻塞
  - 给一个已经关闭的 channel 发送数据，引起 panic
  - 从一个已经关闭的 channel 接收数据，如果缓冲区中为空，则返回一个零值
  - 无缓冲的channel是同步的，而有缓冲的channel是非同步的

https://www.cnblogs.com/paulwhw/p/15585467.html

- - 数组/切片越界

- - 空指针调用。比如访问一个 nil 结构体指针的成员
  - 过早关闭 HTTP 响应体
  - 除以 0
  - 向已经关闭的 channel 发送消息
  - 重复关闭 channel
  - 关闭未初始化的 channel
  - 未初始化 map。注意访问 map 不存在的 key 不会 panic，而是返回 map 类型对应的零值，但是不能直接赋值
  - 跨协程的 panic 处理
  - sync 计数为负数。
  - 类型断言不匹配。`var a interface{} = 1; fmt.Println(a.(string))` 会 panic，建议用 `s,ok := a.(string)`

1. **能实习多久？什么时候能过来实习？**
2. **要理论结合实践**

------

## 胡超

- [📎百度一面.mp3](https://www.yuque.com/attachments/yuque/0/2022/mp3/22219483/1661266876752-25c92e31-eaaf-42dd-b8b5-3fd1e6cbe5ad.mp3)
- https://www.bilibili.com/audio/au3168287?type=1

1. 自我介绍
2. 实习里面碰到的技术难点并且是怎么解决的？

1. 1. es
   2. 为什用es替代mysql，他俩有啥不一样？
   3. 如果你没有搜索的需求，一开始为什么要用数据库呢？
   4. 数据量不大的话也是要走全表扫描的，跟你数据量大不大有什么关系？
   5. mysql里面的模糊搜索跟es里面的模糊搜索有什么不一样的？
   6. 我要模糊搜索一个词，我在es里面没有对其进行分词，那我是不是就搜索不到了？
   7. 你们是怎么判断哪些是要搜的词的，你们能知道世界上所有的词呢？

1. mysql

1. 1. 隔离级别
   2. 可重复读
   3. 版本控制可以解决幻读问题，能不能讲一下是怎么解决的？ 

1. 1. 1. 也就是说每一条数据都要存到历史版本？那具体是怎么来存的？
      2. 如果把每条数据都要存历史版本，如果存到日志里边，那这个日志是不是越来越大，这个历史版本什么时候清掉？

1. redis

1. 1. 怎么做持久化的？

1. 1. 1. rdb
      2. aof

1. 1. 1. 1. 讲讲日志的写入流程？
         2. redi每个操作都要写一下日志，那会对性能有什么影响，然后通过什么方式来避免这种影响？
         3. 宕机为什么会导致数据丢失，丢失的是哪些数据？
         4. 了解操作系统里面的page cache吗？
         5. aof日志写到page cache里边，异步的把内存里的日志数据写入到磁盘/s

1. java

1. 1. aio、bio、nio的区别？https://blog.csdn.net/meism5/article/details/89469101
   2. 阻塞、非阻塞？
   3. 同步、异步？

1. go和java对比

1. 1. go struct和java 类有什么不一样的？
   2. 一个struct想要复用另外一个struct的成员的话怎么做？
   3.  java里面实现接口和go实现接口有啥区别？
   4. go里面struct怎么实现一个接口？写出来

1. 算法

1. 1. [2385. 感染二叉树需要的总时间](https://leetcode.cn/problems/amount-of-time-for-binary-tree-to-be-infected/)