package joh

import (
	"io"
	"neophora/lib/rwio"
	"net/http"
	"net/rpc"
)

// T ...
type T struct {
	Method string
	Codec  func(conn io.ReadWriteCloser) rpc.ServerCodec
}

func (me *T) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	conn := &rwio.T{R: req.Body, W: w}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	rpc.ServeCodec(me.Codec(conn))
}
