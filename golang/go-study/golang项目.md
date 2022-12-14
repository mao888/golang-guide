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

源码是深入学习的一种方式，也是面试求职的亮点之一。源码其实也可以作为开源项目的一种形式，在精不再多