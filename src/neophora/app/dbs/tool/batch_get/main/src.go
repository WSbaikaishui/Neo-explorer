package main

import (
	"log"
	"net/rpc"
	"os"
)

func main() {
	address := os.ExpandEnv("${DBS_ADDRESS}")
	client, err := rpc.Dial("tcp", address)
	if err != nil {
		log.Fatalln(err)
	}
	var ret []interface{}
	err = client.Call("DB.Batch", struct {
		Method string
		Args   []interface{}
	}{
		Method: "Get",
		Args: []interface{}{
			[]byte("block://height/1"),
			[]byte("block://height/10"),
		},
	}, &ret)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(ret)
}
