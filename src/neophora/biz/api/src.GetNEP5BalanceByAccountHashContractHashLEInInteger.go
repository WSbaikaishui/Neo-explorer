package api

import (
	"encoding/json"
	"neophora/lib/type/bins"
	"neophora/lib/type/h160"
	"neophora/var/stderr"
)

// GetNEP5BalanceByAccountHashContractHashLEInInteger ...
// as an example:
//
// ```
// $ curl https://example.neophora.io -d '{"jsonrpc": "2.0","id": 1,"method": "GetNEP5BalanceByAccountHashContractHashLEInInteger","params":{"AccountHash":"30c6c993a33e630056cfd05556361bf128b90fc8", "ContractHashLE":"e8f98440ad0d7a6e76d84fb1c3d3f8a16e162e97"}}'
// {"id":1,"result":0,"error":null}
// ```
func (me *T) GetNEP5BalanceByAccountHashContractHashLEInInteger(args struct {
	AccountHash    h160.T
	ContractHashLE h160.T
}, ret *json.RawMessage) error {
	var result bins.T
	if args.ContractHashLE.Valid() == false {
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
		Keys:   []string{args.AccountHash.Val(), args.ContractHashLE.RevVal()},
	}, &result); err != nil {
		return nil
	}
	if result.Valid() == false {
		return stderr.ErrNotFound
	}
	*ret = json.RawMessage(result.BigString())
	return nil
}
