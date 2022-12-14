// rpc demo/server-http.go

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
	rpc.HandleHTTP()      //	基于HTTP协议
	listener, err := net.Listen("tcp", ":9091")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	http.Serve(listener, nil)
}
