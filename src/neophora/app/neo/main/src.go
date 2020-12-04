package main

import (
	"fmt"
	"log"
	"neophora/biz"
	"neophora/cli"
	"neophora/lib/run"
	"neophora/srv/collector"
	"neophora/srv/daemon"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", os.ExpandEnv("0.0.0.0:${NEO_PORT}"))
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
	address := os.ExpandEnv("${DBS_ADDRESS}")
	addresses := strings.Split(address, " ")
	client := &cli.T{
		Addresses: addresses,
		TryTimes:  3,
	}
	dmn := &daemon.T{
		IP:      os.ExpandEnv("${NEO_IP}"),
		ExecCmd: os.ExpandEnv("${NEO_EXECCMD}"),
	}
	fmt.Sscanf(os.ExpandEnv("${NEO_STARTHEIGHT}"), "%d", &dmn.StartFrom)
	clt := &collector.T{
		Queue: make(chan struct {
			Key   []byte
			Value []byte
		}, 1024),
		Client: client,
	}
	rCLT := &run.T{
		Service: clt,
		Slaves:  4,
	}
	rCLT.Run()
	rDMN := &run.T{
		Service: dmn,
	}
	rDMN.Run()
	rpc.Register(&biz.Version{})
	rpc.Register(&biz.Collector{
		Service: clt,
		Daemon:  dmn,
	})
}
