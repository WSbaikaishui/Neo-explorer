package api

import (
	"neophora/lib/type/h256"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetCoinStateByTransactionHashOutputIndexBlockHeightInHex ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetCoinStateByTransactionHashOutputIndexBlockHeightInHex(args struct {
	TransactionHash h256.T
	OutputIndex     uintval.T
	BlockHeight     uintval.T
}, ret *string) error {
	if args.TransactionHash.Valid() == false {
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
		Target: "uint.cst",
		Index:  "h256.trx-uint.num-uint.hgt",
		Keys:   []string{args.TransactionHash.Val(), args.OutputIndex.Hex(), args.BlockHeight.Hex()},
	}, ret)
}
