package api

import (
	"neophora/lib/type/h256"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetSpentNEOByTransactionHashOutputIndexBlockHeightInHex ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetSpentNEOByTransactionHashOutputIndexBlockHeightInHex(args struct {
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
		Target: "u128.spt",
		Index:  "h256.trx-uint.num-uint.hgt",
		Keys:   []string{args.TransactionHash.Val(), args.OutputIndex.Hex(), args.BlockHeight.Hex()},
	}, ret)
}
