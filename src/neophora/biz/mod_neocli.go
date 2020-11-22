package biz

import (
	"errors"
	"fmt"
	"neophora/cli"
	"neophora/var/stderr"
	"net/url"
)

// NeoCli ...
type NeoCli struct {
	Client *cli.T
}

// GETBLOCK ...
func (me *NeoCli) GETBLOCK(arg []interface{}, ret *interface{}) error {
	switch len(arg) {
	case 1:
		switch key := arg[0].(type) {
		case float64:
			var result map[string]interface{}
			uri := &url.URL{
				Scheme: "bin",
				Host:   "block",
				Path:   fmt.Sprintf("/%.0f", key),
			}
			uristring := uri.String()
			err := me.Client.Call("DB.GetInHexValue", []string{uristring}, &result)
			if err != nil {
				return stderr.ErrUnknown
			}
			*ret = result[uristring]
			return nil
		default:
			return errors.New("invalid args")
		}
	default:
		return errors.New("invalid args")
	}
}

// PING ...
func (me *NeoCli) PING(arg []interface{}, ret *interface{}) error {
	*ret = "pong"
	return nil
}
