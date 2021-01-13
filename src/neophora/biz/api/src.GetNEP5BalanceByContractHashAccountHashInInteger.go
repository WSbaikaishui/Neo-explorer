package api

import (
	"encoding/json"
	"neophora/lib/type/bins"
	"neophora/lib/type/h160"
	"neophora/var/stderr"
)

// GetNEP5BalanceByContractHashAccountHashInInteger ...
// as an example:
//
// ```
// $ curl https://example.neophora.io -d '{"jsonrpc": "2.0","id": 1,"method": "GetNEP5BalanceByContractHashAccountHashInInteger","params":{ "ContractHash":"972e166ea1f8d3c3b14fd8766e7a0dad4084f9e8","AccountHash":"30c6c993a33e630056cfd05556361bf128b90fc8"}}'
// {"id":1,"result":0,"error":null}
// ```
func (me *T) GetNEP5BalanceByContractHashAccountHashInInteger(args struct {
	ContractHash h160.T
	AccountHash  h160.T
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
		Index:  "h160.ctr-h160.act-uint.hgt",
		Keys:   []string{args.ContractHash.Val(), args.AccountHash.Val()},
	}, &result); err != nil {
		return nil
	}
	if result.Valid() == false {
		return stderr.ErrNotFound
	}
	*ret = json.RawMessage(result.BigString())
	return nil
}
