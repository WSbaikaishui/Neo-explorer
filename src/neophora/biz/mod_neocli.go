package biz

import (
	"encoding/hex"
	"fmt"
	"log"
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
		arg = append(arg, 0.0)
	}

	var scheme string
	var host string
	var path string
	var result []byte

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
		scheme = "block"
	case 1.0:
		scheme = "adhocblockinfo"
	default:
		return stderr.ErrInvalidArgs
	}

	uri := &url.URL{
		Scheme: scheme,
		Host:   host,
		Path:   fmt.Sprintf(path, arg[0]),
	}
	uristring := uri.String()
	if err := me.Client.Call("DB.Get", []byte(uristring), &result); err != nil {
		log.Println("[BIZ][NEOCLI][GETBLOCK]", uristring, err)
		return stderr.ErrUnknown
	}

	switch arg[1] {
	case 0.0:
		*ret = hex.EncodeToString(result)
	case 1.0:
		*ret = string(result)
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
		arg = append(arg, 0.0)
	}

	var scheme string
	var host string
	var path string
	var result []byte

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
		scheme = "header"
	case 1.0:
		scheme = "adhocheaderinfo"
	default:
		return stderr.ErrInvalidArgs
	}

	uri := &url.URL{
		Scheme: scheme,
		Host:   host,
		Path:   fmt.Sprintf(path, arg[0]),
	}
	uristring := uri.String()
	if err := me.Client.Call("DB.Get", []byte(uristring), &result); err != nil {
		return stderr.ErrUnknown
	}

	switch arg[1] {
	case 0.0:
		*ret = hex.EncodeToString(result)
	case 1.0:
		*ret = string(result)
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
	var result []byte

	switch arg[0].(type) {
	case float64:
		host = "height"
		path = "/%.0f"
	default:
		return stderr.ErrInvalidArgs
	}

	uri := &url.URL{
		Scheme: "hash",
		Host:   host,
		Path:   fmt.Sprintf(path, arg[0]),
	}
	uristring := uri.String()
	if err := me.Client.Call("DB.Get", []byte(uristring), &result); err != nil {
		return stderr.ErrUnknown
	}

	*ret = hex.EncodeToString(result)
	return nil
}

// GETBLOCKSYSFEE ...
func (me *NeoCli) GETBLOCKSYSFEE(arg []interface{}, ret *interface{}) error {
	if len(arg) != 1 {
		return stderr.ErrInvalidArgs
	}

	var host string
	var path string
	var result []byte

	switch arg[0].(type) {
	case float64:
		host = "height"
		path = "/%.0f"
	default:
		return stderr.ErrInvalidArgs
	}

	uri := &url.URL{
		Scheme: "adhocsysfee",
		Host:   host,
		Path:   fmt.Sprintf(path, arg[0]),
	}
	uristring := uri.String()
	if err := me.Client.Call("DB.Get", []byte(uristring), &result); err != nil {
		return stderr.ErrUnknown
	}

	*ret = string(result)
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

// GETCONNECTIONCOUNT ...
func (me *NeoCli) GETCONNECTIONCOUNT(arg []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// GETMETRICBLOCKTIMESTAMP ...
func (me *NeoCli) GETMETRICBLOCKTIMESTAMP(arg []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// GETNEP5TRANSFERS ...
func (me *NeoCli) GETNEP5TRANSFERS(arg []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// GETNEWADDRESS ...
func (me *NeoCli) GETNEWADDRESS(arg []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// GETRAWMEMPOOL ...
func (me *NeoCli) GETRAWMEMPOOL(arg []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// GETPEERS ...
func (me *NeoCli) GETPEERS(arg []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// GETUNCLAIMEDGAS ...
func (me *NeoCli) GETUNCLAIMEDGAS(arg []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// GETVERSION ...
func (me *NeoCli) GETVERSION(arg []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// GETWALLETHEIGHT ...
func (me *NeoCli) GETWALLETHEIGHT(arg []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// IMPORTPRIVKEY ...
func (me *NeoCli) IMPORTPRIVKEY(arg []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// LISTPLUGINS ...
func (me *NeoCli) LISTPLUGINS(arg []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// LISTADDRESS ...
func (me *NeoCli) LISTADDRESS(arg []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// SENDFROM ...
func (me *NeoCli) SENDFROM(arg []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// SENDTOADDRESS ...
func (me *NeoCli) SENDTOADDRESS(arg []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// SENDMANY ...
func (me *NeoCli) SENDMANY(arg []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}
