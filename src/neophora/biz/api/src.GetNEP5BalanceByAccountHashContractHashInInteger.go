package api

import (
	"encoding/json"
	"neophora/lib/type/bins"
	"neophora/lib/type/h160"
	"neophora/var/stderr"
)

// GetNEP5BalanceByAccountHashContractHashInInteger ...
// as an example:
//
// ```
// $ curl https://example.neophora.io -d '{"jsonrpc": "2.0","id": 1,"method": "GetNEP5BalanceByAccountHashContractHashInInteger","params":{"AccountHash":"30c6c993a33e630056cfd05556361bf128b90fc8", "ContractHash":"972e166ea1f8d3c3b14fd8766e7a0dad4084f9e8"}}'
// {"id":1,"result":0,"error":null}
// ```
func (me *T) GetNEP5BalanceByAccountHashContractHashInInteger(args struct {
	AccountHash  h160.T
	ContractHash h160.T
}, ret *json.RawMessage) error {
	var result bins.T
	if args.ContractHash.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.AccountHash.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if err := me.Data.GetLatestUint64ValInBins(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bigu.bal",
		Index:  "h160.act-h160.ctr-uint.hgt",
		Keys:   []string{args.AccountHash.Val(), args.ContractHash.Val()},
	}, &result); err != nil {
		return nil
	}
	if result.Valid() == false {
		return stderr.ErrNotFound
	}
	*ret = json.RawMessage(result.BigString())
	return nil
}
