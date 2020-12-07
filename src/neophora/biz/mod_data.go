package biz

import (
	"encoding/hex"
	"fmt"
	"neophora/cli"
	"neophora/var/stderr"
	"net/url"
)

// Data ...
type Data struct {
	Client *cli.T
}

// Ping ...
func (me *Data) Ping(arg interface{}, ret *interface{}) error {
	*ret = "pong"
	return nil
}

// GetBlockByHeightInHex ...
func (me *Data) GetBlockByHeightInHex(index uint64, ret *string) error {
	uri := &url.URL{
		Scheme: "block",
		Host:   "height",
		Path:   fmt.Sprintf("/%016x", index),
	}

	var result []byte

	urs := uri.String()
	if err := me.Client.Calls("DB.Get", []byte(urs), &result); err != nil {
		return stderr.ErrUnknown
	}

	*ret = hex.EncodeToString(result)

	return nil
}
