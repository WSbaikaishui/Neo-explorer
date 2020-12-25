package main

import (
	"log"
	"neophora/lib/cli"
	"net/rpc"
	"os"
)

func main() {
	client := &cli.T{
		Address: os.ExpandEnv("${NEODB_ADDRESS}"),
		Pool:    make(chan *rpc.Client, 1),
	}
	var result bool
	if err := client.Call("T.DeleteRange", struct {
		Key []byte
	}{
		Key: []byte(os.ExpandEnv("${KEY}")),
	}, &result); err != nil {
		log.Println(err)
	}
	log.Println(result)
}
