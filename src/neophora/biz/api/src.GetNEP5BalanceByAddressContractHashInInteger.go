package api

import (
	"encoding/json"
	"neophora/lib/type/addr"
	"neophora/lib/type/bins"
	"neophora/lib/type/h160"
	"neophora/var/stderr"
)

// GetNEP5BalanceByAddressContractHashInInteger ...
// as an example:
//
// ```
// $ curl https://example.neophora.io -d '{"jsonrpc": "2.0","id": 1,"method": "GetNEP5BalanceByAddressContractHashInInteger","params":{"Address":"Aa1hdLWLw441vQ94hzpwDj8FY7QEh1zsr6", "ContractHash":"972e166ea1f8d3c3b14fd8766e7a0dad4084f9e8"}}'
// {"id":1,"result":0,"error":null}
// ```
func (me *T) GetNEP5BalanceByAddressContractHashInInteger(args struct {
	Address      addr.T
	ContractHash h160.T
}, ret *json.RawMessage) error {
	var result bins.T
	if args.ContractHash.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.Address.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if err := me.Data.GetLatestUint64ValInBins(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bigu.bal",
		Index:  "h160.act-h160.ctr-uint.hgt",
		Keys:   []string{args.Address.H160(), args.ContractHash.Val()},
	}, &result); err != nil {
		return nil
	}
	if result.Valid() == false {
		return stderr.ErrNotFound
	}
	*ret = json.RawMessage(result.BigString())
	return nil
}
