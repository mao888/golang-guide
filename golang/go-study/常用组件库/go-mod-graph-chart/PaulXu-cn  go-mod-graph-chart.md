Golang 遇到 note: module requires Go 1.xx 解决之道，不升 go
在使用多版本的 golang 的小伙伴，往往会遇到这个问题。本文就如何不升级 go 版本来解决该问题。

## 怎么产生的

1. 同时使用 IDE 和命令行，命令行 go 版本和 IDE 的 go 不是同一个版本。
2. 多人协作同一个项目，别人使用不同版本的 go 加入了一个包且写入了一个较高版本号。
3. go get 时选择了较高版本。

网上的大部分解决方案都是升级自己的 go 版本。那么如果不想升级有办法解决么？

## 不升 Go 版本

例如我遇到了：

```go
# golang.org/x/sys/unix
../../go/pkg/mod/golang.org/x/sys@v0.3.0/unix/syscall.go:83:16: undefined: unsafe.Slice
../../go/pkg/mod/golang.org/x/sys@v0.3.0/unix/syscall_darwin.go:95:8: undefined: unsafe.Slice
../../go/pkg/mod/golang.org/x/sys@v0.3.0/unix/syscall_unix.go:118:7: undefined: unsafe.Slice
../../go/pkg/mod/golang.org/x/sys@v0.3.0/unix/sysvshm_unix.go:33:7: undefined: unsafe.Slice
note: module requires Go 1.17
```

尝试把 go.sum 删掉，再 go mod tidy 还是没用，依然报这个错。

这时看看 go.mod 文件

```go
module github.com/PaulXu-cn/xxx

go 1.15

require (
    github.com/go-faker/faker/v4 v4.0.0-beta.4
    github.com/golang/protobuf v1.5.2
    github.com/snksoft/crc v1.1.0
    github.com/spf13/cobra v1.6.1
    github.com/spf13/viper v1.15.0
)
```

也就是说当前我的 go runtime 是 1.15 的，
是引用了基于 go1.17 的包，需要把这个包降为依赖 go1.15 的即可。那这里的哪个包需要降版本呢？

该项目简单，只有 5 个直接依赖，可以通过依次删除添加测出来，如果有很多依赖的话，又该怎么解决呢？

包后面带 //indirect 是间接依赖，删掉这一行不影响。
参考 [go.dev/ref/mod](https://go.dev/ref/mod)

## 上工具

这里介绍个工具 [gmchart](https://github.com/PaulXu-cn/go-mod-graph-chart) ， go mod 图像化展示工具 —— https://github.com/PaulXu-cn/go-mod-graph-chart

进入工作项目

```go
cd goProejct
```

安装 gmchart

```go
go install github.com/PaulXu-cn/go-mod-graph-chart/gmchart@latest
```

运行

```go
go mod graph | gmchart
go mod graph version v0.5.3
the go mod graph will top in 60s
visit it by http://127.0.0.1:59760
```

go mod graph 是官方工具命令。 可展示出了该项目所有的依赖关系，只不过是文本形式展示，输出的内容多了，人眼看不出啥来。这里借用 gmchart 工具，可以将其依赖关系组织为 树状 渲染 web 页面，也就是和 go 工具一样，跨平台的。

## 利用工具找问题

回到我们刚刚的报错啊，golang.org/x/sys 包依赖了 go1.17.

这里我们搜一下啊

![img](https://cdn.nlark.com/yuque/0/2023/png/22219483/1680752787377-efcdbe56-e044-4de1-84a0-ec30da7a3ec6.png)

呦，一下 84 个，我们再把版本号输入进去，缩小范围

![img](https://cdn.nlark.com/yuque/0/2023/png/22219483/1680752795238-fee1e897-73e5-4b15-91ac-dd6b8d166842.png)

好的，定位到了 1 个，那就是它了，然后呢？

看了看 go.mod , 好像我们也没有直接引用它，要去 go.mod 删也没有得删。

如果是如下情况，带有 indirect 注释的，删除了也不能解决问题！

```go
require(
    golang.org/x/sys v0.3.0 // indirect
)
```

这里有大聪明，建议我去 go.sum 里面去删，这是没用的哈，go mod tidy 一下又回来了。

## 找出直接依赖

1. 用工具找到具体包
2. 在界面中点击包
3. 查看所有引入了该包的 包。这里我们看到了 5 个包，viper 这个包是直接依赖，因此该调整这个包的版本

![img](https://cdn.nlark.com/yuque/0/2023/png/22219483/1680752858017-d63622b8-c11d-49f7-8066-7ab402194039.png)

1. 在 github 上找到 viper https://github.com/spf13/viper
2. 打开 go.mod —— https://github.com/spf13/viper/blob/master/go.mod

![img](https://cdn.nlark.com/yuque/0/2023/png/22219483/1680752912319-c7eb4db3-20bc-4900-9fae-f6b6b093db8d.png)

这里通过查看 viper 的 go.mod ，发现最新的 viper 已是基于 go1.17 , 我的项目 go get 了最新版本的 viper，所以，编译是就会报错 ——note: module requires Go 1.17

这里按理说不会拉取高版本的 viper，但这里是切换了 go 版本，导致了该情况。

1. 项目要求是不高于 go1.15，那就依次便利 viper 的各个 tag。好 ——viper[@1.9.0 ]() 是当前版本最高且要求不高于 go1.15 的。 

![img](https://cdn.nlark.com/yuque/0/2023/png/22219483/1680752952672-aa393448-bbd9-4ef2-927b-1649a27e8b3a.png)

1. 把项目中 go.mod, 依赖 —— viper 版本改为 v1.9.0 报错就解决了。

## 总结

我最近在做项目时，切换 go 版本遇到了该问题，顺手查了下，发现网上的答案都时让其升级 go 版本，其实就是依赖高版本的第三方包，这里借用工具，找出该包，通过降第三方包的方式也能解决该问题。

如果大家遇到同类问题，不想升 go 版本，可以试试改方案。

没有遇到也没关系，收藏一下，某天遇到了不想升 go 可以再翻出来看看

## 参考

https://go.dev/ref/mod

https://github.com/PaulXu-cn/go-mod-graph-chart