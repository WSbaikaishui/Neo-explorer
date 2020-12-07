package main

import (
	"io"
	"log"
	"neophora/biz"
	"neophora/cli"
	"neophora/lib/joh"
	"neophora/lib/scex"
	"net/http"
	"net/rpc"
	"os"
	"strings"
)

func main() {
	address := os.ExpandEnv("0.0.0.0:${RPC_PORT}")
	http.ListenAndServe(address, &joh.T{
		Method: http.MethodPost,
		Codec: func(rwc io.ReadWriteCloser) rpc.ServerCodec {
			ret := &scex.T{}
			ret.Init(rwc)
			ret.SetF(func(v string) string {
				return v
			})
			return ret
		},
	})
}

func init() {
	address := os.ExpandEnv("${DBS_ADDRESS}")
	addresses := strings.Split(address, " ")
	log.Println("[MAIN][INIT][DBS ADDRESS]", addresses)
	rpc.Register(&biz.Version{})
	rpc.Register(&biz.Data{
		Client: &cli.T{
			Addresses: addresses,
			TryTimes:  3,
		},
	})
}
