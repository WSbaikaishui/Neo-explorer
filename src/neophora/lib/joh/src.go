package joh

import (
	"io"
	"neophora/lib/rwio"
	"neophora/lib/scex"
	"net/http"
	"net/rpc"
)

// T ...
type T struct {
	Codec func(conn io.ReadWriteCloser) rpc.ServerCodec
}

func (me *T) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	conn := &rwio.T{R: req.Body, W: w}
	codec := &scex.T{}
	codec.Init(conn)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	rpc.ServeCodec(codec)
}
