package neocli

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"neophora/lib/trans"
	"neophora/var/stderr"
	"net/url"

	"github.com/gomodule/redigo/redis"
)

// T ...
type T struct {
	DB *redis.Pool
}

// GETBLOCK ...
func (me *T) GETBLOCK(args []interface{}, ret *interface{}) error {
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

	reply, err := me.DB.Get().Do("GET", uri.String())
	if err != nil {
		return stderr.ErrUnknown
	}

	if reply == nil {
		return stderr.ErrNotFound
	}

	result, ok := reply.([]byte)
	if ok == false {
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
func (me *T) GETBLOCKHEADER(args []interface{}, ret *interface{}) error {
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

	reply, err := me.DB.Get().Do("GET", uri.String())
	if err != nil {
		return stderr.ErrUnknown
	}

	if reply == nil {
		return stderr.ErrNotFound
	}

	result, ok := reply.([]byte)
	if ok == false {
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
func (me *T) GETRAWTRANSACTION(args []interface{}, ret *interface{}) error {
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

	reply, err := me.DB.Get().Do("GET", uri.String())
	if err != nil {
		return stderr.ErrUnknown
	}

	if reply == nil {
		return stderr.ErrNotFound
	}

	result, ok := reply.([]byte)
	if ok == false {
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
func (me *T) GETAPPLICATIONLOG(args []interface{}, ret *interface{}) error {
	switch len(args) {
	case 1:
	default:
		return stderr.ErrInvalidArgs
	}

	var uri url.URL

	uri.Scheme = "adhoclog"

	switch key := args[0].(type) {
	case string:
		uri.Host = "tx"
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

	reply, err := me.DB.Get().Do("GET", uri.String())
	if err != nil {
		return stderr.ErrUnknown
	}

	if reply == nil {
		return stderr.ErrNotFound
	}

	result, ok := reply.([]byte)
	if ok == false {
		return stderr.ErrUnknown
	}

	if len(result) == 0 {
		return stderr.ErrNotFound
	}

	*ret = json.RawMessage(result)

	return nil
}

// GETSTATEROOT ...
func (me *T) GETSTATEROOT(args []interface{}, ret *interface{}) error {
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

	reply, err := me.DB.Get().Do("GET", uri.String())
	if err != nil {
		return stderr.ErrUnknown
	}

	if reply == nil {
		return stderr.ErrNotFound
	}

	result, ok := reply.([]byte)
	if ok == false {
		return stderr.ErrUnknown
	}

	if len(result) == 0 {
		return stderr.ErrNotFound
	}

	*ret = json.RawMessage(result)

	return nil
}

// GETBLOCKHASH ...
func (me *T) GETBLOCKHASH(args []interface{}, ret *interface{}) error {
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

	reply, err := me.DB.Get().Do("GET", uri.String())
	if err != nil {
		return stderr.ErrUnknown
	}

	if reply == nil {
		return stderr.ErrNotFound
	}

	result, ok := reply.([]byte)
	if ok == false {
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
func (me *T) GETBLOCKSYSFEE(args []interface{}, ret *interface{}) error {
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

	reply, err := me.DB.Get().Do("GET", uri.String())
	if err != nil {
		return stderr.ErrUnknown
	}

	if reply == nil {
		return stderr.ErrNotFound
	}

	result, ok := reply.([]byte)
	if ok == false {
		return stderr.ErrUnknown
	}

	if len(result) == 0 {
		return stderr.ErrNotFound
	}

	*ret = string(result)
	return nil
}

// GETACCOUNTSTATE ...
func (me *T) GETACCOUNTSTATE(args []interface{}, ret *interface{}) error {
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
		uri.Path = fmt.Sprintf("/%s/_", tr.V)
	default:
		return stderr.ErrInvalidArgs
	}

	uri.Path = uri.String()
	uri.Host = "keys"
	uri.Scheme = "index"

	reply, err := me.DB.Get().Do("ZRANGE", uri.String(), -1, -1)
	if err != nil {
		return stderr.ErrUnknown
	}

	if reply == nil {
		return stderr.ErrNotFound
	}

	replies, ok := reply.([]interface{})
	if ok == false {
		return stderr.ErrUnknown
	}

	if len(replies) != 1 {
		return stderr.ErrNotFound
	}

	result, ok := replies[0].([]byte)
	if ok == false {
		return stderr.ErrUnknown
	}

	reply, err = me.DB.Get().Do("GET", result)
	if err != nil {
		return stderr.ErrUnknown
	}

	if reply == nil {
		return stderr.ErrNotFound
	}

	result, ok = reply.([]byte)
	if ok == false {
		return stderr.ErrUnknown
	}

	if len(result) == 0 {
		return stderr.ErrNotFound
	}

	*ret = json.RawMessage(result)
	return nil
}

// GETASSETSTATE ...
func (me *T) GETASSETSTATE(args []interface{}, ret *interface{}) error {
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
		uri.Path = fmt.Sprintf("/%s/_", tr.V)
	default:
		return stderr.ErrInvalidArgs
	}

	uri.Path = uri.String()
	uri.Host = "keys"
	uri.Scheme = "index"

	reply, err := me.DB.Get().Do("ZRANGE", uri.String(), -1, -1)
	if err != nil {
		return stderr.ErrUnknown
	}

	if reply == nil {
		return stderr.ErrNotFound
	}

	replies, ok := reply.([]interface{})
	if ok == false {
		return stderr.ErrUnknown
	}

	if len(replies) != 1 {
		return stderr.ErrNotFound
	}

	result, ok := replies[0].([]byte)
	if ok == false {
		return stderr.ErrUnknown
	}

	reply, err = me.DB.Get().Do("GET", result)
	if err != nil {
		return stderr.ErrUnknown
	}

	if reply == nil {
		return stderr.ErrNotFound
	}

	result, ok = reply.([]byte)
	if ok == false {
		return stderr.ErrUnknown
	}

	if len(result) == 0 {
		return stderr.ErrNotFound
	}

	*ret = json.RawMessage(result)
	return nil
}

// GETBESTBLOCKHASH ...
func (me *T) GETBESTBLOCKHASH(args []interface{}, ret *interface{}) error {
	switch len(args) {
	case 0:
	default:
		return stderr.ErrInvalidArgs
	}

	var uri url.URL

	uri.Scheme = "hash"
	uri.Host = "height"
	uri.Path = "/_"

	uri.Path = uri.String()
	uri.Host = "keys"
	uri.Scheme = "index"

	reply, err := me.DB.Get().Do("ZRANGE", uri.String(), -1, -1)
	if err != nil {
		return stderr.ErrUnknown
	}

	if reply == nil {
		return stderr.ErrNotFound
	}

	replies, ok := reply.([]interface{})
	if ok == false {
		return stderr.ErrUnknown
	}

	if len(replies) != 1 {
		return stderr.ErrNotFound
	}

	result, ok := replies[0].([]byte)
	if ok == false {
		return stderr.ErrUnknown
	}

	reply, err = me.DB.Get().Do("GET", result)
	if err != nil {
		return stderr.ErrUnknown
	}

	if reply == nil {
		return stderr.ErrNotFound
	}

	result, ok = reply.([]byte)
	if ok == false {
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
func (me *T) GETBLOCKCOUNT(args []interface{}, ret *interface{}) error {
	switch len(args) {
	case 0:
	default:
		return stderr.ErrInvalidArgs
	}

	var uri url.URL

	uri.Scheme = "hash"
	uri.Host = "height"
	uri.Path = "/_"

	uri.Path = uri.String()
	uri.Host = "keys"
	uri.Scheme = "index"

	reply, err := me.DB.Get().Do("ZRANGE", uri.String(), -1, -1)
	if err != nil {
		return stderr.ErrUnknown
	}

	if reply == nil {
		return stderr.ErrNotFound
	}

	replies, ok := reply.([]interface{})
	if ok == false {
		return stderr.ErrUnknown
	}

	if len(replies) != 1 {
		return stderr.ErrNotFound
	}

	result, ok := replies[0].([]byte)
	if ok == false {
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
func (me *T) GETCLAIMABLE(args []interface{}, ret *interface{}) error {
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
		uri.Path = fmt.Sprintf("/%s/_", tr.V)
	default:
		return stderr.ErrInvalidArgs
	}

	uri.Path = uri.String()
	uri.Host = "keys"
	uri.Scheme = "index"

	reply, err := me.DB.Get().Do("ZRANGE", uri.String(), -1, -1)
	if err != nil {
		return stderr.ErrUnknown
	}

	if reply == nil {
		return stderr.ErrNotFound
	}

	replies, ok := reply.([]interface{})
	if ok == false {
		return stderr.ErrUnknown
	}

	if len(replies) != 1 {
		return stderr.ErrNotFound
	}

	result, ok := replies[0].([]byte)
	if ok == false {
		return stderr.ErrUnknown
	}

	reply, err = me.DB.Get().Do("GET", result)
	if err != nil {
		return stderr.ErrUnknown
	}

	if reply == nil {
		return stderr.ErrNotFound
	}

	result, ok = reply.([]byte)
	if ok == false {
		return stderr.ErrUnknown
	}

	if len(result) == 0 {
		return stderr.ErrNotFound
	}

	*ret = json.RawMessage(result)
	return nil
}

// GETCONTRACTSTATE ...
func (me *T) GETCONTRACTSTATE(args []interface{}, ret *interface{}) error {
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
		uri.Path = fmt.Sprintf("/%s/_", tr.V)
	default:
		return stderr.ErrInvalidArgs
	}

	uri.Path = uri.String()
	uri.Host = "keys"
	uri.Scheme = "index"

	reply, err := me.DB.Get().Do("ZRANGE", uri.String(), -1, -1)
	if err != nil {
		return stderr.ErrUnknown
	}

	if reply == nil {
		return stderr.ErrNotFound
	}

	replies, ok := reply.([]interface{})
	if ok == false {
		return stderr.ErrUnknown
	}

	if len(replies) != 1 {
		return stderr.ErrNotFound
	}

	result, ok := replies[0].([]byte)
	if ok == false {
		return stderr.ErrUnknown
	}

	reply, err = me.DB.Get().Do("GET", result)
	if err != nil {
		return stderr.ErrUnknown
	}

	if reply == nil {
		return stderr.ErrNotFound
	}

	result, ok = reply.([]byte)
	if ok == false {
		return stderr.ErrUnknown
	}

	if len(result) == 0 {
		return stderr.ErrNotFound
	}

	*ret = json.RawMessage(result)
	return nil
}

// // GETNEP5BALANCES ...
// func (me *T) GETNEP5BALANCES(args []interface{}, ret *interface{}) error {
// 	switch len(args) {
// 	case 1:
// 	default:
// 		return stderr.ErrInvalidArgs
// 	}

// 	var uri url.URL

// 	uri.Scheme = "adhocnep5balances"

// 	switch key := args[0].(type) {
// 	case string:
// 		uri.Host = "account-height"
// 		tr := &trans.T{
// 			V: key,
// 		}
// 		if err := tr.AddressToHash(); err != nil {
// 			return stderr.ErrInvalidArgs
// 		}
// 		if err := tr.BytesToHex(); err != nil {
// 			return stderr.ErrInvalidArgs
// 		}
// 		uri.Path = fmt.Sprintf("/%s/ffffffffffffffff", tr.V)
// 	default:
// 		return stderr.ErrInvalidArgs
// 	}

// 	var result []byte

// 	urs := uri.String()
// 	if err := me.Client.Calls("DB.GetLast", struct {
// 		Key    []byte
// 		Prefix int
// 	}{
// 		Key:    []byte(urs),
// 		Prefix: len(urs) - 16,
// 	}, &result); err != nil {
// 		return stderr.ErrUnknown
// 	}

// 	if len(result) == 0 {
// 		return stderr.ErrNotFound
// 	}

// 	*ret = json.RawMessage(result)
// 	return nil
// }

// // GETSTORAGE ...
// func (me *T) GETSTORAGE(args []interface{}, ret *interface{}) error {
// 	switch len(args) {
// 	case 2:
// 	default:
// 		return stderr.ErrInvalidArgs
// 	}

// 	var uri url.URL

// 	uri.Scheme = "storage"
// 	uri.Host = "dbkey-height"

// 	switch key := args[0].(type) {
// 	case string:
// 		key = strings.ToLower(key)
// 		tr := &trans.T{
// 			V: key,
// 		}
// 		if err := tr.StringToLowerCase(); err != nil {
// 			return stderr.ErrInvalidArgs
// 		}
// 		if err := tr.Remove0xPrefix(); err != nil {
// 			return stderr.ErrInvalidArgs
// 		}
// 		if err := tr.HexToBytes(); err != nil {
// 			return stderr.ErrInvalidArgs
// 		}
// 		if err := tr.BytesReverse(); err != nil {
// 			return stderr.ErrInvalidArgs
// 		}
// 		if err := tr.BytesToHex(); err != nil {
// 			return stderr.ErrInvalidArgs
// 		}
// 		uri.Path = fmt.Sprint(tr.V)
// 	default:
// 		return stderr.ErrInvalidArgs
// 	}

// 	switch key := args[1].(type) {
// 	case string:
// 		uri.Path = fmt.Sprintf("/%s%s/ffffffffffffffff", uri.Path, key)
// 	default:
// 		return stderr.ErrInvalidArgs
// 	}

// 	var result []byte

// 	urs := uri.String()
// 	if err := me.Client.Calls("DB.GetLast", struct {
// 		Key    []byte
// 		Prefix int
// 	}{
// 		Key:    []byte(urs),
// 		Prefix: len(urs) - 16,
// 	}, &result); err != nil {
// 		return stderr.ErrUnknown
// 	}

// 	if len(result) == 0 {
// 		return stderr.ErrNotFound
// 	}

// 	*ret = hex.EncodeToString(result)
// 	return nil
// }

// // PING ...
// func (me *T) PING(args []interface{}, ret *interface{}) error {
// 	*ret = "pong"
// 	return nil
// }

// // DUMPPRIVKEY ...
// func (me *T) DUMPPRIVKEY(args []interface{}, ret *interface{}) error {
// 	return stderr.ErrUnsupportedMethod
// }

// // GETBALANCE ...
// func (me *T) GETBALANCE(args []interface{}, ret *interface{}) error {
// 	return stderr.ErrUnsupportedMethod
// }

// // GETTRANSACTIONHEIGHT ...
// func (me *T) GETTRANSACTIONHEIGHT(args []interface{}, ret *interface{}) error {
// 	switch len(args) {
// 	case 1:
// 	default:
// 		return stderr.ErrInvalidArgs
// 	}

// 	var uri url.URL

// 	uri.Scheme = "height"

// 	switch key := args[0].(type) {
// 	case string:
// 		uri.Host = "tx"
// 		tr := &trans.T{
// 			V: key,
// 		}
// 		if err := tr.StringToLowerCase(); err != nil {
// 			return stderr.ErrInvalidArgs
// 		}
// 		if err := tr.Remove0xPrefix(); err != nil {
// 			return stderr.ErrInvalidArgs
// 		}
// 		if err := tr.HexToBytes(); err != nil {
// 			return stderr.ErrInvalidArgs
// 		}
// 		if err := tr.BytesReverse(); err != nil {
// 			return stderr.ErrInvalidArgs
// 		}
// 		if err := tr.BytesToHex(); err != nil {
// 			return stderr.ErrInvalidArgs
// 		}
// 		uri.Path = fmt.Sprintf("/%s", tr.V)
// 	default:
// 		return stderr.ErrInvalidArgs
// 	}

// 	var result []byte

// 	urs := uri.String()
// 	if err := me.Client.Calls("DB.Get", []byte(urs), &result); err != nil {
// 		return stderr.ErrUnknown
// 	}

// 	if len(result) == 0 {
// 		return stderr.ErrNotFound
// 	}

// 	*ret = json.RawMessage(result)

// 	return nil
// }

// // GETUNSPENTS is ...
// func (me *T) GETUNSPENTS(args []interface{}, ret *interface{}) error {
// 	switch len(args) {
// 	case 1:
// 	default:
// 		return stderr.ErrInvalidArgs
// 	}

// 	var uri url.URL

// 	uri.Scheme = "adhocunspents"

// 	switch key := args[0].(type) {
// 	case string:
// 		uri.Host = "account-height"
// 		tr := &trans.T{
// 			V: key,
// 		}
// 		if err := tr.AddressToHash(); err != nil {
// 			return stderr.ErrInvalidArgs
// 		}
// 		if err := tr.BytesToHex(); err != nil {
// 			return stderr.ErrInvalidArgs
// 		}
// 		uri.Path = fmt.Sprintf("/%s/ffffffffffffffff", tr.V)
// 	default:
// 		return stderr.ErrInvalidArgs
// 	}

// 	var result []byte

// 	urs := uri.String()
// 	if err := me.Client.Calls("DB.GetLast", struct {
// 		Key    []byte
// 		Prefix int
// 	}{
// 		Key:    []byte(urs),
// 		Prefix: len(urs) - 16,
// 	}, &result); err != nil {
// 		return stderr.ErrUnknown
// 	}

// 	if len(result) == 0 {
// 		return stderr.ErrNotFound
// 	}

// 	*ret = json.RawMessage(result)
// 	return nil
// }

// // SENDRAWTRANSACTION ...
// func (me *T) SENDRAWTRANSACTION(args []interface{}, ret *interface{}) error {
// 	data := map[string]interface{}{
// 		"jsonrpc": "2.0",
// 		"id":      rand.Uint32(),
// 		"method":  "sendrawtransaction",
// 		"params":  args,
// 	}
// 	body, err := json.Marshal(data)
// 	if err != nil {
// 		return stderr.ErrInvalidArgs
// 	}
// 	resp, err := http.Post("", "application/json", bytes.NewReader(body))
// 	if err != nil {
// 		return stderr.ErrUnknown
// 	}
// 	defer resp.Body.Close()
// 	decoder := json.NewDecoder(resp.Body)
// 	if err := decoder.Decode(&data); err != nil {
// 		return stderr.ErrUnknown
// 	}
// 	if err := data["error"]; err != nil {
// 		return errors.New(fmt.Sprint(err))
// 	}
// 	*ret = data["result"]
// 	return nil
// }

// CLAIMGAS ...
func (me *T) CLAIMGAS(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// GETCONNECTIONCOUNT ...
func (me *T) GETCONNECTIONCOUNT(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// GETMETRICBLOCKTIMESTAMP ...
func (me *T) GETMETRICBLOCKTIMESTAMP(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// GETUTXOTRANSFERS ...
func (me *T) GETUTXOTRANSFERS(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// GETNEP5TRANSFERS ...
func (me *T) GETNEP5TRANSFERS(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// GETNEWADDRESS ...
func (me *T) GETNEWADDRESS(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// GETRAWMEMPOOL ...
func (me *T) GETRAWMEMPOOL(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// GETPEERS ...
func (me *T) GETPEERS(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// GETUNCLAIMEDGAS ...
func (me *T) GETUNCLAIMEDGAS(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// GETVERSION ...
func (me *T) GETVERSION(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// GETWALLETHEIGHT ...
func (me *T) GETWALLETHEIGHT(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// IMPORTPRIVKEY ...
func (me *T) IMPORTPRIVKEY(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// LISTPLUGINS ...
func (me *T) LISTPLUGINS(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// LISTADDRESS ...
func (me *T) LISTADDRESS(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// SENDFROM ...
func (me *T) SENDFROM(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// SENDTOADDRESS ...
func (me *T) SENDTOADDRESS(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// SENDMANY ...
func (me *T) SENDMANY(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// GETSTATEHEIGHT ...
func (me *T) GETSTATEHEIGHT(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// GETTXOUT ...
func (me *T) GETTXOUT(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// SUBMITBLOCK ...
func (me *T) SUBMITBLOCK(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// GETVALIDATORS ...
func (me *T) GETVALIDATORS(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}
