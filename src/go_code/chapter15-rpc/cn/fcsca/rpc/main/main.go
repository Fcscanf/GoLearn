package main

import (
	"fmt"
	rpc2 "go_code/chapter15-rpc/cn/fcsca/rpc/rpc"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	rpc.Register(rpc2.DemoService{})
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
			continue
		}
		client := jsonrpc.NewClient(conn)

		var result float64
		err = client.Call("DemoService.Div", rpc2.Args{10, 3}, &result)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(result)
		}
		fmt.Println(result, err)

		err = client.Call("DemoService.Div", rpc2.Args{10, 0}, &result)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(result)
		}
		fmt.Println(result, err)
	}
}
