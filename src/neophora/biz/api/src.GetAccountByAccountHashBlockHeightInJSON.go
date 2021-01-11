package api

import (
	"encoding/json"
	"neophora/lib/type/bins"
	"neophora/lib/type/h160"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetAccountByAccountHashBlockHeightInJSON ...
// as an example:
//
// ```
// $ curl https://example.neophora.io -d '{"jsonrpc": "2.0","id": 1,"method": "GetAccountByAccountHashBlockHeightInJSON","params":{"AccountHash":"bf28093023643bb0858a119f0e065b450b14564c","BlockHeight":2400000} }'
// {"id":1,"result":{"balances":[{"asset":"e72d286979ee6cb1b7e65dfddfb2e384100b8d148e7758de42e4168b71792c60","value":"12160.16478122"},{"asset":"9b7cffdaa674beae0f930ebe6085af9093e5fe56b34a5c220ccdcf6efc336fc5","value":"10720"}],"frozen":false,"script_hash":"0x4c56140b455b060e9f118a85b03b6423300928bf","version":0,"votes":[]},"error":null}
// ```
func (me *T) GetAccountByAccountHashBlockHeightInJSON(args struct {
	AccountHash h160.T
	BlockHeight uintval.T
}, ret *json.RawMessage) error {
	if args.AccountHash.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.BlockHeight.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	var result bins.T
	if err := me.Data.GetLastValInBins(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.act",
		Index:  "h160.act-uint.hgt",
		Keys:   []string{args.AccountHash.Val(), args.BlockHeight.Hex()},
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
