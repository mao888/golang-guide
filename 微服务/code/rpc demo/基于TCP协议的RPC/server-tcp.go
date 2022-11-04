package main

import (
	"log"
	"net"
	"net/rpc"
)

func main() {
	service := new(ServiceA)
	rpc.Register(service) // 注册RPC服务
	listener, err := net.Listen("tcp", ":9091")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			return
		}
		rpc.ServeConn(conn)
	}
}
