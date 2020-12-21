package cli

import (
	"net/rpc"
)

// T ...
type T struct {
	Address string
	Pool    chan *rpc.Client
}

// Call ...
func (me *T) Call(method string, args interface{}, reply interface{}) error {
	var client *rpc.Client
	var err error
	select {
	case client = <-me.Pool:
	default:
		client, err = rpc.Dial("tcp", me.Address)
	}
	if err != nil {
		return err
	}
	if err = client.Call(method, args, reply); err != nil {
		client.Close()
		return err
	}
	select {
	case me.Pool <- client:
	default:
		client.Close()
	}
	return nil
}

// Close ...
func (me *T) Close() {
	close(me.Pool)
	for client := range me.Pool {
		client.Close()
	}
}
