package biz

import (
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
func (me *Data) GetBlockByHeightInHex(arg uint, ret *interface{}) error {
	var result map[string]interface{}
	uri := &url.URL{
		Scheme: "block",
		Host:   "height",
		Path:   fmt.Sprintf("/%d", arg),
	}
	uristring := uri.String()
	err := me.Client.Call("DB.GetInHexValue", []string{uristring}, &result)
	if err != nil {
		return stderr.ErrUnknown
	}
	*ret = result[uristring]
	return nil
}
