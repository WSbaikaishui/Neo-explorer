package biz

import (
	"encoding/hex"
	"fmt"
	"neophora/cli"
	"neophora/lib/trans"
	"neophora/var/stderr"
	"net/url"
	"regexp"
	"strings"
)

// NeoCli ...
type NeoCli struct {
	Client *cli.T
}

// GETBLOCK ...
func (me *NeoCli) GETBLOCK(args []interface{}, ret *interface{}) error {
	switch len(args) {
	case 1:
		args = append(args, 0.0)
	case 2:
	default:
		return stderr.ErrInvalidArgs
	}

	var uri url.URL

	switch args[1] {
	case 0.0:
		uri.Scheme = "block"
	case 1.0:
		uri.Scheme = "adhocblockinfo"
	default:
		return stderr.ErrInvalidArgs
	}

	switch key := args[0].(type) {
	case float64:
		uri.Host = "height"
		uri.Path = fmt.Sprintf("/%016X", uint64(key))
	case string:
		uri.Host = "hash"
		key = strings.ToUpper(key)
		matches := modNeoCliVarRegHash.FindStringSubmatch(key)
		if len(matches) != 3 {
			return stderr.ErrInvalidArgs
		}
		key = matches[2]
		uri.Path = fmt.Sprintf("/%s", key)
	default:
		return stderr.ErrInvalidArgs
	}

	var result []byte

	if err := me.Client.Calls("DB.Get", []byte(uri.String()), &result); err != nil {
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
	switch len(args) {
	case 1:
		args = append(args, 0.0)
	case 2:
	default:
		return stderr.ErrInvalidArgs
	}

	var uri url.URL

	switch args[1] {
	case 0.0:
		uri.Scheme = "header"
	case 1.0:
		uri.Scheme = "adhocheaderinfo"
	default:
		return stderr.ErrInvalidArgs
	}

	switch key := args[0].(type) {
	case float64:
		uri.Host = "height"
		uri.Path = fmt.Sprintf("/%016X", uint64(key))
	case string:
		uri.Host = "hash"
		key = strings.ToUpper(key)
		matches := modNeoCliVarRegHash.FindStringSubmatch(key)
		if len(matches) != 3 {
			return stderr.ErrInvalidArgs
		}
		key = matches[2]
		uri.Path = fmt.Sprintf("/%s", key)
	default:
		return stderr.ErrInvalidArgs
	}

	var result []byte

	if err := me.Client.Calls("DB.Get", []byte(uri.String()), &result); err != nil {
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
	switch len(args) {
	case 1:
		args = append(args, 0.0)
	case 2:
	default:
		return stderr.ErrInvalidArgs
	}

	var uri url.URL

	switch args[1] {
	case 0.0:
		uri.Scheme = "tx"
	case 1.0:
		uri.Scheme = "adhoctxinfo"
	default:
		return stderr.ErrInvalidArgs
	}

	switch key := args[0].(type) {
	case string:
		uri.Host = "hash"
		key = strings.ToUpper(key)
		matches := modNeoCliVarRegHash.FindStringSubmatch(key)
		if len(matches) != 3 {
			return stderr.ErrInvalidArgs
		}
		key = matches[2]
		uri.Path = fmt.Sprintf("/%s", key)
	default:
		return stderr.ErrInvalidArgs
	}

	var result []byte

	if err := me.Client.Calls("DB.Get", []byte(uri.String()), &result); err != nil {
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
	switch len(args) {
	case 1:
	default:
		return stderr.ErrInvalidArgs
	}

	var uri url.URL

	uri.Scheme = "adhoclog"

	switch key := args[0].(type) {
	case string:
		uri.Host = "hash"
		key = strings.ToUpper(key)
		matches := modNeoCliVarRegHash.FindStringSubmatch(key)
		if len(matches) != 3 {
			return stderr.ErrInvalidArgs
		}
		key = matches[2]
		uri.Path = fmt.Sprintf("/%s", key)
	default:
		return stderr.ErrInvalidArgs
	}

	var result []byte

	if err := me.Client.Calls("DB.Get", []byte(uri.String()), &result); err != nil {
		return stderr.ErrUnknown
	}

	*ret = string(result)

	return nil
}

// GETSTATEROOT ...
func (me *NeoCli) GETSTATEROOT(args []interface{}, ret *interface{}) error {
	switch len(args) {
	case 1:
	default:
		return stderr.ErrInvalidArgs
	}

	var uri url.URL

	uri.Scheme = "adhocstateroot"

	switch key := args[0].(type) {
	case float64:
		uri.Host = "height"
		uri.Path = fmt.Sprintf("/%016X", uint64(key))
	case string:
		uri.Host = "hash"
		key = strings.ToUpper(key)
		matches := modNeoCliVarRegHash.FindStringSubmatch(key)
		if len(matches) != 3 {
			return stderr.ErrInvalidArgs
		}
		key = matches[2]
		uri.Path = fmt.Sprintf("/%s", key)
	default:
		return stderr.ErrInvalidArgs
	}

	var result []byte

	if err := me.Client.Calls("DB.Get", []byte(uri.String()), &result); err != nil {
		return stderr.ErrUnknown
	}

	*ret = string(result)

	return nil
}

// GETBLOCKHASH ...
func (me *NeoCli) GETBLOCKHASH(args []interface{}, ret *interface{}) error {
	switch len(args) {
	case 1:
	default:
		return stderr.ErrInvalidArgs
	}

	var uri url.URL

	uri.Scheme = "hash"

	switch key := args[0].(type) {
	case float64:
		uri.Host = "height"
		uri.Path = fmt.Sprintf("/%016X", uint64(key))
	default:
		return stderr.ErrInvalidArgs
	}

	var result []byte

	if err := me.Client.Calls("DB.Get", []byte(uri.String()), &result); err != nil {
		return stderr.ErrUnknown
	}

	hash := hex.EncodeToString(result)
	*ret = fmt.Sprintf("0x%s", hash)

	return nil
}

// GETBLOCKSYSFEE ...
func (me *NeoCli) GETBLOCKSYSFEE(args []interface{}, ret *interface{}) error {
	switch len(args) {
	case 1:
	default:
		return stderr.ErrInvalidArgs
	}

	var uri url.URL

	uri.Scheme = "adhocsysfee"

	switch key := args[0].(type) {
	case float64:
		uri.Host = "height"
		uri.Path = fmt.Sprintf("/%016X", uint64(key))
	default:
		return stderr.ErrInvalidArgs
	}

	var result []byte

	if err := me.Client.Calls("DB.Get", []byte(uri.String()), &result); err != nil {
		return stderr.ErrUnknown
	}

	*ret = string(result)
	return nil
}

// GETACCOUNTSTATE ...
func (me *NeoCli) GETACCOUNTSTATE(args []interface{}, ret *interface{}) error {
	switch len(args) {
	case 1:
	default:
		return stderr.ErrInvalidArgs
	}

	var uri url.URL

	uri.Scheme = "adhocaccountstate"

	switch key := args[0].(type) {
	case float64:
		uri.Host = "account-height"
		tr := &trans.T{
			V: key,
		}
		if err := tr.AddressToHash(); err != nil {
			return stderr.ErrInvalidArgs
		}
		if err := tr.BytesToHex(); err != nil {
			return stderr.ErrInvalidArgs
		}
		uri.Path = fmt.Sprintf("/%s/FFFFFFFFFFFFFFFF", tr.V)
	default:
		return stderr.ErrInvalidArgs
	}

	var result []byte

	urs := uri.String()
	if err := me.Client.Calls("DB.GetLast", struct {
		Key    []byte
		Prefix int
	}{
		Key:    []byte(urs),
		Prefix: len(urs) - 16,
	}, &result); err != nil {
		return stderr.ErrUnknown
	}

	*ret = string(result)
	return nil
}

// GETASSETSTATE ...
func (me *NeoCli) GETASSETSTATE(args []interface{}, ret *interface{}) error {
	switch len(args) {
	case 1:
	default:
		return stderr.ErrInvalidArgs
	}

	var uri url.URL

	uri.Scheme = "adhocassetstate"

	switch key := args[0].(type) {
	case string:
		uri.Host = "hash-height"
		key = strings.ToUpper(key)
		matches := modNeoCliVarRegHash.FindStringSubmatch(key)
		if len(matches) != 3 {
			return stderr.ErrInvalidArgs
		}
		key = matches[2]
		uri.Path = fmt.Sprintf("/%s/FFFFFFFFFFFFFFFF", key)
	default:
		return stderr.ErrInvalidArgs
	}

	var result []byte

	urs := uri.String()
	if err := me.Client.Calls("DB.GetLast", struct {
		Key    []byte
		Prefix int
	}{
		Key:    []byte(urs),
		Prefix: len(urs) - 16,
	}, &result); err != nil {
		return stderr.ErrUnknown
	}

	*ret = string(result)
	return nil
}

// GETBESTBLOCKHASH ...
func (me *NeoCli) GETBESTBLOCKHASH(args []interface{}, ret *interface{}) error {
	switch len(args) {
	case 0:
	default:
		return stderr.ErrInvalidArgs
	}

	var uri url.URL

	uri.Scheme = "hash"
	uri.Host = "height"
	uri.Path = "/FFFFFFFFFFFFFFFF"

	var result []byte

	urs := uri.String()
	if err := me.Client.Calls("DB.GetLast", struct {
		Key    []byte
		Prefix int
	}{
		Key:    []byte(urs),
		Prefix: len(urs) - 16,
	}, &result); err != nil {
		return stderr.ErrUnknown
	}

	hash := hex.EncodeToString(result)
	*ret = fmt.Sprintf("0x%s", hash)
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

var (
	modNeoCliVarRegHash = regexp.MustCompile(`^(0X)?([0-9A-F]{64})$`)
)
