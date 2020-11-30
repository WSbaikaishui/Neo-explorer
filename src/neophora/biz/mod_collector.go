package biz

import (
	"encoding/hex"
	"log"
	"neophora/srv/collector"
)

// Collector ...
type Collector struct {
	Service *collector.T
}

// Ping ...
func (me *Collector) Ping(args interface{}, ret *interface{}) error {
	*ret = "pong"
	return nil
}

// Write ...
func (me *Collector) Write(args [2]string, ret *interface{}) error {
	k, v := args[0], args[1]
	key := []byte(k)
	value, err := hex.DecodeString(v)
	if err != nil {
		log.Println("[!!!!][BIZ][DECODE]", k, v)
	}
	me.Service.Add(key, value)
	return nil
}
