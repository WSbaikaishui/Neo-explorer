package api

import (
	"encoding/json"
	"neophora/lib/type/bins"
	"neophora/lib/type/h160"
	"neophora/var/stderr"
)

// GetAccountByAccountHashLEInJSON ...
// as an example:
//
// ```
// $ curl https://example.neophora.io -d '{"jsonrpc": "2.0","id": 1,"method": "GetAccountByAccountHashLEInJSON","params":{"AccountHashLE":"4c56140b455b060e9f118a85b03b6423300928bf"}}'
// {"id":1,"result":{"balances":[{"asset":"e72d286979ee6cb1b7e65dfddfb2e384100b8d148e7758de42e4168b71792c60","value":"1.33647735"}],"frozen":false,"script_hash":"0x4c56140b455b060e9f118a85b03b6423300928bf","version":0,"votes":[]},"error":null}
// ```
func (me *T) GetAccountByAccountHashLEInJSON(args struct {
	AccountHashLE h160.T
}, ret *json.RawMessage) error {
	if args.AccountHashLE.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	var result bins.T
	if err := me.Data.GetLatestUint64ValInBins(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.act",
		Index:  "h160.act-uint.hgt",
		Keys:   []string{args.AccountHashLE.RevVal()},
	}, &result); err != nil {
		return err
	}
	if result.Valid() == false {
		return stderr.ErrNotFound
	}
	js, err := result.JSONViaAccount()
	if err != nil {
		return stderr.ErrNotFound
	}
	*ret = js
	return nil
}
