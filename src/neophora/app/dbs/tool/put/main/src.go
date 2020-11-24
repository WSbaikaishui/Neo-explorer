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

	key := []byte(os.Args[1])
	value := []byte(os.Args[2])
	var ret bool

	switch os.ExpandEnv("${DEC}") {
	case "hex":
		value, err = hex.DecodeString(os.Args[2])
		if err != nil {
			log.Fatalln(err)
		}
	}

	if err := client.Call("DB.Puts", struct {
		Key   []byte
		Value []byte
	}{
		Key:   key,
		Value: value,
	}, &ret); err != nil {
		log.Fatalln(err)
	}

	log.Println(ret)
}
