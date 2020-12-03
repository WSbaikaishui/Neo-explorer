package biz

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"neophora/cli"
	"neophora/lib/trans"
	"neophora/var/stderr"
	"net/url"
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
		uri.Path = fmt.Sprintf("/%016x", uint64(key))
	case string:
		uri.Host = "hash"
		tr := &trans.T{
			V: key,
		}
		if err := tr.StringToLowerCase(); err != nil {
			return stderr.ErrInvalidArgs
		}
		if err := tr.Remove0xPrefix(); err != nil {
			return stderr.ErrInvalidArgs
		}
		if err := tr.HexToBytes(); err != nil {
			return stderr.ErrInvalidArgs
		}
		if err := tr.BytesReverse(); err != nil {
			return stderr.ErrInvalidArgs
		}
		if err := tr.BytesToHex(); err != nil {
			return stderr.ErrInvalidArgs
		}
		uri.Path = fmt.Sprintf("/%s", tr.V)
	default:
		return stderr.ErrInvalidArgs
	}

	var result []byte

	urs := uri.String()
	if err := me.Client.Calls("DB.Get", []byte(urs), &result); err != nil {
		return stderr.ErrUnknown
	}

	if len(result) == 0 {
		return stderr.ErrNotFound
	}

	switch args[1] {
	case 0.0:
		*ret = hex.EncodeToString(result)
	case 1.0:
		*ret = json.RawMessage(result)
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
		uri.Path = fmt.Sprintf("/%016x", uint64(key))
	case string:
		uri.Host = "hash"
		tr := &trans.T{
			V: key,
		}
		if err := tr.StringToLowerCase(); err != nil {
			return stderr.ErrInvalidArgs
		}
		if err := tr.Remove0xPrefix(); err != nil {
			return stderr.ErrInvalidArgs
		}
		if err := tr.HexToBytes(); err != nil {
			return stderr.ErrInvalidArgs
		}
		if err := tr.BytesReverse(); err != nil {
			return stderr.ErrInvalidArgs
		}
		if err := tr.BytesToHex(); err != nil {
			return stderr.ErrInvalidArgs
		}
		uri.Path = fmt.Sprintf("/%s", tr.V)
	default:
		return stderr.ErrInvalidArgs
	}

	var result []byte

	urs := uri.String()
	if err := me.Client.Calls("DB.Get", []byte(urs), &result); err != nil {
		return stderr.ErrUnknown
	}

	if len(result) == 0 {
		return stderr.ErrNotFound
	}

	switch args[1] {
	case 0.0:
		*ret = hex.EncodeToString(result)
	case 1.0:
		*ret = json.RawMessage(result)
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
		tr := &trans.T{
			V: key,
		}
		if err := tr.StringToLowerCase(); err != nil {
			return stderr.ErrInvalidArgs
		}
		if err := tr.Remove0xPrefix(); err != nil {
			return stderr.ErrInvalidArgs
		}
		if err := tr.HexToBytes(); err != nil {
			return stderr.ErrInvalidArgs
		}
		if err := tr.BytesReverse(); err != nil {
			return stderr.ErrInvalidArgs
		}
		if err := tr.BytesToHex(); err != nil {
			return stderr.ErrInvalidArgs
		}
		uri.Path = fmt.Sprintf("/%s", tr.V)
	default:
		return stderr.ErrInvalidArgs
	}

	var result []byte

	urs := uri.String()
	if err := me.Client.Calls("DB.Get", []byte(urs), &result); err != nil {
		return stderr.ErrUnknown
	}

	if len(result) == 0 {
		return stderr.ErrNotFound
	}

	switch args[1] {
	case 0.0:
		*ret = hex.EncodeToString(result)
	case 1.0:
		*ret = json.RawMessage(result)
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
		tr := &trans.T{
			V: key,
		}
		if err := tr.StringToLowerCase(); err != nil {
			return stderr.ErrInvalidArgs
		}
		if err := tr.Remove0xPrefix(); err != nil {
			return stderr.ErrInvalidArgs
		}
		if err := tr.HexToBytes(); err != nil {
			return stderr.ErrInvalidArgs
		}
		if err := tr.BytesReverse(); err != nil {
			return stderr.ErrInvalidArgs
		}
		if err := tr.BytesToHex(); err != nil {
			return stderr.ErrInvalidArgs
		}
		uri.Path = fmt.Sprintf("/%s", tr.V)
	default:
		return stderr.ErrInvalidArgs
	}

	var result []byte

	urs := uri.String()
	if err := me.Client.Calls("DB.Get", []byte(urs), &result); err != nil {
		return stderr.ErrUnknown
	}

	if len(result) == 0 {
		return stderr.ErrNotFound
	}

	*ret = json.RawMessage(result)

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
		uri.Path = fmt.Sprintf("/%016x", uint64(key))
	case string:
		uri.Host = "hash"
		tr := &trans.T{
			V: key,
		}
		if err := tr.StringToLowerCase(); err != nil {
			return stderr.ErrInvalidArgs
		}
		if err := tr.Remove0xPrefix(); err != nil {
			return stderr.ErrInvalidArgs
		}
		if err := tr.HexToBytes(); err != nil {
			return stderr.ErrInvalidArgs
		}
		if err := tr.BytesReverse(); err != nil {
			return stderr.ErrInvalidArgs
		}
		if err := tr.BytesToHex(); err != nil {
			return stderr.ErrInvalidArgs
		}
		uri.Path = fmt.Sprintf("/%s", tr.V)
	default:
		return stderr.ErrInvalidArgs
	}

	var result []byte

	urs := uri.String()
	if err := me.Client.Calls("DB.Get", []byte(urs), &result); err != nil {
		return stderr.ErrUnknown
	}

	if len(result) == 0 {
		return stderr.ErrNotFound
	}

	*ret = json.RawMessage(result)

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
		uri.Path = fmt.Sprintf("/%016x", uint64(key))
	default:
		return stderr.ErrInvalidArgs
	}

	var result []byte

	urs := uri.String()
	if err := me.Client.Calls("DB.Get", []byte(urs), &result); err != nil {
		return stderr.ErrUnknown
	}

	if len(result) == 0 {
		return stderr.ErrNotFound
	}

	tr := &trans.T{
		V: result,
	}

	if err := tr.BytesReverse(); err != nil {
		return stderr.ErrUnknown
	}
	if err := tr.BytesToHex(); err != nil {
		return stderr.ErrUnknown
	}
	*ret = fmt.Sprintf("0x%s", tr.V)

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
		uri.Path = fmt.Sprintf("/%016x", uint64(key))
	default:
		return stderr.ErrInvalidArgs
	}

	var result []byte

	urs := uri.String()
	if err := me.Client.Calls("DB.Get", []byte(urs), &result); err != nil {
		return stderr.ErrUnknown
	}

	if len(result) == 0 {
		return stderr.ErrNotFound
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
	case string:
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
		uri.Path = fmt.Sprintf("/%s/ffffffffffffffff", tr.V)
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

	if len(result) == 0 {
		return stderr.ErrNotFound
	}

	*ret = json.RawMessage(result)
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
		tr := &trans.T{
			V: key,
		}
		if err := tr.StringToLowerCase(); err != nil {
			return stderr.ErrInvalidArgs
		}
		if err := tr.Remove0xPrefix(); err != nil {
			return stderr.ErrInvalidArgs
		}
		if err := tr.HexToBytes(); err != nil {
			return stderr.ErrInvalidArgs
		}
		if err := tr.BytesReverse(); err != nil {
			return stderr.ErrInvalidArgs
		}
		if err := tr.BytesToHex(); err != nil {
			return stderr.ErrInvalidArgs
		}
		uri.Path = fmt.Sprintf("/%s/ffffffffffffffff", tr.V)
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

	if len(result) == 0 {
		return stderr.ErrNotFound
	}

	*ret = json.RawMessage(result)
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
	uri.Path = "/ffffffffffffffff"

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

	if len(result) == 0 {
		return stderr.ErrNotFound
	}

	tr := &trans.T{
		V: result,
	}
	if err := tr.BytesReverse(); err != nil {
		return stderr.ErrUnknown
	}
	if err := tr.BytesToHex(); err != nil {
		return stderr.ErrUnknown
	}
	*ret = fmt.Sprintf("0x%s", tr.V)
	return nil
}

// GETBLOCKCOUNT ...
func (me *NeoCli) GETBLOCKCOUNT(args []interface{}, ret *interface{}) error {
	switch len(args) {
	case 0:
	default:
		return stderr.ErrInvalidArgs
	}

	var uri url.URL

	uri.Scheme = "hash"
	uri.Host = "height"
	uri.Path = "/ffffffffffffffff"

	var result []byte

	urs := uri.String()
	if err := me.Client.Calls("DB.GetLastKey", struct {
		Key    []byte
		Prefix int
	}{
		Key:    []byte(urs),
		Prefix: len(urs) - 16,
	}, &result); err != nil {
		return stderr.ErrUnknown
	}

	if len(result) == 0 {
		return stderr.ErrNotFound
	}

	key, err := url.Parse(string(result))
	if err != nil {
		return stderr.ErrUnknown
	}

	var count uint64

	fmt.Sscanf(key.Path, "/%x", &count)

	*ret = count
	return nil
}

// GETCLAIMABLE ...
func (me *NeoCli) GETCLAIMABLE(args []interface{}, ret *interface{}) error {
	switch len(args) {
	case 1:
	default:
		return stderr.ErrInvalidArgs
	}

	var uri url.URL

	uri.Scheme = "adhocclaimable"

	switch key := args[0].(type) {
	case string:
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
		uri.Path = fmt.Sprintf("/%s/ffffffffffffffff", tr.V)
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

	if len(result) == 0 {
		return stderr.ErrNotFound
	}

	*ret = json.RawMessage(result)
	return nil
}

// GETCONTRACTSTATE ...
func (me *NeoCli) GETCONTRACTSTATE(args []interface{}, ret *interface{}) error {
	switch len(args) {
	case 1:
	default:
		return stderr.ErrInvalidArgs
	}

	var uri url.URL

	uri.Scheme = "adhoccontractstate"

	switch key := args[0].(type) {
	case string:
		uri.Host = "hash-height"
		tr := &trans.T{
			V: key,
		}
		if err := tr.StringToLowerCase(); err != nil {
			return stderr.ErrInvalidArgs
		}
		if err := tr.Remove0xPrefix(); err != nil {
			return stderr.ErrInvalidArgs
		}
		if err := tr.HexToBytes(); err != nil {
			return stderr.ErrInvalidArgs
		}
		if err := tr.BytesReverse(); err != nil {
			return stderr.ErrInvalidArgs
		}
		if err := tr.BytesToHex(); err != nil {
			return stderr.ErrInvalidArgs
		}
		uri.Path = fmt.Sprintf("/%s/ffffffffffffffff", tr.V)
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

	if len(result) == 0 {
		return stderr.ErrNotFound
	}

	*ret = json.RawMessage(result)
	return nil
}

// GETNEP5BALANCES ...
func (me *NeoCli) GETNEP5BALANCES(args []interface{}, ret *interface{}) error {
	switch len(args) {
	case 1:
	default:
		return stderr.ErrInvalidArgs
	}

	var uri url.URL

	uri.Scheme = "adhocnep5balances"

	switch key := args[0].(type) {
	case string:
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
		uri.Path = fmt.Sprintf("/%s/ffffffffffffffff", tr.V)
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

	if len(result) == 0 {
		return stderr.ErrNotFound
	}

	*ret = json.RawMessage(result)
	return nil
}

// GETSTORAGE ...
func (me *NeoCli) GETSTORAGE(args []interface{}, ret *interface{}) error {
	switch len(args) {
	case 2:
	default:
		return stderr.ErrInvalidArgs
	}

	var uri url.URL

	uri.Scheme = "storage"
	uri.Host = "dbkey-height"

	switch key := args[0].(type) {
	case string:
		key = strings.ToLower(key)
		tr := &trans.T{
			V: key,
		}
		if err := tr.StringToLowerCase(); err != nil {
			return stderr.ErrInvalidArgs
		}
		if err := tr.Remove0xPrefix(); err != nil {
			return stderr.ErrInvalidArgs
		}
		if err := tr.HexToBytes(); err != nil {
			return stderr.ErrInvalidArgs
		}
		if err := tr.BytesReverse(); err != nil {
			return stderr.ErrInvalidArgs
		}
		if err := tr.BytesToHex(); err != nil {
			return stderr.ErrInvalidArgs
		}
		uri.Path = fmt.Sprint(tr.V)
	default:
		return stderr.ErrInvalidArgs
	}

	switch key := args[1].(type) {
	case string:
		uri.Path = fmt.Sprintf("/%s%s/ffffffffffffffff", uri.Path, key)
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

	if len(result) == 0 {
		return stderr.ErrNotFound
	}

	*ret = hex.EncodeToString(result)
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

// GETTRANSACTIONHEIGHT is ...
func (me *NeoCli) GETTRANSACTIONHEIGHT(args []interface{}, ret *interface{}) error {
	switch len(args) {
	case 1:
	default:
		return stderr.ErrInvalidArgs
	}

	var uri url.URL

	uri.Scheme = "height"

	switch key := args[0].(type) {
	case string:
		uri.Host = "hash"
		tr := &trans.T{
			V: key,
		}
		if err := tr.StringToLowerCase(); err != nil {
			return stderr.ErrInvalidArgs
		}
		if err := tr.Remove0xPrefix(); err != nil {
			return stderr.ErrInvalidArgs
		}
		if err := tr.HexToBytes(); err != nil {
			return stderr.ErrInvalidArgs
		}
		if err := tr.BytesReverse(); err != nil {
			return stderr.ErrInvalidArgs
		}
		if err := tr.BytesToHex(); err != nil {
			return stderr.ErrInvalidArgs
		}
		uri.Path = fmt.Sprintf("/%s", tr.V)
	default:
		return stderr.ErrInvalidArgs
	}

	var result []byte

	urs := uri.String()
	if err := me.Client.Calls("DB.Get", []byte(urs), &result); err != nil {
		return stderr.ErrUnknown
	}

	if len(result) == 0 {
		return stderr.ErrNotFound
	}

	*ret = json.RawMessage(result)

	return nil
}

// GETUNSPENTS is ...
func (me *NeoCli) GETUNSPENTS(args []interface{}, ret *interface{}) {
	switch len(args) {
	case 1:
	default:
		return stderr.ErrInvalidArgs
	}

	var uri url.URL

	uri.Scheme = "adhocunspents"

	switch key := args[0].(type) {
	case string:
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
		uri.Path = fmt.Sprintf("/%s/ffffffffffffffff", tr.V)
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

	if len(result) == 0 {
		return stderr.ErrNotFound
	}

	*ret = json.RawMessage(result)
	return nil
}

// GETCONNECTIONCOUNT ...
func (me *NeoCli) GETCONNECTIONCOUNT(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// GETMETRICBLOCKTIMESTAMP ...
func (me *NeoCli) GETMETRICBLOCKTIMESTAMP(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// GETUTXOTRANSFERS ...
func (me *NeoCli) GETUTXOTRANSFERS(args []interface{}, ret *interface{}) error {
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

// GETSTATEHEIGHT ...
func (me *NeoCli) GETSTATEHEIGHT(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// GETTXOUT ...
func (me *NeoCli) GETTXOUT(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// SUBMITBLOCK ...
func (me *NeoCli) SUBMITBLOCK(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// GETVALIDATORS ...
func (me *NeoCli) GETVALIDATORS(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}
