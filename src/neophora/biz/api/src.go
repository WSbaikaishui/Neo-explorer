package api

import (
	"neophora/biz/data"
)

// T ...
type T struct {
	Data *data.T
}

// Ping ...
func (me *T) Ping(args struct{}, ret *string) error {
	*ret = "pong"
	return nil
}
