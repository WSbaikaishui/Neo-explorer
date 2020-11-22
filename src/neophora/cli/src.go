package cli

import (
	"net/rpc/jsonrpc"
)

// T ...
type T struct {
	Address string
}

// Call ...
func (me *T) Call(method string, args interface{}, res interface{}) error {
	client, err := jsonrpc.Dial("tcp", me.Address)
	if err != nil {
		return err
	}
	defer client.Close()
	return client.Call(method, args, res)
}
