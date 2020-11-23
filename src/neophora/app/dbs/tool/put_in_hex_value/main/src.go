package main

import (
	"log"
	"net/rpc/jsonrpc"
	"os"
)

func main() {
	address := os.ExpandEnv("${DBS_ADDRESS}")
	client, err := jsonrpc.Dial("tcp", address)
	if err != nil {
		log.Fatalln(err)
	}
	var ret interface{}
	err = client.Call("DB.PutInHexValue", map[string]interface{}{
		os.Args[1]: os.Args[2],
	}, &ret)
	if err != nil {
		log.Fatalln(err)
	}
	if ret != "ok" {
		log.Fatalln(err)
	}
	log.Println("ok")
}
