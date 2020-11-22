package main

import (
	"io"
	"neophora/biz"
	"neophora/cli"
	"neophora/lib/joh"
	"neophora/lib/scex"
	"net/http"
	"net/rpc"
	"os"
)

func main() {
	address := os.ExpandEnv("0.0.0.0:${CLI_PORT}")
	http.ListenAndServe(address, &joh.T{
		Method: http.MethodPost,
		Codec: func(rwc io.ReadWriteCloser) rpc.ServerCodec {
			ret := &scex.T{}
			ret.Init(rwc)
			return ret
		},
	})
}

func init() {
	address := os.ExpandEnv("${DBS_ADDRESS}")
	rpc.Register(&biz.NeoCli{
		Client: &cli.T{Address: address},
	})
}
