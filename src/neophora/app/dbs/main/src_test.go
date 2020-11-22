package main

import (
	"log"
	"net/rpc/jsonrpc"
	"os"
	"testing"
	"time"
)

func TestPing(t *testing.T) {
	client, err := jsonrpc.Dial("tcp", address)
	if err != nil {
		t.Error(err)
	}
	var ret interface{}
	client.Call("DB.Ping", nil, &ret)
	if ret != "pong" {
		t.Error(ret)
	}
}
func TestWriteRead(t *testing.T) {
	{
		client, err := jsonrpc.Dial("tcp", address)
		if err != nil {
			t.Error(err)
		}
		var ret interface{}
		err = client.Call("DB.PutInHexValue", map[string]interface{}{
			"block://height/1": "11111111",
		}, &ret)
		if err != nil {
			t.Error(err)
		}
		if ret != "ok" {
			t.Error(ret)
		}
	}
	{
		client, err := jsonrpc.Dial("tcp", address)
		if err != nil {
			t.Error(err)
		}
		var ret map[string]interface{}
		err = client.Call("DB.GetInHexValue", []interface{}{
			"block://height/1",
		}, &ret)
		if err != nil {
			t.Error(err)
		}
		if ret["block://height/1"] != "11111111" {
			t.Error(ret)
		}
	}
}

func init() {
	go main()
	for i := time.Duration(0); i < time.Second; i++ {
		time.Sleep(i)
		client, err := jsonrpc.Dial("tcp", address)
		if err == nil {
			client.Close()
			return
		}
	}
	log.Fatalln("NET LISTEN ERROR")
}

var address = os.ExpandEnv("127.0.0.1:${DBS_PORT}")
