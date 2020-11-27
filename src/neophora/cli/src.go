package cli

import (
	"math/rand"
	"neophora/var/stderr"
	"net/rpc"
)

// T ...
type T struct {
	Addresses []string
	TryTimes  int
}

// Calls ...
func (me *T) Calls(method string, args interface{}, res interface{}) error {
	for i := 0; i < me.TryTimes; i++ {
		if err := me.Call(method, args, res); err == nil {
			return nil
		}
	}
	return stderr.ErrUnknown
}

// Call ...
func (me *T) Call(method string, args interface{}, res interface{}) error {
	address := me.Addresses[rand.Intn(len(me.Addresses))]
	client, err := rpc.Dial("tcp", address)
	if err != nil {
		return err
	}
	defer client.Close()
	return client.Call(method, args, res)
}
