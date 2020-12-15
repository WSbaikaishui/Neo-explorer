package neocli

import (
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"neophora/biz/data"
	"neophora/lib/bq"
	"neophora/lib/trans"
	"neophora/lib/transex"
	"neophora/var/stderr"
	"net/url"
)

// T ...
type T struct {
	Data *data.T
	BQ   *bq.T
}

// Getblock ...
func (me *T) Getblock(args []interface{}, ret *interface{}) error {
	switch len(args) {
	case 1:
		args = append(args, 0.0)
	case 2:
	default:
		return stderr.ErrInvalidArgs
	}

	var scheme string
	switch args[1] {
	case 0.0:
		scheme = "block"
	case 1.0:
		scheme = "adhocblockinfo"
	default:
		return stderr.ErrInvalidArgs
	}

	var result []byte
	switch key := args[0].(type) {
	case float64:
		if err := me.Data.GetArgs(struct {
			Target string
			Index  string
			Keys   []string
		}{
			Target: scheme,
			Index:  "height",
			Keys:   []string{fmt.Sprintf("%016x", uint64(key))},
		}, &result); err != nil {
			return err
		}
	case string:
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
		if err := me.Data.GetArgs(struct {
			Target string
			Index  string
			Keys   []string
		}{
			Target: scheme,
			Index:  "hash",
			Keys:   []string{tr.V.(string)},
		}, &result); err != nil {
			return err
		}
	default:
		return stderr.ErrInvalidArgs
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

	var scheme string
	switch args[1] {
	case 0.0:
		scheme = "header"
	case 1.0:
		scheme = "adhocheaderinfo"
	default:
		return stderr.ErrInvalidArgs
	}

	var result []byte
	switch key := args[0].(type) {
	case float64:
		if err := me.Data.GetArgs(struct {
			Target string
			Index  string
			Keys   []string
		}{
			Target: scheme,
			Index:  "height",
			Keys:   []string{fmt.Sprintf("%016x", uint64(key))},
		}, &result); err != nil {
			return err
		}
	case string:
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
		if err := me.Data.GetArgs(struct {
			Target string
			Index  string
			Keys   []string
		}{
			Target: scheme,
			Index:  "hash",
			Keys:   []string{tr.V.(string)},
		}, &result); err != nil {
			return err
		}
	default:
		return stderr.ErrInvalidArgs
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

// Getrawtransaction ...
func (me *T) Getrawtransaction(args []interface{}, ret *interface{}) error {
	switch len(args) {
	case 1:
		args = append(args, 0.0)
	case 2:
	default:
		return stderr.ErrInvalidArgs
	}

	var tr transex.T

	var result []byte
	switch key := args[0].(type) {
	case string:
		tr.V = key
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
		if err := me.Data.GetArgs(struct {
			Target string
			Index  string
			Keys   []string
		}{
			Target: "tx",
			Index:  "hash",
			Keys:   []string{tr.V.(string)},
		}, &result); err != nil {
			return err
		}
	default:
		return stderr.ErrInvalidArgs
	}

	if len(result) == 0 {
		return stderr.ErrNotFound
	}

	switch args[1] {
	case 0.0:
		*ret = hex.EncodeToString(result)
	case 1.0:
		tr.V = result
		if err := tr.BytesToJSONViaTX(); err != nil {
			return stderr.ErrNotFound
		}
		*ret = tr.V
	}

	return nil
}

// Getapplicationlog ...
func (me *T) Getapplicationlog(args []interface{}, ret *interface{}) error {
	if len(args) != 1 {
		return stderr.ErrInvalidArgs
	}

	var result []byte
	switch key := args[0].(type) {
	case string:
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
		if err := me.Data.GetArgs(struct {
			Target string
			Index  string
			Keys   []string
		}{
			Target: "adhoclog",
			Index:  "tx",
			Keys:   []string{tr.V.(string)},
		}, &result); err != nil {
			return err
		}
	default:
		return stderr.ErrInvalidArgs
	}

	if len(result) == 0 {
		return stderr.ErrNotFound
	}

	*ret = json.RawMessage(result)
	return nil
}

// Getstateroot ...
func (me *T) Getstateroot(args []interface{}, ret *interface{}) error {
	if len(args) != 1 {
		return stderr.ErrInvalidArgs
	}

	var result []byte
	switch key := args[0].(type) {
	case float64:
		if err := me.Data.GetArgs(struct {
			Target string
			Index  string
			Keys   []string
		}{
			Target: "adhocstateroot",
			Index:  "height",
			Keys:   []string{fmt.Sprintf("%016x", uint64(key))},
		}, &result); err != nil {
			return err
		}
	case string:
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
		if err := me.Data.GetArgs(struct {
			Target string
			Index  string
			Keys   []string
		}{
			Target: "adhocstateroot",
			Index:  "hash",
			Keys:   []string{tr.V.(string)},
		}, &result); err != nil {
			return err
		}
	default:
		return stderr.ErrInvalidArgs
	}

	if len(result) == 0 {
		return stderr.ErrNotFound
	}

	*ret = json.RawMessage(result)
	return nil
}

// Getblockhash ...
func (me *T) Getblockhash(args []interface{}, ret *interface{}) error {
	if len(args) != 1 {
		return stderr.ErrInvalidArgs
	}

	var result []byte
	switch key := args[0].(type) {
	case float64:
		if err := me.Data.GetArgs(struct {
			Target string
			Index  string
			Keys   []string
		}{
			Target: "hash",
			Index:  "height",
			Keys:   []string{fmt.Sprintf("%016x", uint64(key))},
		}, &result); err != nil {
			return err
		}
	default:
		return stderr.ErrInvalidArgs
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

// Getblocksysfee ...
func (me *T) Getblocksysfee(args []interface{}, ret *interface{}) error {
	if len(args) != 1 {
		return stderr.ErrInvalidArgs
	}

	var result []byte
	switch key := args[0].(type) {
	case float64:
		if err := me.Data.GetArgs(struct {
			Target string
			Index  string
			Keys   []string
		}{
			Target: "adhocsysfee",
			Index:  "height",
			Keys:   []string{fmt.Sprintf("%016x", uint64(key))},
		}, &result); err != nil {
			return err
		}
	default:
		return stderr.ErrInvalidArgs
	}

	if len(result) == 0 {
		return stderr.ErrNotFound
	}

	*ret = string(result)
	return nil
}

// Getaccountstate ...
func (me *T) Getaccountstate(args []interface{}, ret *interface{}) error {
	if len(args) != 1 {
		return stderr.ErrInvalidArgs
	}

	tr := &trans.T{
		V: args[0],
	}
	if err := tr.AddressToHash(); err != nil {
		return stderr.ErrInvalidArgs
	}
	if err := tr.BytesToHex(); err != nil {
		return stderr.ErrInvalidArgs
	}

	var result []byte
	if err := me.Data.GetLast(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "adhocaccountstate",
		Index:  "hash-height",
		Keys:   []string{tr.V.(string), "_"},
	}, &result); err != nil {
		return err
	}

	if len(result) == 0 {
		return stderr.ErrNotFound
	}

	*ret = json.RawMessage(result)
	return nil
}

// Getassetstate ...
func (me *T) Getassetstate(args []interface{}, ret *interface{}) error {
	if len(args) != 1 {
		return stderr.ErrInvalidArgs
	}

	tr := &trans.T{
		V: args[0],
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

	var result []byte
	if err := me.Data.GetLast(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "adhocassetstate",
		Index:  "hash-height",
		Keys:   []string{tr.V.(string), "_"},
	}, &result); err != nil {
		return err
	}

	if len(result) == 0 {
		return stderr.ErrNotFound
	}

	*ret = json.RawMessage(result)
	return nil
}

// Getclaimable ...
func (me *T) Getclaimable(args []interface{}, ret *interface{}) error {
	if len(args) != 1 {
		return stderr.ErrInvalidArgs
	}

	tr := &trans.T{
		V: args[0],
	}
	if err := tr.AddressToHash(); err != nil {
		return stderr.ErrInvalidArgs
	}
	if err := tr.BytesToHex(); err != nil {
		return stderr.ErrInvalidArgs
	}

	var result []byte
	if err := me.Data.GetLast(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "adhocclaimable",
		Index:  "hash-height",
		Keys:   []string{tr.V.(string), "_"},
	}, &result); err != nil {
		return err
	}

	if len(result) == 0 {
		return stderr.ErrNotFound
	}

	*ret = json.RawMessage(result)
	return nil
}

// Getunspents ...
func (me *T) Getunspents(args []interface{}, ret *interface{}) error {
	if len(args) != 1 {
		return stderr.ErrInvalidArgs
	}

	tr := &trans.T{
		V: args[0],
	}
	if err := tr.AddressToHash(); err != nil {
		return stderr.ErrInvalidArgs
	}
	if err := tr.BytesToHex(); err != nil {
		return stderr.ErrInvalidArgs
	}

	var result []byte
	if err := me.Data.GetLast(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "adhocunspents",
		Index:  "hash-height",
		Keys:   []string{tr.V.(string), "_"},
	}, &result); err != nil {
		return err
	}

	if len(result) == 0 {
		return stderr.ErrNotFound
	}

	*ret = json.RawMessage(result)
	return nil
}

// Getnep5balances ...
func (me *T) Getnep5balances(args []interface{}, ret *interface{}) error {
	if len(args) != 1 {
		return stderr.ErrInvalidArgs
	}

	tr := &trans.T{
		V: args[0],
	}
	if err := tr.AddressToHash(); err != nil {
		return stderr.ErrInvalidArgs
	}
	if err := tr.BytesToHex(); err != nil {
		return stderr.ErrInvalidArgs
	}

	var result []byte
	if err := me.Data.GetLast(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "adhocnep5balances",
		Index:  "hash-height",
		Keys:   []string{tr.V.(string), "_"},
	}, &result); err != nil {
		return err
	}

	if len(result) == 0 {
		return stderr.ErrNotFound
	}

	*ret = json.RawMessage(result)
	return nil
}

// Getcontractstate ...
func (me *T) Getcontractstate(args []interface{}, ret *interface{}) error {
	if len(args) != 1 {
		return stderr.ErrInvalidArgs
	}

	tr := &trans.T{
		V: args[0],
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

	var result []byte
	if err := me.Data.GetLast(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "adhoccontractstate",
		Index:  "hash-height",
		Keys:   []string{tr.V.(string), "_"},
	}, &result); err != nil {
		return err
	}

	switch len(args) {
	case 1:
	default:
		return stderr.ErrInvalidArgs
	}

	if len(result) == 0 {
		return stderr.ErrNotFound
	}

	*ret = json.RawMessage(result)
	return nil
}

// Getstorage ...
func (me *T) Getstorage(args []interface{}, ret *interface{}) error {
	if len(args) != 2 {
		return stderr.ErrInvalidArgs
	}

	tr := &trans.T{
		V: args[0],
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

	switch key := args[1].(type) {
	case string:
		tr.V = tr.V.(string) + key
	default:
		return stderr.ErrInvalidArgs
	}

	if err := tr.HexToBytes(); err != nil {
		return stderr.ErrInvalidArgs
	}
	if err := tr.BytesToHash(); err != nil {
		return stderr.ErrInvalidArgs
	}
	if err := tr.BytesToHex(); err != nil {
		return stderr.ErrInvalidArgs
	}

	var result []byte
	if err := me.Data.GetLast(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "storage",
		Index:  "hash-height",
		Keys:   []string{tr.V.(string), "_"},
	}, &result); err != nil {
		return err
	}

	if len(result) == 0 {
		return stderr.ErrNotFound
	}

	*ret = hex.EncodeToString(result)
	return nil
}

// Getblockcount ...
func (me *T) Getblockcount(args []interface{}, ret *interface{}) error {
	if len(args) != 0 {
		return stderr.ErrInvalidArgs
	}

	var result []byte
	if err := me.Data.GetLastKey(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "hash",
		Index:  "height",
		Keys:   []string{"_"},
	}, &result); err != nil {
		return err
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

// Getbestblockhash ...
func (me *T) Getbestblockhash(args []interface{}, ret *interface{}) error {
	if len(args) != 0 {
		return stderr.ErrInvalidArgs
	}

	var result []byte
	if err := me.Data.GetLast(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "hash",
		Index:  "height",
		Keys:   []string{"_"},
	}, &result); err != nil {
		return err
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

// Gettransactionheight ...
func (me *T) Gettransactionheight(args []interface{}, ret *interface{}) error {
	if len(args) != 1 {
		return stderr.ErrInvalidArgs
	}

	tr := &trans.T{
		V: args[0],
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

	var result []byte
	if err := me.Data.GetArgs(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "height",
		Index:  "tx",
		Keys:   []string{tr.V.(string)},
	}, &result); err != nil {
		return err
	}

	if len(result) != 8 {
		return stderr.ErrNotFound
	}

	*ret = binary.BigEndian.Uint64(result)
	return nil
}

// Sendrawtransaction ...
func (me *T) Sendrawtransaction(args []interface{}, ret *interface{}) error {
	if len(args) != 1 {
		return stderr.ErrInvalidArgs
	}
	tr := &trans.T{
		V: args[0],
	}
	if err := tr.HexToBytes(); err != nil {
		return stderr.ErrInvalidArgs
	}
	if err := me.BQ.Push(tr.V.([]byte)); err != nil {
		return err
	}
	*ret = "ok"
	return nil
}

// Dumpprivkey ...
func (me *T) Dumpprivkey(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// Getproof ...
func (me *T) Getproof(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// Getbalance ...
func (me *T) Getbalance(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// Invokescript ...
func (me *T) Invokescript(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// Invokefunction ...
func (me *T) Invokefunction(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// Claimgas ...
func (me *T) Claimgas(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// Getconnectioncount ...
func (me *T) Getconnectioncount(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// Getmetricblocktimestamp ...
func (me *T) Getmetricblocktimestamp(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// Getutxotransfers ...
func (me *T) Getutxotransfers(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// Getnep5transfers ...
func (me *T) Getnep5transfers(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// Getnewaddress ...
func (me *T) Getnewaddress(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// Getrawmempool ...
func (me *T) Getrawmempool(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// Getpeers ...
func (me *T) Getpeers(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// Getunclaimedgas ...
func (me *T) Getunclaimedgas(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// Getunclaimed ...
func (me *T) Getunclaimed(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// Getversion ...
func (me *T) Getversion(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// Getwalletheight ...
func (me *T) Getwalletheight(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// Importprivkey ...
func (me *T) Importprivkey(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// Listplugins ...
func (me *T) Listplugins(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// Listaddress ...
func (me *T) Listaddress(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// Sendfrom ...
func (me *T) Sendfrom(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// Sendtoaddress ...
func (me *T) Sendtoaddress(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// Sendmany ...
func (me *T) Sendmany(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// Getstateheight ...
func (me *T) Getstateheight(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// Gettxout ...
func (me *T) Gettxout(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// Submitblock ...
func (me *T) Submitblock(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// Getvalidators ...
func (me *T) Getvalidators(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// Validateaddress ...
func (me *T) Validateaddress(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// Verifyproof ...
func (me *T) Verifyproof(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}
