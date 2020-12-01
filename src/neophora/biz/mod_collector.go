package biz

import (
	"encoding/hex"
	"log"
	"neophora/srv/collector"
	"neophora/srv/daemon"
)

// Collector ...
type Collector struct {
	Service *collector.T
	Daemon  *daemon.T
}

// Ping ...
func (me *Collector) Ping(args interface{}, ret *interface{}) error {
	*ret = "pong"
	return nil
}

// HeartBeat ...
func (me *Collector) HeartBeat(args interface{}, ret *interface{}) error {
	me.Daemon.Call()
	return nil
}

// Commit ...
func (me *Collector) Commit(args int, ret *interface{}) error {
	log.Println("[BIZ][COLLECTOR][COMMIT]", args)
	me.Daemon.Commit(args)
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
	me.Daemon.Call()
	return nil
}
