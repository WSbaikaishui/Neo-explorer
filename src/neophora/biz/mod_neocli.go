package biz

import (
	"encoding/hex"
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
		arg = append(arg, 0.0)
	}

	var scheme string
	var host string
	var path string
	var result []byte

	switch key := arg[0].(type) {
	case float64:
		host = "height"
		path = fmt.Sprintf("/%016X", uint64(key))
	case string:
		host = "hash"
		path = fmt.Sprintf("/%s", key)
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
		Path:   path,
	}
	uristring := uri.String()
	if err := me.Client.Calls("DB.Get", []byte(uristring), &result); err != nil {
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

	switch key := arg[0].(type) {
	case float64:
		host = "height"
		path = fmt.Sprintf("/%016X", uint64(key))
	case string:
		host = "hash"
		path = fmt.Sprintf("/%s", key)
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
		Path:   path,
	}
	uristring := uri.String()
	if err := me.Client.Calls("DB.Get", []byte(uristring), &result); err != nil {
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

// GETRAWTRANSACTION ...
func (me *NeoCli) GETRAWTRANSACTION(arg []interface{}, ret *interface{}) error {
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

	switch key := arg[0].(type) {
	case string:
		host = "hash"
		path = fmt.Sprintf("/%s", key)
	default:
		return stderr.ErrInvalidArgs
	}

	switch arg[1] {
	case 0.0:
		scheme = "tx"
	case 1.0:
		scheme = "adhoctxinfo"
	default:
		return stderr.ErrInvalidArgs
	}

	uri := &url.URL{
		Scheme: scheme,
		Host:   host,
		Path:   path,
	}
	uristring := uri.String()
	if err := me.Client.Calls("DB.Get", []byte(uristring), &result); err != nil {
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

// GETAPPLICATIONLOG ...
func (me *NeoCli) GETAPPLICATIONLOG(arg []interface{}, ret *interface{}) error {
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

	switch key := arg[0].(type) {
	case string:
		host = "hash"
		path = fmt.Sprintf("/%s", key)
	default:
		return stderr.ErrInvalidArgs
	}

	scheme = "adhoclog"

	uri := &url.URL{
		Scheme: scheme,
		Host:   host,
		Path:   path,
	}
	uristring := uri.String()
	if err := me.Client.Calls("DB.Get", []byte(uristring), &result); err != nil {
		return stderr.ErrUnknown
	}

	*ret = string(result)

	return nil
}

// GETSTATEROOT ...
func (me *NeoCli) GETSTATEROOT(arg []interface{}, ret *interface{}) error {
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

	switch key := arg[0].(type) {
	case float64:
		host = "height"
		path = fmt.Sprintf("/%016X", uint64(key))
	case string:
		host = "hash"
		path = fmt.Sprintf("/%s", key)
	default:
		return stderr.ErrInvalidArgs
	}

	scheme = "adhocstateroot"

	uri := &url.URL{
		Scheme: scheme,
		Host:   host,
		Path:   path,
	}
	uristring := uri.String()
	if err := me.Client.Calls("DB.Get", []byte(uristring), &result); err != nil {
		return stderr.ErrUnknown
	}

	*ret = string(result)

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

	switch key := arg[0].(type) {
	case float64:
		host = "height"
		path = fmt.Sprintf("/%016X", uint64(key))
	default:
		return stderr.ErrInvalidArgs
	}

	uri := &url.URL{
		Scheme: "hash",
		Host:   host,
		Path:   path,
	}
	uristring := uri.String()
	if err := me.Client.Calls("DB.Get", []byte(uristring), &result); err != nil {
		return stderr.ErrUnknown
	}

	hash := hex.EncodeToString(result)
	*ret = fmt.Sprintf("0x%s", hash)
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

	switch key := arg[0].(type) {
	case float64:
		host = "height"
		path = fmt.Sprintf("/%016X", uint64(key))
	default:
		return stderr.ErrInvalidArgs
	}

	uri := &url.URL{
		Scheme: "adhocsysfee",
		Host:   host,
		Path:   path,
	}
	uristring := uri.String()
	if err := me.Client.Calls("DB.Get", []byte(uristring), &result); err != nil {
		return stderr.ErrUnknown
	}

	*ret = string(result)
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
