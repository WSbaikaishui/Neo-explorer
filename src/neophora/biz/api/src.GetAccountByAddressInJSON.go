package api

import (
	"encoding/json"
	"neophora/lib/type/addr"
	"neophora/lib/type/bins"
	"neophora/var/stderr"
)

// GetAccountByAddressInJSON ...
// as an example:
//
// ```
// $ curl https://example.neophora.io -d '{"jsonrpc": "2.0","id": 1,"method": "GetAccountByAddressInJSON","params":{"Address":"AZCcft1uYtmZXxzHPr5tY7L6M85zG7Dsrv"}}'
// {"id":1,"result":{"balances":[{"asset":"9b7cffdaa674beae0f930ebe6085af9093e5fe56b34a5c220ccdcf6efc336fc5","value":"1"}],"frozen":false,"script_hash":"0x4be4b57a3835ba5a7ec63f2c83476e58130428bf","version":0,"votes":[]},"error":null}
// ```
func (me *T) GetAccountByAddressInJSON(args struct {
	Address addr.T
}, ret *json.RawMessage) error {
	if args.Address.Valid() == false {
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
		Keys:   []string{args.Address.H160()},
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
