package api

import (
	"neophora/lib/type/h160"
	"neophora/lib/type/h256"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetUTXOByAccountHashLETransactionHashLEOutputIndexInHex ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetUTXOByAccountHashLETransactionHashLEOutputIndexInHex(args struct {
	AccountHashLE     h160.T
	TransactionHashLE h256.T
	OutputIndex       uintval.T
}, ret *string) error {
	if args.AccountHashLE.Valid() == false {
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
		Keys:   []string{args.AccountHashLE.RevVal(), args.TransactionHashLE.RevVal(), args.OutputIndex.Hex()},
	}, ret)
}
