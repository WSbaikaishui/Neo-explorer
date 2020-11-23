package biz

import (
	"encoding/json"
	"fmt"
	"neophora/cli"
	"neophora/var/stderr"
	"net/url"
)

// NeoCli ...
type NeoCli struct {
	Client *cli.T
}

// PING ...
func (me *NeoCli) PING(arg []interface{}, ret *interface{}) error {
	*ret = "pong"
	return nil
}

// GETBLOCK ...
func (me *NeoCli) GETBLOCK(arg []interface{}, ret *interface{}) error {
	if len(arg) == 0 {
		return stderr.ErrInvalidArgs
	}
	if len(arg) > 2 {
		return stderr.ErrInvalidArgs
	}
	if len(arg) == 1 {
		arg = append(arg, 0)
	}

	var host string
	var path string

	switch arg[0].(type) {
	case float64:
		host = "height"
		path = "/%.0f"
	case string:
		host = "hash"
		path = "/%s"
	default:
		return stderr.ErrInvalidArgs
	}

	switch arg[1] {
	case 0.0:
		var result map[string]interface{}
		uri := &url.URL{
			Scheme: "block",
			Host:   host,
			Path:   fmt.Sprintf(path, arg[0]),
		}
		uristring := uri.String()
		err := me.Client.Call("DB.GetInHexValue", []string{uristring}, &result)
		if err != nil {
			return stderr.ErrUnknown
		}
		*ret = result[uristring]
	case 1.0:
		var result map[string]string
		uri := &url.URL{
			Scheme: "adhocblockinfo",
			Host:   host,
			Path:   fmt.Sprintf(path, arg[0]),
		}
		uristring := uri.String()
		err := me.Client.Call("DB.GetInStringValue", []string{uristring}, &result)
		if err != nil {
			return stderr.ErrUnknown
		}
		*ret = json.RawMessage(result[uristring])
	default:
		return stderr.ErrInvalidArgs
	}

	return nil
}

// GETBLOCKHEADER ...
func (me *NeoCli) GETBLOCKHEADER(arg []interface{}, ret *interface{}) error {
	if len(arg) == 0 {
		return stderr.ErrInvalidArgs
	}
	if len(arg) > 2 {
		return stderr.ErrInvalidArgs
	}
	if len(arg) == 1 {
		arg = append(arg, 0)
	}

	var host string
	var path string

	switch arg[0].(type) {
	case string:
		host = "hash"
		path = "/%s"
	default:
		return stderr.ErrInvalidArgs
	}

	switch arg[1] {
	case 0.0:
		var result map[string]interface{}
		uri := &url.URL{
			Scheme: "header",
			Host:   host,
			Path:   fmt.Sprintf(path, arg[0]),
		}
		uristring := uri.String()
		err := me.Client.Call("DB.GetInHexValue", []string{uristring}, &result)
		if err != nil {
			return stderr.ErrUnknown
		}
		*ret = result[uristring]
	case 1.0:
		var result map[string]string
		uri := &url.URL{
			Scheme: "adhocheaderinfo",
			Host:   host,
			Path:   fmt.Sprintf(path, arg[0]),
		}
		uristring := uri.String()
		err := me.Client.Call("DB.GetInStringValue", []string{uristring}, &result)
		if err != nil {
			return stderr.ErrUnknown
		}
		*ret = json.RawMessage(result[uristring])
	default:
		return stderr.ErrInvalidArgs
	}

	return nil
}

// GETBLOCKHASH ...
func (me *NeoCli) GETBLOCKHASH(arg []interface{}, ret *interface{}) error {
	if len(arg) != 1 {
		return stderr.ErrInvalidArgs
	}

	var host string
	var path string

	switch arg[0].(type) {
	case float64:
		host = "height"
		path = "/%.0f"
	default:
		return stderr.ErrInvalidArgs
	}

	var result map[string]interface{}
	uri := &url.URL{
		Scheme: "hash",
		Host:   host,
		Path:   fmt.Sprintf(path, arg[0]),
	}
	uristring := uri.String()
	err := me.Client.Call("DB.GetInHexValue", []string{uristring}, &result)
	if err != nil {
		return stderr.ErrUnknown
	}
	*ret = result[uristring]
	return nil
}

// GETBLOCKSYSFEE ...
func (me *NeoCli) GETBLOCKSYSFEE(arg []interface{}, ret *interface{}) error {
	if len(arg) != 1 {
		return stderr.ErrInvalidArgs
	}

	var host string
	var path string

	switch arg[0].(type) {
	case float64:
		host = "height"
		path = "/%.0f"
	default:
		return stderr.ErrInvalidArgs
	}

	var result map[string]string
	uri := &url.URL{
		Scheme: "adhocsysfee",
		Host:   host,
		Path:   fmt.Sprintf(path, arg[0]),
	}
	uristring := uri.String()
	err := me.Client.Call("DB.GetInHexValue", []string{uristring}, &result)
	if err != nil {
		return stderr.ErrUnknown
	}
	*ret = result[uristring]
	return nil
}

// GETAPPLICATIONLOG ...
func (me *NeoCli) GETAPPLICATIONLOG(arg []interface{}, ret *interface{}) error {
	*ret = "pong"
	return nil
}

// GETACCOUNTSTATE ...
func (me *NeoCli) GETACCOUNTSTATE(arg []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// CLAIMGAS ...
func (me *NeoCli) CLAIMGAS(arg []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// DUMPPRIVKEY ...
func (me *NeoCli) DUMPPRIVKEY(arg []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// GETBALANCE ...
func (me *NeoCli) GETBALANCE(arg []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}
