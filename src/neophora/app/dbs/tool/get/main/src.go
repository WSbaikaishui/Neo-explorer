package main

import (
	"encoding/hex"
	"fmt"
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

	key := []byte(os.Args[1])
	var ret []byte

	if err := client.Call("DB.Get", key, &ret); err != nil {
		log.Fatalln(err)
	}
	switch os.ExpandEnv("${ENC}") {
	case "hex":
		fmt.Println(hex.EncodeToString(ret))
	case "str":
		fmt.Println(string(ret))
	default:
		fmt.Println(ret)
	}
}
