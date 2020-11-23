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
	var ret map[string]interface{}
	err = client.Call("DB.GetInStringValue", []interface{}{
		os.Args[1],
	}, &ret)
	if err != nil {
		log.Fatalln(err)
	}
	for k, v := range ret {
		log.Println(k, ":", v)
	}
}
