# protocol buffers使用指南

protobuf是一种高效的数据格式，平台无关、语言无关、可扩展，可用于 RPC 系统和持续数据存储系统。

## protocol buffers

### protobuf介绍

[Protobuf](https://developers.google.com/protocol-buffers)全称`Protocol Buffer`，是 Google 公司于2008年开源的一种语言无关、平台无关、可扩展的用于序列化结构化数据——类似于XML，但比XML更小、更快、更简单，它可用于（数据）通信协议、数据存储等。你只需要定义一次你想要的数据结构，然后你就可以使用特殊生成的源代码来轻松地从各种数据流和各种语言中写入和读取你的结构化数据。目前 Protobuf 被广泛用作微服务中的通信协议。

### protobuf语法

[protobuf v3语法官方文档](https://developers.google.com/protocol-buffers/docs/proto3) [protobuf v3中文语法指南](https://www.liwenzhou.com/posts/Go/Protobuf3-language-guide-zh)

## protobuf 编译器指南

### 安装 protobuf

从官方仓库：https://github.com/google/protobuf/releases 下载适合你平台的预编译好的二进制文件（`protoc-<version>-<platform>.zip`）。

- 适用Windows 64位[protoc-3.20.1-win64.zip](https://github.com/protocolbuffers/protobuf/releases/download/v3.20.1/protoc-3.20.1-win64.zip)
- 适用于Mac Intel 64位[protoc-3.20.1-osx-x86_64.zip](https://github.com/protocolbuffers/protobuf/releases/download/v3.20.1/protoc-3.20.1-osx-x86_64.zip)
- 适用于Mac ARM 64位[protoc-3.20.1-osx-aarch_64.zip](https://github.com/protocolbuffers/protobuf/releases/download/v3.20.1/protoc-3.20.1-osx-aarch_64.zip)
- 适用于Linux 64位[protoc-3.20.1-linux-x86_64.zip](https://github.com/protocolbuffers/protobuf/releases/download/v3.20.1/protoc-3.20.1-linux-x86_64.zip)

例如，我使用 Intel 芯片的 Mac 系统则下载 `protoc-3.20.1-osx-x86_64.zip` 文件，解压之后得到如下内容。

![protoc](https://www.liwenzhou.com/images/Go/grpc/protoc.png)

其中：

- bin 目录下的 protoc 是可执行文件。
- include 目录下的是 google 定义的`.proto`文件，我们`import "google/protobuf/timestamp.proto"`就是从此处导入。

我们需要将下载得到的可执行文件`protoc`所在的 bin 目录加到我们电脑的环境变量中。

## 生成Go代码

### 编译器调用

protocol buffer编译器需要一个插件来根据提供的proto文件生成 Go 代码，Go1.16+请使用下面的命令安装插件。

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

这个命令将在 `$GOBIN` 中安装一个 `protocol-gen-go` 的二进制文件。我们需要确保 `$GOBIN` 在你的环境变量中，protocol buffer编译器才能找到它（可以通过`go env`命令查看`$GOPATH`）。

当使用`go_out` 标志调用 protoc 时，protocol buffer编译器将生成 Go 代码。protocol buffer编译器会将生成的Go代码输出到命令行参数`go_out`指定的位置。`go_out`标志的参数是你希望编译器编写 Go 输出的目录。编译器为每个`.proto` 文件输入创建一个源文件。输出文件的名称是通过将`.proto` 扩展名替换为`.pb.go` 而创建的。

生成的`.pb.go`文件放置的目录取决于编译器标志。有以下几种输出模式:

- `paths=import`：输出文件放在以 Go 包的导入路径命名的目录中。例如，`protos/buzz.proto`文件中带有`example.com/project/protos/fizz`的导入路径，则输出的生成文件会保存在`example.com/project/protos/fizz/buzz.pb.go`。如果未指定路径标志，这就是默认输出模式。
- `module=$PREFIX`：输出文件放在以 Go 包的导入路径命名的目录中，但是从输出文件名中删除了指定的目录前缀。例如，输入文件 `pros/buzz.proto`，其导入路径为 `example.com/project/protos/fizz` 并指定`example.com/project`为`module`前缀，结果会产生一个名为 `pros/fizz/buzz.pb.go` 的输出文件。在module路径之外生成任何 Go 包都会导致错误。此模式对于将生成的文件直接输出到 Go 模块非常有用。
- `paths=source_relative`：输出文件与输入文件放在相同的相对目录中。例如，一个`protos/buzz.proto`输入文件会产生一个位于`protos/buzz.pb.go`的输出文件。

在调用`protoc`时，通过传递 `go_opt` 标志来提供特定于 `protocol-gen-go` 的标志位参数。可以传递多个`go_opt`标志位参数。例如，当执行下面的命令时：

```bash
protoc --proto_path=src --go_out=out --go_opt=paths=source_relative foo.proto bar/baz.proto
```

编译器将从 `src` 目录中读取输入文件 `foo.proto` 和 `bar/baz.proto`，并将输出文件 `foo.pb.go` 和 `bar/baz.pb.go` 写入 `out` 目录。如果需要，编译器会自动创建嵌套的输出子目录，但不会创建输出目录本身。

### package

为了生成 Go 代码，必须为每个 `.proto` 文件（包括那些被生成的 `.proto` 文件传递依赖的文件）提供 Go 包的导入路径。有两种方法可以指定 Go 导入路径：

- 通过在 `.proto` 文件中声明它。
- 通过在调用 `protoc` 时在命令行上声明它。

我们建议在 `.proto` 文件中声明它，以便 `.proto` 文件的 Go 包可以与 `.proto` 文件本身集中标识，并简化调用 `protoc` 时传递的标志集。 如果给定 `.proto` 文件的 Go 导入路径由 `.proto` 文件本身和命令行提供，则后者优先于前者。

Go 导入路径是在 `.proto` 文件中指定的，通过声明带有 Go 包的完整导入路径的 `go_package` 选项来创建 proto 文件。用法示例：

```bash
option go_package = "example.com/project/protos/fizz";
```

调用编译器时，可以在命令行上指定 Go 导入路径，方法是传递一个或多个 `M${PROTO_FILE}=${GO_IMPORT_PATH}` 标志位。用法示例：

```bash
protoc --proto_path=src \
  --go_opt=Mprotos/buzz.proto=examples.com/project/protos/fizz \
  --go_opt=Mprotos/bar.proto=examples.com/project/protos/foo \
  protos/buzz.proto protos/bar.proto
```

由于所有 `.proto` 文件到其 Go 导入路径的映射可能非常大，这种指定 Go 导入路径的模式通常由控制整个依赖树的某些构建工具（例如 Bazel）执行。 如果给定的 `.proto` 文件有重复条目，则指定的最后一个条目优先。

对于 `go_package` 选项和 `M` 标志位，它们的值可以包含一个显式的包名称，该名称与导入路径之间用分号分隔。 例如：`“example.com/protos/foo;package_name”`。 不鼓励这种用法，因为默认情况下包名称将以合理的方式从导入路径派生。

导入路径用于确定一个 `.proto` 文件导入另一个 `.proto` 文件时必须生成哪些导入语句。 例如，如果 `a.proto`导入 `b.proto`，则生成的 `a.pb.go` 文件需要导入包含生成的 `b.pb.go` 文件的 Go 包（除非两个文件在同一个包中）。 导入路径也用于构造输出文件名。 有关详细信息，请参阅上面的“编译器调用”部分。

Go 导入路径和 `.proto` 文件中的`package`说明符之间没有关联。 后者仅与 protobuf 命名空间相关，而前者仅与 Go 命名空间相关。 此外，Go 导入路径和 `.proto` 导入路径之间没有关联。

## Go语言使用protoc示例

我们新建一个名为`demo`的项目，并且将项目中定义的`.proto`文件都保存在`proto`目录下。

本文后续的操作命令默认都在`demo`目录下执行。

### 普通编译

下面的示例中我们将定义一个单独的`proto`文件并进行编译。

#### 定义proto

新建一个`price.proto`文件。

```protobuf
// proto/book/price.proto

syntax = "proto3";

package book;

// 声明生成Go代码的导入路径（import path）
option go_package = "github.com/Q1mi/demo/proto/book";

message Price {
    int64 market_price = 1;  // 建议使用下划线的命名方式
    int64 sale_price = 2;
}
```

我们在这个文件中使用`option go_package = "github.com/Q1mi/demo/proto/book"`语句声明了生成的Go代码的导入路径。

项目当前的目录结构如下：

```bash
demo
└── proto
    └── book
        └── price.proto
```

#### 生成代码

假设我们想把最终生成的Go代码还保存在`proto`文件夹中，那么就可以执行下面的命令。

```bash
protoc --proto_path=proto --go_out=proto --go_opt=paths=source_relative book/price.proto
```

其中：

- `--proto_path=proto` 表示从proto目录下读取proto文件。
- `--go_out=proto` 表示生成的Go代码保存的路径。
- `--go_opt=paths=source_relative` 表示输出文件与输入文件放在相同的相对目录中。
- `book/price.proto` 表示在proto目录下的`book/price.proto`文件。

此外，`--proto_path`有一个别名`-I`，上述编译命令也可以这样写。

```bash
protoc -I=proto --go_out=proto --go_opt=paths=source_relative book/price.proto
```

执行上述命令将会在`proto`目录下生成`book/price.pb.go`文件。

```bash
demo
└── proto
    └── book
        ├── price.pb.go
        └── price.proto
```

此处如果不指定`--proto_path`参数那么编译命令可以简写为:

```bash
protoc --go_out=. --go_opt=paths=source_relative proto/book/price.proto
```

上面的命令都是将代码生成到`demo/proto`目录，如果想要将生成的Go代码保存在其他文件夹中（例如`pb`文件夹），那么我们需要先在`demo`目录下创建一个`pb`文件夹。然后在命令行通过`--go_out=pb`指定生成的Go代码保存的路径。完整命令如下：

```bash
protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative book/price.proto
```

执行上面的命令便会在`demo/pb`文件夹下生成Go代码。

```bash
demo
├── pb
│   └── book
│       └── price.pb.go
└── proto
    └── book
        ├── price.pb.go
        └── price.proto
```

### import同目录下protobuf文件

随着业务的复杂度上升，我们可能会定义多个`.proto`源文件，然后根据需要引入其他的protobuf文件。

在这个示例中，我们在`demo/proto/book`目录下新建一个`book.proto`文件，它通过`import "book/price.proto";`语句引用了同目录下的`price.proto`文件。

```protobuf
// demo/proto/book/book.proto

syntax = "proto3";

// 声明protobuf中的包名
package book;

// 声明生成的Go代码的导入路径
option go_package = "github.com/Q1mi/demo/proto/book";

// 引入同目录下的protobuf文件（注意起始位置为proto_path的下层）
import "book/price.proto";

message Book {
    string title = 1;
    Price price = 2;
}
```

编译命令如下：

```bash
protoc --proto_path=proto --go_out=proto --go_opt=paths=source_relative book/book.proto book/price.proto
```

这里有几点需要注意：

1. 因为我们通过编译命令指定`--proto_path=proto`，所以import导入语句需要从`demo/proto`文件夹的下层目录`book`这一层开始写。
2. 因为导入的`price.proto`与`book.proto`同属于一个`package book;`，所以可以直接使用`price`作为类型，无需添加 package 前缀（即无需写成`book.price`）。

上述编译命令最终会生成`demo/proto/book/book.pb.go`文件。

```bash
demo
└── proto
    └── book
        ├── book.pb.go
        ├── book.proto
        ├── price.pb.go
        └── price.proto
```

### import其他目录下文件

我们在`demo/proto`目录下新建了一个`author`文件夹，用来存放与 author 相关的 protobuf 文件。例如我们定义一个表示作者信息的`author.proto`文件，其内容如下：

```protobuf
// demo/proto/author/author.proto

syntax = "proto3";

// 声明protobuf中的包名
package author;

// 声明生成的Go代码的导入路径
option go_package = "github.com/Q1mi/demo/proto/author";

message Info {
    string name = 1;
}
```

此时的目录结构：

```bash
demo
└── proto
    ├── author
    │   └── author.proto
    └── book
        ├── book.pb.go
        ├── book.proto
        ├── price.pb.go
        └── price.proto
```

假设我们的 book 需要增加一个作者信息的字段——`authorInfo`，这时我们需要在`demo/proto/book/book.proto`中导入其他目录下的 `author.proto` 文件。具体改动如下。

```protobuf
// proto/proto/book/book.proto

syntax = "proto3";

// 声明protobuf中的包名
package book;

// 声明生成的Go代码的导入路径
option go_package = "github.com/Q1mi/demo/proto/book";

// 引入同目录下的protobuf文件（注意起始位置为proto_path的下层）
import "book/price.proto";
// 引入其他目录下的protobuf文件
import "author/author.proto";

message Book {
    string title = 1;
    Price price = 2;
    author.Info authorInfo = 3;  // 需要带package前缀
}
```

我们通过`import "author/author.proto";`导入了`author`包的`author.proto`文件，所以在`book`包下使用`Info`类型时需要添加其包名前缀即`author.Info`。

编译命令如下：

```bash
protoc --proto_path=proto --go_out=proto --go_opt=paths=source_relative book/book.proto book/price.proto author/author.proto
```

此时的目录结构：

```bash
demo
└── proto
    ├── author
    │   ├── author.pb.go
    │   └── author.proto
    └── book
        ├── book.pb.go
        ├── book.proto
        ├── price.pb.go
        └── price.proto
```

### import google proto文件

有时候我们也需要在我们定义的 protobuf 文件中使用 Google 定义的类型，例如`Timestamp`、`Any`等。

例如我们要为我们的 book 添加出版日期——`date`字段，就可以通过 `import "google/protobuf/timestamp.proto";`导入并使用`Timestamp`类型了。

修改后的`book.proto`文件内容如下：

```protobuf
// demo/proto/book/book.proto

syntax = "proto3";

// 声明protobuf中的包名
package book;

// 声明生成的Go代码的导入路径
option go_package = "github.com/Q1mi/demo/proto/book";

// 引入同目录下的protobuf文件（注意起始位置为proto_path的下层）
import "book/price.proto";
// 引入其他目录下的protobuf文件
import "author/author.proto";
// 引入google/protobuf/timestamp.proto文件
import "google/protobuf/timestamp.proto";

message Book {
    string title = 1;
    Price price =2;
    author.Info authorInfo = 3;  // 需要带package前缀
    // Timestamp是大写T!大写T!大写T!
    google.protobuf.Timestamp date = 4;  // 注意包名前缀
}
```

那么这个 `google/protobuf/timestamp.proto` 是从哪里导入的呢？

通常我们下载 [protobuf](https://github.com/google/protobuf/releases)编译器的时候会解压得到如下文件：

![protoc](https://www.liwenzhou.com/images/Go/grpc/protoc.png)

其中：

- bin 目录下的 protoc 是可执行文件。
- include 目录下的是 google 定义的`.proto`文件，我们`import "google/protobuf/timestamp.proto"`就是从此处导入。

我们需要将下载得到的可执行文件`protoc`所在的 bin 目录加到我们电脑的环境变量中。

如果你不是通过这种方式安装的 protobuf 那么你也可以手动将 [Google 定义的protobuf文件](https://github.com/protocolbuffers/protobuf)下载到本地（git clone或者go get，protobuf文件在src下），然后通过 `--proto_path`指定其路径。

```bash
protoc --proto_path=/Users/liwenzhou/workspace/go/pkg/mod/github.com/protocolbuffers/protobuf@v3.21.2+incompatible/src/ --proto_path=proto --go_out=proto --go_opt=paths=source_relative book/book.proto book/price.proto author/author.proto
```

或者你还可以简单粗暴的把下载好的 protobuf 文件拷贝到你项目的 proto 目录下。

```bash
demo
└── proto
    ├── author
    │   ├── author.pb.go
    │   └── author.proto
    ├── book
    │   ├── book.pb.go
    │   ├── book.proto
    │   ├── price.pb.go
    │   └── price.proto
    └── google
        └── protobuf
            └── timestamp.proto
```

然后执行下面的编译命令：

```bash
protoc --proto_path=proto --go_out=proto --go_opt=paths=source_relative book/book.proto book/price.proto author/author.proto
```

## 生成gRPC代码

由于通常我们都是配合 gRPC 来使用 protobuf ，所以我们也需要基于`.proto`文件生成Go代码的同时生成 gRPC 代码。

要想生成 gRPC 代码就需要先安装 `protoc-gen-go-grpc` 插件。

```bash
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

上述命令会默认将插件安装到`$GOPATH/bin`，为了`protoc`编译器能找到这些插件，请确保你的`$GOPATH/bin`在环境变量中。

假设我们现在要提供一个创建书籍的 RPC 方法，那么我在`book.proto`中添加如下定义。

```protobuf
// demo/proto/book/book.proto

// ...省略...

service BookService{
    rpc Create(Book)returns(Book);
}
```

然后在 protoc 的编译命令添加 gRPC相关输出的参数，完整命令如下。

```bash
protoc --proto_path=proto --go_out=proto --go_opt=paths=source_relative --go-grpc_out=proto --go-grpc_opt=paths=source_relative book/book.proto book/price.proto author/author.proto
```

上述命令就会生成`book_grpc.pb.go`文件。

```bash
demo
└── proto
    ├── author
    │   ├── author.pb.go
    │   └── author.proto
    └── book
        ├── book.pb.go
        ├── book.proto
        ├── book_grpc.pb.go
        ├── price.pb.go
        └── price.proto
```

### gRPC-Gateway

[gRPC-Gateway](https://github.com/grpc-ecosystem/grpc-gateway) 也是日常开发中比较常用的一个工具，它同样也是根据 protobuf 生成相应的代码。

### 安装工具

```bash
go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
```

### 为protobuf文件添加注释

我们在`book.proto`文件中添加如下注释。

```protobuf
// demo/proto/book/book.proto

syntax = "proto3";

// 声明protobuf中的包名
package book;

// 声明生成的Go代码的导入路径
option go_package = "github.com/Q1mi/demo/proto/book";

// 引入同目录下的protobuf文件（注意起始位置为proto_path的下层）
import "book/price.proto";
// 引入其他目录下的protobuf文件
import "author/author.proto";
// 引入google/protobuf/timestamp.proto文件
import "google/protobuf/timestamp.proto";
// 引入google/api/annotations.proto文件
import "google/api/annotations.proto";

message Book {
    string title = 1;
    Price price = 2;
    author.Info authorInfo = 3;  // 需要带package前缀
    // Timestamp是大写T!大写T!大写T!
    google.protobuf.Timestamp date = 4;  // 注意包名前缀
}

service BookService{
    rpc Create(Book)returns(Book){
        option (google.api.http) = {
            post: "/v1/book"
            body: "*"
        };
    };
}
```

此时，我们又引入了`google/api/annotations.proto` ，这个文件是由googleapi定义在https://github.com/googleapis/googleapis。

想要在项目中引入上述protobuf源文件可以像上面引入`timestamp.proto`文件一样将这个库下载到本地然后通过`--proto_path`指定，或者直接把用到的 protobuf 源文件拷贝到我们的项目中。

本示例就采用第二种方法把此处用到的`google/api/annotations.proto`文件和`http.proto`文件拷贝到项目的`google/api`目录下（`annotations.proto`文件中引入了`http.proto`文件）。

此时的项目目录如下：

```bash
demo
└── proto
    ├── author
    │   ├── author.pb.go
    │   └── author.proto
    ├── book
    │   ├── book.pb.go
    │   ├── book.proto
    │   ├── book_grpc.pb.go
    │   ├── price.pb.go
    │   └── price.proto
    └── google
        └── api
            ├── annotations.proto
            └── http.proto
```

### 编译

这一次编译命令在之前的基础上要继续加上 gRPC-Gateway相关的 `--grpc-gateway_out=proto --grpc-gateway_opt=paths=source_relative` 参数。

完整的编译命令如下：

```bash
protoc --proto_path=proto --go_out=proto --go_opt=paths=source_relative --go-grpc_out=proto --go-grpc_opt=paths=source_relative --grpc-gateway_out=proto --grpc-gateway_opt=paths=source_relative book/book.proto book/price.proto author/author.proto
```

最终会编译得到一个`book.pb.gw.go`文件。

```bash
demo
└── proto
    ├── author
    │   ├── author.pb.go
    │   └── author.proto
    ├── book
    │   ├── book.pb.go
    │   ├── book.pb.gw.go
    │   ├── book.proto
    │   ├── book_grpc.pb.go
    │   ├── price.pb.go
    │   └── price.proto
    └── google
        └── api
            ├── annotations.proto
            └── http.proto
```

为了方便编译可以在项目下定义`Makefile`。

```makefile
.PHONY: gen help

PROTO_DIR=proto

gen:
	protoc \
	--proto_path=$(PROTO_DIR) \
	--go_out=$(PROTO_DIR) \
	--go_opt=paths=source_relative \
	--go-grpc_out=$(PROTO_DIR) \
	--go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=$(PROTO_DIR) \
	--grpc-gateway_opt=paths=source_relative \
	$(shell find $(PROTO_DIR) -iname "*.proto")

help:
	@echo "make gen - 生成pb及grpc代码"
```

后续想要编译只需在项目目录下执行`make gen`即可。

## 管理 protobuf

在企业的项目开发中，我们通常会把 protobuf 文件存储到一个单独的代码库中，并在具体项目中通过`git submodule`引入。这样做的好处是能够将 protobuf 文件统一管理和维护，避免因 protobuf 文件改动导致的问题。

本文示例代码已上传至github仓库：https://github.com/Q1mi/demo，请点击查看完整代码。