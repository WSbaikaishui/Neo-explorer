package main

import (
	"log"
	"neophora/biz"
	"neophora/dat"
	"net"
	"net/rpc"
	"os"
)

func main() {
	path := os.ExpandEnv("${DBS_DBPATH}")
	db := &dat.T{}
	err := db.Init(path)
	if err != nil {
		log.Fatalln("[MAIN][Init]", err)
	}
	defer db.Close()

	rpc.Register(&biz.DB{DB: db})

	address := os.ExpandEnv("0.0.0.0:${DBS_PORT}")
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalln("[MAIN][Listen]", err)
	}

	log.Println("[MAIN][Listen]", address)
	rpc.DefaultServer.Accept(listener)
}
