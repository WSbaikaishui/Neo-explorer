package api

import (
	"encoding/json"
	"neophora/lib/type/addr"
	"neophora/lib/type/bins"
	"neophora/lib/type/h160"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetNEP5BalanceByAddressContractHashBlockHeightInInteger ...
// as an example:
//
// ```
// $ curl https://example.neophora.io -d '{"jsonrpc": "2.0","id": 1,"method": "GetNEP5BalanceByAddressContractHashBlockHeightInInteger","params":{"Address":"Aa1hdLWLw441vQ94hzpwDj8FY7QEh1zsr6", "ContractHash":"972e166ea1f8d3c3b14fd8766e7a0dad4084f9e8","BlockHeight":2400003}}'
// {"id":1,"result":1694840000000000,"error":null}
// ```
func (me *T) GetNEP5BalanceByAddressContractHashBlockHeightInInteger(args struct {
	Address      addr.T
	ContractHash h160.T
	BlockHeight  uintval.T
}, ret *json.RawMessage) error {
	var result bins.T
	if args.ContractHash.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.Address.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.BlockHeight.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if err := me.Data.GetLastValInBins(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bigu.bal",
		Index:  "h160.act-h160.ctr-uint.hgt",
		Keys:   []string{args.Address.H160(), args.ContractHash.Val(), args.BlockHeight.Hex()},
	}, &result); err != nil {
		return nil
	}
	if result.Valid() == false {
		return stderr.ErrNotFound
	}
	*ret = json.RawMessage(result.BigString())
	return nil
}
