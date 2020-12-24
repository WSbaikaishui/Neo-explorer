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
// TODO
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
