本文介绍了RPC的概念以及Go语言中标准库rpc的基本使用。

## 什么是RPC

RPC（Remote Procedure Call），即远程过程调用。它允许像调用本地服务一样调用远程服务。

RPC是一种服务器-客户端（Client/Server）模式，经典实现是一个通过发送请求-接受回应进行信息交互的系统。

首先与RPC（远程过程调用）相对应的是本地调用。

### 本地调用

```go
package main

import "fmt"

func add(x, y int)int{
	return x + y
}

func main(){
	// 调用本地函数add
	a := 10
	b := 20
	ret := add(x, y)
	fmt.Println(ret)
}
```

将上述程序编译成二进制文件——`app1`后运行，会输出结果30。

在`app1`程序中本地调用`add`函数的执行流程，可以理解为以下四个步骤。

1. 将变量 a 和 b 的值分别压入堆栈上
2. 执行 add 函数，从堆栈中获取 a 和 b 的值，并将它们分配给 x 和 y
3. 计算 x + y 的值并将其保存到堆栈中
4. 退出 add 函数并将 x + y 的值赋给 ret

### RPC调用

本地过程调用发生在同一进程中——定义`add`函数的代码和调用`add`函数的代码共享同一个内存空间，所以调用能够正常执行。![本地调用add函数](https://www.liwenzhou.com/images/Go/rpc/rpc01.png)但是我们无法直接在另一个程序——`app2`中调用`add`函数，因为它们是两个程序——内存空间是相互隔离的。（app1和app2可能部署在同一台服务器上也可能部署在互联网的不同服务器上。）![跨程序调用add函数](https://www.liwenzhou.com/images/Go/rpc/rpc02.png)RPC就是为了解决类似远程、跨内存空间、的函数/方法调用的。要实现RPC就需要解决以下三个问题。

1. 如何确定要执行的函数？ 在本地调用中，函数主体通过函数指针函数指定，然后调用 add 函数，编译器通过函数指针函数自动确定 add 函数在内存中的位置。但是在 RPC 中，调用不能通过函数指针完成，因为它们的内存地址可能完全不同。因此，调用方和被调用方都需要维护一个{ function <-> ID }映射表，以确保调用正确的函数。
2. 如何表达参数？ 本地过程调用中传递的参数是通过堆栈内存结构实现的，但 RPC 不能直接使用内存传递参数，因此参数或返回值需要在传输期间序列化并转换成字节流，反之亦然。
3. 如何进行网络传输？ 函数的调用方和被调用方通常是通过网络连接的，也就是说，function ID 和序列化字节流需要通过网络传输，因此，只要能够完成传输，调用方和被调用方就不受某个网络协议的限制。.例如，一些 RPC 框架使用 TCP 协议，一些使用 HTTP。

以往实现跨服务调用的时候，我们会采用RESTful API的方式，被调用方会对外提供一个HTTP接口，调用方按要求发起HTTP请求并接收API接口返回的响应数据。下面的示例是将`add`函数包装成一个RESTful API。

### HTTP调用RESTful API

首先，我们编写一个基于HTTP的server服务，它将接收其他程序发来的HTTP请求，执行特定的程序并将结果返回。

```go
// server/main.go

package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type addParam struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type addResult struct {
	Code int `json:"code"`
	Data int `json:"data"`
}

func add(x, y int) int {
	return x + y
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	// 解析参数
	b, _ := ioutil.ReadAll(r.Body)
	var param addParam
	json.Unmarshal(b, &param)
	// 业务逻辑
	ret := add(param.X, param.Y)
	// 返回响应
	respBytes , _ := json.Marshal(addResult{Code: 0, Data: ret})
	w.Write(respBytes)
}

func main() {
	http.HandleFunc("/add", addHandler)
	log.Fatal(http.ListenAndServe(":9090", nil))
}
```

我们编写一个客户端来请求上述HTTP服务，传递x和y两个整数，等待返回结果。

```go
// client/main.go

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type addParam struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type addResult struct {
	Code int `json:"code"`
	Data int `json:"data"`
}

func main() {
	// 通过HTTP请求调用其他服务器上的add服务
	url := "http://127.0.0.1:9090/add"
	param := addParam{
		X: 10,
		Y: 20,
	}
	paramBytes, _ := json.Marshal(param)
	resp, _ := http.Post(url, "application/json", bytes.NewReader(paramBytes))
	defer resp.Body.Close()

	respBytes, _ := ioutil.ReadAll(resp.Body)
	var respData addResult
	json.Unmarshal(respBytes, &respData)
	fmt.Println(respData.Data) // 30
}
```

这种模式是我们目前比较常见的跨服务或跨语言之间基于RESTful API的服务调用模式。 既然使用API调用也能实现类似远程调用的目的，为什么还要用RPC呢？

使用 RPC 的目的是让我们调用远程方法像调用本地方法一样无差别。并且基于RESTful API通常是基于HTTP协议，传输数据采用JSON等文本协议，相较于RPC 直接使用TCP协议，传输数据多采用二进制协议来说，RPC通常相比RESTful API性能会更好。

RESTful API多用于前后端之间的数据传输，而目前微服务架构下各个微服务之间多采用RPC调用。

## net/rpc

### 基础RPC示例

Go语言的 rpc 包提供对通过网络或其他 i/o 连接导出的对象方法的访问，服务器注册一个对象，并把它作为服务对外可见（服务名称就是类型名称）。注册后，对象的导出方法将支持远程访问。服务器可以注册不同类型的多个对象(服务) ，但是不支持注册同一类型的多个对象。

在下面的代码中我们定义一个`ServiceA`类型，并为其定义了一个可导出的`Add`方法。

```go
// rpc demo/service.go

package main

type Args struct {
	X, Y int
}

// ServiceA 自定义一个结构体类型
type ServiceA struct{}

// Add 为ServiceA类型增加一个可导出的Add方法
func (s *ServiceA) Add(args *Args, reply *int) error {
	*reply = args.X + args.Y
	return nil
}
```

通过下面的代码将上面定义的`ServiceA`类型注册为一个服务，其Add方法就支持RPC调用了。

```go
// rpc demo/server.go

package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	service := new(ServiceA)
	rpc.Register(service) // 注册RPC服务
	rpc.HandleHTTP()      // 基于HTTP协议
	l, e := net.Listen("tcp", ":9091")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil)
}
```

此时，client 端便能看到一个拥有“Add”方法的“ServiceA”服务，想要调用这个服务需要使用下面的代码先连接到server端再执行远程调用。

```go
// rpc demo/client.go

package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	// 建立HTTP连接
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:9091")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// 同步调用
	args := &Args{10, 20}
	var reply int
	err = client.Call("ServiceA.Add", args, &reply)
	if err != nil {
		log.Fatal("ServiceA.Add error:", err)
	}
	fmt.Printf("ServiceA.Add: %d+%d=%d\n", args.X, args.Y, reply)

	// 异步调用
	var reply2 int
	divCall := client.Go("ServiceA.Add", args, &reply2, nil)
	replyCall := <-divCall.Done // 接收调用结果
	fmt.Println(replyCall.Error)
	fmt.Println(reply2)
}
```

执行上述程序，查看 RPC 调用的结果。

先启动 server 端。

```bash
go run server.go service.go
```

再启动 client 端。

```bash
go run client.go service.go
```

会看到如下输出结果。

```bash
ServiceA.Add: 10+20=30
<nil>
30
```

这个RPC调用过程可以简化如下图所示。![rpc调用图示](https://www.liwenzhou.com/images/Go/rpc/rpc03.png)

### 基于TCP协议的RPC

当然 rpc 包也支持直接使用 TCP 协议而不使用HTTP协议。

server 端代码修改如下。

```go
// rpc_demo/server2.go

package main

import (
	"log"
	"net"
	"net/rpc"
)

func main() {
	service := new(ServiceA)
	rpc.Register(service) // 注册RPC服务
	l, e := net.Listen("tcp", ":9091")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	for {
		conn, _ := l.Accept()
		rpc.ServeConn(conn)
	}
}
```

client 端代码修改如下。

```go
// rpc demo/client2.go

package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	// 建立TCP连接
	client, err := rpc.Dial("tcp", "127.0.0.1:9091")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// 同步调用
	args := &Args{10, 20}
	var reply int
	err = client.Call("ServiceA.Add", args, &reply)
	if err != nil {
		log.Fatal("ServiceA.Add error:", err)
	}
	fmt.Printf("ServiceA.Add: %d+%d=%d\n", args.X, args.Y, reply)

	// 异步调用
	var reply2 int
	divCall := client.Go("ServiceA.Add", args, &reply2, nil)
	replyCall := <-divCall.Done // 接收调用结果
	fmt.Println(replyCall.Error)
	fmt.Println(reply2)
}
```

### 使用JSON协议的RPC

rpc 包默认使用的是 gob 协议对传输数据进行序列化/反序列化，比较有局限性。下面的代码将尝试使用 JSON 协议对传输数据进行序列化与反序列化。

server 端代码修改如下。

```go
// rpc demo/server3.go

package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	service := new(ServiceA)
	rpc.Register(service) // 注册RPC服务
	l, e := net.Listen("tcp", ":9091")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	for {
		conn, _ := l.Accept()
		// 使用JSON协议
		rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
```

client 端代码修改如下。

```go
// rpc demo/client3.go

package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	// 建立TCP连接
	conn, err := net.Dial("tcp", "127.0.0.1:9091")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	// 使用JSON协议
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	// 同步调用
	args := &Args{10, 20}
	var reply int
	err = client.Call("ServiceA.Add", args, &reply)
	if err != nil {
		log.Fatal("ServiceA.Add error:", err)
	}
	fmt.Printf("ServiceA.Add: %d+%d=%d\n", args.X, args.Y, reply)

	// 异步调用
	var reply2 int
	divCall := client.Go("ServiceA.Add", args, &reply2, nil)
	replyCall := <-divCall.Done // 接收调用结果
	fmt.Println(replyCall.Error)
	fmt.Println(reply2)
}
```

### Python调用RPC

下面的代码演示了如何使用 python client 远程调用上面 Go server中 serviceA的Add方法。

```python
import socket
import json

request = {
    "id": 0,
    "params": [{"x":10, "y":20}],  # 参数要对应上Args结构体
    "method": "ServiceA.Add"
}

client = socket.create_connection(("127.0.0.1", 9091),5)
client.sendall(json.dumps(request).encode())

rsp = client.recv(1024)
rsp = json.loads(rsp.decode())
print(rsp)
```

## RPC原理

RPC 让远程调用就像本地调用一样，其调用过程可拆解为以下步骤。

![rpc](https://www.liwenzhou.com/images/Go/rpc/rpc04.png)

① 服务调用方（client）以本地调用方式调用服务；

② client stub接收到调用后负责将方法、参数等组装成能够进行网络传输的消息体；

③ client stub找到服务地址，并将消息发送到服务端；

④ server 端接收到消息；

⑤ server stub收到消息后进行解码；

⑥ server stub根据解码结果调用本地的服务；

⑦ 本地服务执行并将结果返回给server stub；

⑧ server stub将返回结果打包成能够进行网络传输的消息体；

⑨ 按地址将消息发送至调用方；

⑩ client 端接收到消息；

⑪ client stub收到消息并进行解码；

⑫ 调用方得到最终结果。

使用RPC框架的目标是只需要关心第1步和最后1步，中间的其他步骤统统封装起来，让使用者无需关心。例如社区中各式RPC框架（grpc、thrift等）就是为了让RPC调用更方便。