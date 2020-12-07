package main

import (
	"log"
	"neophora/biz"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
)

func main() {
	listener, err := net.Listen("tcp", os.ExpandEnv("0.0.0.0:${RPC_PORT}"))
	if err != nil {
		log.Fatalln(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("[MAIN][Accept]", err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}

func init() {
	rpc.Register(&biz.Version{})
	rpc.Register(&biz.Data{})
}
