package api

import (
	"neophora/lib/type/h256"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetCoinStateByTransactionHashLEOutputIndexInHex ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetCoinStateByTransactionHashLEOutputIndexInHex(args struct {
	TransactionHashLE h256.T
	OutputIndex       uintval.T
}, ret *string) error {
	if args.TransactionHashLE.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.OutputIndex.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	return me.Data.GetLatestUint64ValInHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.cst",
		Index:  "h256.trx-uint.num-uint.hgt",
		Keys:   []string{args.TransactionHashLE.RevVal(), args.OutputIndex.Hex()},
	}, ret)
}
