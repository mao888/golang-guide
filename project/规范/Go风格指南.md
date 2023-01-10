# Go风格指南

精读：

- https://golang.org/doc/effective_go.html
- https://github.com/golang/go/wiki/CodeReviewComments

以下是一些额外的原则，以及重申重要问题。

### 何时panic？

Effective Go提到不要panic。申请方的代码允许panic，但必须严格依照如下准则：

- panic可以发生在程序初始化时间。
- 初始化完成以后，panic不能让服务器崩溃。可以在HTTP Handler内的逻辑panic。任何这些panic将会以500返回结果。
- panic只能是由于两种事情发生：

1. 1. 错误的原因是软件工程师的判断错误，比如程序员相信某个struct可以marshal成json，不可能出错，那么万一出错，可以panic。
   2. 错误是由于核心服务出错，比如从couchdb或者elasticsearch查询结果，服务器返回异常结果，这个时候可以panic。

- panic不能由于以下原因：

1. 1. 数据录入不正确。
   2. 最终用户提供的参数不正确。
   3. 任何其他不是程序员直接引发的错误。

- 总体而言，导致panic的错误必须是软件工程师或者运维工程师可以修复的错误。
- 如果你不确定，不要panic。

### 空slice声明

当你要创建一个slice，并不知道它的长度的话使用如下风格：

```go
var someSlice []MyStruct
for ... {
    someSlice = append(someSlice, ...)
}
```



不要使用someSlice := []MyStruct{}或者someSlice := make(...)（除非你能预判大小）。

这样的好处是如果这个slice最后是0长度（没有任何append），我们省了一次内存分配。

### 函数参数太多的时候，考虑写Params struct。

例如

```go
func GetMajorList(province, types, atProvince, fieldOfStudyKey, batch, slug string, year, offset, limit, minimun, maximum int) { ... }
```


这个函数接受很多参数，而且类型都一样。这样调用的时候，参数的位置很容易搞错，也不容易读懂。
可以改成

```go
type GetMajorListParam struct {     
    Province string     
    Types string     
    AtProvince string     
    FieldOfStudyKey string,     
    Batch string     
    Slug string     
    Year int     
    Offset int     
    ... } 
func GetMajorList(p GetMajorListParam) { ... } 
```

### 不要吃掉错误。

如果一个函数返回有错误，你必须选择其中一种：

1. 把错误作为结果从函数返回。
2. panic，见上文讨论何时panic。
3. 如果是数据录入错误，可以考虑调用sentry.DataError。
4. 如果是后台go rountine，可以调用sentry.Error。
5. 在handler中, return lu.Error(err)
6. 如果错误是预计得到的，比如预计到某个条码可能在数据库不存在，正确检查以后，进行相应处理。

以下处理方法就是吃掉错误：

1. 完全不检查错误。
2. fmt.Print(err)，然后不管。
3. glog.Error(err)，然后不管。
4. tr.Print，然后不管。
5. 其他这里没提到的，同时上面“你必须选择其中一种”也没提到的。

### 如何log

[glog.Info](http://glog.info/), glog.Error, glog.Warning，log.Print, fmt.Print：
调试的时候可以随便用，但是提交代码审查前，要把代码删掉（不要注释掉）。
如果一个程序启动期间只会执行少数几次，可以用glog记录一些事件。不要用glog
在handler里面记录东西。如果要记录，可以用trace记录。
glov.V(1)：如果经常需要调试的，又想提交到代码里，这个可以用。开发环境会显示V(1)记录的信息。
另外，可以使用trace.Debugf来log一些调试信息，这会记录到trace，同时记录到V(1)，这样可以在开发环境看log，生产环境看trace。

### 启动go routine（可能后期加在go基础库中）

不要直接用go xxx启动goroutine。而是使用[concurrent.Go](https://ma.applysquare.net/eng/go/blob/master/pkg/base/concurrent/go.go).
原因是直接用go xxx会有如下问题：

1. 如果没有recover，该go routine内的panic会导致整个服务器退出。
2. 很容易忘记go routine外的trace不能传入go routine内，导致错误。

concurrent.Go帮助正确的进行recover，以及建立trace。

### 如何使用context.Context

从07/16/2017开始trace.T被弃用。取而代之的是context.Context。trace的功能还有保留，但是在base/trace这个包内部被实现，
外部调用全部使用context.Context。

更变如下:

1. 之前tr trace.T，现在ctx context.Context。
2. 之前当作参数传入的tr，现在传入ctx。
3. 之前tr.Printf(msg string, args...), 现在trace.Printf(ctx, msg string, args...)
4. 之前tr.SetError(),现在trace.MarkFailed(ctx)
5. 之前trace.EventLog被弃用，现在使用通用trace。

需要注意的：

1. 命名: context.Context 需要被命名为ctx。
2. 参数: ctx context.Context 需要是函数的第一个参数。（如果不是，lint也会提醒你改成第一个）
3. 和之前的tr一样ctx不能在goroutine之间传递。单独的goroutine必须建立单独的context，具体如何创建参见"如何建立context"
4. 除少数特例之外，context需要作为参数在函数间传递，不要储存在struct内部。

如何得到context:

1. 如果在lu handler里面，直接添加ctx context.Context参数。
2. 如果在graphql resolve里面，调用p.(*gql.Params).Ctx或者p.(graphql.ResolveParams).Context。之前的buddy.Tr被弃用。

如何建立context:

1. 使用concurrent.Go，它会建立带trace的context。直接调用它建立的context.
2. 建立不带trace的context，使用context.Background()。这个相当于trace.Noop
3. 手动建立带trace的context, 使用trace.WithTrace(context.Background(), name string)。name参数的格式是family/method，如果不用/分割，family为name, method为"run"。
   这个相当于之前的trace.New(family, method string)。
4. 如果caller函数需要context, 而call site还没有context传入，可以用context.TODO()。这个相当于之前的trace.TODO。记得之后重构函数将context正确传入。

### 何时使用jsonutil.Merge() （可能后期加在go基础库中）

需要注意的：
大多数情况不推荐使用merge方法，当你使用merge方式时，对相应字段并没有验证，这样会导致数据的一系列问题。

1. 当graphql传入的struct字段较少时，可不使用merge，最优做法应为将model各字段值取出核对后将数据赋值保存
2. 当字段较多时，应先核对在使用merge方法，或者前端传入完整的struct，核对后保存
3. 在dynamicconfig/spec.go中有使用Validate相关的代码，可对Schema进行验证后进行merge

### 使用ES搜索时须知

1. 创建indexer时应对dynamic设置相应值，如下：
   { "mappings": { "my_type": { "dynamic": "false", "properties": { "title": { "type": "string"}, } } } }
2. 当 Elasticsearch 遇到文档中以前未遇到的字段，它用 dynamic mapping 来确定字段的数据类型并自动把新的字段添加到类型映射，
   有时这是想要的行为有时又不希望这样。所以可对dynamic进行配置：
   true
   动态添加新的字段--缺省
   false
   忽略新的字段
   strict
   如果遇到新字段抛出异常
3. 可接受的选项为以上字段，当没有特殊要求时应将其设置成false
4. 同上也应对_all进行设置
   _all字段是把所有其它字段中的值，以空格为分隔符组成一个大字符串，然后被分析和索引，但是不存储。
   也就是说它能被查询，但不能被取回显示。_all字段在查询时占用更多的CPU和占用更多的磁盘空间，
   如果确实不需要它可以完全的关闭它或者基于每字段定制。
   例如：
   { "mappings": { "type_1": { "properties": {...} }, "type_2": { "_all": { "enabled": false }, "properties": {...} } } }
5. type_1中的_all字段是enabled,type_2中的_all字段是disabled,所以一般不特别使用时，将enabled设置成false
6. 参考文档：https://www.elastic.co/guide/en/elasticsearch/reference/current/mapping-all-field.html

### 筛选slice中元素

```go
for _, term := range terms {
if ... {
continue
}
if ...{
continue
}
validTerms = append(validTerms, term)
}
```

### 关于正则表达式的使用

尽量不使用正则表达式，如果使用的话，可以声明全局变量var reg = regexp.MustCompile(具体内容)（相当于一个常量），提高代码运行速度

### 关于字符串的拼接

如果出现大量字符串的拼接，应使用bytes.Buffer
例如a=a+b 可以写成 buf:= bytes.NewBufferString(a)
buf.WriteString(b)
a=buf.String()

### 创建Map

创建一个map时 当多个不同的key会对应相同的value时，建议value的类型使用指针（例如：entity—>*entity）。
这样好处是，相同的Entity只要有个一个，不同的key通过指针指过去，节省内存。