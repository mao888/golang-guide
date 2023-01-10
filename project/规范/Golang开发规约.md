|      | **分类** | **描述**                                                     | **重要性** | **说明**                                                     |
| ---- | -------- | ------------------------------------------------------------ | ---------- | ------------------------------------------------------------ |
|      | **分类** | **描述**                                                     | **重要性** | **说明**                                                     |
|      | 工具     | 必须在工程中使用go fmt、goimports、gometalinter或者golangci-lint工具对代码试试检查 | 高         |                                                              |
|      | 命名     | 目录名必须为全小写单词，允许加中划线‘-’组合方式，但是头尾不能为中划线。 | 高         |                                                              |
|      | 命名     | 包名必须全部为小写单词，无下划线，越短越好。尽量不要与标准库重名。 | 高         |                                                              |
|      | 命名     | 文件名必须为小写单词，允许加下划线‘_’组合方式，但是头尾不能为下划线。 | 高         |                                                              |
|      | 命名     | 常量&枚举名采用大小写混排的驼峰模式（Golang官方要求），不允许出现下划线。 | 高         |                                                              |
|      | 命名     | 按照功能来区分，而不是将所有类型都分在一组，并建议将公共常量置于私有常量之前。 | 高         |                                                              |
|      | 命名     | 如果是枚举类型的常量，需要先创建相应类型。                   | 高         |                                                              |
|      | 命名     | 变量名称一般遵循驼峰法，并且不允许出现下划线，当遇到特有名词时，需要遵循以下规则：如果变量为私有，且特有名词为首个单词，则使用小写，如：apiClient其它情况都应当使用该名词原有的写法，如 APIClient、repoID、UserID | 高         |                                                              |
|      | 导包     | 导包需分组，至少分为三组，第一组为系统包，第二组为自身应用包，第三组为私有仓库，第四组为第三方包                                                                          import (<br/>    "context"<br/>    "strconv"<br/>    "fmt"<br/>    "net/http"<br/>    "time"<br/> <br/>    "filebeatx-web/interval/app/api/consts"<br/>    "filebeatx-web/interval/app/api/handler"<br/>    "filebeatx-web/interval/app/api/middleware"<br/>    "filebeatx-web/interval/app/api/repository"<br/>    "filebeatx-web/interval/app/api/service"<br/> <br/>    "[gitlab.ftsview.com/OpenPlatform/open-go-lib/ftlog](http://gitlab.ftsview.com/OpenPlatform/open-go-lib/ftlog)"<br/> <br/>    "[github.com/gin-contrib/expvar](http://github.com/gin-contrib/expvar)"<br/>    "[github.com/gin-contrib/pprof](http://github.com/gin-contrib/pprof)"<br/>    "[github.com/gin-gonic/gin](http://github.com/gin-gonic/gin)"<br/>    "[github.com/spf13/viper](http://github.com/spf13/viper)"<br/>    dicentity "jryg-dictionary/model/entity"<br/>    "[github.com/spf13/viper](http://github.com/spf13/viper)"<br/>    "[go.uber.org/zap](http://go.uber.org/zap)"<br/>) | 高         |                                                              |
|      | 变量     | 一般在使用前初始化即可，当一个逻辑块需要较多变量且不易于理解时，需使用 var() 提前初始化。 | 高         |                                                              |
|      | 函数     | 当入参需要 context 参数时，第一个参数应为 ctx context.Context。 当出参需要 error 参数时，最后一个参数应为 error 类型的返回值。 | 高         |                                                              |
|      | 语法     | 当明确expr为bool类型时，禁止使用==或!=与true/false比较，应该使用expr或!expr 判断某个整数表达式expr是否为零时，禁止使用!expr，应该使用expr == 0 | 高         |                                                              |
|      | 语法     | embedding只用于"is a"的语义下，而不用于"has a"的语义下 一个定义内，多于一个的embedding尽量少用。 | 高         |                                                              |
|      | 语法     | 除非出现不可恢复的程序错误，不要使用panic，用多返回值和error。 | 高         |                                                              |
|      | 语法     | 如果临界区内的逻辑较复杂、无法完全避免panic的发生，则要求适用defer来调用Unlock，即使在临界区过程中发生了panic，也会在函数退出时调用Unlock释放锁 | 高         |                                                              |
|      | 语法     | 除非特殊原因，不建议使用unsafe                               | 高         |                                                              |
|      | 语法     | 代码中禁止使用魔鬼数字。                                     | 高         |                                                              |
|      | 语法     | 如果你利用 iota 来使用自定义的整数枚举类型，务必要为其添加 String() 方法。 | 高         |                                                              |
|      | 语法     | 对于主要功能模块抽象模块接口，通过interface提供对外功能。    | 高         |                                                              |
|      | 语法     | 一个文件只定义一个init函数，一个包内的如果存在多个init函数，不能有任何的依赖关系。 | 高         | 如果包内有多个init，每个init的执行顺序是不确定的。           |
|      | 语法     | defer会消耗更多的系统资源，不建议用于频繁调用的方法中，避免在for循环中使用defer。 | 高         |                                                              |
|      | 语法     | 确保每个goroutine都能退出。                                  | 高         | 启动goroutine就相当于启动了一个线程，如果不设置线程退出的条件就相当于这个线程失去了控制，占用的资源将无法回收，导致内存泄露 |
|      | 语法     | 确保对channel是否需要关闭检查，已防止死循环                  | 高         | for { select { case <-cc: //【错误】当channel cc被关闭后如果不做检查则造成死循环 fmt.Println("continue") case <-time.After(5 * time.Second): fmt.Println("timeout") } } |
|      | 语法     | 禁止局部变量与全局变量同名（禁止变量同名）。                 | 高         |                                                              |
|      | 内存优化 | 创建结构体时尽量按照占用字符顺序定义。（防止内存对齐产生的内存碎片）int8：1字节int16：2字节int32：4字节 int：8字节（64位系统） 4字节（32位系统）int64：8字节string：16字节map：8字节slice：24字节array：16字节指针：8字节（64位系统） 4字节（32位系统） | 中         |                                                              |
|      | 内存优化 | 尽量不要在for{ switch case } 中使用time.After来处理超时等。应该使用time.NewTimer(time.Second)反例：for {  switch xxx {  case x:  //正常  case <- time.After(time.Second):  //异常  } }正例：timer := time.NewTimer(time.Second)for {timer.Reset(time.Second)switch xxx { case x:  //正常  case <- timer.C:  //异常  } } | 中         |                                                              |
|      | 性能优化 | 当为小结构体并且不需要对其进行修改时，使用值传递。在测试中验证当结构体大小在3000字节左右时传值和传指针的处理能力差不多。Sizeof : 3072 goos: windows goarch: amd64 pkg: micro-me/testing BenchmarkByPointer-12 30000000 59.0 ns/op BenchmarkByValue-12 30000000 53.6 ns/op PASS | 低         |                                                              |
|      | 性能优化 | 注意select的default使用，select添加default之后其中的case不会阻塞。 | 高         |                                                              |
|      | 性能优化 | 减少[]byte和string之间的转换，尽量使用[]byte来处理字符。     | 低         |                                                              |
|      | 性能优化 | make申请slice/map时，根据预估大小来申请合适内存。            | 低         |                                                              |
|      | 性能优化 | 字符串拼接优先考虑bytes.Buffer。                             | 低         |                                                              |
|      | 可读性   | 代码注释率要保持再15%左右                                    | 中         |                                                              |
|      |          |                                                              |            |                                                              |
|      |          |                                                              |            |                                                              |
|      |          |                                                              |            |                                                              |