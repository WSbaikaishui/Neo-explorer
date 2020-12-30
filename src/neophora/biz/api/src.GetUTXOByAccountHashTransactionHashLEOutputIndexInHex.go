package api

import (
	"neophora/lib/type/h160"
	"neophora/lib/type/h256"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetUTXOByAccountHashTransactionHashLEOutputIndexInHex ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetUTXOByAccountHashTransactionHashLEOutputIndexInHex(args struct {
	AccountHash       h160.T
	TransactionHashLE h256.T
	OutputIndex       uintval.T
}, ret *string) error {
	if args.AccountHash.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.TransactionHashLE.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.OutputIndex.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	return me.Data.GetArgsInHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.utx",
		Index:  "h160.act-h256.trx-uint.num",
		Keys:   []string{args.AccountHash.Val(), args.TransactionHashLE.RevVal(), args.OutputIndex.Hex()},
	}, ret)
}
