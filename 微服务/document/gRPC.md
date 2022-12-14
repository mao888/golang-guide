# gRPC教程

RPC算是近些年比较火热的概念了，随着微服务架构的兴起，RPC的应用越来越广泛。本文介绍了RPC和gRPC的相关概念，并且通过详细的代码示例介绍了gRPC的基本使用。

## gRPC是什么

`gRPC`是一种现代化开源的高性能RPC框架，能够运行于任意环境之中。最初由谷歌进行开发。它使用HTTP/2作为传输协议。

> 快速了解HTTP/2就戳[HTTP/2相比HTTP/1.x有哪些重大改进？](https://www.zhihu.com/question/34074946)

在gRPC里，客户端可以像调用本地方法一样直接调用其他机器上的服务端应用程序的方法，帮助你更容易创建分布式应用程序和服务。与许多RPC系统一样，gRPC是基于定义一个服务，指定一个可以远程调用的带有参数和返回类型的的方法。在服务端程序中实现这个接口并且运行gRPC服务处理客户端调用。在客户端，有一个stub提供和服务端相同的方法。![grpc](https://www.liwenzhou.com/images/Go/grpc/grpc.svg)

## 为什么要用gRPC

使用gRPC， 我们可以一次性的在一个`.proto`文件中定义服务并使用任何支持它的语言去实现客户端和服务端，反过来，它们可以应用在各种场景中，从Google的服务器到你自己的平板电脑—— gRPC帮你解决了不同语言及环境间通信的复杂性。使用`protocol buffers`还能获得其他好处，包括高效的序列化，简单的IDL以及容易进行接口更新。总之一句话，使用gRPC能让我们更容易编写跨语言的分布式代码。

> IDL（Interface description language）是指接口描述语言，是用来描述软件组件接口的一种计算机语言，是跨平台开发的基础。IDL通过一种中立的方式来描述接口，使得在不同平台上运行的对象和用不同语言编写的程序可以相互通信交流；比如，一个组件用C++写成，另一个组件用Go写成。

## 安装gRPC

### 安装gRPC

在你的项目目录下执行以下命令，获取 gRPC 作为项目依赖。

```bash
go get google.golang.org/grpc@latest
```

### 安装Protocol Buffers v3

安装用于生成gRPC服务代码的协议编译器，最简单的方法是从下面的链接：https://github.com/google/protobuf/releases下载适合你平台的预编译好的二进制文件（`protoc-<version>-<platform>.zip`）。

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

### 安装插件

因为本文我们是使用Go语言做开发，接下来执行下面的命令安装`protoc`的Go插件：

安装go语言插件：

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
```

该插件会根据`.proto`文件生成一个后缀为`.pb.go`的文件，包含所有`.proto`文件中定义的类型及其序列化方法。

安装grpc插件：

```bash
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

该插件会生成一个后缀为`_grpc.pb.go`的文件，其中包含：

- 一种接口类型(或存根) ，供客户端调用的服务方法。
- 服务器要实现的接口类型。

上述命令会默认将插件安装到`$GOPATH/bin`，为了`protoc`编译器能找到这些插件，请确保你的`$GOPATH/bin`在环境变量中。

[protocol-buffers 官方Go教程](https://developers.google.com/protocol-buffers/docs/gotutorial)

### 检查

依次执行以下命令检查一下是否开发环境都准备完毕。

1. 确认 protoc 安装完成。

   ```bash
   ❯ protoc --version
   libprotoc 3.20.1
   ```

2. 确认 protoc-gen-go 安装完成。

   ```bash
   ❯ protoc-gen-go --version
   protoc-gen-go v1.28.0
   ```

   如果这里提示`protoc-gen-go`不是可执行的程序，请确保你的 GOPATH 下的 bin 目录在你电脑的环境变量中。

3. 确认 protoc-gen-go-grpc 安装完成。

   ```bash
   ❯ protoc-gen-go-grpc --version
   protoc-gen-go-grpc 1.2.0
   ```

   如果这里提示`protoc-gen-go-grpc`不是可执行的程序，请确保你的 GOPATH 下的 bin 目录在你电脑的环境变量中。

## gRPC的开发方式

把大象放进冰箱分几步？

1. 把冰箱门打开。
2. 把大象放进去。
3. 把冰箱门带上。

gRPC开发同样分三步：

### 编写`.proto`文件定义服务

像许多 RPC 系统一样，gRPC 基于定义服务的思想，指定可以通过参数和返回类型远程调用的方法。默认情况下，gRPC 使用 [protocol buffers](https://developers.google.com/protocol-buffers)作为接口定义语言(IDL)来描述服务接口和有效负载消息的结构。可以根据需要使用其他的IDL代替。

例如，下面使用 protocol buffers 定义了一个`HelloService`服务。

```protobuf
service HelloService {
  rpc SayHello (HelloRequest) returns (HelloResponse);
}

message HelloRequest {
  string greeting = 1;
}

message HelloResponse {
  string reply = 1;
}
```

在gRPC中你可以定义四种类型的服务方法。

- 普通 rpc，客户端向服务器发送一个请求，然后得到一个响应，就像普通的函数调用一样。

```protobuf
  rpc SayHello(HelloRequest) returns (HelloResponse);
```

- 服务器流式 rpc，其中客户端向服务器发送请求，并获得一个流来读取一系列消息。客户端从返回的流中读取，直到没有更多的消息。gRPC 保证在单个 RPC 调用中的消息是有序的。

```protobuf
  rpc LotsOfReplies(HelloRequest) returns (stream HelloResponse);
```

- 客户端流式 rpc，其中客户端写入一系列消息并将其发送到服务器，同样使用提供的流。一旦客户端完成了消息的写入，它就等待服务器读取消息并返回响应。同样，gRPC 保证在单个 RPC 调用中对消息进行排序。

```protobuf
  rpc LotsOfGreetings(stream HelloRequest) returns (HelloResponse);
```

- 双向流式 rpc，其中双方使用读写流发送一系列消息。这两个流独立运行，因此客户端和服务器可以按照自己喜欢的顺序读写: 例如，服务器可以等待接收所有客户端消息后再写响应，或者可以交替读取消息然后写入消息，或者其他读写组合。每个流中的消息是有序的。

### 生成指定语言的代码

在 `.proto` 文件中的定义好服务之后，gRPC 提供了生成客户端和服务器端代码的 protocol buffers 编译器插件。

我们使用这些插件可以根据需要生成`Java`、`Go`、`C++`、`Python`等语言的代码。我们通常会在客户端调用这些 API，并在服务器端实现相应的 API。

- 在服务器端，服务器实现服务声明的方法，并运行一个 gRPC 服务器来处理客户端发来的调用请求。gRPC 底层会对传入的请求进行解码，执行被调用的服务方法，并对服务响应进行编码。
- 在客户端，客户端有一个称为存根（stub）的本地对象，它实现了与服务相同的方法。然后，客户端可以在本地对象上调用这些方法，将调用的参数包装在适当的 protocol buffers 消息类型中—— gRPC 在向服务器发送请求并返回服务器的 protocol buffers 响应之后进行处理。

### 编写业务逻辑代码

gRPC 帮我们解决了 RPC 中的服务调用、数据传输以及消息编解码，我们剩下的工作就是要编写业务逻辑代码。

在服务端编写业务代码实现具体的服务方法，在客户端按需调用这些方法。

## gRPC入门示例

### 编写proto代码

`Protocol Buffers`是一种与语言无关，平台无关的可扩展机制，用于序列化结构化数据。使用`Protocol Buffers`可以一次定义结构化的数据，然后可以使用特殊生成的源代码轻松地在各种数据流中使用各种语言编写和读取结构化数据。

关于`Protocol Buffers`的教程可以查看[Protocol Buffers V3中文指南](https://www.liwenzhou.com/posts/Go/Protobuf3-language-guide-zh/)，本文后续内容默认读者熟悉`Protocol Buffers`。

```protobuf
syntax = "proto3"; // 版本声明，使用Protocol Buffers v3版本

option go_package = "xx";  // 指定生成的Go代码在你项目中的导入路径

package pb; // 包名


// 定义服务
service Greeter {
    // SayHello 方法
    rpc SayHello (HelloRequest) returns (HelloResponse) {}
}

// 请求消息
message HelloRequest {
    string name = 1;
}

// 响应消息
message HelloResponse {
    string reply = 1;
}
```

### 编写Server端Go代码

我们新建一个`hello_server`项目，在项目根目录下执行`go mod init hello_server`。

再新建一个`pb`文件夹，将上面的 proto 文件保存为`hello.proto`，将`go_package`按如下方式修改。

```protobuf
// ...

option go_package = "hello_server/pb";

// ...
```

此时，项目的目录结构为：

```bash
hello_server
├── go.mod
├── go.sum
├── main.go
└── pb
    └── hello.proto
```

在项目根目录下执行以下命令，根据`hello.proto`生成 go 源码文件。

```bash
protoc --go_out=. --go_opt=paths=source_relative \
--go-grpc_out=. --go-grpc_opt=paths=source_relative \
pb/hello.proto
```

**注意** 如果你的终端不支持`\`符（例如某些同学的Windows），那么你就复制粘贴下面不带`\`的命令执行。

```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pb/hello.proto
```

生成后的go源码文件会保存在pb文件夹下。

```bash
hello_server
├── go.mod
├── go.sum
├── main.go
└── pb
    ├── hello.pb.go
    ├── hello.proto
    └── hello_grpc.pb.go
```

将下面的内容添加到`hello_server/main.go`中。

```go
package main

import (
	"context"
	"fmt"
	"hello_server/pb"
	"net"

	"google.golang.org/grpc"
)

// hello server

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Reply: "Hello " + in.Name}, nil
}

func main() {
	// 监听本地的8972端口
	lis, err := net.Listen("tcp", ":8972")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()                  // 创建gRPC服务器
	pb.RegisterGreeterServer(s, &server{}) // 在gRPC服务端注册服务
	// 启动服务
	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}
```

编译并执行 `http_server`：

```bash
go build
./server
```

### 编写Client端Go代码

我们新建一个`hello_client`项目，在项目根目录下执行`go mod init hello_client`。

再新建一个`pb`文件夹，将上面的 proto 文件保存为`hello.proto`，将`go_package`按如下方式修改。

```protobuf
// ...

option go_package = "hello_client/pb";

// ...
```

在项目根目录下执行以下命令，根据`hello.proto`在`http_client`项目下生成 go 源码文件。

```bash
protoc --go_out=. --go_opt=paths=source_relative \
--go-grpc_out=. --go-grpc_opt=paths=source_relative \
pb/hello.proto
```

**注意** 如果你的终端不支持`\`符（例如某些同学的Windows），那么你就复制粘贴下面不带`\`的命令执行。

```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pb/hello.proto
```

此时，项目的目录结构为：

```bash
http_client
├── go.mod
├── go.sum
├── main.go
└── pb
    ├── hello.pb.go
    ├── hello.proto
    └── hello_grpc.pb.go
```

在`http_client/main.go`文件中按下面的代码调用`http_server`提供的 `SayHello` RPC服务。

```go
package main

import (
	"context"
	"flag"
	"log"
	"time"

	"hello_client/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// hello_client

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "127.0.0.1:8972", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	// 连接到server端，此处禁用安全传输
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// 执行RPC调用并打印收到的响应数据
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetReply())
}
```

保存后将`http_client`编译并执行：

```bash
go build
./hello_client -name=七米
```

得到以下输出结果，说明RPC调用成功。

```bash
2022/05/15 00:31:52 Greeting: Hello 七米
```

### gRPC跨语言调用

接下来，我们演示一下如何使用gRPC实现跨语言的RPC调用。

我们使用`Python`语言编写`Client`，然后向上面使用`go`语言编写的`server`发送RPC请求。

python下安装 grpc：

```bash
python -m pip install grpcio
```

安装gRPC tools：

```bash
python -m pip install grpcio-tools
```

### 生成Python代码

新建一个`py_client`目录，将`hello.proto`文件保存到`py_client/pb/`目录下。 在`py_client`目录下执行以下命令，生成python源码文件。

```bash
cd py_cleint
python3 -m grpc_tools.protoc -Ipb --python_out=. --grpc_python_out=. pb/hello.proto
```

### 编写Python版RPC客户端

将下面的代码保存到`py_client/client.py`文件中。

```python
from __future__ import print_function

import logging

import grpc
import hello_pb2
import hello_pb2_grpc


def run():
    # NOTE(gRPC Python Team): .close() is possible on a channel and should be
    # used in circumstances in which the with statement does not fit the needs
    # of the code.
    with grpc.insecure_channel('127.0.0.1:8972') as channel:
        stub = hello_pb2_grpc.GreeterStub(channel)
        resp = stub.SayHello(hello_pb2.HelloRequest(name='q1mi'))
    print("Greeter client received: " + resp.reply)


if __name__ == '__main__':
    logging.basicConfig()
    run()
```

此时项目的目录结构图如下：

```bash
py_client
├── client.py
├── hello_pb2.py
├── hello_pb2_grpc.py
└── pb
    └── hello.proto
```

### Python RPC 调用

执行`client.py`调用go语言的`SayHello`RPC服务。

```bash
❯ python3 client.py
Greeter client received: Hello q1mi
```

这里我们就实现了，使用 Python 代码编写的client去调用Go语言版本的server了。

点击右边的链接查看完整代码：[gRPC_demo完整代码](https://github.com/Q1mi/gRPC_demo)

## gRPC流式示例

在上面的示例中，客户端发起了一个RPC请求到服务端，服务端进行业务处理并返回响应给客户端，这是gRPC最基本的一种工作方式（Unary RPC）。除此之外，依托于HTTP2，gRPC还支持流式RPC（Streaming RPC）。

### 服务端流式RPC

客户端发出一个RPC请求，服务端与客户端之间建立一个单向的流，服务端可以向流中写入多个响应消息，最后主动关闭流；而客户端需要监听这个流，不断获取响应直到流关闭。应用场景举例：客户端向服务端发送一个股票代码，服务端就把该股票的实时数据源源不断的返回给客户端。

我们在此编写一个使用多种语言打招呼的方法，客户端发来一个用户名，服务端分多次返回打招呼的信息。

1.定义服务

```protobuf
// 服务端返回流式数据
rpc LotsOfReplies(HelloRequest) returns (stream HelloResponse);
```

修改`.proto`文件后，需要重新使用 protocol buffers编译器生成客户端和服务端代码。

2.服务端需要实现 `LotsOfReplies` 方法。

```go
// LotsOfReplies 返回使用多种语言打招呼
func (s *server) LotsOfReplies(in *pb.HelloRequest, stream pb.Greeter_LotsOfRepliesServer) error {
	words := []string{
		"你好",
		"hello",
		"こんにちは",
		"안녕하세요",
	}

	for _, word := range words {
		data := &pb.HelloResponse{
			Reply: word + in.GetName(),
		}
		// 使用Send方法返回多个数据
		if err := stream.Send(data); err != nil {
			return err
		}
	}
	return nil
}
```

3.客户端调用`LotsOfReplies` 并将收到的数据依次打印出来。

```go
func runLotsOfReplies(c pb.GreeterClient) {
	// server端流式RPC
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stream, err := c.LotsOfReplies(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("c.LotsOfReplies failed, err: %v", err)
	}
	for {
		// 接收服务端返回的流式数据，当收到io.EOF或错误时退出
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("c.LotsOfReplies failed, err: %v", err)
		}
		log.Printf("got reply: %q\n", res.GetReply())
	}
}
```

执行程序后会得到如下输出结果。

```bash
2022/05/21 14:36:20 got reply: "你好七米"
2022/05/21 14:36:20 got reply: "hello七米"
2022/05/21 14:36:20 got reply: "こんにちは七米"
2022/05/21 14:36:20 got reply: "안녕하세요七米"
```

### 客户端流式RPC

客户端传入多个请求对象，服务端返回一个响应结果。典型的应用场景举例：物联网终端向服务器上报数据、大数据流式计算等。

在这个示例中，我们编写一个多次发送人名，服务端统一返回一个打招呼消息的程序。

1.定义服务

```protobuf
// 客户端发送流式数据
rpc LotsOfGreetings(stream HelloRequest) returns (HelloResponse);
```

修改`.proto`文件后，需要重新使用 protocol buffers编译器生成客户端和服务端代码。

2.服务端实现`LotsOfGreetings`方法。

```go
// LotsOfGreetings 接收流式数据
func (s *server) LotsOfGreetings(stream pb.Greeter_LotsOfGreetingsServer) error {
	reply := "你好："
	for {
		// 接收客户端发来的流式数据
		res, err := stream.Recv()
		if err == io.EOF {
			// 最终统一回复
			return stream.SendAndClose(&pb.HelloResponse{
				Reply: reply,
			})
		}
		if err != nil {
			return err
		}
		reply += res.GetName()
	}
}  
```

3.客户端调用`LotsOfGreetings`方法，向服务端发送流式请求数据，接收返回值并打印。

```go
func runLotsOfGreeting(c pb.GreeterClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// 客户端流式RPC
	stream, err := c.LotsOfGreetings(ctx)
	if err != nil {
		log.Fatalf("c.LotsOfGreetings failed, err: %v", err)
	}
	names := []string{"七米", "q1mi", "沙河娜扎"}
	for _, name := range names {
		// 发送流式数据
		err := stream.Send(&pb.HelloRequest{Name: name})
		if err != nil {
			log.Fatalf("c.LotsOfGreetings stream.Send(%v) failed, err: %v", name, err)
		}
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("c.LotsOfGreetings failed: %v", err)
	}
	log.Printf("got reply: %v", res.GetReply())
}
```

执行上述函数将得到如下数据结果。

```bash
2022/05/21 14:57:31 got reply: 你好：七米q1mi沙河娜扎
```

### 双向流式RPC

双向流式RPC即客户端和服务端均为流式的RPC，能发送多个请求对象也能接收到多个响应对象。典型应用示例：聊天应用等。

我们这里还是编写一个客户端和服务端进行人机对话的双向流式RPC示例。

1.定义服务

```protobuf
// 双向流式数据
rpc BidiHello(stream HelloRequest) returns (stream HelloResponse);
```

修改`.proto`文件后，需要重新使用 protocol buffers编译器生成客户端和服务端代码。

2.服务端实现`BidiHello`方法。

```go
// BidiHello 双向流式打招呼
func (s *server) BidiHello(stream pb.Greeter_BidiHelloServer) error {
	for {
		// 接收流式请求
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		reply := magic(in.GetName()) // 对收到的数据做些处理

		// 返回流式响应
		if err := stream.Send(&pb.HelloResponse{Reply: reply}); err != nil {
			return err
		}
	}
}
```

这里我们还定义了一个处理数据的`magic`函数，其内容如下。

```go
// magic 一段价值连城的“人工智能”代码
func magic(s string) string {
	s = strings.ReplaceAll(s, "吗", "")
	s = strings.ReplaceAll(s, "吧", "")
	s = strings.ReplaceAll(s, "你", "我")
	s = strings.ReplaceAll(s, "？", "!")
	s = strings.ReplaceAll(s, "?", "!")
	return s
}
```

3.客户端调用`BidiHello`方法，一边从终端获取输入的请求数据发送至服务端，一边从服务端接收流式响应。

```go
func runBidiHello(c pb.GreeterClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	// 双向流模式
	stream, err := c.BidiHello(ctx)
	if err != nil {
		log.Fatalf("c.BidiHello failed, err: %v", err)
	}
	waitc := make(chan struct{})
	go func() {
		for {
			// 接收服务端返回的响应
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("c.BidiHello stream.Recv() failed, err: %v", err)
			}
			fmt.Printf("AI：%s\n", in.GetReply())
		}
	}()
	// 从标准输入获取用户输入
	reader := bufio.NewReader(os.Stdin) // 从标准输入生成读对象
	for {
		cmd, _ := reader.ReadString('\n') // 读到换行
		cmd = strings.TrimSpace(cmd)
		if len(cmd) == 0 {
			continue
		}
		if strings.ToUpper(cmd) == "QUIT" {
			break
		}
		// 将获取到的数据发送至服务端
		if err := stream.Send(&pb.HelloRequest{Name: cmd}); err != nil {
			log.Fatalf("c.BidiHello stream.Send(%v) failed: %v", cmd, err)
		}
	}
	stream.CloseSend()
	<-waitc
}
```

将服务端和客户端的代码都运行起来，就可以实现简单的对话程序了。

```bash
hello
AI：hello
你吃饭了吗?
AI：我吃饭了!
你会写代码吗
AI：我会写代码
可以和你约会吗?
AI：可以和我约会!
现在可以吗?
AI：现在可以!
走吧?
AI：走!
```

## metadata

元数据（[metadata](https://pkg.go.dev/google.golang.org/grpc/metadata)）是指在处理RPC请求和响应过程中需要但又不属于具体业务（例如身份验证详细信息）的信息，采用键值对列表的形式，其中键是`string`类型，值通常是`[]string`类型，但也可以是二进制数据。gRPC中的 metadata 类似于我们在 HTTP headers中的键值对，元数据可以包含认证token、请求标识和监控标签等。

metadata中的键是大小写不敏感的，由字母、数字和特殊字符`-`、`_`、`.`组成并且不能以`grpc-`开头（gRPC保留自用），二进制值的键名必须以`-bin`结尾。

元数据对 gRPC 本身是不可见的，我们通常是在应用程序代码或中间件中处理元数据，我们不需要在`.proto`文件中指定元数据。

如何访问元数据取决于具体使用的编程语言。 在Go语言中我们是用[google.golang.org/grpc/metadata](https://pkg.go.dev/google.golang.org/grpc/metadata)这个库来操作metadata。

metadata 类型定义如下：

```go
type MD map[string][]string
```

元数据可以像普通map一样读取。注意，这个 map 的值类型是`[]string`，因此用户可以使用一个键附加多个值。

### 创建新的metadata

常用的创建MD的方法有以下两种。

第一种方法是使用函数 `New` 基于`map[string]string` 创建元数据:

```go
md := metadata.New(map[string]string{"key1": "val1", "key2": "val2"})
```

另一种方法是使用`Pairs`。具有相同键的值将合并到一个列表中:

```go
md := metadata.Pairs(
    "key1", "val1",
    "key1", "val1-2", // "key1"的值将会是 []string{"val1", "val1-2"}
    "key2", "val2",
)
```

注意: 所有的键将自动转换为小写，因此“ kEy1”和“ Key1”将是相同的键，它们的值将合并到相同的列表中。这种情况适用于 `New` 和 `Pair`。

### 元数据中存储二进制数据

在元数据中，键始终是字符串。但是值可以是字符串或二进制数据。要在元数据中存储二进制数据值，只需在密钥中添加“-bin”后缀。在创建元数据时，将对带有“-bin”后缀键的值进行编码:

```go
md := metadata.Pairs(
    "key", "string value",
    "key-bin", string([]byte{96, 102}), // 二进制数据在发送前会进行(base64) 编码
                                        // 收到后会进行解码
)
```

### 从请求上下文中获取元数据

可以使用 `FromIncomingContext` 可以从RPC请求的上下文中获取元数据:

```go
func (s *server) SomeRPC(ctx context.Context, in *pb.SomeRequest) (*pb.SomeResponse, err) {
    md, ok := metadata.FromIncomingContext(ctx)
    // do something with metadata
}
```

### 发送和接收元数据-客户端

#### 发送metadata

有两种方法可以将元数据发送到服务端。推荐的方法是使用 `AppendToOutgoingContext` 将 kv 对附加到context。无论context中是否已经有元数据都可以使用这个方法。如果先前没有元数据，则添加元数据; 如果context中已经存在元数据，则将 kv 对合并进去。

```go
// 创建带有metadata的context
ctx := metadata.AppendToOutgoingContext(ctx, "k1", "v1", "k1", "v2", "k2", "v3")

// 添加一些 metadata 到 context (e.g. in an interceptor)
ctx := metadata.AppendToOutgoingContext(ctx, "k3", "v4")

// 发起普通RPC请求
response, err := client.SomeRPC(ctx, someRequest)

// 或者发起流式RPC请求
stream, err := client.SomeStreamingRPC(ctx)
```

或者，可以使用 `NewOutgoingContext` 将元数据附加到context。但是，这将替换context中的任何已有的元数据，因此必须注意保留现有元数据(如果需要的话)。这个方法比使用 `AppendToOutgoingContext` 要慢。这方面的一个例子如下:

```go
// 创建带有metadata的context
md := metadata.Pairs("k1", "v1", "k1", "v2", "k2", "v3")
ctx := metadata.NewOutgoingContext(context.Background(), md)

// 添加一些metadata到context (e.g. in an interceptor)
send, _ := metadata.FromOutgoingContext(ctx)
newMD := metadata.Pairs("k3", "v3")
ctx = metadata.NewOutgoingContext(ctx, metadata.Join(send, newMD))

// 发起普通RPC请求
response, err := client.SomeRPC(ctx, someRequest)

// 或者发起流式RPC请求
stream, err := client.SomeStreamingRPC(ctx)
```

#### 接收metadata

客户端可以接收的元数据包括header和trailer。

> trailer可以用于服务器希望在处理请求后给客户端发送任何内容，例如在流式RPC中只有等所有结果都流到客户端后才能计算出负载信息，这时候就不能使用headers（header在数据之前，trailer在数据之后）。

引申：[HTTP trailer](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Headers/Trailer)

##### 普通调用

可以使用 [CallOption](https://godoc.org/google.golang.org/grpc#CallOption) 中的 [Header](https://godoc.org/google.golang.org/grpc#Header) 和 [Trailer](https://godoc.org/google.golang.org/grpc#Trailer) 函数来获取普通RPC调用发送的header和trailer:

```go
var header, trailer metadata.MD // 声明存储header和trailer的变量
r, err := client.SomeRPC(
    ctx,
    someRequest,
    grpc.Header(&header),    // 将会接收header
    grpc.Trailer(&trailer),  // 将会接收trailer
)

// do something with header and trailer
```

##### 流式调用

流式调用包括：

- 客户端流式
- 服务端流式
- 双向流式

使用接口 [ClientStream](https://godoc.org/google.golang.org/grpc#ClientStream) 中的 `Header` 和 `Trailer` 函数，可以从返回的流中接收 Header 和 Trailer:

```go
stream, err := client.SomeStreamingRPC(ctx)

// 接收 header
header, err := stream.Header()

// 接收 trailer
trailer := stream.Trailer()
```

### 发送和接收元数据-服务器端

#### 接收metadata

要读取客户端发送的元数据，服务器需要从 RPC 上下文检索它。如果是普通RPC调用，则可以使用 RPC 处理程序的上下文。对于流调用，服务器需要从流中获取上下文。

##### 普通调用

```go
func (s *server) SomeRPC(ctx context.Context, in *pb.someRequest) (*pb.someResponse, error) {
    md, ok := metadata.FromIncomingContext(ctx)
    // do something with metadata
}
```

##### 流式调用

```go
func (s *server) SomeStreamingRPC(stream pb.Service_SomeStreamingRPCServer) error {
    md, ok := metadata.FromIncomingContext(stream.Context()) // get context from stream
    // do something with metadata
}
```

#### 发送metadata

##### 普通调用

在普通调用中，服务器可以调用 [grpc](https://godoc.org/google.golang.org/grpc) 模块中的 [SendHeader](https://godoc.org/google.golang.org/grpc#SendHeader) 和 [SetTrailer](https://godoc.org/google.golang.org/grpc#SetTrailer) 函数向客户端发送header和trailer。这两个函数将context作为第一个参数。它应该是 RPC 处理程序的上下文或从中派生的上下文：

```go
func (s *server) SomeRPC(ctx context.Context, in *pb.someRequest) (*pb.someResponse, error) {
    // 创建和发送 header
    header := metadata.Pairs("header-key", "val")
    grpc.SendHeader(ctx, header)
    // 创建和发送 trailer
    trailer := metadata.Pairs("trailer-key", "val")
    grpc.SetTrailer(ctx, trailer)
}
```

##### 流式调用

对于流式调用，可以使用接口 [ServerStream](https://godoc.org/google.golang.org/grpc#ServerStream) 中的 `SendHeader` 和 `SetTrailer` 函数发送header和trailer:

```go
func (s *server) SomeStreamingRPC(stream pb.Service_SomeStreamingRPCServer) error {
    // 创建和发送 header
    header := metadata.Pairs("header-key", "val")
    stream.SendHeader(header)
    // 创建和发送 trailer
    trailer := metadata.Pairs("trailer-key", "val")
    stream.SetTrailer(trailer)
}
```

### 普通RPC调用metadata示例

#### client端的metadata操作

下面的代码片段演示了client端如何设置和获取metadata。

```go
// unaryCallWithMetadata 普通RPC调用客户端metadata操作
func unaryCallWithMetadata(c pb.GreeterClient, name string) {
	fmt.Println("--- UnarySayHello client---")
	// 创建metadata
	md := metadata.Pairs(
		"token", "app-test-q1mi",
		"request_id", "1234567",
	)
	// 基于metadata创建context.
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	// RPC调用
	var header, trailer metadata.MD
	r, err := c.SayHello(
		ctx,
		&pb.HelloRequest{Name: name},
		grpc.Header(&header),   // 接收服务端发来的header
		grpc.Trailer(&trailer), // 接收服务端发来的trailer
	)
	if err != nil {
		log.Printf("failed to call SayHello: %v", err)
		return
	}
	// 从header中取location
	if t, ok := header["location"]; ok {
		fmt.Printf("location from header:\n")
		for i, e := range t {
			fmt.Printf(" %d. %s\n", i, e)
		}
	} else {
		log.Printf("location expected but doesn't exist in header")
		return
	}
   // 获取响应结果
	fmt.Printf("got response: %s\n", r.Reply)
	// 从trailer中取timestamp
	if t, ok := trailer["timestamp"]; ok {
		fmt.Printf("timestamp from trailer:\n")
		for i, e := range t {
			fmt.Printf(" %d. %s\n", i, e)
		}
	} else {
		log.Printf("timestamp expected but doesn't exist in trailer")
	}
}
```

#### server端metadata操作

下面的代码片段演示了server端如何设置和获取metadata。

```go
// UnarySayHello 普通RPC调用服务端metadata操作
func (s *server) UnarySayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	// 通过defer中设置trailer.
	defer func() {
		trailer := metadata.Pairs("timestamp", strconv.Itoa(int(time.Now().Unix())))
		grpc.SetTrailer(ctx, trailer)
	}()

	// 从客户端请求上下文中读取metadata.
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.DataLoss, "UnarySayHello: failed to get metadata")
	}
	if t, ok := md["token"]; ok {
		fmt.Printf("token from metadata:\n")
		if len(t) < 1 || t[0] != "app-test-q1mi" {
			return nil, status.Error(codes.Unauthenticated, "认证失败")
		}
	}

	// 创建和发送header.
	header := metadata.New(map[string]string{"location": "BeiJing"})
	grpc.SendHeader(ctx, header)

	fmt.Printf("request received: %v, say hello...\n", in)

	return &pb.HelloResponse{Reply: in.Name}, nil
}
```

### 流式RPC调用metadata示例

这里以双向流式RPC为例演示客户端和服务端如何进行metadata操作。

#### client端的metadata操作

下面的代码片段演示了client端在服务端流式RPC模式下如何设置和获取metadata。

```go
// bidirectionalWithMetadata 流式RPC调用客户端metadata操作
func bidirectionalWithMetadata(c pb.GreeterClient, name string) {
	// 创建metadata和context.
	md := metadata.Pairs("token", "app-test-q1mi")
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	// 使用带有metadata的context执行RPC调用.
	stream, err := c.BidiHello(ctx)
	if err != nil {
		log.Fatalf("failed to call BidiHello: %v\n", err)
	}

	go func() {
		// 当header到达时读取header.
		header, err := stream.Header()
		if err != nil {
			log.Fatalf("failed to get header from stream: %v", err)
		}
		// 从返回响应的header中读取数据.
		if l, ok := header["location"]; ok {
			fmt.Printf("location from header:\n")
			for i, e := range l {
				fmt.Printf(" %d. %s\n", i, e)
			}
		} else {
			log.Println("location expected but doesn't exist in header")
			return
		}

		// 发送所有的请求数据到server.
		for i := 0; i < 5; i++ {
			if err := stream.Send(&pb.HelloRequest{Name: name}); err != nil {
				log.Fatalf("failed to send streaming: %v\n", err)
			}
		}
		stream.CloseSend()
	}()

	// 读取所有的响应.
	var rpcStatus error
	fmt.Printf("got response:\n")
	for {
		r, err := stream.Recv()
		if err != nil {
			rpcStatus = err
			break
		}
		fmt.Printf(" - %s\n", r.Reply)
	}
	if rpcStatus != io.EOF {
		log.Printf("failed to finish server streaming: %v", rpcStatus)
		return
	}

	// 当RPC结束时读取trailer
	trailer := stream.Trailer()
	// 从返回响应的trailer中读取metadata.
	if t, ok := trailer["timestamp"]; ok {
		fmt.Printf("timestamp from trailer:\n")
		for i, e := range t {
			fmt.Printf(" %d. %s\n", i, e)
		}
	} else {
		log.Printf("timestamp expected but doesn't exist in trailer")
	}
}
```

#### server端的metadata操作

下面的代码片段演示了server端在服务端流式RPC模式下设置和操作metadata。

```go
// BidirectionalStreamingSayHello 流式RPC调用客户端metadata操作
func (s *server) BidirectionalStreamingSayHello(stream pb.Greeter_BidiHelloServer) error {
	// 在defer中创建trailer记录函数的返回时间.
	defer func() {
		trailer := metadata.Pairs("timestamp", strconv.Itoa(int(time.Now().Unix())))
		stream.SetTrailer(trailer)
	}()

	// 从client读取metadata.
	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		return status.Errorf(codes.DataLoss, "BidirectionalStreamingSayHello: failed to get metadata")
	}

	if t, ok := md["token"]; ok {
		fmt.Printf("token from metadata:\n")
		for i, e := range t {
			fmt.Printf(" %d. %s\n", i, e)
		}
	}

	// 创建和发送header.
	header := metadata.New(map[string]string{"location": "X2Q"})
	stream.SendHeader(header)

	// 读取请求数据发送响应数据.
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		fmt.Printf("request received %v, sending reply\n", in)
		if err := stream.Send(&pb.HelloResponse{Reply: in.Name}); err != nil {
			return err
		}
	}
}
```

## 错误处理

### gRPC code

类似于HTTP定义了一套响应状态码，gRPC也定义有一些状态码。Go语言中此状态码由[codes](https://pkg.go.dev/google.golang.org/grpc/codes)定义，本质上是一个uint32。

```go
type Code uint32
```

使用时需导入`google.golang.org/grpc/codes`包。

```go
import "google.golang.org/grpc/codes"
```

目前已经定义的状态码有如下几种。

|        Code        |  值  |                             含义                             |
| :----------------: | :--: | :----------------------------------------------------------: |
|         OK         |  0   |                           请求成功                           |
|      Canceled      |  1   |                          操作已取消                          |
|      Unknown       |  2   | 未知错误。如果从另一个地址空间接收到的状态值属 于在该地址空间中未知的错误空间，则可以返回此错误的示例。 没有返回足够的错误信息的API引发的错误也可能会转换为此错误 |
|  InvalidArgument   |  3   | 表示客户端指定的参数无效。 请注意，这与 FailedPrecondition 不同。 它表示无论系统状态如何都有问题的参数（例如，格式错误的文件名）。 |
|  DeadlineExceeded  |  4   | 表示操作在完成之前已过期。对于改变系统状态的操作，即使操作成功完成，也可能会返回此错误。 例如，来自服务器的成功响应可能已延迟足够长的时间以使截止日期到期。 |
|      NotFound      |  5   |        表示未找到某些请求的实体（例如，文件或目录）。        |
|   AlreadyExists    |  6   |            创建实体的尝试失败，因为实体已经存在。            |
|  PermissionDenied  |  7   | 表示调用者没有权限执行指定的操作。 它不能用于拒绝由耗尽某些资源引起的（使用 ResourceExhausted ）。 如果无法识别调用者，也不能使用它（使用 Unauthenticated ）。 |
| ResourceExhausted  |  8   | 表示某些资源已耗尽，可能是每个用户的配额，或者整个文件系统空间不足 |
| FailedPrecondition |  9   | 指示操作被拒绝，因为系统未处于操作执行所需的状态。 例如，要删除的目录可能是非空的，rmdir 操作应用于非目录等。 |
|      Aborted       |  10  | 表示操作被中止，通常是由于并发问题，如排序器检查失败、事务中止等。 |
|     OutOfRange     |  11  |                 表示尝试超出有效范围的操作。                 |
|   Unimplemented    |  12  |            表示此服务中未实施或不支持/启用操作。             |
|      Internal      |  13  | 意味着底层系统预期的一些不变量已被破坏。 如果你看到这个错误，则说明问题很严重。 |
|    Unavailable     |  14  | 表示服务当前不可用。这很可能是暂时的情况，可以通过回退重试来纠正。 请注意，重试非幂等操作并不总是安全的。 |
|      DataLoss      |  15  |                 表示不可恢复的数据丢失或损坏                 |
|  Unauthenticated   |  16  |            表示请求没有用于操作的有效身份验证凭据            |
|      _maxCode      |  17  |                              -                               |

### gRPC Status

Go语言使用的gRPC Status 定义在[google.golang.org/grpc/status](https://pkg.go.dev/google.golang.org/grpc/status)，使用时需导入。

```go
import "google.golang.org/grpc/status"
```

RPC服务的方法应该返回 `nil` 或来自`status.Status`类型的错误。客户端可以直接访问错误。

#### 创建错误

当遇到错误时，gRPC服务的方法函数应该创建一个 `status.Status`。通常我们会使用 `status.New`函数并传入适当的`status.Code`和错误描述来生成一个`status.Status`。调用`status.Err`方法便能将一个`status.Status`转为`error`类型。也存在一个简单的`status.Error`方法直接生成`error`。下面是两种方式的比较。

```go
// 创建status.Status
st := status.New(codes.NotFound, "some description")
err := st.Err()  // 转为error类型

// vs.

err := status.Error(codes.NotFound, "some description")
```

#### 为错误添加其他详细信息

在某些情况下，可能需要为服务器端的特定错误添加详细信息。`status.WithDetails`就是为此而存在的，它可以添加任意多个`proto.Message`，我们可以使用`google.golang.org/genproto/googleapis/rpc/errdetails`中的定义或自定义的错误详情。

```go
st := status.New(codes.ResourceExhausted, "Request limit exceeded.")
ds, _ := st.WithDetails(
	// proto.Message
)
return nil, ds.Err()
```

然后，客户端可以通过首先将普通`error`类型转换回`status.Status`，然后使用`status.Details`来读取这些详细信息。

```go
s := status.Convert(err)
for _, d := range s.Details() {
	// ...
}
```

### 代码示例

我们现在要为`hello`服务设置访问限制，每个`name`只能调用一次`SayHello`方法，超过此限制就返回一个请求超过限制的错误。

#### 服务端

使用map存储每个name的请求次数，超过1次则返回错误，并且记录错误详情。

```go
package main

import (
	"context"
	"fmt"
	"hello_server/pb"
	"net"
	"sync"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// grpc server

type server struct {
	pb.UnimplementedGreeterServer
	mu    sync.Mutex     // count的并发锁
	count map[string]int // 记录每个name的请求次数
}

// SayHello 是我们需要实现的方法
// 这个方法是我们对外提供的服务
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.count[in.Name]++ // 记录用户的请求次数
	// 超过1次就返回错误
	if s.count[in.Name] > 1 {
		st := status.New(codes.ResourceExhausted, "Request limit exceeded.")
		ds, err := st.WithDetails(
			&errdetails.QuotaFailure{
				Violations: []*errdetails.QuotaFailure_Violation{{
					Subject:     fmt.Sprintf("name:%s", in.Name),
					Description: "限制每个name调用一次",
				}},
			},
		)
		if err != nil {
			return nil, st.Err()
		}
		return nil, ds.Err()
	}
	// 正常返回响应
	reply := "hello " + in.GetName()
	return &pb.HelloResponse{Reply: reply}, nil
}

func main() {
	// 启动服务
	l, err := net.Listen("tcp", ":8972")
	if err != nil {
		fmt.Printf("failed to listen, err:%v\n", err)
		return
	}
	s := grpc.NewServer() // 创建grpc服务
	// 注册服务，注意初始化count
	pb.RegisterGreeterServer(s, &server{count: make(map[string]int)})
	// 启动服务
	err = s.Serve(l)
	if err != nil {
		fmt.Printf("failed to serve,err:%v\n", err)
		return
	}
}
```

#### 客户端

当服务端返回错误时，尝试从错误中获取detail信息。

```go
package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc/status"
	"hello_client/pb"
	"log"
	"time"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// grpc 客户端
// 调用server端的 SayHello 方法

var name = flag.String("name", "七米", "通过-name告诉server你是谁")

func main() {
	flag.Parse() // 解析命令行参数

	// 连接server
	conn, err := grpc.Dial("127.0.0.1:8972", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("grpc.Dial failed,err:%v", err)
		return
	}
	defer conn.Close()
	// 创建客户端
	c := pb.NewGreeterClient(conn) // 使用生成的Go代码
	// 调用RPC方法
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		s := status.Convert(err)        // 将err转为status
		for _, d := range s.Details() { // 获取details
			switch info := d.(type) {
			case *errdetails.QuotaFailure:
				fmt.Printf("Quota failure: %s\n", info)
			default:
				fmt.Printf("Unexpected type: %s\n", info)
			}
		}
		fmt.Printf("c.SayHello failed, err:%v\n", err)
		return
	}
	// 拿到了RPC响应
	log.Printf("resp:%v\n", resp.GetReply())
}
```

## 加密或认证

### 无加密认证

在上面的示例中，我们都没有为我们的 gRPC 配置加密或认证，属于不安全的连接（insecure connection）。

Client端：

```go
conn, _ := grpc.Dial("127.0.0.1:8972", grpc.WithTransportCredentials(insecure.NewCredentials()))
client := pb.NewGreeterClient(conn)
```

Server端：

```go
s := grpc.NewServer()
lis, _ := net.Listen("tcp", "127.0.0.1:8972")
// error handling omitted
s.Serve(lis)
```

### 使用服务器身份验证 SSL/TLS

gRPC 内置支持 SSL/TLS，可以通过 SSL/TLS 证书建立安全连接，对传输的数据进行加密处理。

这里我们演示如何使用自签名证书进行server端加密。

#### 生成证书

##### 生成私钥

执行下面的命令生成私钥文件——`server.key`。

```bash
openssl ecparam -genkey -name secp384r1 -out server.key
```

这里生成的是ECC私钥，当然你也可以使用RSA。

##### 生成自签名的证书

> Go1.15之后x509弃用Common Name改用SANs。

当出现如下错误时，需要提供SANs信息。

```bash
transport: authentication handshake failed: x509: certificate relies on legacy Common Name field, use SANs or temporarily enable Common Name matching with GODEBUG=x509ignoreCN=0
```

为了在证书中添加SANs信息，我们将下面自定义配置保存到`server.cnf`文件中。

```conf
[ req ]
default_bits       = 4096
default_md		= sha256
distinguished_name = req_distinguished_name
req_extensions     = req_ext

[ req_distinguished_name ]
countryName                 = Country Name (2 letter code)
countryName_default         = CN
stateOrProvinceName         = State or Province Name (full name)
stateOrProvinceName_default = BEIJING
localityName                = Locality Name (eg, city)
localityName_default        = BEIJING
organizationName            = Organization Name (eg, company)
organizationName_default    = DEV
commonName                  = Common Name (e.g. server FQDN or YOUR name)
commonName_max              = 64
commonName_default          = liwenzhou.com

[ req_ext ]
subjectAltName = @alt_names

[alt_names]
DNS.1   = localhost
DNS.2   = liwenzhou.com
IP      = 127.0.0.1
```

执行下面的命令生成自签名证书——`server.crt`。

```bash
openssl req -nodes -new -x509 -sha256 -days 3650 -config server.cnf -extensions 'req_ext' -key server.key -out server.crt
```

#### 建立安全连接

Server端使用`credentials.NewServerTLSFromFile`函数分别加载证书`server.cert`和秘钥`server.key`。

```go
creds, _ := credentials.NewServerTLSFromFile(certFile, keyFile)
s := grpc.NewServer(grpc.Creds(creds))
lis, _ := net.Listen("tcp", "127.0.0.1:8972")
// error handling omitted
s.Serve(lis)
```

而client端使用上一步生成的证书文件——`server.cert`建立安全连接。

```go
creds, _ := credentials.NewClientTLSFromFile(certFile, "")
conn, _ := grpc.Dial("127.0.0.1:8972", grpc.WithTransportCredentials(creds))
// error handling omitted
client := pb.NewGreeterClient(conn)
// ...
```

除了这种自签名证书的方式外，生产环境对外通信时通常需要使用受信任的CA证书。

## 拦截器（中间件）

gRPC 为在每个 ClientConn/Server 基础上实现和安装拦截器提供了一些简单的 API。 拦截器拦截每个 RPC 调用的执行。用户可以使用拦截器进行日志记录、身份验证/授权、指标收集以及许多其他可以跨 RPC 共享的功能。

在 gRPC 中，拦截器根据拦截的 RPC 调用类型可以分为两类。第一个是普通拦截器（一元拦截器），它拦截普通RPC 调用。另一个是流拦截器，它处理流式 RPC 调用。而客户端和服务端都有自己的普通拦截器和流拦截器类型。因此，在 gRPC 中总共有四种不同类型的拦截器。

### 客户端端拦截器

#### 普通拦截器/一元拦截器

[UnaryClientInterceptor](https://godoc.org/google.golang.org/grpc#UnaryClientInterceptor) 是客户端一元拦截器的类型，它的函数前面如下：

```go
func(ctx context.Context, method string, req, reply interface{}, cc *ClientConn, invoker UnaryInvoker, opts ...CallOption) error
```

一元拦截器的实现通常可以分为三个部分: 调用 RPC 方法之前（预处理）、调用 RPC 方法（RPC调用）和调用 RPC 方法之后（调用后）。

- 预处理：用户可以通过检查传入的参数(如 RPC 上下文、方法字符串、要发送的请求和 CallOptions 配置)来获得有关当前 RPC 调用的信息。
- RPC调用：预处理完成后，可以通过执行`invoker`执行 RPC 调用。
- 调用后：一旦调用者返回应答和错误，用户就可以对 RPC 调用进行后处理。通常，它是关于处理返回的响应和错误的。 若要在 `ClientConn` 上安装一元拦截器，请使用`DialOptionWithUnaryInterceptor`的`DialOption`配置 Dial 。

#### 流拦截器

[StreamClientInterceptor](https://godoc.org/google.golang.org/grpc#StreamClientInterceptor)是客户端流拦截器的类型。它的函数签名是

```go
func(ctx context.Context, desc *StreamDesc, cc *ClientConn, method string, streamer Streamer, opts ...CallOption) (ClientStream, error)
```

流拦截器的实现通常包括预处理和流操作拦截。

- 预处理：类似于上面的一元拦截器。
- 流操作拦截：流拦截器并没有事后进行 RPC 方法调用和后处理，而是拦截了用户在流上的操作。首先，拦截器调用传入的`streamer`以获取 `ClientStream`，然后包装 `ClientStream` 并用拦截逻辑重载其方法。最后，拦截器将包装好的 `ClientStream` 返回给用户进行操作。

若要为 `ClientConn` 安装流拦截器，请使用`WithStreamInterceptor`的 DialOption 配置 Dial。

### server端拦截器

服务器端拦截器与客户端类似，但提供的信息略有不同。

#### 普通拦截器/一元拦截器

[UnaryServerInterceptor](https://godoc.org/google.golang.org/grpc#UnaryServerInterceptor)是服务端的一元拦截器类型，它的函数签名是

```go
func(ctx context.Context, req interface{}, info *UnaryServerInfo, handler UnaryHandler) (resp interface{}, err error)
```

服务端一元拦截器具体实现细节和客户端版本的类似。

若要为服务端安装一元拦截器，请使用 `UnaryInterceptor` 的`ServerOption`配置 `NewServer`。

#### 流拦截器

[StreamServerInterceptor](https://godoc.org/google.golang.org/grpc#StreamServerInterceptor)是服务端流式拦截器的类型，它的签名如下：

```go
func(srv interface{}, ss ServerStream, info *StreamServerInfo, handler StreamHandler) error
```

实现细节类似于客户端流拦截器部分。

若要为服务端安装流拦截器，请使用 `StreamInterceptor` 的`ServerOption`来配置 `NewServer`。

### 拦截器示例

下面将演示一个完整的拦截器示例，我们为一元RPC和流式RPC服务都添加上拦截器。

我们首先定义一个名为`valid`的校验函数。

```go
// valid 校验认证信息.
func valid(authorization []string) bool {
	if len(authorization) < 1 {
		return false
	}
	token := strings.TrimPrefix(authorization[0], "Bearer ")
	// 执行token认证的逻辑
	// 这里是为了演示方便简单判断token是否与"some-secret-token"相等
	return token == "some-secret-token"
}
```

#### 客户端拦截器定义

##### 一元拦截器

```go
// unaryInterceptor 客户端一元拦截器
func unaryInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	var credsConfigured bool
	for _, o := range opts {
		_, ok := o.(grpc.PerRPCCredsCallOption)
		if ok {
			credsConfigured = true
			break
		}
	}
	if !credsConfigured {
		opts = append(opts, grpc.PerRPCCredentials(oauth.NewOauthAccess(&oauth2.Token{
			AccessToken: "some-secret-token",
		})))
	}
	start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	end := time.Now()
	fmt.Printf("RPC: %s, start time: %s, end time: %s, err: %v\n", method, start.Format("Basic"), end.Format(time.RFC3339), err)
	return err
}
```

其中，`grpc.PerRPCCredentials()`函数指明每个 RPC 请求使用的凭据，它接收一个`credentials.PerRPCCredentials`接口类型的参数。`credentials.PerRPCCredentials`接口的定义如下：

```go
type PerRPCCredentials interface {
	// GetRequestMetadata 获取当前请求的元数据,如果需要则会设置token。
	// 传输层在每个请求上调用，并且数据会被填充到headers或其他context。
	GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error)
	// RequireTransportSecurity 指示该 Credentials 的传输是否需要需要 TLS 加密
	RequireTransportSecurity() bool
}
```

而示例代码中使用的`oauth.NewOauthAccess()`是内置oauth包提供的一个函数，用来返回包含给定token的`PerRPCCredentials`。

```go
// NewOauthAccess constructs the PerRPCCredentials using a given token.
func NewOauthAccess(token *oauth2.Token) credentials.PerRPCCredentials {
	return oauthAccess{token: *token}
}

func (oa oauthAccess) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	ri, _ := credentials.RequestInfoFromContext(ctx)
	if err := credentials.CheckSecurityLevel(ri.AuthInfo, credentials.PrivacyAndIntegrity); err != nil {
		return nil, fmt.Errorf("unable to transfer oauthAccess PerRPCCredentials: %v", err)
	}
	return map[string]string{
		"authorization": oa.token.Type() + " " + oa.token.AccessToken,
	}, nil
}

func (oa oauthAccess) RequireTransportSecurity() bool {
	return true
}
```

##### 流式拦截器

自定义一个`ClientStream`类型。

```go
type wrappedStream struct {
	grpc.ClientStream
}
```

`wrappedStream`重写`grpc.ClientStream`接口的`RecvMsg`和`SendMsg`方法。

```go
func (w *wrappedStream) RecvMsg(m interface{}) error {
	logger("Receive a message (Type: %T) at %v", m, time.Now().Format(time.RFC3339))
	return w.ClientStream.RecvMsg(m)
}

func (w *wrappedStream) SendMsg(m interface{}) error {
	logger("Send a message (Type: %T) at %v", m, time.Now().Format(time.RFC3339))
	return w.ClientStream.SendMsg(m)
}

func newWrappedStream(s grpc.ClientStream) grpc.ClientStream {
	return &wrappedStream{s}
}
```

> 这里的`wrappedStream`嵌入了`grpc.ClientStream`接口类型，然后又重新实现了一遍`grpc.ClientStream`接口的方法。

下面就定义一个流式拦截器，最后返回上面定义的`wrappedStream`。

```go
// streamInterceptor 客户端流式拦截器
func streamInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	var credsConfigured bool
	for _, o := range opts {
		_, ok := o.(*grpc.PerRPCCredsCallOption)
		if ok {
			credsConfigured = true
			break
		}
	}
	if !credsConfigured {
		opts = append(opts, grpc.PerRPCCredentials(oauth.NewOauthAccess(&oauth2.Token{
			AccessToken: "some-secret-token",
		})))
	}
	s, err := streamer(ctx, desc, cc, method, opts...)
	if err != nil {
		return nil, err
	}
	return newWrappedStream(s), nil
}
```

#### 服务端拦截器定义

##### 一元拦截器

服务端定义一个一元拦截器，对从请求元数据中获取的`authorization`进行校验。

```go
// unaryInterceptor 服务端一元拦截器
func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// authentication (token verification)
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "missing metadata")
	}
	if !valid(md["authorization"]) {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token")
	}
	m, err := handler(ctx, req)
	if err != nil {
		fmt.Printf("RPC failed with error %v\n", err)
	}
	return m, err
}
```

#### 流拦截器

同样为流RPC也定义一个从元数据中获取认证信息的流式拦截器。

```go
// streamInterceptor 服务端流拦截器
func streamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	// authentication (token verification)
	md, ok := metadata.FromIncomingContext(ss.Context())
	if !ok {
		return status.Errorf(codes.InvalidArgument, "missing metadata")
	}
	if !valid(md["authorization"]) {
		return status.Errorf(codes.Unauthenticated, "invalid token")
	}

	err := handler(srv, newWrappedStream(ss))
	if err != nil {
		fmt.Printf("RPC failed with error %v\n", err)
	}
	return err
}
```

#### 注册拦截器

客户端注册拦截器

```go
conn, err := grpc.Dial("127.0.0.1:8972",
	grpc.WithTransportCredentials(creds),
	grpc.WithUnaryInterceptor(unaryInterceptor),
	grpc.WithStreamInterceptor(streamInterceptor),
)
```

服务端注册拦截器

```go
s := grpc.NewServer(
	grpc.Creds(creds),
	grpc.UnaryInterceptor(unaryInterceptor),
	grpc.StreamInterceptor(streamInterceptor),
)
```

### go-grpc-middleware

社区中有很多开源的常用的grpc中间件——[go-grpc-middleware](https://github.com/grpc-ecosystem/go-grpc-middleware)，大家可以根据需要选择使用。