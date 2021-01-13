package api

import (
	"encoding/json"
	"neophora/lib/type/bins"
	"neophora/lib/type/h160"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetNEP5BalanceByAccountHashLEContractHashLEBlockHeightInInteger ...
// as an example:
//
// ```
// $ curl https://example.neophora.io -d '{"jsonrpc": "2.0","id": 1,"method": "GetNEP5BalanceByAccountHashLEContractHashLEBlockHeightInInteger","params":{"AccountHashLE":"c80fb928f11b365655d0cf5600633ea393c9c630", "ContractHashLE":"e8f98440ad0d7a6e76d84fb1c3d3f8a16e162e97", "BlockHeight":2400003}}'
// {"id":1,"result":1694840000000000,"error":null}
// ```
func (me *T) GetNEP5BalanceByAccountHashLEContractHashLEBlockHeightInInteger(args struct {
	AccountHashLE  h160.T
	ContractHashLE h160.T
	BlockHeight    uintval.T
}, ret *json.RawMessage) error {
	var result bins.T
	if args.ContractHashLE.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.AccountHashLE.Valid() == false {
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
		Keys:   []string{args.AccountHashLE.RevVal(), args.ContractHashLE.RevVal(), args.BlockHeight.Hex()},
	}, &result); err != nil {
		return nil
	}
	if result.Valid() == false {
		return stderr.ErrNotFound
	}
	*ret = json.RawMessage(result.BigString())
	return nil
}
