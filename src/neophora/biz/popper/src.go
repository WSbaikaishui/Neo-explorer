package popper

import (
	"gopkg.in/neophora/lib/bq"
	"net/http"
)

// T ...
type T struct {
	BQ *bq.T
}

func (me *T) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	result := me.BQ.Pop()
	if result == nil {
		w.WriteHeader(http.StatusNotFound)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
