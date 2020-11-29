package biz

import (
	"encoding/hex"
	"fmt"
	"neophora/cli"
	"neophora/lib/trans"
	"neophora/var/stderr"
	"net/url"
)

// NeoCli ...
type NeoCli struct {
	Client *cli.T
}

// GETBLOCK ...
func (me *NeoCli) GETBLOCK(args []interface{}, ret *interface{}) error {
	if len(args) == 0 {
		return stderr.ErrInvalidArgs
	}
	if len(args) > 2 {
		return stderr.ErrInvalidArgs
	}
	if len(args) == 1 {
		args = append(args, 0.0)
	}

	var scheme string
	var host string
	var path string
	var result []byte

	switch key := args[0].(type) {
	case float64:
		host = "height"
		path = fmt.Sprintf("/%016X", uint64(key))
	case string:
		host = "hash"
		path = fmt.Sprintf("/%s", key)
	default:
		return stderr.ErrInvalidArgs
	}

	switch args[1] {
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

	switch args[1] {
	case 0.0:
		*ret = hex.EncodeToString(result)
	case 1.0:
		*ret = string(result)
	}

	return nil
}

// GETBLOCKHEADER ...
func (me *NeoCli) GETBLOCKHEADER(args []interface{}, ret *interface{}) error {
	if len(args) == 0 {
		return stderr.ErrInvalidArgs
	}
	if len(args) > 2 {
		return stderr.ErrInvalidArgs
	}
	if len(args) == 1 {
		args = append(args, 0.0)
	}

	var scheme string
	var host string
	var path string
	var result []byte

	switch key := args[0].(type) {
	case float64:
		host = "height"
		path = fmt.Sprintf("/%016X", uint64(key))
	case string:
		host = "hash"
		path = fmt.Sprintf("/%s", key)
	default:
		return stderr.ErrInvalidArgs
	}

	switch args[1] {
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

	switch args[1] {
	case 0.0:
		*ret = hex.EncodeToString(result)
	case 1.0:
		*ret = string(result)
	}

	return nil
}

// GETRAWTRANSACTION ...
func (me *NeoCli) GETRAWTRANSACTION(args []interface{}, ret *interface{}) error {
	if len(args) == 0 {
		return stderr.ErrInvalidArgs
	}
	if len(args) > 2 {
		return stderr.ErrInvalidArgs
	}
	if len(args) == 1 {
		args = append(args, 0.0)
	}

	var scheme string
	var host string
	var path string
	var result []byte

	switch key := args[0].(type) {
	case string:
		host = "hash"
		path = fmt.Sprintf("/%s", key)
	default:
		return stderr.ErrInvalidArgs
	}

	switch args[1] {
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

	switch args[1] {
	case 0.0:
		*ret = hex.EncodeToString(result)
	case 1.0:
		*ret = string(result)
	}

	return nil
}

// GETAPPLICATIONLOG ...
func (me *NeoCli) GETAPPLICATIONLOG(args []interface{}, ret *interface{}) error {
	if len(args) == 0 {
		return stderr.ErrInvalidArgs
	}
	if len(args) > 2 {
		return stderr.ErrInvalidArgs
	}
	if len(args) == 1 {
		args = append(args, 0.0)
	}

	var scheme string
	var host string
	var path string
	var result []byte

	switch key := args[0].(type) {
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
func (me *NeoCli) GETSTATEROOT(args []interface{}, ret *interface{}) error {
	if len(args) == 0 {
		return stderr.ErrInvalidArgs
	}
	if len(args) > 2 {
		return stderr.ErrInvalidArgs
	}
	if len(args) == 1 {
		args = append(args, 0.0)
	}

	var scheme string
	var host string
	var path string
	var result []byte

	switch key := args[0].(type) {
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
func (me *NeoCli) GETBLOCKHASH(args []interface{}, ret *interface{}) error {
	if len(args) != 1 {
		return stderr.ErrInvalidArgs
	}

	var host string
	var path string
	var result []byte

	switch key := args[0].(type) {
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
func (me *NeoCli) GETBLOCKSYSFEE(args []interface{}, ret *interface{}) error {
	if len(args) != 1 {
		return stderr.ErrInvalidArgs
	}

	var host string
	var path string
	var result []byte

	switch key := args[0].(type) {
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
func (me *NeoCli) GETACCOUNTSTATE(args []interface{}, ret *interface{}) error {
	if len(args) != 1 {
		return stderr.ErrInvalidArgs
	}

	var host string
	var path string
	var result []byte

	switch key := args[0].(type) {
	case string:
		host = "account-height"
		tr := &trans.T{
			V: key,
		}
		if err := tr.AddressToHash(); err != nil {
			return stderr.ErrInvalidArgs
		}
		if err := tr.BytesToHex(); err != nil {
			return stderr.ErrInvalidArgs
		}
		path = fmt.Sprintf("/%s/FFFFFFFFFFFFFFFF", tr.V)
	default:
		return stderr.ErrInvalidArgs
	}

	uri := &url.URL{
		Scheme: "adhocaccountstate",
		Host:   host,
		Path:   path,
	}
	uristring := uri.String()
	if err := me.Client.Calls("DB.GetLast", struct {
		Key    []byte
		Prefix int
	}{
		Key:    []byte(uristring),
		Prefix: len(uristring) - 16,
	}, &result); err != nil {
		return stderr.ErrUnknown
	}

	*ret = string(result)
	return nil
}

// GETASSETSTATE ...
func (me *NeoCli) GETASSETSTATE(args []interface{}, ret *interface{}) error {
	if len(args) != 1 {
		return stderr.ErrInvalidArgs
	}

	var host string
	var path string
	var result []byte

	switch key := args[0].(type) {
	case string:
		host = "hash-height"
		path = fmt.Sprintf("/%s/FFFFFFFFFFFFFFFF", key)
	default:
		return stderr.ErrInvalidArgs
	}

	uri := &url.URL{
		Scheme: "adhocassetstate",
		Host:   host,
		Path:   path,
	}
	uristring := uri.String()
	if err := me.Client.Calls("DB.GetLast", struct {
		Key    []byte
		Prefix int
	}{
		Key:    []byte(uristring),
		Prefix: len(uristring) - 16,
	}, &result); err != nil {
		return stderr.ErrUnknown
	}

	*ret = string(result)
	return nil
}

// PING ...
func (me *NeoCli) PING(args []interface{}, ret *interface{}) error {
	*ret = "pong"
	return nil
}

// CLAIMGAS ...
func (me *NeoCli) CLAIMGAS(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// DUMPPRIVKEY ...
func (me *NeoCli) DUMPPRIVKEY(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// GETBALANCE ...
func (me *NeoCli) GETBALANCE(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// GETCONNECTIONCOUNT ...
func (me *NeoCli) GETCONNECTIONCOUNT(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// GETMETRICBLOCKTIMESTAMP ...
func (me *NeoCli) GETMETRICBLOCKTIMESTAMP(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// GETNEP5TRANSFERS ...
func (me *NeoCli) GETNEP5TRANSFERS(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// GETNEWADDRESS ...
func (me *NeoCli) GETNEWADDRESS(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// GETRAWMEMPOOL ...
func (me *NeoCli) GETRAWMEMPOOL(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// GETPEERS ...
func (me *NeoCli) GETPEERS(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// GETUNCLAIMEDGAS ...
func (me *NeoCli) GETUNCLAIMEDGAS(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// GETVERSION ...
func (me *NeoCli) GETVERSION(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// GETWALLETHEIGHT ...
func (me *NeoCli) GETWALLETHEIGHT(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// IMPORTPRIVKEY ...
func (me *NeoCli) IMPORTPRIVKEY(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// LISTPLUGINS ...
func (me *NeoCli) LISTPLUGINS(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// LISTADDRESS ...
func (me *NeoCli) LISTADDRESS(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// SENDFROM ...
func (me *NeoCli) SENDFROM(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// SENDTOADDRESS ...
func (me *NeoCli) SENDTOADDRESS(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// SENDMANY ...
func (me *NeoCli) SENDMANY(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}
