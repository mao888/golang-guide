## 什么是 RPC ？
- RPC (Remote Procedure Call)即**远程过程调用**，是分布式系统常见的一种通信方法。它允许程序调用另一个地址空间（通常是共享网络的另一台机器上）的过程或函数，而不用程序员显式编码这个远程调用的细节。
- 除 RPC 之外，常见的多系统数据交互方案还有分布式消息队列、HTTP 请求调用、数据库和分布式缓存等。
- 其中 RPC 和 HTTP 调用是没有经过中间件的，它们是端到端系统的直接数据交互。

**简单的说**

- RPC就是从一台机器（客户端）上通过参数传递的方式调用另一台机器（服务器）上的一个函数或方法（可以统称为服务）并得到返回的结果。
- RPC会隐藏底层的通讯细节（不需要直接处理Socket通讯或Http通讯）。
- 客户端发起请求，服务器返回响应（类似于Http的工作方式）RPC在使用形式上像调用本地函数（或方法）一样去调用远程的函数（或方法）。

## 底层通信协议是什么协议
常用的RPC通信协议有TCP、HTTP、gRPC（基于HTTP/2）等。
其中，gRPC是目前较为流行的选择，提供了高效的二进制传输和多语言支持。 

## rpc服务这一块的服务注册发现是怎么做的 
服务注册发现通常借助服务注册中心实现，如Etcd、Consul、Zookeeper等。服务启动时会将自身信息（如IP、端口）注册到服务中心，调用方通过服务中心发现服务，并获取调用信息。  
## 为什么我们要用RPC?
RPC 的主要目标是让构建分布式应用更容易，在提供强大的远程调用能力时不损失本地调用的语义简洁性。为实现该目标，RPC 框架需提供一种透明调用机制让使用者不必显式的区分本地调用和远程调用。
## RPC需要解决的三个问题
RPC要达到的目标：远程调用时，要能够像本地调用一样方便，让调用者感知不到远程调用的逻辑。
![image.png](https://cdn.nlark.com/yuque/0/2022/png/22219483/1661878965339-53175ed8-3ccf-4904-ac5a-b4f348076316.png#averageHue=%23cbf2cc&clientId=u1ae3b37b-b084-4&from=paste&id=u8e9df112&originHeight=487&originWidth=1000&originalType=url&ratio=1&rotation=0&showTitle=false&size=122010&status=done&style=none&taskId=u0227fdbf-661c-48c4-96eb-eadcad7c82c&title=)

- **Call ID映射**。我们怎么告诉远程机器我们要**调用哪个函数呢**？在本地调用中，函数体是直接通过函数指针来指定的，我们调用具体函数，编译器就自动帮我们调用它相应的函数指针。但是在远程调用中，是无法调用函数指针的，因为两个进程的地址空间是完全不一样。所以，在RPC中，**所有的函数都必须有自己的一个ID**。这个ID在所有进程中都是唯一确定的。客户端在做远程过程调用时，必须附上这个ID。然后我们还需要在客户端和服务端分别维护一个 {函数 <--> Call ID} 的对应表。两者的表不一定需要完全相同，但相同的函数对应的Call ID必须相同。当客户端需要进行远程调用时，它就查一下这个表，找出相应的Call ID，然后把它传给服务端，服务端也通过查表，来确定客户端需要调用的函数，然后执行相应函数的代码。
- **序列化和反序列化**。客户端怎么把参数值传给远程的函数呢？在本地调用中，我们只需要把参数压到栈里，然后让函数自己去栈里读就行。但是在远程过程调用时，客户端跟服务端是不同的进程，**不能通过内存来传递参数**。甚至有时候客户端和服务端使用的都**不是同一种语言**（比如服务端用C++，客户端用Java或者Python）。这时候就需要客户端把参数先转成一个字节流，传给服务端后，再把字节流转成自己能读取的格式。这个过程叫序列化和反序列化。同理，从服务端返回的值也需要序列化反序列化的过程。
- **网络传输**。远程调用往往是基于网络的，客户端和服务端是通过网络连接的。所有的数据都需要通过网络传输，因此就需要有一个网络传输层。网络传输层需要把Call ID和序列化后的参数字节流传给服务端，然后再把序列化后的调用结果传回客户端。只要能完成这两者的，都可以作为传输层使用。因此，它所使用的协议其实是不限的，能完成传输就行。尽管大部分RPC框架都使用TCP协议，但其实UDP也可以，而gRPC干脆就用了HTTP2。Java的Netty也属于这层的东西。
## 实现高可用RPC框架需要考虑到的问题

- 既然系统采用分布式架构，那一个服务势必会有多个实例，要解决**如何获取实例的问题**。所以需要一个服务注册中心，比如在Dubbo中，就可以使用Zookeeper作为注册中心，在调用时，从Zookeeper获取服务的实例列表，再从中选择一个进行调用；
- 如何选择实例呢？就要考虑负载均衡，例如dubbo提供了4种负载均衡策略；
- 如果每次都去注册中心查询列表，效率很低，那么就要加缓存；
- 客户端总不能每次调用完都等着服务端返回数据，所以就要支持异步调用；
- 服务端的接口修改了，老的接口还有人在用，这就需要版本控制；
- 服务端总不能每次接到请求都马上启动一个线程去处理，于是就需要线程池；
## 理论结构模型

![image.png](https://cdn.nlark.com/yuque/0/2022/png/22219483/1661878967950-98e4e8cf-0d4c-4da7-866c-64c87a244124.png#averageHue=%23f3f1ef&clientId=u1ae3b37b-b084-4&from=paste&id=uf9a30b96&originHeight=411&originWidth=620&originalType=url&ratio=1&rotation=0&showTitle=false&size=35534&status=done&style=none&taskId=ub618257c-6224-4313-95a1-0572e13f8ee&title=)
RPC 服务端通过RpcServer去导出（export）远程接口方法，而客户端通过RpcClient去导入（import）远程接口方法。客户端像调用本地方法一样去调用远程接口方法，**RPC 框架提供接口的代理实现**，实际的调用将委托给代理RpcProxy。代理封装调用信息并将调用转交给RpcInvoker去实际执行。在客户端的RpcInvoker通过连接器RpcConnector去维持与服务端的通道RpcChannel，并使用RpcProtocol执行协议编码（encode）并将编码后的请求消息通过通道发送给服务端。
RPC 服务端接收器RpcAcceptor接收客户端的调用请求，同样使用RpcProtocol执行协议解码（decode）。
解码后的调用信息传递给RpcProcessor去控制处理调用过程，最后再委托调用给RpcInvoker去实际执行并返回调用结果。

## 主流的RPC框架
服务治理型

- **dubbo**：是阿里巴巴公司开源的一个Java高性能优秀的服务框架，使得应用可通过高性能的 RPC 实现服务的输出和输入功能，可以和 Spring框架无缝集成。dubbo 已经与12年年底停止维护升级。
- **dubbox**：是当当团队基于dubbo升级的一个版本。是一个分布式的服务架构，可直接用于生产环境作为SOA服务框架。dubbox资源链接
- **motan**：是新浪微博开源的一个Java框架。它诞生的比较晚，起于2013年，2016年5月开源。Motan 在微博平台中已经广泛应用，每天为数百个服务完成近千亿次的调用。motan资源链接
## RPC模式
RPC采用客户端/服务端的模式，通过request-response消息模式实现
![image.png](https://cdn.nlark.com/yuque/0/2022/png/22219483/1661878909836-5b7df260-2dbb-41a7-a656-63e0efcee9ad.png#averageHue=%23c7f0a5&clientId=u1ae3b37b-b084-4&from=paste&id=u37bc2f79&originHeight=419&originWidth=720&originalType=url&ratio=1&rotation=0&showTitle=false&size=205868&status=done&style=none&taskId=uacaf5c46-5d8c-45c0-ad2d-c27d1065756&title=)

## RPC的三个过程

1. **通讯协议**  比如：你需要找人在国外干活，那么你可以直接飞过去或者打电话或者通过互联网的形式，去找人，这个找人的过程就是通讯协议 
2. **寻址**：既然要找人干活，肯定要知道地址在哪，飞过去需要找到详细地址，打电话需要知道电话号码，互联网需要知道IP是多少 
3. **数据序列化**：就是说，语言需要互通，才能够让别人干活，之间需要一个大家都懂的语言去交流
## 为什么要使用RPC
1：服务化/微服务 
2：分布式系统架构 
3：服务可重用 
4：系统间交互调用
## RPC和其他协议的区别
RMI远程方法调用是RPC的一种具体实现，webservice、restfull都是RPC，只是消息的组织形式、消息协议不同
## RPC使用场景
和MQ做对比： MQ有一个中间节点queue，可以存储消息 
RPC的特性： 同步调用，对于需要等待返回结果的场景，可以使用RPC 
消息MQ的特性： 

1. 异步单向的消息，不需要等待消息处理完成 如果需要同步得到结果的场景，RPC比较适合，
2. 如果希望使用简单，RPC也适合，RPC操作基于接口，操作简单，使用的方式模拟本地方法的调用，异步的方式编程比较复杂



![image.png](https://cdn.nlark.com/yuque/0/2022/png/22219483/1661878909811-66f72243-9d2e-4250-a6c3-25c99a51a72d.png#averageHue=%23d2dbe4&clientId=u1ae3b37b-b084-4&from=paste&id=u3dbc6336&originHeight=241&originWidth=1026&originalType=url&ratio=1&rotation=0&showTitle=false&size=153332&status=done&style=none&taskId=ucf7b54fd-e328-40d0-bea2-50cd4424dd1&title=)

## RPC的流程
![image.png](https://cdn.nlark.com/yuque/0/2022/png/22219483/1661878909760-55bd7dc1-cbb3-47d4-b563-fe532a054fb5.png#averageHue=%23f6f684&clientId=u1ae3b37b-b084-4&from=paste&id=u3742879f&originHeight=153&originWidth=458&originalType=url&ratio=1&rotation=0&showTitle=false&size=56691&status=done&style=none&taskId=uafb0d100-d316-41e7-84b5-2d0688a5a54&title=)

1：客户端处理过程中调用client sub，就像调用本地方法一样，传入参数 
2：client sub将参数编组为消息，然后通过系统调用向服务端发送消息 
3：客户端本地的操作系统将消息从客户端发送到服务端 
4：服务端将接收到的数据包传递给server sub 
5：server sub将接收到的数据解组为参数 
6：server sub再调用服务端的过程，过程执行的结果以反方向的相同步骤响应给客户端
sub(存根) ：分布式计算中的存根是一段代码，它转换在远程过程调用（RPC）期间client和server之间传递的参数
需要处理的问题： 1：client sub、server sub的开发 2：参数的编组和解组 3：消息如何发送 4：过程结果如何表示、异常情况如何处理 5：如何实现安全的访问控制
## RPC核心概念术语
1：client， 客户端 
2：server，服务端 
3：calls，请求 
4：replier，响应 
5：services，一个网络服务由一个或者多个远程程序集构成 
6：programs，一个远程程序实现一个或多个远程过程 
7：procedures，过程、过程的参数、结果在程序协议说明书中定义说明 
8：version，为兼容程序协议变更，一个服务端可能支持多个版本的远程程序
## RPC协议
RPC调用过程中需要将消息进行编组然后发送，接收方需要解组消息为参数，过程处理结果也需要经过编组、解组；消息由哪些部分构成以及消息的表示形式就构成了消息协议。 RPC协议规定请求消息、响应消息的格式，在TCP之上我们可以选用或者自定义消息协议来实现RPC的交互
## RPC框架
封装好了参数编组、消息解组、底层网络通信的RPC程序开发框架，可以直接在此基础上编写，只关注过程代码 java领域中常用的RPC框架有： 传统的webservice框架：apache CXF、apache Axis2 新兴的微服务框架：Dubbo、springcloud、apache Thrift、ICE、GRPC等
## 服务暴露
远程提供者需要以某种形式提供服务调用相关的信息，包括但不限于服务接口定义、数据结构或者中间态的服务定义文件，web service的WSDL文件；服务调用者需要通过一定的途径获取远程服务调用相关的信息
![image.png](https://cdn.nlark.com/yuque/0/2022/png/22219483/1661878910396-395a5ab4-3a1b-4f10-8ab2-60bbccfd9de9.png#averageHue=%23d1d4de&clientId=u1ae3b37b-b084-4&from=paste&id=u612a7368&originHeight=184&originWidth=720&originalType=url&ratio=1&rotation=0&showTitle=false&size=94817&status=done&style=none&taskId=ue57af141-84ac-4ed2-8572-c0f0c951767&title=)
## 远程代理对象
服务调用者使用的服务实际上是远程服务的本地代理，说白了就是通过动态代理实现 java中至少提供了两种动态代码的生成，一种是jdk动态代理，一种是字节码生成： 动态代理比字节码生成使用起来更加方便，但是性能上没有字节码生成好，字节码生成在代码可读性上要差一些
## 通信
RPC框架的通信与具体的协议无关，RPC可基于HTTP或者TCP协议
![image.png](https://cdn.nlark.com/yuque/0/2022/png/22219483/1661878910515-480c8ffa-9154-497b-9111-be41f163bc83.png#averageHue=%23c2c4ce&clientId=u1ae3b37b-b084-4&from=paste&id=u8d08c6b0&originHeight=338&originWidth=720&originalType=url&ratio=1&rotation=0&showTitle=false&size=273158&status=done&style=none&taskId=u57ee5c4e-1b87-4576-84ec-3481ec1cd49&title=)
## 序列化
传输方式和序列化会直接影响RPC的性能
![image.png](https://cdn.nlark.com/yuque/0/2022/png/22219483/1661878910721-6847f3bc-f525-401c-9571-da5c63401c17.png#averageHue=%23cbcdd9&clientId=u1ae3b37b-b084-4&from=paste&id=uc680bc17&originHeight=287&originWidth=720&originalType=url&ratio=1&rotation=0&showTitle=false&size=148943&status=done&style=none&taskId=u69da2d8a-a74f-4406-9e2c-f92316be7a3&title=)
