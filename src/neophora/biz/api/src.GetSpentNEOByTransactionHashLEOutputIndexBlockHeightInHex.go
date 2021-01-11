package api

import (
	"neophora/lib/type/h256"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetSpentNEOByTransactionHashLEOutputIndexBlockHeightInHex ...
// as an example:
//
// ```
// $ curl https://example.neophora.io -d '{"jsonrpc": "2.0","id": 1,"method": "GetSpentNEOByTransactionHashLEOutputIndexBlockHeightInHex","params":{"TransactionHashLE":"22e1a1617f0d53d6f8e504d6d663d5c33f07698b4e3cb0006019710761dd0e16","OutputIndex":0, "BlockHeight":2475663}}'
// {"id":1,"result":"0000000000249f1d000000000025c68f","error":null}
// ```
func (me *T) GetSpentNEOByTransactionHashLEOutputIndexBlockHeightInHex(args struct {
	TransactionHashLE h256.T
	OutputIndex       uintval.T
	BlockHeight       uintval.T
}, ret *string) error {
	if args.TransactionHashLE.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.OutputIndex.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.BlockHeight.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	return me.Data.GetLastValInHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "u128.spt",
		Index:  "h256.trx-uint.num-uint.hgt",
		Keys:   []string{args.TransactionHashLE.RevVal(), args.OutputIndex.Hex(), args.BlockHeight.Hex()},
	}, ret)
}
