# protobuf中使用oneof、WrapValue和FieldMask

本文介绍了在Go语言中如何使用`oneof`字段以及如何通过使用`google/protobuf/wrappers.proto`中定义的类型区分默认值和没有传值；最后演示了Go语言中借助[fieldmask-utils](https://github.com/mennanov/fieldmask-utils)库使用`google/protobuf/field_mask.proto`实现部分更新的方法。

## oneof

如果你有一条包含多个字段的消息，并且最多同时设置其中一个字段，那么你可以通过使用`oneof`来实现并节省内存。

`oneof`字段类似于常规字段，只不过`oneof`中的所有字段共享内存，而且最多可以同时设置一个字段。设置其中的任何成员都会自动清除所有其他成员。

可以在`oneof`中添加除了map字段和repeated字段外的任何类型的字段。

### protobuf 定义

假设我的博客系统支持为读者朋友们发送博客更新的通知信息，系统支持通过邮件和短信两个方式发送通知。但每一次只允许使用一种方式发送通知。

在这个场景下我们就可以使用`oneof`字段来定义通知的方式——`notice_way`。

```protobuf
// 通知读者的消息
message NoticeReaderRequest{
    string msg = 1;
    oneof notice_way{
        string email = 2;
        string phone = 3;
    }
}
```

### client端代码

Go语言创建`oneof`字段的client端示例代码。

```go
// 使用邮件通知的请求消息
noticeReq := proto.NoticeReaderRequest{
	Msg: "李文周的博客更新啦~",
	NoticeWay: &proto.NoticeReaderRequest_Email{
		Email: "123@xx.com",
	},
}
// 使用短信通知的请求消息
noticeReq2 := proto.NoticeReaderRequest{
	Msg: "李文周的博客更新啦~",
	NoticeWay: &proto.NoticeReaderRequest_Phone{
		Phone: "123456789",
	},
}
```

### server端代码

Go语言操作`oneof`字段的server端示例代码。下面的代码中使用`switch case`的方式，根据请求消息中的通知类型选择执行不同的业务逻辑。

```go
// ... liwenzhou.com ...

// 根据`NoticeWay`的不同而执行不同的操作
switch v := noticeReq.NoticeWay.(type) {
case *proto.NoticeReaderRequest_Email:
	noticeWithEmail(v)
case *proto.NoticeReaderRequest_Phone:
	noticeWithPhone(v)
}

// ... liwenzhou.com ...

// 发送通知相关的功能函数
func noticeWithEmail(in *proto.NoticeReaderRequest_Email) {
	fmt.Printf("notice reader by email:%v\n", in.Email)
}

func noticeWithPhone(in *proto.NoticeReaderRequest_Phone) {
	fmt.Printf("notice reader by phone:%v\n", in.Phone)
}
```

## WrapValue

protobuf v3在删除`required`的同时把`optional`也一起删除了（[v3.15.0又加回来了](https://github.com/protocolbuffers/protobuf/blob/v3.15.0/docs/field_presence.md#how-to-enable-explicit-presence-in-proto3)），这使得我们没办法轻易判断某些字段究竟是未赋值还是其被赋值为零值。

例如，当我们有如下消息定义时，我们拿到一个book消息，当`book.Price = 0`时我们没办法区分`book.Price`字段是未赋值还是被赋值为0。

```protobuf
message Book {
    string title = 1;
    string author = 2;
    int64 price = 3;
}
```

### protobuf 定义

类似这种场景推荐使用`google/protobuf/wrappers.proto`中定义的WrapValue，本质上就是使用自定义message代替基本类型。

```protobuf
// google/protobuf/wrappers.proto

// ...

// Wrapper message for `float`.
//
// The JSON representation for `FloatValue` is JSON number.
message FloatValue {
  // The float value.
  float value = 1;
}

// Wrapper message for `int64`.
//
// The JSON representation for `Int64Value` is JSON string.
message Int64Value {
  // The int64 value.
  int64 value = 1;
}

// ... 
```

在这个示例中，我们就可以使用`Int64Value`代替`int64`，修改后的`protobuf`文件如下。

```protobuf
message Book {
    string title = 1;
    string author = 2;
    google.protobuf.Int64Value price = 3;
}
```

### client端代码

使用了`wrappers.proto`中定义的包装类型后，我们在赋值的时候就需要额外包一层。

```go
import "google.golang.org/protobuf/types/known/wrapperspb"

book := proto.Book{
	Title: "《跟七米学Go语言》",
	Price: &wrapperspb.Int64Value{Value: 9900},
}
```

### server端代码

WrapValue本质上类似于标准库sql中定义的`sql.NullInt64`、`sql.NullString`，即将基本数据类型包装为一个结构体类型。在使用时通过判断某个字段是否为nil（空指针）来区分该字段是否被赋值。

```go
if book.GetPrice() == nil {  // price没赋值
	fmt.Println("book with no price")
} else {
	fmt.Printf("book with price:%v\n", book.GetPrice().GetValue())
}
```

## v3.15.0+使用optional

Protobuf **v3.15.0** 版本之后又支持使用`optional`显式指定字段为可选。

下面的示例中，我们使用`optional`标识`price`为可选字段。

```protobuf
message Book {
    string title = 1;
    string author = 2;
    //google.protobuf.Int64Value price = 3;
    optional int64 price = 3;  // 使用optional
}
```

修改了`proto`文件后，重新编译。

### client端代码

现在`price`字段就是`*int64`类型了，我们需要使用`google.golang.org/protobuf/proto`包提供的系列函数完成赋值操作。

```go
import "google.golang.org/protobuf/proto"

book := proto.Book{
	Title: "《跟七米学Go语言》",
	Price: proto.Int64(9900),
}
```

### server端代码

如果需要判断`price`字段是否赋值，可以判断是否为`nil`。

```go
if book.Price == nil {  // price没赋值
	fmt.Println("book with no price")
} else {
	fmt.Printf("book with price:%v\n", book.GetPrice())
}
```

## FieldMask

假设现在需要实现一个更新书籍信息接口，我们可能会定义如下更新书籍的消息。

```protobuf
message UpdateBookRequest {
    // 操作人 
    string op = 1;
    // 要更新的书籍信息
    Book book = 2;
}
```

但是如果我们的`Book`中定义有很多很多字段时，我们不太可能每次请求都去全量更新`Book`的每个字段，因为通常每次操作只会更新1到2个字段。

那么我们该如何确定每次更新操作涉及到了哪些具体字段呢？

答案是使用`google/protobuf/field_mask.proto`，它能够记录在一次更新请求中涉及到的具体字段路径。

为了实现一个支持部分更新的接口，我们把`UpdateBookRequest`消息修改如下。

```protobuf
message UpdateBookRequest {
    // 操作人 
    string op = 1;
    // 要更新的书籍信息
    Book book = 2;

    // 要更新的字段
    google.protobuf.FieldMask update_mask = 3;
}
```

### client端代码

我们通过`paths`记录本次更新的字段路径，如果是嵌套的消息类型则通过`x.y`的方式标识。

```go
import "google.golang.org/protobuf/types/known/fieldmaskpb"

paths := []string{"title", "read"} // 记录更新的字段路径
updateReq := proto.UpdateBookRequest{
	Book: &proto.Book{
		Title: "《跟七米学Go语言》",
		Read:  true,
	},
	UpdateMask: &fieldmaskpb.FieldMask{Paths: paths},
}
```

### server端代码

在收到更新消息后，我们需要根据`UpdateMask`字段中记录的更新路径去读取更新数据。这里借助第三方库[github.com/mennanov/fieldmask-utils](https://github.com/mennanov/fieldmask-utils)实现。

```go
import "github.com/golang/protobuf/protoc-gen-go/generator"
import fieldmask_utils "github.com/mennanov/fieldmask-utils"

mask, _ := fieldmask_utils.MaskFromProtoFieldMask(updateReq.UpdateMask, generator.CamelCase)
var bookDst = make(map[string]interface{})
// 将数据读取到map[string]interface{}
// fieldmask-utils支持读取到结构体等，更多用法可查看文档。
fieldmask_utils.StructToMap(mask, updateReq.Book, bookDst)
// do update with bookDst
fmt.Printf("bookDst:%#v\n", bookDst)
```

**2022-11-20更新**：由于`github.com/golang/protobuf/protoc-gen-go/generator`包已弃用，而`MaskFromProtoFieldMask`函数（签名如下）

```go
func MaskFromProtoFieldMask(fm *field_mask.FieldMask, naming func(string) string) (Mask, error)
```

接收的`naming`参数本质上是一个将字段掩码字段名映射到 Go 结构中使用的名称的函数，它必须根据你的实际需求实现。

例如在我们这个示例中，还可以使用`github.com/iancoleman/strcase`包提供的`ToCamel`方法：

```go
import "github.com/iancoleman/strcase"
import fieldmask_utils "github.com/mennanov/fieldmask-utils"

mask, _ := fieldmask_utils.MaskFromProtoFieldMask(updateReq.UpdateMask, strcase.ToCamel)
var bookDst = make(map[string]interface{})
// 将数据读取到map[string]interface{}
// fieldmask-utils支持读取到结构体等，更多用法可查看文档。
fieldmask_utils.StructToMap(mask, updateReq.Book, bookDst)
// do update with bookDst
fmt.Printf("bookDst:%#v\n", bookDst)
```

参考资料：

- https://cloud.google.com/apis/design/standard_methods
- https://github.com/mennanov/fieldmask-utils
- https://github.com/iancoleman/strcase