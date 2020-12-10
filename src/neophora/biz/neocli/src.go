package neocli

import (
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"neophora/biz/data"
	"neophora/lib/trans"
	"neophora/var/stderr"
	"net/url"
)

// T ...
type T struct {
	Data *data.T
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
		if err := me.Data.GETARGS(struct {
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
		if err := me.Data.GETARGS(struct {
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
		if err := me.Data.GETARGS(struct {
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
		if err := me.Data.GETARGS(struct {
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

// GETRAWTRANSACTION ...
func (me *T) GETRAWTRANSACTION(args []interface{}, ret *interface{}) error {
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
		scheme = "tx"
	case 1.0:
		scheme = "adhoctxinfo"
	default:
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
		if err := me.Data.GETARGS(struct {
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

// GETAPPLICATIONLOG ...
func (me *T) GETAPPLICATIONLOG(args []interface{}, ret *interface{}) error {
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
		if err := me.Data.GETARGS(struct {
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

// GETSTATEROOT ...
func (me *T) GETSTATEROOT(args []interface{}, ret *interface{}) error {
	if len(args) != 1 {
		return stderr.ErrInvalidArgs
	}

	var result []byte
	switch key := args[0].(type) {
	case float64:
		if err := me.Data.GETARGS(struct {
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
		if err := me.Data.GETARGS(struct {
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

// GETBLOCKHASH ...
func (me *T) GETBLOCKHASH(args []interface{}, ret *interface{}) error {
	if len(args) != 1 {
		return stderr.ErrInvalidArgs
	}

	var result []byte
	switch key := args[0].(type) {
	case float64:
		if err := me.Data.GETARGS(struct {
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

// GETBLOCKSYSFEE ...
func (me *T) GETBLOCKSYSFEE(args []interface{}, ret *interface{}) error {
	if len(args) != 1 {
		return stderr.ErrInvalidArgs
	}

	var result []byte
	switch key := args[0].(type) {
	case float64:
		if err := me.Data.GETARGS(struct {
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

// GETACCOUNTSTATE ...
func (me *T) GETACCOUNTSTATE(args []interface{}, ret *interface{}) error {
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
	if err := me.Data.GETLAST(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "adhocaccountstate",
		Index:  "account-height",
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

// GETASSETSTATE ...
func (me *T) GETASSETSTATE(args []interface{}, ret *interface{}) error {
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
	if err := me.Data.GETLAST(struct {
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

// GETCLAIMABLE ...
func (me *T) GETCLAIMABLE(args []interface{}, ret *interface{}) error {
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
	if err := me.Data.GETLAST(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "adhocclaimable",
		Index:  "account-height",
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

// GETUNSPENTS ...
func (me *T) GETUNSPENTS(args []interface{}, ret *interface{}) error {
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
	if err := me.Data.GETLAST(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "adhocunspents",
		Index:  "account-height",
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

// GETNEP5BALANCES ...
func (me *T) GETNEP5BALANCES(args []interface{}, ret *interface{}) error {
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
	if err := me.Data.GETLAST(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "adhocnep5balances",
		Index:  "account-height",
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

// GETCONTRACTSTATE ...
func (me *T) GETCONTRACTSTATE(args []interface{}, ret *interface{}) error {
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
	if err := me.Data.GETLAST(struct {
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

// GETSTORAGE ...
func (me *T) GETSTORAGE(args []interface{}, ret *interface{}) error {
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
	if err := me.Data.GETLAST(struct {
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

// GETBLOCKCOUNT ...
func (me *T) GETBLOCKCOUNT(args []interface{}, ret *interface{}) error {
	if len(args) != 0 {
		return stderr.ErrInvalidArgs
	}

	var result []byte
	if err := me.Data.GETLASTKEY(struct {
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

// GETBESTBLOCKHASH ...
func (me *T) GETBESTBLOCKHASH(args []interface{}, ret *interface{}) error {
	if len(args) != 0 {
		return stderr.ErrInvalidArgs
	}

	var result []byte
	if err := me.Data.GETLAST(struct {
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

// GETTRANSACTIONHEIGHT ...
func (me *T) GETTRANSACTIONHEIGHT(args []interface{}, ret *interface{}) error {
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
	if err := me.Data.GETARGS(struct {
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

// DUMPPRIVKEY ...
func (me *T) DUMPPRIVKEY(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// GETPROOF ...
func (me *T) GETPROOF(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// GETBALANCE ...
func (me *T) GETBALANCE(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// INVOKESCRIPT ...
func (me *T) INVOKESCRIPT(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// INVOKEFUNCTION ...
func (me *T) INVOKEFUNCTION(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

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

// GETUNCLAIMED ...
func (me *T) GETUNCLAIMED(args []interface{}, ret *interface{}) error {
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

// ValidateAddress ...
func (me *T) ValidateAddress(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}

// VerifyProof ...
func (me *T) VerifyProof(args []interface{}, ret *interface{}) error {
	return stderr.ErrUnsupportedMethod
}
