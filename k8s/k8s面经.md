## etcd及其特点
etcd是CoreOS团队发起的开源项目，是一个管理配置信息和服务发现（service discovery）的项目，它的目标是构建一个高可用的分布式键值（key-value）数据库，基于Go语言实现。
特点：

- 简单：支持REST风格的HTTP+JSON API
- 安全：支持HTTPS方式的访问
- 快速：支持并发1k/s的写操作
- 可靠：支持分布式结构，基于Raft的一致性算法，Raft是一套通过选举主节点来实现分布式系统一致性的算法。
## etcd适应的场景
etcd基于其优秀的特点，可广泛的应用于以下场景：

- 服务发现（Service Discovery）：服务发现主要解决在同一个分布式集群中的进程或服务，要如何才能找到对方并建立连接。本质上来说，服务发现就是想要了解集群中是否有进程在监听UDP或TCP端口，并且通过名字就可以查找和连接。
- 消息发布与订阅：在分布式系统中，最适用的一种组件间通信方式就是消息发布与订阅。即构建一个配置共享中心，数据提供者在这个配置中心发布消息，而消息使用者则订阅他们关心的主题，一旦主题有消息发布，就会实时通知订阅者。通过这种方式可以做到分布式系统配置的集中式管理与动态更新。应用中用到的一些配置信息放到etcd上进行集中管理。
- 负载均衡：在分布式系统中，为了保证服务的高可用以及数据的一致性，通常都会把数据和服务部署多份，以此达到对等服务，即使其中的某一个服务失效了，也不影响使用。etcd本身分布式架构存储的信息访问支持负载均衡。etcd集群化以后，每个etcd的核心节点都可以处理用户的请求。所以，把数据量小但是访问频繁的消息数据直接存储到etcd中也可以实现负载均衡的效果。
- 分布式通知与协调：与消息发布和订阅类似，都用到了etcd中的Watcher机制，通过注册与异步通知机制，实现分布式环境下不同系统之间的通知与协调，从而对数据变更做到实时处理。
- 分布式锁：因为etcd使用Raft算法保持了数据的强一致性，某次操作存储到集群中的值必然是全局一致的，所以很容易实现分布式锁。锁服务有两种使用方式，一是保持独占，二是控制时序。
- 集群监控与Leader竞选：通过etcd来进行监控实现起来非常简单并且实时性强。
## 什么是k8s
k8s是一个全新的基于容器技术的分布式系统支撑平台。是Google开源的容器集群管理系统（谷歌内部：Borg）。在Docker技术的基础上，为容器化的应用提供部署运行、资源调度、服务发现和动态伸缩等一系列完整功能，提高了大规模容器集群管理的便捷性。并且具有完备的集群管理能力，多层次的安全防护和准入机制、多租户应用支撑能力、透明的服务注册和发现机制、內建智能负载均衡器、强大的故障发现和自我修复能力、服务滚动升级和在线扩容能力、可扩展的资源自动调度机制以及多粒度的资源配额管理能力。
## k8s和Docker的关系
Docker提供容器的生命周期管理和Docker镜像构建运行时容器。它的主要优点是将将软件/应用程序运行所需的设置和依赖项打包到一个容器中，从而实现了可移植性等优点。
k8s用于关联和编排在多个主机上运行的容器。
## Minikube、Kubectl、Kubelet分别是什么
Minikube是一种可以在本地轻松运行一个单节点k8s群集的工具。
Kubectl是一个命令行工具，可以使用该工具控制k8s集群管理器，如检查群集资源，创建、删除和更新组件，查看应用程序。
Kubelet是一个代理服务，它在每个节点上运行，并使从服务器与主服务器通信。
## k8s常见的部署方式
常见的k8s部署方式有：

- kubeadm，也是推荐的一种部署方式；
- 二进制；
- minikube，在本地轻松运行一个单节点k8s群集的工具。
## k8s如何实现集群管理
在集群管理方面，k8s将集群中的机器划分为一个Master节点和一群工作节点Node。其中，在Master节点运行着集群管理相关的一组进程kube-apiserver、kube-controller-manager和kube-scheduler，这些进程实现了整个集群的资源管理、Pod调度、弹性伸缩、安全控制、系统监控和纠错等管理能力，并且都是全自动完成的。
## k8s的优势、适应场景及其特点
k8s作为一个完备的分布式系统支撑平台，其主要优势：

- 容器编排
- 轻量级
- 开源
- 弹性伸缩
- 负载均衡

k8s常见场景：

- 快速部署应用
- 快速扩展应用
- 无缝对接新的应用功能
- 节省资源，优化硬件资源的使用

k8s相关特点：

- 可移植：支持公有云、私有云、混合云、多重云（multi-cloud）。
- 可扩展: 模块化,、插件化、可挂载、可组合。
- 自动化: 自动部署、自动重启、自动复制、自动伸缩/扩展。
## k8s的缺点或当前的不足之处
k8s当前存在的缺点（不足）如下：

- 安装过程和配置相对困难复杂。
- 管理服务相对繁琐。
- 运行和编译需要很多时间。
- 它比其他替代品更昂贵。
- 对于简单的应用程序来说，可能不需要涉及k8s即可满足。
## k8s相关基础概念

- Master：k8s集群的管理节点，负责管理集群，提供集群的资源数据访问入口。拥有etcd存储服务（可选），运行Api Server进程，Controller Manager服务进程及Scheduler服务进程。
- Node（worker）：Node（worker）是k8s集群架构中运行Pod的服务节点，是k8s集群操作的单元，用来承载被分配Pod的运行，是Pod运行的宿主机。运行Docker Eninge服务，守护进程kunelet及负载均衡器kube-proxy。
- Pod：运行于Node节点上，若干相关容器的组合。Pod内包含的容器运行在同一宿主机上，使用相同的网络命名空间、IP地址和端口，能够通过localhost进行通信。Pod是k8s进行创建、调度和管理的最小单位，它提供了比容器更高层次的抽象，使得部署和管理更加灵活。一个Pod可以包含一个容器或者多个相关容器。
- Label：k8s中的Label实质是一系列的Key/Value键值对，其中key与value可自定义。Label可以附加到各种资源对象上，如Node、Pod、Service、RC等。一个资源对象可以定义任意数量的Label，同一个Label也可以被添加到任意数量的资源对象上去。k8s通过Label Selector（标签选择器）查询和筛选资源对象。
- Replication Controller：Replication Controller用来管理Pod的副本，保证集群中存在指定数量的Pod副本。集群中副本的数量大于指定数量，则会停止指定数量之外的多余容器数量。反之，则会启动少于指定数量个数的容器，保证数量不变。Replication Controller是实现弹性伸缩、动态扩容和滚动升级的核心。
- Deployment：Deployment在内部使用了RS来实现目的，Deployment相当于RC的一次升级，其最大的特色为可以随时获知当前Pod的部署进度。
- HPA（Horizontal Pod Autoscaler）：Pod的横向自动扩容，也是k8s的一种资源，通过追踪分析RC控制的所有Pod目标的负载变化情况，来确定是否需要针对性的调整Pod副本数量。
- Service：Service定义了Pod的逻辑集合和访问该集合的策略，是真实服务的抽象。Service提供了一个统一的服务访问入口以及服务代理和发现机制，关联多个相同Label的Pod，用户不需要了解后台Pod是如何运行。
- Volume：Volume是Pod中能够被多个容器访问的共享目录，k8s中的Volume是定义在Pod上，可以被一个或多个Pod中的容器挂载到某个目录下。
- Namespace：Namespace用于实现多租户的资源隔离，可将集群内部的资源对象分配到不同的Namespace中，形成逻辑上的不同项目、小组或用户组，便于不同的Namespace在共享使用整个集群的资源的同时还能被分别管理。
## k8s集群相关组件
k8s Master控制组件，调度管理整个系统（集群），包含如下组件：

- k8s API Server：作为k8s系统的入口，其封装了核心对象的增删改查操作，以RESTful API接口方式提供给外部客户和内部组件调用，集群内各个功能模块之间数据交互和通信的中心枢纽。
- k8s Scheduler：为新建立的Pod进行节点（Node）选择（即分配机器），负责集群的资源调度。
- k8s Controller：负责执行各种控制器，目前已经提供了很多控制器来保证k8s的正常运行。
- Replication Controller：管理维护Replication Controller，关联Replication Controller和Pod，保证Replication Controller定义的副本数量与实际运行Pod数量一致。
- Node Controller：管理维护Node，定期检查Node的健康状态，标识出（失效|未失效）的Node节点。
- Namespace Controller：管理维护Namespace，定期清理无效的Namespace，包括Namesapce下的API对象，比如Pod、Service等。
- Service Controller：管理维护Service，提供负载以及服务代理。
- EndPoints Controller：管理维护Endpoints，关联Service和Pod，创建Endpoints为Service的后端，当Pod发生变化时，实时更新Endpoints。
- Service Account Controller：管理维护Service Account，为每个Namespace创建默认的Service Account，同时为Service Account创建Service Account Secret。
- Persistent Volume Controller：管理维护Persistent Volume和Persistent Volume Claim，为新的Persistent Volume Claim分配Persistent Volume进行绑定，为释放的Persistent Volume执行清理回收。
- Daemon Set Controller：管理维护Daemon Set，负责创建Daemon Pod，保证指定的Node上正常的运行Daemon Pod。
- Deployment Controller：管理维护Deployment，关联Deployment和Replication Controller，保证运行指定数量的Pod。当Deployment更新时，控制实现Replication Controller和Pod的更新。
- Job Controller：管理维护Job，为Jod创建一次性任务Pod，保证完成Job指定完成的任务数目
- Pod Autoscaler Controller：实现Pod的自动伸缩，定时获取监控数据，进行策略匹配，当满足条件时执行Pod的伸缩动作。
## k8s的网络原理是什么样子
Kubernetes（k8s）的网络原理主要基于Pod之间的通信和集群内外的通信。k8s通过CNI（容器网络接口）插件实现Pod间的网络隔离与通信，每个Pod都会被分配一个唯一的IP地址，Pod间通过虚拟网络进行通信。集群内外的通信则通过Service资源实现，Service通过ClusterIP、NodePort或LoadBalancer等方式暴露服务，使得外部可以访问集群内的服务。

在面试中讨论Kubernetes（K8s）的网络原理时，可以从以下几个方面进行阐述：
### K8s网络概述
Kubernetes的网络模型设计目标是提供一个简单、高效、可扩展的网络解决方案，以确保集群内的Pod之间以及Pod与外部网络之间的通信顺畅。K8s网络的核心在于为集群内的每个Pod分配一个唯一的IP地址，并且假设这些Pod位于一个可以直接连通的扁平网络空间中。
### K8s网络组件与原理
#### Pod网络

- **Pod IP**：每个Pod在K8s集群中被分配一个唯一的IP地址（Pod IP），这个地址由K8s的网络插件（如Calico、Flannel等）负责配置和管理。Pod内的所有容器共享同一个网络命名空间，因此它们可以直接通过localhost进行通信。
- **Veth设备对**：为了实现不同Pod之间的通信，K8s使用Veth设备对（虚拟网络接口对）将Pod连接到宿主机上的虚拟网桥（如docker0或cni0）。这样，不同Pod之间的通信就可以通过宿主机的网络进行路由和转发。
#### 2. Service网络

- **Cluster IP**：Service是K8s中用于抽象一组Pod的资源对象，它为这组Pod提供了一个统一的访问入口（Cluster IP）。Cluster IP是一个虚拟IP地址，它不会配置在任何主机或容器的网络接口上，而是通过Node上的kube-proxy组件配置为iptables或ipvs规则，将发往该地址的流量调度到后端的Pod对象上。
- **NodePort和LoadBalancer**：除了Cluster IP外，Service还可以通过NodePort和LoadBalancer类型暴露给外部网络。NodePort类型会在集群的每个节点上暴露一个端口，外部客户端可以通过`<Node IP>:<NodePort>`访问Service。LoadBalancer类型则进一步封装了NodePort，在支持负载均衡器的环境中，它会自动创建一个外部负载均衡器，并将流量分发到集群内的节点上。
#### 3. 网络插件

- **CNI（Container Network Interface）**：CNI是K8s的网络接口规范，它定义了容器如何与网络交互。CNI插件（如Calico、Flannel等）实现了这个接口，为Pod提供网络连接。
- **Calico**：Calico是一个基于BGP协议的网络解决方案，它提供了简单的网络配置和强大的网络策略功能。Calico通过BGP协议在集群内部交换路由信息，确保网络路由的透明性和高效性。
- **Flannel**：Flannel是另一个流行的网络插件，它使用Overlay网络为Pod提供跨主机的通信能力。Flannel支持多种后端实现（如UDP、VXLAN、Host-gw等），其中最常用的是VXLAN模式，它通过虚拟隧道技术在不同节点之间传输数据包。
### K8s网络通信流程
以一次典型的Service访问为例，网络通信流程大致如下：

1. 客户端发起对Service的访问请求（如通过Cluster IP或NodePort）。
2. 请求到达Node上的kube-proxy组件。
3. kube-proxy根据iptables或ipvs规则将请求转发到后端的某个Pod上。
4. Pod收到请求后进行处理，并将响应返回给客户端。

在这个过程中，K8s的网络插件（如Calico、Flannel）负责配置和管理Pod网络，确保不同Pod之间以及Pod与外部网络之间的通信顺畅。同时，kube-proxy组件作为Service的透明代理和负载均衡器，负责将请求转发到正确的Pod上。
### 总结
Kubernetes的网络原理涉及多个组件和机制，包括Pod网络、Service网络、网络插件以及网络通信流程等。了解这些原理和组件对于深入理解K8s的网络架构至关重要。在面试中，可以结合具体场景和问题深入阐述这些原理和组件的作用和实现方式。
## service之间是如何调用的
Service之间的调用通常通过Service的ClusterIP进行。在k8s集群中，每个Service都会被分配一个ClusterIP，Pod通过访问这个ClusterIP来调用Service。Service会根据其背后的Pod列表进行负载均衡，将请求分发到具体的Pod上。
在Kubernetes（K8s）环境中，Service之间的调用是微服务架构中的关键组成部分。Service之间的调用主要通过Service对象和Kubernetes集群内置的DNS服务来实现。以下是详细的调用过程：
### Service对象

- **定义与创建**：Service是Kubernetes中的一种抽象，它定义了一组Pod的访问策略，并为这些Pod提供了一个稳定的访问入口。通过创建Service对象，可以为Pod集合分配一个虚拟IP地址（Cluster IP）和一个或多个端口。
- **选择器（Selector）**：Service通过标签选择器（label selector）来关联一组Pod，只有标签与选择器匹配的Pod才会被Service代理。
### DNS服务

- **DNS解析**：Kubernetes集群内置了DNS服务，允许通过Service的名称进行DNS解析，获取到Service的Cluster IP地址。这使得Service之间的调用可以通过名称而非IP地址进行，提高了系统的可维护性和可移植性。
- **命名空间**：在同一命名空间中，Service可以直接通过名称调用其他Service，无需指定完整的DNS名称（如省略`.svc.cluster.local`后缀）。跨命名空间的调用则需要使用完整的DNS名称（即`<service-name>.<namespace>.svc.cluster.local`）。
### 调用流程

- **服务发现**：当需要调用某个Service时，客户端（如另一个Pod中的应用程序）首先通过DNS解析获取到目标Service的Cluster IP地址。
- **负载均衡**：Kubernetes的Service会将进入的流量根据负载均衡策略分发到后端的一组Pod上。这确保了服务的高可用性和可扩展性。
- **通信**：客户端与目标Pod之间通过Cluster IP和端口进行通信。在Kubernetes集群内部，这种通信是通过Overlay网络实现的，它确保了不同节点上的Pod能够相互通信。
### 特殊类型的Service

- **Headless Service**：Headless Service是一种特殊类型的Service，它不会分配Cluster IP地址，而是直接返回所有后端Pod的IP地址列表。这使得客户端可以直接与Pod进行通信，适用于需要直接访问Pod IP的场景。
### 示例
假设有两个Service，分别名为`frontend`和`backend`，且它们位于同一命名空间中。在`frontend` Service中的应用程序可以通过DNS解析`backend` Service的名称来获取其Cluster IP地址，并直接通过HTTP或gRPC等方式调用`backend` Service提供的API。
### 总结
Kubernetes中的Service之间的调用主要依赖于Service对象和集群内置的DNS服务。通过创建Service对象并配置DNS解析，可以实现微服务之间的灵活调用和通信。这种机制提高了系统的可维护性、可移植性和可扩展性。

## 怎么实现负载均衡的呢？
k8s通过Service实现负载均衡。Service会根据其背后的Pod列表和负载均衡算法（如轮询、最少连接等）将请求分发到具体的Pod上。此外，k8s还支持通过Ingress资源实现更复杂的路由和负载均衡策略。

1. Kubernetes通过Service对象和kube-proxy组件来实现负载均衡。
2. Service对象定义了一组Pod的访问策略，并为它们提供了一个稳定的访问入口，包括一个虚拟IP地址（Cluster IP）和一个或多个端口。
3. kube-proxy是集群中每个节点上运行的网络代理，它负责监控Service和Endpoint对象的变动，并根据这些变动来动态更新节点上的iptables或IPVS规则。

当外部请求到达节点的Service端口时，iptables或IPVS规则会将请求重定向到后端Pod的IP和端口上。
在iptables模式下，默认的负载均衡算法是随机或轮询，但也可以通过自定义iptables规则来设置其他算法。
而在IPVS模式下，则支持更多种类的负载均衡算法，如最小连接数、源地址哈希等。
通过这种方式，Kubernetes实现了对Pod的灵活访问和负载均衡，确保了服务的高可用性和可扩展性。这也是Kubernetes作为容器编排平台的一大优势。”
## k8s RC的机制
Replication Controller用来管理Pod的副本，保证集群中存在指定数量的Pod副本。当定义了RC并提交至k8s集群中之后，Master节点上的Controller Manager组件获悉，并同时巡检系统中当前存活的目标Pod，并确保目标Pod实例的数量刚好等于此RC的期望值，若存在过多的Pod副本在运行，系统会停止一些Pod，反之则自动创建一些Pod。
## k8s Replica Set和Replication Controller之间有什么区别
Replica Set和Replication Controller类似，都是确保在任何给定时间运行指定数量的Pod副本。不同之处在于RS使用基于集合的选择器，而Replication Controller使用基于权限的选择器。
## kube-proxy的作用
kube-proxy运行在所有节点上，它监听apiserver中service和endpoint的变化情况，创建路由规则以提供服务IP和负载均衡功能。简单理解此进程是Service的透明代理兼负载均衡器，其核心功能是将到某个Service的访问请求转发到后端的多个Pod实例上。
## kube-proxy iptables的原理
k8s从1.2版本开始，将iptables作为kube-proxy的默认模式。iptables模式下的kube-proxy不再起到Proxy的作用，其核心功能：通过API Server的Watch接口实时跟踪Service与Endpoint的变更信息，并更新对应的iptables规则，Client的请求流量则通过iptables的NAT机制“直接路由”到目标Pod。
## kube-proxy ipvs的原理
IPVS在k8s1.11中升级为GA稳定版。IPVS则专门用于高性能负载均衡，并使用更高效的数据结构（Hash表），允许几乎无限的规模扩张，因此被kube-proxy采纳为最新模式。
在IPVS模式下，使用iptables的扩展ipset，而不是直接调用iptables来生成规则链。iptables规则链是一个线性的数据结构，ipset则引入了带索引的数据结构，因此当规则很多时，也可以很高效地查找和匹配。
可以将ipset简单理解为一个IP（段）的集合，这个集合的内容可以是IP地址、IP网段、端口等，iptables可以直接添加规则对这个“可变的集合”进行操作，这样做的好处在于可以大大减少iptables规则的数量，从而减少性能损耗。
## kube-proxy ipvs和iptables的异同
iptables与IPVS都是基于Netfilter实现的，但因为定位不同，二者有着本质的差别：iptables是为防火墙而设计的；IPVS则专门用于高性能负载均衡，并使用更高效的数据结构（Hash表），允许几乎无限的规模扩张。
与iptables相比，IPVS拥有以下明显优势：

- 为大型集群提供了更好的可扩展性和性能；
- 支持比iptables更复杂的复制均衡算法（最小负载、最少连接、加权等）；
- 支持服务器健康检查和连接重试等功能；
- 可以动态修改ipset的集合，即使iptables的规则正在使用这个集合。
## k8s中什么是静态Pod
静态Pod是由kubelet进行管理的仅存在于特定Node的Pod上，他们不能通过API Server进行管理，无法与ReplicationController、Deployment或者DaemonSet进行关联，并且kubelet无法对他们进行健康检查。静态Pod总是由kubelet进行创建，并且总是在kubelet所在的Node上运行。
## k8s中Pod可能位于的状态

- Pending：API Server已经创建该Pod，且Pod内还有一个或多个容器的镜像没有创建，包括正在下载镜像的过程。
- Running：Pod内所有容器均已创建，且至少有一个容器处于运行状态、正在启动状态或正在重启状态。
- Succeeded：Pod内所有容器均成功执行退出，且不会重启。
- Failed：Pod内所有容器均已退出，但至少有一个容器退出为失败状态。
- Unknown：由于某种原因无法获取该Pod状态，可能由于网络通信不畅导致。
## k8s创建一个Pod的主要流程？
k8s中创建一个Pod涉及多个组件之间联动，主要流程如下：

- 客户端提交Pod的配置信息（可以是yaml文件定义的信息）到kube-apiserver。
- Apiserver收到指令后，通知给controller-manager创建一个资源对象。
- Controller-manager通过api-server将Pod的配置信息存储到etcd数据中心中。
- Kube-scheduler检测到Pod信息会开始调度预选，会先过滤掉不符合Pod资源配置要求的节点，然后开始调度调优，主要是挑选出更适合运行Pod的节点，然后将Pod的资源配置单发送到Node节点上的kubelet组件上。
- Kubelet根据scheduler发来的资源配置单运行Pod，运行成功后，将Pod的运行信息返回给scheduler，scheduler将返回的Pod运行状况的信息存储到etcd数据中心。
## k8s中Pod的重启策略
Pod重启策略（RestartPolicy）应用于Pod内的所有容器，并且仅在Pod所处的Node上由kubelet进行判断和重启操作。当某个容器异常退出或者健康检查失败时，kubelet将根据RestartPolicy的设置来进行相应操作。
Pod的重启策略包括Always、OnFailure和Never，默认值为Always。

- Always：当容器失效时，由kubelet自动重启该容器；
- OnFailure：当容器终止运行且退出码不为0时，由kubelet自动重启该容器；
- Never：不论容器运行状态如何，kubelet都不会重启该容器。

同时Pod的重启策略与控制方式关联，当前可用于管理Pod的控制器包括ReplicationController、Job、DaemonSet及直接管理kubelet管理（静态Pod）。
不同控制器的重启策略限制如下：

- RC和DaemonSet：必须设置为Always，需要保证该容器持续运行；
- Job：OnFailure或Never，确保容器执行完成后不再重启；
- kubelet：在Pod失效时重启，不论将RestartPolicy设置为何值，也不会对Pod进行健康检查。
## k8s中Pod的健康检查方式
对Pod的健康检查可以通过两类探针来检查：LivenessProbe和ReadinessProbe。

- LivenessProbe探针：用于判断容器是否存活（running状态），如果LivenessProbe探针探测到容器不健康，则kubelet将杀掉该容器，并根据容器的重启策略做相应处理。若一个容器不包含LivenessProbe探针，kubelet认为该容器的LivenessProbe探针返回值用于是“Success”。
- ReadineeProbe探针：用于判断容器是否启动完成（ready状态）。如果ReadinessProbe探针探测到失败，则Pod的状态将被修改。Endpoint Controller将从Service的Endpoint中删除包含该容器所在Pod的Eenpoint。
- startupProbe探针：启动检查机制，应用一些启动缓慢的业务，避免业务长时间启动而被上面两类探针kill掉。
## k8s Pod的LivenessProbe探针的常见方式
kubelet定期执行LivenessProbe探针来诊断容器的健康状态，通常有以下三种方式：

- ExecAction：在容器内执行一个命令，若返回码为0，则表明容器健康。
- TCPSocketAction：通过容器的IP地址和端口号执行TCP检查，若能建立TCP连接，则表明容器健康。
- HTTPGetAction：通过容器的IP地址、端口号及路径调用HTTP Get方法，若响应的状态码大于等于200且小于400，则表明容器健康。
## k8s Pod的常见调度方式
k8s中，Pod通常是容器的载体，主要有如下常见调度方式：

- Deployment或RC：该调度策略主要功能就是自动部署一个容器应用的多份副本，以及持续监控副本的数量，在集群内始终维持用户指定的副本数量。
- NodeSelector：定向调度，当需要手动指定将Pod调度到特定Node上，可以通过Node的标签（Label）和Pod的nodeSelector属性相匹配。
- NodeAffinity亲和性调度：亲和性调度机制极大的扩展了Pod的调度能力，目前有两种节点亲和力表达：
   - requiredDuringSchedulingIgnoredDuringExecution：硬规则，必须满足指定的规则，调度器才可以调度Pod至Node上（类似nodeSelector，语法不同）。
   - preferredDuringSchedulingIgnoredDuringExecution：软规则，优先调度至满足的Node的节点，但不强求，多个优先级规则还可以设置权重值。
- Taints和Tolerations（污点和容忍）：
   - Taint：使Node拒绝特定Pod运行；
   - Toleration：为Pod的属性，表示Pod能容忍（运行）标注了Taint的Node。
## k8s初始化容器（init container）
init container的运行方式与应用容器不同，它们必须先于应用容器执行完成，当设置了多个init container时，将按顺序逐个运行，并且只有前一个init container运行成功后才能运行后一个init container。当所有init container都成功运行后，k8s才会初始化Pod的各种信息，并开始创建和运行应用容器。
## k8s deployment升级过程

- 初始创建Deployment时，系统创建了一个ReplicaSet，并按用户的需求创建了对应数量的Pod副本。
- 当更新Deployment时，系统创建了一个新的ReplicaSet，并将其副本数量扩展到1，然后将旧ReplicaSet缩减为2。
- 之后，系统继续按照相同的更新策略对新旧两个ReplicaSet进行逐个调整。
- 最后，新的ReplicaSet运行了对应个新版本Pod副本，旧的ReplicaSet副本数量则缩减为0。
## k8s deployment升级策略
在Deployment的定义中，可以通过spec.strategy指定Pod更新的策略，目前支持两种策略：Recreate（重建）和RollingUpdate（滚动更新），默认值为RollingUpdate。

- Recreate：设置spec.strategy.type=Recreate，表示Deployment在更新Pod时，会先杀掉所有正在运行的Pod，然后创建新的Pod。
- RollingUpdate：设置spec.strategy.type=RollingUpdate，表示Deployment会以滚动更新的方式来逐个更新Pod。同时，可以通过设置spec.strategy.rollingUpdate下的两个参数（maxUnavailable和maxSurge）来控制滚动更新的过程。
## k8s DaemonSet类型的资源特性
DaemonSet资源对象会在每个k8s集群中的节点上运行，并且每个节点只能运行一个Pod，这是它和Deployment资源对象的最大也是唯一的区别。因此，在定义yaml文件中，不支持定义replicas。
它的一般使用场景如下：

- 在去做每个节点的日志收集工作。
- 监控每个节点的的运行状态。
## k8s自动扩容机制
k8s使用Horizontal Pod Autoscaler（HPA）的控制器实现基于CPU使用率进行自动Pod扩缩容的功能。HPA控制器周期性地监测目标Pod的资源性能指标，并与HPA资源对象中的扩缩容条件进行对比，在满足条件时对Pod副本数量进行调整。
k8s中的某个Metrics Server（Heapster或自定义Metrics Server）持续采集所有Pod副本的指标数据。HPA控制器通过Metrics Server的API（Heapster的API或聚合API）获取这些数据，基于用户定义的扩缩容规则进行计算，得到目标Pod副本数量。
当目标Pod副本数量与当前副本数量不同时，HPA控制器就向Pod的副本控制器（Deployment、RC或ReplicaSet）发起scale操作，调整Pod的副本数量，完成扩缩容操作。
## k8s Service类型
通过创建Service，可以为一组具有相同功能的容器应用提供一个统一的入口地址，并且将请求负载分发到后端的各个容器应用上。其主要类型有：

- ClusterIP：虚拟的服务IP地址，该地址用于k8s集群内部的Pod访问，在Node上kube-proxy通过设置的iptables规则进行转发；
- NodePort：使用宿主机的端口，使能够访问各Node的外部客户端通过Node的IP地址和端口号就能访问服务；
- LoadBalancer：使用外接负载均衡器完成到服务的负载分发，需要在spec.status.loadBalancer字段指定外部负载均衡器的IP地址，通常用于公有云。、
## k8s Service分发后端的策略
Service负载分发的策略有：RoundRobin和SessionAffinity

- RoundRobin：默认为轮询模式，即轮询将请求转发到后端的各个Pod上。
- SessionAffinity：基于客户端IP地址进行会话保持的模式，即第1次将某个客户端发起的请求转发到后端的某个Pod上，之后从相同的客户端发起的请求都将被转发到后端相同的Pod上。
## k8s Headless Service
在某些应用场景中，若需要人为指定负载均衡器，不使用Service提供的默认负载均衡的功能，或者应用程序希望知道属于同组服务的其他实例。k8s提供了Headless Service来实现这种功能，即不为Service设置ClusterIP（入口IP地址），仅通过Label Selector将后端的Pod列表返回给调用的客户端。
## k8s外部如何访问集群内的服务
对于k8s，集群外的客户端默认情况，无法通过Pod的IP地址或者Service的虚拟IP地址：虚拟端口号进行访问。通常可以通过以下方式进行访问k8s集群内的服务：

- 映射Pod到物理机：将Pod端口号映射到宿主机，即在Pod中采用hostPort方式，以使客户端应用能够通过物理机访问容器应用。
- 映射Service到物理机：将Service端口号映射到宿主机，即在Service中采用nodePort方式，以使客户端应用能够通过物理机访问容器应用。
- 映射Sercie到LoadBalancer：通过设置LoadBalancer映射到云服务商提供的LoadBalancer地址。这种用法仅用于在公有云服务提供商的云平台上设置Service的场景。
## k8s ingress
k8s的Ingress资源对象，用于将不同URL的访问请求转发到后端不同的Service，以实现HTTP层的业务路由机制。
k8s使用了Ingress策略和Ingress Controller，两者结合并实现了一个完整的Ingress负载均衡器。使用Ingress进行负载分发时，Ingress Controller基于Ingress规则将客户端请求直接转发到Service对应的后端Endpoint（Pod）上，从而跳过kube-proxy的转发功能，kube-proxy不再起作用，全过程为：ingress controller + ingress 规则 ----> services。
同时当Ingress Controller提供的是对外服务，则实际上实现的是边缘路由器的功能。
## k8s镜像的下载策略
k8s的镜像下载策略有三种：Always、Never、IFNotPresent。

- Always：镜像标签为latest时，总是从指定的仓库中获取镜像。
- Never：禁止从仓库中下载镜像，也就是说只能使用本地镜像。
- IfNotPresent：仅当本地没有对应镜像时，才从目标仓库中下载。默认的镜像下载策略是：当镜像标签是latest时，默认策略是Always；当镜像标签是自定义时（也就是标签不是latest），那么默认策略是IfNotPresent。
## k8s的负载均衡器
负载均衡器是暴露服务的最常见和标准方式之一。
根据工作环境使用两种类型的负载均衡器，即内部负载均衡器或外部负载均衡器。内部负载均衡器自动平衡负载并使用所需配置分配容器，而外部负载均衡器将流量从外部负载引导至后端容器。
## k8s各模块如何与API Server通信
k8s API Server作为集群的核心，负责集群各功能模块之间的通信。集群内的各个功能模块通过API Server将信息存入etcd，当需要获取和操作这些数据时，则通过API Server提供的REST接口（用GET、LIST或WATCH方法）来实现，从而实现各模块之间的信息交互。
如kubelet进程与API Server的交互：每个Node上的kubelet每隔一个时间周期，就会调用一次API Server的REST接口报告自身状态，API Server在接收到这些信息后，会将节点状态信息更新到etcd中。
如kube-controller-manager进程与API Server的交互：kube-controller-manager中的Node Controller模块通过API Server提供的Watch接口实时监控Node的信息，并做相应处理。
如kube-scheduler进程与API Server的交互：Scheduler通过API Server的Watch接口监听到新建Pod副本的信息后，会检索所有符合该Pod要求的Node列表，开始执行Pod调度逻辑，在调度成功后将Pod绑定到目标节点上。
## k8s Scheduler作用及实现原理
k8s Scheduler是负责Pod调度的重要功能模块，k8s Scheduler在整个系统中承担了“承上启下”的重要功能，“承上”是指它负责接收Controller Manager创建的新Pod，为其调度至目标Node；“启下”是指调度完成后，目标Node上的kubelet服务进程接管后继工作，负责Pod接下来生命周期。
k8s Scheduler的作用是将待调度的Pod（API新创建的Pod、Controller Manager为补足副本而创建的Pod等）按照特定的调度算法和调度策略绑定（Binding）到集群中某个合适的Node上，并将绑定信息写入etcd中。
在整个调度过程中涉及三个对象，分别是待调度Pod列表、可用Node列表，以及调度算法和策略。
k8s Scheduler通过调度算法调度为待调度Pod列表中的每个Pod从Node列表中选择一个最适合的Node来实现Pod的调度。随后，目标节点上的kubelet通过API Server监听到k8s Scheduler产生的Pod绑定事件，然后获取对应的Pod清单，下载Image镜像并启动容器。
## k8s Scheduler使用哪两种算法将Pod绑定到worker节点
k8s Scheduler根据如下两种调度算法将 Pod 绑定到最合适的工作节点：

- 预选（Predicates）：输入是所有节点，输出是满足预选条件的节点。kube-scheduler根据预选策略过滤掉不满足策略的Nodes。如果某节点的资源不足或者不满足预选策略的条件则无法通过预选。如“Node的label必须与Pod的Selector一致”。
- 优选（Priorities）：输入是预选阶段筛选出的节点，优选会根据优先策略为通过预选的Nodes进行打分排名，选择得分最高的Node。例如，资源越富裕、负载越小的Node可能具有越高的排名。
## k8s kubelet的作用
在k8s集群中，在每个Node（又称Worker）上都会启动一个kubelet服务进程。该进程用于处理Master下发到本节点的任务，管理Pod及Pod中的容器。每个kubelet进程都会在API Server上注册节点自身的信息，定期向Master汇报节点资源的使用情况，并通过cAdvisor监控容器和节点资源。
## k8s kubelet监控Worker节点资源是使用什么组件来实现的
kubelet使用cAdvisor对worker节点资源进行监控。在k8s系统中，cAdvisor已被默认集成到kubelet组件内，当kubelet服务启动时，它会自动启动cAdvisor服务，然后cAdvisor会实时采集所在节点的性能指标及在节点上运行的容器的性能指标。
## k8s如何保证集群的安全性
k8s通过一系列机制来实现集群的安全控制，主要有如下不同的维度：

- 基础设施方面：保证容器与其所在宿主机的隔离；
- 权限方面：
   - 最小权限原则：合理限制所有组件的权限，确保组件只执行它被授权的行为，通过限制单个组件的能力来限制它的权限范围。
   - 用户权限：划分普通用户和管理员的角色。
- 集群方面：
   - API Server的认证授权：k8s集群中所有资源的访问和变更都是通过k8s API Server来实现的，因此需要建议采用更安全的HTTPS或Token来识别和认证客户端身份（Authentication），以及随后访问权限的授权（Authorization）环节。
   - API Server的授权管理：通过授权策略来决定一个API调用是否合法。对合法用户进行授权并且随后在用户访问时进行鉴权，建议采用更安全的RBAC方式来提升集群安全授权。
- 敏感数据引入Secret机制：对于集群敏感数据建议使用Secret方式进行保护。
- AdmissionControl（准入机制）：对k8s api的请求过程中，顺序为：先经过认证 & 授权，然后执行准入操作，最后对目标对象进行操作。
## k8s准入机制
在对集群进行请求时，每个准入控制代码都按照一定顺序执行。如果有一个准入控制拒绝了此次请求，那么整个请求的结果将会立即返回，并提示用户相应的error信息。
准入控制（AdmissionControl）准入控制本质上为一段准入代码，在对k8s api的请求过程中，顺序为：先经过认证 & 授权，然后执行准入操作，最后对目标对象进行操作。常用组件（控制代码）如下：

- AlwaysAdmit：允许所有请求
- AlwaysDeny：禁止所有请求，多用于测试环境。
- ServiceAccount：它将serviceAccounts实现了自动化，它会辅助serviceAccount做一些事情，比如如果pod没有serviceAccount属性，它会自动添加一个default，并确保pod的serviceAccount始终存在。
- LimitRanger：观察所有的请求，确保没有违反已经定义好的约束条件，这些条件定义在namespace中LimitRange对象中。
- NamespaceExists：观察所有的请求，如果请求尝试创建一个不存在的namespace，则这个请求被拒绝。
## k8s RBAC及其特点（优势）
RBAC是基于角色的访问控制，是一种基于个人用户的角色来管理对计算机或网络资源的访问的方法。
相对于其他授权模式，RBAC具有如下优势：

- 对集群中的资源和非资源权限均有完整的覆盖。
- 整个RBAC完全由几个API对象完成， 同其他API对象一样， 可以用kubectl或API进行操作。
- 可以在运行时进行调整，无须重新启动API Server。
## k8s Secret作用
Secret对象，主要作用是保管私密数据，比如密码、OAuth Tokens、SSH Keys等信息。将这些私密信息放在Secret对象中比直接放在Pod或Docker Image中更安全，也更便于使用和分发。
## k8s Secret有哪些使用方式
创建完secret之后，可通过如下三种方式使用：

- 在创建Pod时，通过为Pod指定Service Account来自动使用该Secret。
- 通过挂载该Secret到Pod来使用它。
- 在Docker镜像下载时使用，通过指定Pod的spc.ImagePullSecrets来引用它。
## k8s PodSecurityPolicy机制
k8s PodSecurityPolicy是为了更精细地控制Pod对资源的使用方式以及提升安全策略。在开启PodSecurityPolicy准入控制器后，k8s默认不允许创建任何Pod，需要创建PodSecurityPolicy策略和相应的RBAC授权策略（Authorizing Policies），Pod才能创建成功。
## k8s PodSecurityPolicy机制能实现哪些安全策略
在PodSecurityPolicy对象中可以设置不同字段来控制Pod运行时的各种安全策略，常见的有：

- 特权模式：privileged是否允许Pod以特权模式运行。
- 宿主机资源：控制Pod对宿主机资源的控制，如hostPID：是否允许Pod共享宿主机的进程空间。
- 用户和组：设置运行容器的用户ID（范围）或组（范围）。
- 提升权限：AllowPrivilegeEscalation：设置容器内的子进程是否可以提升权限，通常在设置非root用户（MustRunAsNonRoot）时进行设置。
- SELinux：进行SELinux的相关配置。
## k8s网络模型
k8s网络模型中每个Pod都拥有一个独立的IP地址，并假定所有Pod都在一个可以直接连通的、扁平的网络空间中。所以不管它们是否运行在同一个Node（宿主机）中，都要求它们可以直接通过对方的IP进行访问。设计这个原则的原因是，用户不需要额外考虑如何建立Pod之间的连接，也不需要考虑如何将容器端口映射到主机端口等问题。
同时为每个Pod都设置一个IP地址的模型使得同一个Pod内的不同容器会共享同一个网络命名空间，也就是同一个Linux网络协议栈。这就意味着同一个Pod内的容器可以通过localhost来连接对方的端口。
在k8s的集群里，IP是以Pod为单位进行分配的。一个Pod内部的所有容器共享一个网络堆栈（相当于一个网络命名空间，它们的IP地址、网络设备、配置等都是共享的）。
## k8s CNI模型
CNI提供了一种应用容器的插件化网络解决方案，定义对容器网络进行操作和配置的规范，通过插件的形式对CNI接口进行实现。CNI仅关注在创建容器时分配网络资源，和在销毁容器时删除网络资源。在CNI模型中只涉及两个概念：容器和网络。

- 容器（Container）：是拥有独立Linux网络命名空间的环境，例如使用Docker或rkt创建的容器。容器需要拥有自己的Linux网络命名空间，这是加入网络的必要条件。
- 网络（Network）：表示可以互连的一组实体，这些实体拥有各自独立、唯一的IP地址，可以是容器、物理机或者其他网络设备（比如路由器）等。

对容器网络的设置和操作都通过插件（Plugin）进行具体实现，CNI插件包括两种类型：CNI Plugin和IPAM（IP Address Management）Plugin。CNI Plugin负责为容器配置网络资源，IPAM Plugin负责对容器的IP地址进行分配和管理。IPAM Plugin作为CNI Plugin的一部分，与CNI Plugin协同工作。
## k8s网络策略
为实现细粒度的容器间网络访问隔离策略，k8s引入Network Policy。
Network Policy的主要功能是对Pod间的网络通信进行限制和准入控制，设置允许访问或禁止访问的客户端Pod列表。Network Policy定义网络策略，配合策略控制器（Policy Controller）进行策略的实现。
## k8s网络策略原理
Network Policy的工作原理主要为：policy controller需要实现一个API Listener，监听用户设置的Network Policy定义，并将网络访问规则通过各Node的Agent进行实际设置（Agent则需要通过CNI网络插件实现）。
## k8s中flannel的作用
Flannel可以用于k8s底层网络的实现，主要作用有：

- 它能协助k8s，给每一个Node上的Docker容器都分配互相不冲突的IP地址。
- 它能在这些IP地址之间建立一个覆盖网络（Overlay Network），通过这个覆盖网络，将数据包原封不动地传递到目标容器内。
## k8s Calico网络组件实现原理
Calico是一个基于BGP的纯三层的网络方案，与OpenStack、k8s、AWS、GCE等云平台都能够良好地集成。
Calico在每个计算节点都利用Linux Kernel实现了一个高效的vRouter来负责数据转发。每个vRouter都通过BGP协议把在本节点上运行的容器的路由信息向整个Calico网络广播，并自动设置到达其他节点的路由转发规则。
Calico保证所有容器之间的数据流量都是通过IP路由的方式完成互联互通的。Calico节点组网时可以直接利用数据中心的网络结构（L2或者L3），不需要额外的NAT、隧道或者Overlay Network，没有额外的封包解包，能够节约CPU运算，提高网络效率。
## k8s共享存储的作用
k8s对于有状态的容器应用或者对数据需要持久化的应用，因此需要更加可靠的存储来保存应用产生的重要数据，以便容器应用在重建之后仍然可以使用之前的数据。因此需要使用共享存储。
## k8s数据持久化的方式有哪些
k8s通过数据持久化来持久化保存重要数据，常见的方式有：

- EmptyDir（空目录）：没有指定要挂载宿主机上的某个目录，直接由Pod内保部映射到宿主机上。类似于docker中的manager volume。
- 场景：
   - 只需要临时将数据保存在磁盘上，比如在合并/排序算法中；
   - 作为两个容器的共享存储。
- 特性：
   - 同个pod里面的不同容器，共享同一个持久化目录，当pod节点删除时，volume的数据也会被删除。
   - emptyDir的数据持久化的生命周期和使用的pod一致，一般是作为临时存储使用。
- Hostpath：将宿主机上已存在的目录或文件挂载到容器内部。类似于docker中的bind mount挂载方式。
- 特性：增加了Pod与节点之间的耦合。

PersistentVolume（简称PV）：如基于NFS服务的PV，也可以基于GFS的PV。它的作用是统一数据持久化目录，方便管理。
## k8s PV和PVC
PV是对底层网络共享存储的抽象，将共享存储定义为一种“资源”。
PVC则是用户对存储资源的一个“申请”。
## k8s PV生命周期内的阶段
某个PV在生命周期中可能处于以下4个阶段（Phaes）之一。

- Available：可用状态，还未与某个PVC绑定。
- Bound：已与某个PVC绑定。
- Released：绑定的PVC已经删除，资源已释放，但没有被集群回收。
- Failed：自动资源回收失败。
## k8s所支持的存储供应模式
k8s支持两种资源的存储供应模式：静态模式（Static）和动态模式（Dynamic）。

- 静态模式：集群管理员手工创建许多PV，在定义PV时需要将后端存储的特性进行设置。
- 动态模式：集群管理员无须手工创建PV，而是通过StorageClass的设置对后端存储进行描述，标记为某种类型。此时要求PVC对存储的类型进行声明，系统将自动完成PV的创建及与PVC的绑定。
## k8s CSI模型
k8s CSI是k8s推出与容器对接的存储接口标准，存储提供方只需要基于标准接口进行存储插件的实现，就能使用k8s的原生存储机制为容器提供存储服务。CSI使得存储提供方的代码能和k8s代码彻底解耦，部署也与k8s核心组件分离，显然，存储插件的开发由提供方自行维护，就能为k8s用户提供更多的存储功能，也更加安全可靠。
CSI包括CSI Controller和CSI Node：

- CSI Controller的主要功能是提供存储服务视角对存储资源和存储卷进行管理和操作。
- CSI Node的主要功能是对主机（Node）上的Volume进行管理和操作。
## k8s Worker节点加入集群的过程
通常需要对Worker节点进行扩容，从而将应用系统进行水平扩展。主要过程如下：

- 在该Node上安装Docker、kubelet和kube-proxy服务；
- 然后配置kubelet和kubeproxy的启动参数，将Master URL指定为当前k8s集群Master的地址，最后启动这些服务；
- 通过kubelet默认的自动注册机制，新的Worker将会自动加入现有的k8s集群中；
- k8s Master在接受了新Worker的注册之后，会自动将其纳入当前集群的调度范围。
## k8s Pod如何实现对节点的资源控制
k8s集群里的节点提供的资源主要是计算资源，计算资源是可计量的能被申请、分配和使用的基础资源。当前k8s集群中的计算资源主要包括CPU、GPU及Memory。CPU与Memory是被Pod使用的，因此在配置Pod时可以通过参数CPU Request及Memory Request为其中的每个容器指定所需使用的CPU与Memory量，k8s会根据Request的值去查找有足够资源的Node来调度此Pod。
通常，一个程序所使用的CPU与Memory是一个动态的量，确切地说，是一个范围，跟它的负载密切相关：负载增加时，CPU和Memory的使用量也会增加。
## k8s Requests和Limits如何影响Pod的调度
当一个Pod创建成功时，k8s调度器（Scheduler）会为该Pod选择一个节点来执行。对于每种计算资源（CPU和Memory）而言，每个节点都有一个能用于运行Pod的最大容量值。调度器在调度时，首先要确保调度后该节点上所有Pod的CPU和内存的Requests总和，不超过该节点能提供给Pod使用的CPU和Memory的最大容量值。
## k8s Metric Service
在k8s从1.10版本后采用Metrics Server作为默认的性能数据采集和监控，主要用于提供核心指标（Core Metrics），包括Node、Pod的CPU和内存使用指标。
对其他自定义指标（Custom Metrics）的监控则由Prometheus等组件来完成。
## k8s中，如何使用EFK实现日志的统一管理
在k8s集群环境中，通常一个完整的应用或服务涉及组件过多，建议对日志系统进行集中化管理，通常采用EFK实现。
EFK是 Elasticsearch、Fluentd 和 Kibana 的组合，其各组件功能如下：

- Elasticsearch：是一个搜索引擎，负责存储日志并提供查询接口；
- Fluentd：负责从 k8s 搜集日志，每个Node节点上面的Fluentd监控并收集该节点上面的系统日志，并将处理过后的日志信息发送给Elasticsearch；
- Kibana：提供了一个 Web GUI，用户可以浏览和搜索存储在 Elasticsearch 中的日志。

通过在每台Node上部署一个以DaemonSet方式运行的Fluentd来收集每台Node上的日志。Fluentd将Docker日志目录/var/lib/docker/containers和/var/log目录挂载到Pod中，然后Pod会在Node节点的/var/log/pods目录中创建新的目录，可以区别不同的容器日志输出，该目录下有一个日志文件链接到/var/lib/docker/contianers目录下的容器日志输出。
## k8s如何进行优雅的节点关机维护
由于k8s节点运行大量Pod，因此在进行关机维护之前，建议先使用kubectl drain将该节点的Pod进行驱逐，然后进行关机维护。
## k8s集群联邦
k8s集群联邦可以将多个k8s集群作为一个集群进行管理。因此，可以在一个数据中心/云中创建多个k8s集群，并使用集群联邦在一个地方控制/管理所有集群。
## Helm及其优势
Helm是k8s的软件包管理工具。类似Ubuntu中使用的APT、CentOS中使用的yum 或者Python中的 pip 一样。
Helm能够将一组k8s资源打包统一管理, 是查找、共享和使用为k8s构建的软件的最佳方式。
Helm中通常每个包称为一个Chart，一个Chart是一个目录（一般情况下会将目录进行打包压缩，形成name-version.tgz格式的单一文件，方便传输和存储）。
在k8s中部署一个可以使用的应用，需要涉及到很多的 k8s 资源的共同协作。使用Helm则具有如下优势：

- 统一管理、配置和更新这些分散的k8s的应用资源文件；
- 分发和复用一套应用模板；
- 将应用的一系列资源当做一个软件包管理。
- 对于应用发布者而言，可以通过 Helm 打包应用、管理应用依赖关系、管理应用版本并发布应用到软件仓库。

对于使用者而言，使用Helm后不用需要编写复杂的应用部署文件，可以以简单的方式在k8s上查找、安装、升级、回滚、卸载应用程序。
## k8s是什么？请说出你的了解？
答：Kubenetes是一个针对容器应用，进行自动部署，弹性伸缩和管理的开源系统。主要功能是生产环境中的容器编排。
K8S是Google公司推出的，它来源于由Google公司内部使用了15年的Borg系统，集结了Borg的精华。
## K8s架构的组成是什么？
答：和大多数分布式系统一样，K8S集群至少需要一个主节点（Master）和多个计算节点（Node）。

- 主节点主要用于暴露API，调度部署和节点的管理；
- 计算节点运行一个容器运行环境，一般是docker环境（类似docker环境的还有rkt），同时运行一个K8s的代理（kubelet）用于和master通信。计算节点也会运行一些额外的组件，像记录日志，节点监控，服务发现等等。计算节点是k8s集群中真正工作的节点。
#### K8S架构细分：
1、Master节点（默认不参加实际工作）：

- Kubectl：客户端命令行工具，作为整个K8s集群的操作入口；
- Api Server：在K8s架构中承担的是“桥梁”的角色，作为资源操作的唯一入口，它提供了认证、授权、访问控制、API注册和发现等机制。客户端与k8s群集及K8s内部组件的通信，都要通过Api Server这个组件；
- Controller-manager：负责维护群集的状态，比如故障检测、自动扩展、滚动更新等；
- Scheduler：负责资源的调度，按照预定的调度策略将pod调度到相应的node节点上；
- Etcd：担任数据中心的角色，保存了整个群集的状态；

2、Node节点：

- Kubelet：负责维护容器的生命周期，同时也负责Volume和网络的管理，一般运行在所有的节点，是Node节点的代理，当Scheduler确定某个node上运行pod之后，会将pod的具体信息（image，volume）等发送给该节点的kubelet，kubelet根据这些信息创建和运行容器，并向master返回运行状态。（自动修复功能：如果某个节点中的容器宕机，它会尝试重启该容器，若重启无效，则会将该pod杀死，然后重新创建一个容器）；
- Kube-proxy：Service在逻辑上代表了后端的多个pod。负责为Service提供cluster内部的服务发现和负载均衡（外界通过Service访问pod提供的服务时，Service接收到的请求后就是通过kube-proxy来转发到pod上的）；
- container-runtime：是负责管理运行容器的软件，比如docker
- Pod：是k8s集群里面最小的单位。每个pod里边可以运行一个或多个container（容器），如果一个pod中有两个container，那么container的USR（用户）、MNT（挂载点）、PID（进程号）是相互隔离的，UTS（主机名和域名）、IPC（消息队列）、NET（网络栈）是相互共享的。我比较喜欢把pod来当做豌豆夹，而豌豆就是pod中的container；
## 容器和主机部署应用的区别是什么？
答：容器的中心思想就是秒级启动；一次封装、到处运行；这是主机部署应用无法达到的效果，但同时也更应该注重容器的数据持久化问题。
另外，容器部署可以将各个服务进行隔离，互不影响，这也是容器的另一个核心概念。
## 请你说一下kubenetes针对pod资源对象的健康监测机制？
答：K8s中对于pod资源对象的健康状态检测，提供了三类probe（探针）来执行对pod的健康监测：
1） `livenessProbe`探针
可以根据用户自定义规则来判定pod是否健康，如果livenessProbe探针探测到容器不健康，则kubelet会根据其重启策略来决定是否重启，如果一个容器不包含livenessProbe探针，则kubelet会认为容器的livenessProbe探针的返回值永远成功。
2） `ReadinessProbe`探针
同样是可以根据用户自定义规则来判断pod是否健康，如果探测失败，控制器会将此pod从对应service的endpoint列表中移除，从此不再将任何请求调度到此Pod上，直到下次探测成功。
3） `startupProbe`探针
启动检查机制，应用一些启动缓慢的业务，避免业务长时间启动而被上面两类探针kill掉，这个问题也可以换另一种方式解决，就是定义上面两类探针机制时，初始化时间定义的长一些即可。
每种探测方法能支持以下几个相同的检查参数，用于设置控制检查时间：

- `initialDelaySeconds`：初始第一次探测间隔，用于应用启动的时间，防止应用还没启动而健康检查失败
- `periodSeconds`：检查间隔，多久执行probe检查，默认为10s；
- `timeoutSeconds`：检查超时时长，探测应用timeout后为失败；
- `successThreshold`：成功探测阈值，表示探测多少次为健康正常，默认探测1次。

上面两种探针都支持以下三种探测方法：
1）Exec：通过执行命令的方式来检查服务是否正常，比如使用cat命令查看pod中的某个重要配置文件是否存在，若存在，则表示pod健康。反之异常。
Exec探测方式的yaml文件语法如下：
```yaml
spec:
  containers:
  - name: liveness
    image: k8s.gcr.io/busybox
    args:
    - /bin/sh
    - -c
    - touch /tmp/healthy; sleep 30; rm -rf /tmp/healthy; sleep 600
    livenessProbe:         #选择livenessProbe的探测机制
      exec:                      #执行以下命令
        command:
        - cat
        - /tmp/healthy
      initialDelaySeconds: 5          #在容器运行五秒后开始探测
      periodSeconds: 5               #每次探测的时间间隔为5秒
```
在上面的配置文件中，探测机制为在容器运行5秒后，每隔五秒探测一次，如果cat命令返回的值为“0”，则表示健康，如果为非0，则表示异常。
2）Httpget：通过发送http/htps请求检查服务是否正常，返回的状态码为200-399则表示容器健康（注http get类似于命令`curl -I`）。
Httpget探测方式的yaml文件语法如下：
```yaml
spec:
  containers:
  - name: liveness
    image: k8s.gcr.io/liveness
    livenessProbe:              #采用livenessProbe机制探测
      httpGet:                  #采用httpget的方式
    scheme:HTTP         #指定协议，也支持https
        path: /healthz          #检测是否可以访问到网页根目录下的healthz网页文件
        port: 8080              #监听端口是8080
      initialDelaySeconds: 3     #容器运行3秒后开始探测
      periodSeconds: 3                #探测频率为3秒
```
上述配置文件中，探测方式为项容器发送HTTP GET请求，请求的是8080端口下的healthz文件，返回任何大于或等于200且小于400的状态码表示成功。任何其他代码表示异常。
3）tcpSocket：通过容器的IP和Port执行TCP检查，如果能够建立TCP连接，则表明容器健康，这种方式与HTTPget的探测机制有些类似，tcpsocket健康检查适用于TCP业务。
tcpSocket探测方式的yaml文件语法如下：
```yaml
spec:
  containers:
  - name: goproxy
    image: k8s.gcr.io/goproxy:0.1
    ports:
- containerPort: 8080
#这里两种探测机制都用上了，都是为了和容器的8080端口建立TCP连接
    readinessProbe:
      tcpSocket:
        port: 8080
      initialDelaySeconds: 5
      periodSeconds: 10
    livenessProbe:
      tcpSocket:
        port: 8080
      initialDelaySeconds: 15
      periodSeconds: 20
```
在上述的yaml配置文件中，两类探针都使用了，在容器启动5秒后，kubelet将发送第一个readinessProbe探针，这将连接容器的8080端口，如果探测成功，则该pod为健康，十秒后，kubelet将进行第二次连接。
除了readinessProbe探针外，在容器启动15秒后，kubelet将发送第一个livenessProbe探针，仍然尝试连接容器的8080端口，如果连接失败，则重启容器。
探针探测的结果无外乎以下三者之一：

- Success：Container通过了检查；
- Failure：Container没有通过检查；
- Unknown：没有执行检查，因此不采取任何措施（通常是我们没有定义探针检测，默认为成功）。

若觉得上面还不够透彻，可以移步其官网文档：
[https://k8s.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/)
## 如何控制滚动更新过程？
答：可以通过下面的命令查看到更新时可以控制的参数：
```bash
[root@master yaml]# kubectl explain deploy.spec.strategy.rollingUpdate
```
`maxSurge`： 此参数控制滚动更新过程，副本总数超过预期pod数量的上限。可以是百分比，也可以是具体的值。默认为1。
（上述参数的作用就是在更新过程中，值若为3，那么不管三七二一，先运行三个pod，用于替换旧的pod，以此类推）
`maxUnavailable`：此参数控制滚动更新过程中，不可用的Pod的数量。
（这个值和上面的值没有任何关系，举个例子：我有十个pod，但是在更新的过程中，我允许这十个pod中最多有三个不可用，那么就将这个参数的值设置为3，在更新的过程中，只要不可用的pod数量小于或等于3，那么更新过程就不会停止）。
## K8s中镜像的下载策略是什么？
答：可通过命令“kubectl explain pod.spec.containers”来查看imagePullPolicy这行的解释。
K8s的镜像下载策略有三种：Always、Never、IFNotPresent；

- Always：镜像标签为latest时，总是从指定的仓库中获取镜像；
- Never：禁止从仓库中下载镜像，也就是说只能使用本地镜像；
- IfNotPresent：仅当本地没有对应镜像时，才从目标仓库中下载。
- 默认的镜像下载策略是：当镜像标签是latest时，默认策略是Always；当镜像标签是自定义时（也就是标签不是latest），那么默认策略是IfNotPresent。
## image的状态有哪些？

- Running：Pod所需的容器已经被成功调度到某个节点，且已经成功运行，
- Pending：APIserver创建了pod资源对象，并且已经存入etcd中，但它尚未被调度完成或者仍然处于仓库中下载镜像的过程
- Unknown：APIserver无法正常获取到pod对象的状态，通常是其无法与所在工作节点的kubelet通信所致。
## pod的重启策略是什么？
答：可以通过命令“kubectl explain pod.spec”查看pod的重启策略。（restartPolicy字段）

- Always：但凡pod对象终止就重启，此为默认策略。
- OnFailure：仅在pod对象出现错误时才重启
## Service这种资源对象的作用是什么？
答：用来给相同的多个pod对象提供一个固定的统一访问接口，常用于服务发现和服务访问。
## 版本回滚相关的命令？
```bash
[root@master httpd-web]# kubectl apply -f httpd2-deploy1.yaml  --record  
#运行yaml文件，并记录版本信息；
[root@master httpd-web]# kubectl rollout history deployment httpd-devploy1  
#查看该deployment的历史版本
[root@master httpd-web]# kubectl rollout undo deployment httpd-devploy1 --to-revision=1    
#执行回滚操作，指定回滚到版本1
#在yaml文件的spec字段中，可以写以下选项（用于限制最多记录多少个历史版本）：
spec:
  revisionHistoryLimit: 5            
#这个字段通过 kubectl explain deploy.spec  命令找到revisionHistoryLimit   <integer>行获得
```
## 标签与标签选择器的作用是什么？
标签：是当相同类型的资源对象越来越多的时候，为了更好的管理，可以按照标签将其分为一个组，为的是提升资源对象的管理效率。
标签选择器：就是标签的查询过滤条件。目前API支持两种标签选择器：

- 基于等值关系的，如：“=”、“”“”、“！=”（注：“”也是等于的意思，yaml文件中的matchLabels字段）；
- 基于集合的，如：in、notin、exists（yaml文件中的matchExpressions字段）；

注：in:在这个集合中；notin：不在这个集合中；exists：要么全在（exists）这个集合中，要么都不在（notexists）；
使用标签选择器的操作逻辑：

- 在使用基于集合的标签选择器同时指定多个选择器之间的逻辑关系为“与”操作（比如：- {key: name,operator: In,values: [zhangsan,lisi]} ，那么只要拥有这两个值的资源，都会被选中）；
- 使用空值的标签选择器，意味着每个资源对象都被选中（如：标签选择器的键是“A”，两个资源对象同时拥有A这个键，但是值不一样，这种情况下，如果使用空值的标签选择器，那么将同时选中这两个资源对象）
- 空的标签选择器（注意不是上面说的空值，而是空的，都没有定义键的名称），将无法选择出任何资源；

在基于集合的选择器中，使用“In”或者“Notin”操作时，其values可以为空，但是如果为空，这个标签选择器，就没有任何意义了。
两种标签选择器类型（基于等值、基于集合的书写方法）：
```yaml
selector:
  matchLabels:           #基于等值
    app: nginx
  matchExpressions:         #基于集合
    - {key: name,operator: In,values: [zhangsan,lisi]}     #key、operator、values这三个字段是固定的
    - {key: age,operator: Exists,values:}   #如果指定为exists，那么values的值一定要为空
```
## 常用的标签分类有哪些？
标签分类是可以自定义的，但是为了能使他人可以达到一目了然的效果，一般会使用以下一些分类：

- 版本类标签（release）：stable（稳定版）、canary（金丝雀版本，可以将其称之为测试版中的测试版）、beta（测试版）；
- 环境类标签（environment）：dev（开发）、qa（测试）、production（生产）、op（运维）；
- 应用类（app）：ui、as、pc、sc；
- 架构类（tier）：frontend（前端）、backend（后端）、cache（缓存）；
- 分区标签（partition）：customerA（客户A）、customerB（客户B）；
- 品控级别（Track）：daily（每天）、weekly（每周）。
## 有几种查看标签的方式？
答：常用的有以下三种查看方式：
```bash
[root@master ~]# kubectl get pod --show-labels    #查看pod，并且显示标签内容
[root@master ~]# kubectl get pod -L env,tier      #显示资源对象标签的值
[root@master ~]# kubectl get pod -l env,tier      #只显示符合键值资源对象的pod，而“-L”是显示所有的pod
```
## 添加、修改、删除标签的命令？
```bash
#对pod标签的操作
[root@master ~]# kubectl label pod label-pod abc=123     #给名为label-pod的pod添加标签
[root@master ~]# kubectl label pod label-pod abc=456 --overwrite       #修改名为label-pod的标签
[root@master ~]# kubectl label pod label-pod abc-             #删除名为label-pod的标签
[root@master ~]# kubectl get pod --show-labels
 
#对node节点的标签操作   
[root@master ~]# kubectl label nodes node01 disk=ssd      #给节点node01添加disk标签
[root@master ~]# kubectl label nodes node01 disk=sss –overwrite    #修改节点node01的标签
[root@master ~]# kubectl label nodes node01 disk-         #删除节点node01的disk标签
```
## DaemonSet资源对象的特性？
DaemonSet这种资源对象会在每个k8s集群中的节点上运行，并且每个节点只能运行一个pod，这是它和deployment资源对象的最大也是唯一的区别。所以，在其yaml文件中，不支持定义replicas，除此之外，与Deployment、RS等资源对象的写法相同。
它的一般使用场景如下：

- 在去做每个节点的日志收集工作；
- 监控每个节点的的运行状态；
## 说说你对Job这种资源对象的了解？
答：Job与其他服务类容器不同，Job是一种工作类容器（一般用于做一次性任务）。使用常见不多，可以忽略这个问题。
```yaml
#提高Job执行效率的方法：
spec:
  parallelism: 2           #一次运行2个
  completions: 8           #最多运行8个
  template:
metadata:
```
## 描述一下pod的生命周期有哪些状态？

- Pending：表示pod已经被同意创建，正在等待kube-scheduler选择合适的节点创建，一般是在准备镜像；
- Running：表示pod中所有的容器已经被创建，并且至少有一个容器正在运行或者是正在启动或者是正在重启；
- Succeeded：表示所有容器已经成功终止，并且不会再启动；
- Failed：表示pod中所有容器都是非0（不正常）状态退出；
- Unknown：表示无法读取Pod状态，通常是kube-controller-manager无法与Pod通信。
## 创建一个pod的流程是什么？

- 客户端提交Pod的配置信息（可以是yaml文件定义好的信息）到kube-apiserver；
- Apiserver收到指令后，通知给controller-manager创建一个资源对象；
- Controller-manager通过api-server将pod的配置信息存储到ETCD数据中心中；
- Kube-scheduler检测到pod信息会开始调度预选，会先过滤掉不符合Pod资源配置要求的节点，然后开始调度调优，主要是挑选出更适合运行pod的节点，然后将pod的资源配置单发送到node节点上的kubelet组件上。
- Kubelet根据scheduler发来的资源配置单运行pod，运行成功后，将pod的运行信息返回给scheduler，scheduler将返回的pod运行状况的信息存储到etcd数据中心。
## 删除一个Pod会发生什么事情？
答：Kube-apiserver会接受到用户的删除指令，默认有30秒时间等待优雅退出，超过30秒会被标记为死亡状态，此时Pod的状态Terminating，kubelet看到pod标记为Terminating就开始了关闭Pod的工作；
关闭流程如下：

- pod从service的endpoint列表中被移除；
- 如果该pod定义了一个停止前的钩子，其会在pod内部被调用，停止钩子一般定义了如何优雅的结束进程；
- 进程被发送TERM信号（kill -14）
- 当超过优雅退出的时间后，Pod中的所有进程都会被发送SIGKILL信号（kill -9）。
## K8s的Service是什么？
答：Pod每次重启或者重新部署，其IP地址都会产生变化，这使得pod间通信和pod与外部通信变得困难，这时候，就需要Service为pod提供一个固定的入口。
Service的Endpoint列表通常绑定了一组相同配置的pod，通过负载均衡的方式把外界请求分配到多个pod上
## k8s是怎么进行服务注册的？
答：Pod启动后会加载当前环境所有Service信息，以便不同Pod根据Service名进行通信。
## k8s集群外流量怎么访问Pod？
答：可以通过Service的NodePort方式访问，会在所有节点监听同一个端口，比如：30000，访问节点的流量会被重定向到对应的Service上面。
## k8s数据持久化的方式有哪些？
答：
#### 1）EmptyDir（空目录）：
没有指定要挂载宿主机上的某个目录，直接由Pod内保部映射到宿主机上。类似于docker中的manager volume。
主要使用场景：

- 只需要临时将数据保存在磁盘上，比如在合并/排序算法中；
- 作为两个容器的共享存储，使得第一个内容管理的容器可以将生成的数据存入其中，同时由同一个webserver容器对外提供这些页面。

emptyDir的特性：
同个pod里面的不同容器，共享同一个持久化目录，当pod节点删除时，volume的数据也会被删除。如果仅仅是容器被销毁，pod还在，则不会影响volume中的数据。
总结来说：emptyDir的数据持久化的生命周期和使用的pod一致。一般是作为临时存储使用。
#### 2）Hostpath：
将宿主机上已存在的目录或文件挂载到容器内部。类似于docker中的bind mount挂载方式。
这种数据持久化方式，运用场景不多，因为它增加了pod与节点之间的耦合。
一般对于k8s集群本身的数据持久化和docker本身的数据持久化会使用这种方式，可以自行参考apiService的yaml文件，位于：/etc/k8s/main…目录下。
#### 3）PersistentVolume（简称PV）：
基于NFS服务的PV，也可以基于GFS的PV。它的作用是统一数据持久化目录，方便管理。
在一个PV的yaml文件中，可以对其配置PV的大小，指定PV的访问模式：

- `ReadWriteOnce`：只能以读写的方式挂载到单个节点；
- `ReadOnlyMany`：能以只读的方式挂载到多个节点；
- `ReadWriteMany`：能以读写的方式挂载到多个节点。以及指定pv的回收策略：
- `recycle`：清除PV的数据，然后自动回收；
- `Retain`：需要手动回收；
- `delete`：删除云存储资源，云存储专用；

PS：这里的回收策略指的是在PV被删除后，在这个PV下所存储的源文件是否删除）。
若需使用PV，那么还有一个重要的概念：PVC，PVC是向PV申请应用所需的容量大小，K8s集群中可能会有多个PV，PVC和PV若要关联，其定义的访问模式必须一致。定义的storageClassName也必须一致，若群集中存在相同的（名字、访问模式都一致）两个PV，那么PVC会选择向它所需容量接近的PV去申请，或者随机申请。
