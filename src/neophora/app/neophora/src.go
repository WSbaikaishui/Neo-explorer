package main

import (
	"log"
	"neophora/biz/api"
	"neophora/biz/data"
	"neophora/lib/cli"
	"neophora/lib/joh"
	"net/http"
	"net/rpc"
	"os"
	"strconv"
)

func main() {
	address := os.ExpandEnv("0.0.0.0:${NEOPHORA_PORT}")
	log.Println("[LISTEN]", address)
	http.ListenAndServe(address, &joh.T{})
}

func init() {
	address := os.ExpandEnv("${NEODB_ADDRESS}")
	poolsize, err := strconv.Atoi(os.ExpandEnv("${NEODB_POOLSIZE}"))
	if err != nil {
		log.Fatalln(err)
	}
	client := &cli.T{
		Address: address,
		Pool:    make(chan *rpc.Client, poolsize),
	}
	defer client.Close()
	rpc.Register(&api.T{
		Data: &data.T{
			Client: client,
		},
	})
}
