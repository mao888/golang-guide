### Go官方基础指南-中文

https://tour.go-zh.org/welcome/1

这个是一个官方的入门教程，或者说只是一个大概了解的教程，只介绍了一些简单的东西，并且没有太多的说明。不过这个教程支持在线执行代码，还是很不错的，这个时候你都不需要有本地的开发环境。不用想太多，现在就开始，把这个教程从头到尾看一遍，练习一遍，遇到不明白的地方也不要纠结，继续向后看就行了。

### 开发环境

Go 的安装非常的简单，没有太多的依赖，如果是 Linux 下安装基本上下载一个二进制包，解压配置上一个环境变量、GOROOT 既可以了，具体的可以查看官方的安装方法： [下载安装文档](https://go.dev/doc/install)

### 开发工具

这里推荐VSCode或GoLand，更多Go编辑器和IDE可参考[官方](https://github.com/golang/go/wiki/IDEsAndTextEditorPlugins)

### 一套基础视频+一套文档教程

[Go 编程基础](https://learnku.com/docs/go-fundamental-programming) （视频）

[Go 入门指南](https://learnku.com/docs/the-way-to-go)（教程）[仓库地址](https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/directory.md)



[李文周博客教程和视频](https://www.qfgolang.com/)

[Golang中国教程和视频](https://www.liwenzhou.com/archives/)

### 将标准库过一遍

至少要把常用的全都看一遍，如 strings /strconv/http 等，如果有能力可以将它们都记住，如果记忆力不太好，至少也要知道有什么，用到的时候通过手册可以快速找到。 

1. 官方标准库： https://golang.org/pkg/
2. 中文版的标准库： https://studygolang.com/static/pkgdoc/main.html
3. 推荐 https://github.com/astaxie/gopkg 和[《Go 语言标准库》The Golang Standard Library by Example](https://studygolang.com/static/pkgdoc/main.html) ，有关于标准库的详细说明和示例，通过这两个文档库学习起来会容易一些，等全都明白了要使用的时候可以去查看上面的文档。

### 基础书籍推荐

纸质书：Go语言学习笔记、HeadFirst Go语言程序设计 

电子书：[Go By Example 中文版](https://gobyexample-cn.github.io/)、[跟着单元测试学习 Go](https://studygolang.gitbook.io/learn-go-with-tests/)

### 写案例

这个时候一般都已经入门，可以试着写点东西，比如写一个博客、小系统，或者去学习一个框架，提升自己Go Web和Go Api的开发能力。

https://www.yuque.com/go/doc/68565648

### 初级项目案例（实现或参考实现其中之一）

1. 使用Go生成GitHub上面项目的star趋势图  https://github.com/caarlos0/starcharts
2. 使用Go写的吃豆人小游戏，每一步都有详细的描述和代码实现  https://github.com/danicat/pacgo
3. 微信 web 版 API 的 Go 实现，模拟微信网页版的登录／联系人／消息收发等功能 https://github.com/songtianyi/wechat-go

### Go Web开发

基础知识掌握之后，可以上手做一些 web 应用，进一步了解更多的 Go 语言相关框架以及生产环境中的常用中间件，推荐书籍《Go Web 编程》

1. **Gin框架**

官方文档都有中文，照着 demo 敲一下，了解下怎么处理 HTTP 请求的

1. **ORM 框架 Gorm**

有官方中文文档，照着 demo 敲一下基本上两天就能掌握了，后面遇到不会的再来查

1. [**Casbin**](https://casbin.org/docs/zh-CN/get-started)

能够基于Casbin进行的控制处理

### Web项目推荐（实现或参考实现其中之一）

1. https://www.yuque.com/go/doc/68565677： 一个 Go 语言入门项目，旨在让初学者花尽可能短的时间，通过尽可能详细的步骤，历经 17 个 demo，最终一步步构建出一个生产级的 Go 后端服务器。从开发准备到 API 设计，再到 API 实现、测试和部署，每一步都详细介绍了如何去构建
2. [gin-vue-admin](https://github.com/flipped-aurora/gin-vue-admin)：使用 Gin 框架构建的后台管理系统，有详细的中文文档，并且配套教学视频https://github.com/flipped-aurora/gin-vue-admin
3. [ferry](https://github.com/lanyulei/ferry)：基于 Gin + Vue + Element UI 前后端分离的工单系统。https://github.com/lanyulei/ferry
4. [Go-admi](http://github.com/go-admin-team/go-admin)：Gin + Vue + Element UI 的前后端分离权限管理系统，有详细中文文档和配套视频教程 http://github.com/go-admin-team/go-admin

对于 web 项目的学习，可能有同学觉得项目太庞杂，根本不知道怎么下手。我想建议的是，可以在本地把项目跑起来，然后断点调试一个 HTTP 请求的整体流程，搞懂了一个接口，其他的大同小异。

### https://www.zhihu.com/question/399923003

### 进阶项目（实现或参考实现其中之一）

1. 一个 Go 语言实现的快速、稳定、内嵌的 k-v 存储引擎 [rosedb](https://github.com/flower-corp/rosedb) https://github.com/flower-corp/rosedb

视频：[space.bilibili.com/26194591](https://space.bilibili.com/26194591) （作者手把手带你用go实现一个数据库）

1. [gochat](http://github.com/LockGit/gochat) http://github.com/LockGit/gochat：一个 Go 语言实现的轻量级 im 系统，对网络方面熟悉或者感兴趣的可以看看。http://github.com/LockGit/gochat
2. 7DaysGolang：[7days-golang](http://github.com/geektutu/7days-golang)：http://github.com/geektutu/7days-golang，7 天使用 Go 从零实现 web 框架、分布式缓存、ORM 框架，、RPC 框架，代码量不多，但是质量挺不错的 http://github.com/geektutu/7days-golang
3. 云捷Go：https://gitee.com/yyz116/vigo 我翻过很多优秀的开源项目，一直没找到类似于若依开发思路的快速
   开发框架。而这个项目就是试图在用go写了一套类似若依的后台系统。这个框架可以用于所有的web应用程序，
   如网站管理后台，网站会员中心，CMS, CRM，OA。所有前端后台代码封装过后十分精简易上手，出错概率
   低。该框架以GoFrame为web服务框架，架构思路沿袭着若依的以辅助生成重复代码为主，没有过度封装，生成
   的代码可以快速修改适应不同的需求，适应每个开发者自己的习惯和风格，是很好的框架参考样本。
4. Filber:是一个受Express：https://github.com/expressjs/express启发的Web 框架，构建在Fasthttp：https://github.com/valyala/fasthttp之上，这是Go中最快的HTTP引擎。旨在简化快速开发的工作，同时考虑到零内存分配和性能。该项目借鉴 nodej's 框架的思路是很有启发性的，同时我们也能基于 Fasthttp 在项目和简历做出一些突出点。
   项目地址：https://aithub.com/gofiber/fiber
5. novel-fpg：https://github.com/black-currant/novel-fpg是一个小说开源项目，前端基于Flutter，后端是Python + Go双端(各自基于Flask、Gin实现了-遍），包含了用户注册、登录、iwt鉴权、签到、任务、书架、阅读器、购买章节、搜索书籍、绑定第三方账号、设置等功能。该项目业务功能比较丰富，并且基于 Gin 框架开发，业务开发能力不强的同学，可以学习-下
6. 源码系列：Go 源码、Gin 框架源码、Gorm 源码、zap 源码、标准库源码（如ioutill包、http、log、Timer)

码是深入学习的一种方式，也是面试求职的亮点之一。源码其实也可以作为开源项目的一种形式，在精不再多

源

### 进阶书籍和资料

1. 纸质书：go语言圣经、Go专家编程、Go语言高级编程、Go语言设计与实现
2. 在线版：[Go语言设计与实现](https://draveness.me/golang/)：https://draveness.me/golang/、[Go语言高级编程 ](https://chai2010.cn/advanced-go-programming-book/)[官方《Effective Go》 中文版](https://learnku.com/docs/effective-go/2020)：https://chai2010.cn/advanced-go-programming-book/ 进阶 - 技巧规范篇
3.  [Go2编程指南](https://github.com/chai2010/go2-book#go2编程指南)：[https://github.com/chai2010/go2-book#go2%E7%BC%96%E7%A8%8B%E6%8C%87%E5%8D%97](https://github.com/chai2010/go2-book#go2编程指南) 进阶 - 新版本讲解
4. [Go语法树入门——开启自制编程语言和编译器之旅！](https://github.com/chai2010/go-ast-book) ：https://github.com/chai2010/go-ast-book深入高级
5.  [幼麟实验室（图解 Go和操作系统等内容，深入底层）](https://space.bilibili.com/567195437)：https://space.bilibili.com/567195437

### 面试

[Go 面试题](http://www.topgoer.cn/docs/gomianshiti/mianshiti)

------

### [rabbitmq](https://github.com/rabbitmq)

- http://g0d_v.gitee.io/rabbitmq-cn/#/?id=quothello-worldquot
- https://github.com/rabbitmq/amqp091-go
- https://github.com/streadway/amqp
- Documentation

- - [Godoc API reference](http://godoc.org/github.com/rabbitmq/amqp091-go)
  - [RabbitMQ tutorials in Go](https://github.com/rabbitmq/rabbitmq-tutorials/tree/master/go) currently use a different client. They will be switched to use this client eventually

- https://rabbitmq.com/getstarted.html
- [RabbitMQ Go客户端教程1——HelloWorld（翻译）](https://www.liwenzhou.com/posts/Go/go_rabbitmq_tutorials_01/)
- [RabbitMQ Go客户端教程2——任务队列（翻译）](https://www.liwenzhou.com/posts/Go/go_rabbitmq_tutorials_02/)
- [RabbitMQ Go客户端教程3——发布/订阅（翻译）](https://www.liwenzhou.com/posts/Go/go_rabbitmq_tutorials_03/)
- [RabbitMQ Go客户端教程4——路由（翻译）](https://www.liwenzhou.com/posts/Go/go_rabbitmq_tutorials_04/)
- [RabbitMQ Go客户端教程5——topic（翻译）](https://www.liwenzhou.com/posts/Go/go_rabbitmq_tutorials_05/)
- [RabbitMQ Go客户端教程6——RPC（翻译）](https://www.liwenzhou.com/posts/Go/go_rabbitmq_tutorials_06/)

------

### 微服务(暂时先不学，后期根据时间再做打算)

目前 Go 在微服务中的应用也比较广泛，但说实话，微服务是一个太庞大的话题，你不可能把每一个核心的问题都能够搞清楚，而且也没条件，或许只能在公司的具体的微服务生产环境中，才能够对相关的概念有更加深刻的体会。推荐微服务概述的基础书籍《微服务设计》、《微服务架构设计模式》，可以帮助你理解微服务的建模、集成、测试、部署和监控的一些基础知识。



 推荐 Go 语言的微服务框架 GoKit、GoMicro、go-zero、kratos，可以随便选择一个，理解其基本的用法、设计等等。其中 go-zero 和 kratos 是国内开源的，因此都有比较详细的中文文档。

 一个在线学习的资料：https://ewanvalentine.io/microservices-in-golang-part-1/手把手实现一个简单的 Go 微服务项目，你可以通过这个项目来学习微服务的相关知识，并且有中文版。

#### go-zero

- https://github.com/zeromicro/go-zero
- [go-zero官方文档](https://go-zero.dev/cn/)
- https://go-zero.dev/cn/docs/goctl/goctl
- https://github.com/zeromicro/zero-doc/blob/main/go-zero.dev/cn/api-grammar.md
- https://www.cnblogs.com/haima/p/16057786.html
- https://github.com/zeromicro/go-zero-demo
- 视频教程-https://space.bilibili.com/389552232/channel/seriesdetail?sid=2122723

####  RPC

- [https://www.topgoer.com/%E5%BE%AE%E6%9C%8D%E5%8A%A1/RPC.html](https://www.topgoer.com/微服务/RPC.html)
- [https://www.topgoer.com/%E5%BE%AE%E6%9C%8D%E5%8A%A1/gRPC/Protobuf%E8%AF%AD%E6%B3%95.html?h=proto](https://www.topgoer.com/微服务/gRPC/Protobuf语法.html?h=proto)
- https://blog.csdn.net/qianfeng_dashuju/article/details/109196739
- https://blog.csdn.net/weixin_45413603/article/details/121514386

------