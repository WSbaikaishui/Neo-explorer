package main

import (
	"encoding/hex"
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
	var ret []byte
	err = client.Call("DB.Get", []byte(os.Args[1]), &ret)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(hex.EncodeToString(ret))
}
