# cron

## 简介

[`cron`](https://github.com/robfig/cron)一个用于管理定时任务的库，用 Go 实现 Linux 中`crontab`这个命令的效果。之前我们也介绍过一个类似的 Go 库——[`gron`](https://darjun.github.io/2020/04/20/godailylib/gron/)。`gron`代码小巧，用于学习是比较好的。但是它功能相对简单些，并且已经不维护了。如果有定时任务需求，还是建议使用`cron`。

## 快速使用

文本代码使用 Go Modules。

创建目录并初始化：

```cmd
$ mkdir cron && cd cron
$ go mod init github.com/darjun/go-daily-lib/cron
```

安装`cron`，目前最新稳定版本为 v3：

```cmd
$ go get -u github.com/robfig/cron/v3
```

使用：

```golang
package main

import (
  "fmt"
  "time"

  "github.com/robfig/cron/v3"
)

func main() {
  c := cron.New()

  c.AddFunc("@every 1s", func() {
    fmt.Println("tick every 1 second")
  })

  c.Start()
  time.Sleep(time.Second * 5)
}
```

使用非常简单，创建`cron`对象，这个对象用于管理定时任务。

调用`cron`对象的`AddFunc()`方法向管理器中添加定时任务。`AddFunc()`接受两个参数，参数 1 以字符串形式指定触发时间规则，参数 2 是一个无参的函数，每次触发时调用。`@every 1s`表示每秒触发一次，`@every`后加一个时间间隔，表示每隔多长时间触发一次。例如`@every 1h`表示每小时触发一次，`@every 1m2s`表示每隔 1 分 2 秒触发一次。`time.ParseDuration()`支持的格式都可以用在这里。

调用`c.Start()`启动定时循环。

注意一点，因为`c.Start()`启动一个新的 goroutine 做循环检测，我们在代码最后加了一行`time.Sleep(time.Second * 5)`防止主 goroutine 退出。

运行效果，每隔 1s 输出一行字符串：

```cmd
$ go run main.go 
tick every 1 second
tick every 1 second
tick every 1 second
tick every 1 second
tick every 1 second
```

## 时间格式

与Linux 中`crontab`命令相似，`cron`库支持用 **5** 个空格分隔的域来表示时间。这 5 个域含义依次为：

- `Minutes`：分钟，取值范围`[0-59]`，支持特殊字符`* / , -`；
- `Hours`：小时，取值范围`[0-23]`，支持特殊字符`* / , -`；
- `Day of month`：每月的第几天，取值范围`[1-31]`，支持特殊字符`* / , - ?`；
- `Month`：月，取值范围`[1-12]`或者使用月份名字缩写`[JAN-DEC]`，支持特殊字符`* / , -`；
- `Day of week`：周历，取值范围`[0-6]`或名字缩写`[JUN-SAT]`，支持特殊字符`* / , - ?`。

注意，月份和周历名称都是不区分大小写的，也就是说`SUN/Sun/sun`表示同样的含义（都是周日）。

特殊字符含义如下：

- `*`：使用`*`的域可以匹配任何值，例如将月份域（第 4 个）设置为`*`，表示每个月；
- `/`：用来指定范围的**步长**，例如将小时域（第 2 个）设置为`3-59/15`表示第 3 分钟触发，以后每隔 15 分钟触发一次，因此第 2 次触发为第 18 分钟，第 3 次为 33 分钟。。。直到分钟大于 59；
- `,`：用来列举一些离散的值和多个范围，例如将周历的域（第 5 个）设置为`MON,WED,FRI`表示周一、三和五；
- `-`：用来表示范围，例如将小时的域（第 1 个）设置为`9-17`表示上午 9 点到下午 17 点（包括 9 和 17）；
- `?`：只能用在月历和周历的域中，用来代替`*`，表示每月/周的任意一天。

了解规则之后，我们可以定义任意时间：

- `30 * * * *`：分钟域为 30，其他域都是`*`表示任意。每小时的 30 分触发；
- `30 3-6,20-23 * * *`：分钟域为 30，小时域的`3-6,20-23`表示 3 点到 6 点和 20 点到 23 点。3,4,5,6,20,21,22,23 时的 30 分触发；
- `0 0 1 1 *`：1（第 4 个） 月 1（第 3 个） 号的 0（第 2 个） 时 0（第 1 个） 分触发。

记熟了这几个域的顺序，再多练习几次很容易就能掌握格式。熟悉规则了之后，就能熟练使用`crontab`命令了。

```golang
func main() {
  c := cron.New()

  c.AddFunc("30 * * * *", func() {
    fmt.Println("Every hour on the half hour")
  })

  c.AddFunc("30 3-6,20-23 * * *", func() {
    fmt.Println("On the half hour of 3-6am, 8-11pm")
  })

  c.AddFunc("0 0 1 1 *", func() {
    fmt.Println("Jun 1 every year")
  })

  c.Start()

  for {
    time.Sleep(time.Second)
  }
}
```

### 预定义时间规则

为了方便使用，`cron`预定义了一些时间规则：

- `@yearly`：也可以写作`@annually`，表示每年第一天的 0 点。等价于`0 0 1 1 *`；
- `@monthly`：表示每月第一天的 0 点。等价于`0 0 1 * *`；
- `@weekly`：表示每周第一天的 0 点，注意第一天为周日，即周六结束，周日开始的那个 0 点。等价于`0 0 * * 0`；
- `@daily`：也可以写作`@midnight`，表示每天 0 点。等价于`0 0 * * *`；
- `@hourly`：表示每小时的开始。等价于`0 * * * *`。

例如：

```golang
func main() {
  c := cron.New()

  c.AddFunc("@hourly", func() {
    fmt.Println("Every hour")
  })

  c.AddFunc("@daily", func() {
    fmt.Println("Every day on midnight")
  })

  c.AddFunc("@weekly", func() {
    fmt.Println("Every week")
  })

  c.Start()

  for {
    time.Sleep(time.Second)
  }
}
```

上面代码只是演示用法，实际运行可能要等待非常长的时间才能有输出。

### 固定时间间隔

`cron`支持固定时间间隔，格式为：

```golang
@every <duration>
```

含义为每隔`duration`触发一次。`<duration>`会调用`time.ParseDuration()`函数解析，所以`ParseDuration`支持的格式都可以。例如`1h30m10s`。在快速开始部分，我们已经演示了`@every`的用法了，这里就不赘述了。

## 时区

默认情况下，所有时间都是基于当前时区的。当然我们也可以指定时区，有 2 两种方式：

- 在时间字符串前面添加一个`CRON_TZ=` + 具体时区，具体时区的格式在之前[`carbon`](https://darjun.github.io/2020/02/14/godailylib/carbon/)的文章中有详细介绍。东京时区为`Asia/Tokyo`，纽约时区为`America/New_York`；
- 创建`cron`对象时增加一个时区选项`cron.WithLocation(location)`，`location`为`time.LoadLocation(zone)`加载的时区对象，`zone`为具体的时区格式。或者调用已创建好的`cron`对象的`SetLocation()`方法设置时区。

示例：

```golang
func main() {
  nyc, _ := time.LoadLocation("America/New_York")
  c := cron.New(cron.WithLocation(nyc))
  c.AddFunc("0 6 * * ?", func() {
    fmt.Println("Every 6 o'clock at New York")
  })

  c.AddFunc("CRON_TZ=Asia/Tokyo 0 6 * * ?", func() {
    fmt.Println("Every 6 o'clock at Tokyo")
  })

  c.Start()

  for {
    time.Sleep(time.Second)
  }
}
```

## `Job`接口

除了直接将无参函数作为回调外，`cron`还支持`Job`接口：

```golang
// cron.go
type Job interface {
  Run()
}
```

我们定义一个实现接口`Job`的结构：

```golang
type GreetingJob struct {
  Name string
}

func (g GreetingJob) Run() {
  fmt.Println("Hello ", g.Name)
}
```

调用`cron`对象的`AddJob()`方法将`GreetingJob`对象添加到定时管理器中：

```golang
func main() {
  c := cron.New()
  c.AddJob("@every 1s", GreetingJob{"dj"})
  c.Start()

  time.Sleep(5 * time.Second)
}
```

运行效果：

```cmd
$ go run main.go 
Hello  dj
Hello  dj
Hello  dj
Hello  dj
Hello  dj
```

使用自定义的结构可以让任务携带状态（`Name`字段）。

实际上`AddFunc()`方法内部也调用了`AddJob()`方法。首先，`cron`基于`func()`类型定义一个新的类型`FuncJob`：

```golang
// cron.go
type FuncJob func()
```

然后让`FuncJob`实现`Job`接口：

```golang
// cron.go
func (f FuncJob) Run() {
  f()
}
```

在`AddFunc()`方法中，将传入的回调转为`FuncJob`类型，然后调用`AddJob()`方法：

```golang
func (c *Cron) AddFunc(spec string, cmd func()) (EntryID, error) {
  return c.AddJob(spec, FuncJob(cmd))
}
```

## 线程安全

`cron`会创建一个新的 goroutine 来执行触发回调。如果这些回调需要并发访问一些资源、数据，我们需要显式地做同步。

## 自定义时间格式

`cron`支持灵活的时间格式，如果默认的格式不能满足要求，我们可以自己定义时间格式。时间规则字符串需要`cron.Parser`对象来解析。我们先来看看默认的解析器是如何工作的。

首先定义各个域：

```golang
// parser.go
const (
  Second         ParseOption = 1 << iota
  SecondOptional                        
  Minute                                
  Hour                                  
  Dom                                   
  Month                                 
  Dow                                   
  DowOptional                           
  Descriptor                            
)
```

除了`Minute/Hour/Dom(Day of month)/Month/Dow(Day of week)`外，还可以支持`Second`。相对顺序都是固定的：

```golang
// parser.go
var places = []ParseOption{
  Second,
  Minute,
  Hour,
  Dom,
  Month,
  Dow,
}

var defaults = []string{
  "0",
  "0",
  "0",
  "*",
  "*",
  "*",
}
```

默认的时间格式使用 5 个域。

我们可以调用`cron.NewParser()`创建自己的`Parser`对象，以位格式传入使用哪些域，例如下面的`Parser`使用 6 个域，支持`Second`（秒）：

```golang
parser := cron.NewParser(
  cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor,
)
```

调用`cron.WithParser(parser)`创建一个选项传入构造函数`cron.New()`，使用时就可以指定秒了：

```golang
c := cron.New(cron.WithParser(parser))
c.AddFunc("1 * * * * *", func () {
  fmt.Println("every 1 second")
})
c.Start()
```

这里时间格式必须使用 6 个域，顺序与上面的`const`定义一致。

因为上面的时间格式太常见了，`cron`定义了一个便捷的函数：

```golang
// option.go
func WithSeconds() Option {
  return WithParser(NewParser(
    Second | Minute | Hour | Dom | Month | Dow | Descriptor,
  ))
}
```

注意`Descriptor`表示对`@every/@hour`等的支持。有了`WithSeconds()`，我们不用手动创建`Parser`对象了：

```golang
c := cron.New(cron.WithSeconds())
```

## 选项

`cron`对象创建使用了选项模式，我们前面已经介绍了 3 个选项：

- `WithLocation`：指定时区；
- `WithParser`：使用自定义的解析器；
- `WithSeconds`：让时间格式支持秒，实际上内部调用了`WithParser`。

`cron`还提供了另外两种选项：

- `WithLogger`：自定义`Logger`；
- `WithChain`：Job 包装器。

### `WithLogger`

`WithLogger`可以设置`cron`内部使用我们自定义的`Logger`：

```golang
func main() {
  c := cron.New(
    cron.WithLogger(
      cron.VerbosePrintfLogger(log.New(os.Stdout, "cron: ", log.LstdFlags))))
  c.AddFunc("@every 1s", func() {
    fmt.Println("hello world")
  })
  c.Start()

  time.Sleep(5 * time.Second)
}
```

上面调用`cron.VerbosPrintfLogger()`包装`log.Logger`，这个`logger`会详细记录`cron`内部的调度过程：

```cmd
$ go run main.go
cron: 2020/06/26 07:09:14 start
cron: 2020/06/26 07:09:14 schedule, now=2020-06-26T07:09:14+08:00, entry=1, next=2020-06-26T07:09:15+08:00
cron: 2020/06/26 07:09:15 wake, now=2020-06-26T07:09:15+08:00
cron: 2020/06/26 07:09:15 run, now=2020-06-26T07:09:15+08:00, entry=1, next=2020-06-26T07:09:16+08:00
hello world
cron: 2020/06/26 07:09:16 wake, now=2020-06-26T07:09:16+08:00
cron: 2020/06/26 07:09:16 run, now=2020-06-26T07:09:16+08:00, entry=1, next=2020-06-26T07:09:17+08:00
hello world
cron: 2020/06/26 07:09:17 wake, now=2020-06-26T07:09:17+08:00
cron: 2020/06/26 07:09:17 run, now=2020-06-26T07:09:17+08:00, entry=1, next=2020-06-26T07:09:18+08:00
hello world
cron: 2020/06/26 07:09:18 wake, now=2020-06-26T07:09:18+08:00
hello world
cron: 2020/06/26 07:09:18 run, now=2020-06-26T07:09:18+08:00, entry=1, next=2020-06-26T07:09:19+08:00
cron: 2020/06/26 07:09:19 wake, now=2020-06-26T07:09:19+08:00
hello world
cron: 2020/06/26 07:09:19 run, now=2020-06-26T07:09:19+08:00, entry=1, next=2020-06-26T07:09:20+08:0
```

我们看看默认的`Logger`是什么样的：

```golang
// logger.go
var DefaultLogger Logger = PrintfLogger(log.New(os.Stdout, "cron: ", log.LstdFlags))

func PrintfLogger(l interface{ Printf(string, ...interface{}) }) Logger {
  return printfLogger{l, false}
}

func VerbosePrintfLogger(l interface{ Printf(string, ...interface{}) }) Logger {
  return printfLogger{l, true}
}

type printfLogger struct {
  logger  interface{ Printf(string, ...interface{}) }
  logInfo bool
}
```

### `WithChain`

Job 包装器可以在执行实际的`Job`前后添加一些逻辑：

- 捕获`panic`；
- 如果`Job`上次运行还未结束，推迟本次执行;
- 如果`Job`上次运行还未介绍，跳过本次执行；
- 记录每个`Job`的执行情况。

我们可以将`Chain`类比为 Web 处理器的中间件。实际上就是在`Job`的执行逻辑外在封装一层逻辑。我们的封装逻辑需要写成一个函数，传入一个`Job`类型，返回封装后的`Job`。`cron`为这种函数定义了一个类型`JobWrapper`：

```golang
// chain.go
type JobWrapper func(Job) Job
```

然后使用一个`Chain`对象将这些`JobWrapper`组合到一起：

```golang
type Chain struct {
  wrappers []JobWrapper
}

func NewChain(c ...JobWrapper) Chain {
  return Chain{c}
}
```

调用`Chain`对象的`Then(job)`方法应用这些`JobWrapper`，返回最终的`Job：

```golang
func (c Chain) Then(j Job) Job {
  for i := range c.wrappers {
    j = c.wrappers[len(c.wrappers)-i-1](j)
  }
  return j
}
```

注意应用`JobWrapper`的顺序。

### 内置`JobWrapper`

`cron`内置了 3 个用得比较多的`JobWrapper`：

- `Recover`：捕获内部`Job`产生的 panic；
- `DelayIfStillRunning`：触发时，如果上一次任务还未执行完成（耗时太长），则等待上一次任务完成之后再执行；
- `SkipIfStillRunning`：触发时，如果上一次任务还未完成，则跳过此次执行。

下面分别介绍。

#### `Recover`

先看看如何使用：

```golang
type panicJob struct {
  count int
}

func (p *panicJob) Run() {
  p.count++
  if p.count == 1 {
    panic("oooooooooooooops!!!")
  }

  fmt.Println("hello world")
}

func main() {
  c := cron.New()
  c.AddJob("@every 1s", cron.NewChain(cron.Recover(cron.DefaultLogger)).Then(&panicJob{}))
  c.Start()

  time.Sleep(5 * time.Second)
}
```

`panicJob`在第一次触发时，触发了`panic`。因为有`cron.Recover()`保护，后续任务还能执行：

```cmd
go run main.go 
cron: 2020/06/27 14:02:00 panic, error=oooooooooooooops!!!, stack=...
goroutine 18 [running]:
github.com/robfig/cron/v3.Recover.func1.1.1(0x514ee0, 0xc0000044a0)
        D:/code/golang/pkg/mod/github.com/robfig/cron/v3@v3.0.1/chain.go:45 +0xbc
panic(0x4cf380, 0x513280)
        C:/Go/src/runtime/panic.go:969 +0x174
main.(*panicJob).Run(0xc0000140e8)
        D:/code/golang/src/github.com/darjun/go-daily-lib/cron/recover/main.go:17 +0xba
github.com/robfig/cron/v3.Recover.func1.1()
        D:/code/golang/pkg/mod/github.com/robfig/cron/v3@v3.0.1/chain.go:53 +0x6f
github.com/robfig/cron/v3.FuncJob.Run(0xc000070390)
        D:/code/golang/pkg/mod/github.com/robfig/cron/v3@v3.0.1/cron.go:136 +0x2c
github.com/robfig/cron/v3.(*Cron).startJob.func1(0xc00005c0a0, 0x514d20, 0xc000070390)
        D:/code/golang/pkg/mod/github.com/robfig/cron/v3@v3.0.1/cron.go:312 +0x68
created by github.com/robfig/cron/v3.(*Cron).startJob
        D:/code/golang/pkg/mod/github.com/robfig/cron/v3@v3.0.1/cron.go:310 +0x7a
hello world
hello world
hello world
hello world
```

我们看看`cron.Recover()`的实现，很简单：

```golang
// cron.go
func Recover(logger Logger) JobWrapper {
  return func(j Job) Job {
    return FuncJob(func() {
      defer func() {
        if r := recover(); r != nil {
          const size = 64 << 10
          buf := make([]byte, size)
          buf = buf[:runtime.Stack(buf, false)]
          err, ok := r.(error)
          if !ok {
            err = fmt.Errorf("%v", r)
          }
          logger.Error(err, "panic", "stack", "...\n"+string(buf))
        }
      }()
      j.Run()
    })
  }
}
```

就是在执行内层的`Job`逻辑前，添加`recover()`调用。如果`Job.Run()`执行过程中有`panic`。这里的`recover()`会捕获到，输出调用堆栈。

#### `DelayIfStillRunning`

还是先看如何使用：

```golang
type delayJob struct {
  count int
}

func (d *delayJob) Run() {
  time.Sleep(2 * time.Second)
  d.count++
  log.Printf("%d: hello world\n", d.count)
}

func main() {
  c := cron.New()
  c.AddJob("@every 1s", cron.NewChain(cron.DelayIfStillRunning(cron.DefaultLogger)).Then(&delayJob{}))
  c.Start()

  time.Sleep(10 * time.Second)
}
```

上面我们在`Run()`中增加了一个 2s 的延迟，输出中间隔变为 2s，而不是定时的 1s：

```golang
$ go run main.go 
2020/06/27 14:11:16 1: hello world
2020/06/27 14:11:18 2: hello world
2020/06/27 14:11:20 3: hello world
2020/06/27 14:11:22 4: hello world
```

看看源码：

```golang
// chain.go
func DelayIfStillRunning(logger Logger) JobWrapper {
  return func(j Job) Job {
    var mu sync.Mutex
    return FuncJob(func() {
      start := time.Now()
      mu.Lock()
      defer mu.Unlock()
      if dur := time.Since(start); dur > time.Minute {
        logger.Info("delay", "duration", dur)
      }
      j.Run()
    })
  }
}
```

首先定义一个该任务共用的互斥锁`sync.Mutex`，每次执行任务前获取锁，执行结束之后释放锁。所以在上一个任务结束前，下一个任务获取锁是无法成功的，从而保证的任务的串行执行。

#### `SkipIfStillRunning`

还是先看看如何使用：

```golang
type skipJob struct {
  count int32
}

func (d *skipJob) Run() {
  atomic.AddInt32(&d.count, 1)
  log.Printf("%d: hello world\n", d.count)
  if atomic.LoadInt32(&d.count) == 1 {
    time.Sleep(2 * time.Second)
  }
}

func main() {
  c := cron.New()
  c.AddJob("@every 1s", cron.NewChain(cron.SkipIfStillRunning(cron.DefaultLogger)).Then(&skipJob{}))
  c.Start()

  time.Sleep(10 * time.Second)
}
```

输出：

```golang
$ go run main.go
2020/06/27 14:22:07 1: hello world
2020/06/27 14:22:10 2: hello world
2020/06/27 14:22:11 3: hello world
2020/06/27 14:22:12 4: hello world
2020/06/27 14:22:13 5: hello world
2020/06/27 14:22:14 6: hello world
2020/06/27 14:22:15 7: hello world
2020/06/27 14:22:16 8: hello world
```

注意观察时间，第一个与第二个输出之间相差 3s，因为跳过了两次执行。

注意`DelayIfStillRunning`与`SkipIfStillRunning`是有本质上的区别的，前者`DelayIfStillRunning`只要时间足够长，所有的任务都会按部就班地完成，只是可能前一个任务耗时过长，导致后一个任务的执行时间推迟了一点。`SkipIfStillRunning`会跳过一些执行。

看看源码：

```golang
func SkipIfStillRunning(logger Logger) JobWrapper {
  return func(j Job) Job {
    var ch = make(chan struct{}, 1)
    ch <- struct{}{}
    return FuncJob(func() {
      select {
      case v := <-ch:
        j.Run()
        ch <- v
      default:
        logger.Info("skip")
      }
    })
  }
}
```

定义一个该任务共用的缓存大小为 1 的通道`chan struct{}`。执行任务时，从通道中取值，如果成功，执行，否则跳过。执行完成之后再向通道中发送一个值，确保下一个任务能执行。初始发送一个值到通道中，保证第一个任务的执行。

## 总结

`cron`实现比较小巧，且优雅，代码行数也不多，非常值得一看！

大家如果发现好玩、好用的 Go 语言库，欢迎到 [golang-guide](https://github.com/mao888/golang-guide/tree/main/golang/go-study/%E5%B8%B8%E7%94%A8%E7%BB%84%E4%BB%B6%E5%BA%93) 上提交 issue

## 参考

1. cron GitHub：https://github.com/robfig/cron
2. Cron Documentation here: https://godoc.org/github.com/robfig/cron
3. Go 每日一库之 carbon：https://darjun.github.io/2020/02/14/godailylib/carbon/
4. Go 每日一库之 gron：https://darjun.github.io/2020/04/20/godailylib/gron/
5. Go 每日一库 GitHub：https://github.com/darjun/go-daily-lib