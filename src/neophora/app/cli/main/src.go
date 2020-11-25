package main

import (
	"fmt"
	"io"
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
	address := os.ExpandEnv("0.0.0.0:${CLI_PORT}")
	http.ListenAndServe(address, &joh.T{
		Method: http.MethodPost,
		Codec: func(rwc io.ReadWriteCloser) rpc.ServerCodec {
			ret := &scex.T{}
			ret.Init(rwc)
			ret.SetF(func(v string) string {
				return fmt.Sprintf("NeoCli.%s", strings.ToUpper(v))
			})
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
