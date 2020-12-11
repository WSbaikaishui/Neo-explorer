package main

import (
	"log"
	"neophora/biz/data"
	"neophora/biz/neocli"
	"neophora/biz/popper"
	"neophora/lib/bq"
	"neophora/lib/joh"
	"net/http"
	"net/rpc"
	"os"
	"strconv"

	"github.com/gomodule/redigo/redis"
)

func main() {
	address := os.ExpandEnv("0.0.0.0:${NEORPC_PORT}")
	log.Println("[LISTEN]", address)
	http.ListenAndServe(address, &joh.T{})
}

func init() {
	netowrk := os.ExpandEnv("${REDIS_NETWORK}")
	address := os.ExpandEnv("${REDIS_ADDRESS}")
	maxidle, err := strconv.Atoi(os.ExpandEnv("${REDIS_MAXIDLE}"))
	if err != nil {
		log.Fatalln(err)
	}
	db := redis.NewPool(func() (redis.Conn, error) {
		return redis.Dial(netowrk, address)
	}, maxidle)
	txs := &bq.T{}
	rpc.Register(&neocli.T{
		Data: &data.T{
			DB: db,
		},
		BQ: txs,
	})

	go http.ListenAndServe(os.ExpandEnv("0.0.0.0:${NEORPC_POPPERPORT}"), &popper.T{
		BQ: txs,
	})
}
