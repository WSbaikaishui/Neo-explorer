package api

import (
	"neophora/lib/type/h160"
	"neophora/lib/type/h256"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetUTXOByAccountHashTransactionHashOutputIndexInHex ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetUTXOByAccountHashTransactionHashOutputIndexInHex(args struct {
	AccountHash     h160.T
	TransactionHash h256.T
	OutputIndex     uintval.T
}, ret *string) error {
	if args.AccountHash.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.TransactionHash.Valid() == false {
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
		Keys:   []string{args.AccountHash.Val(), args.TransactionHash.Val(), args.OutputIndex.Hex()},
	}, ret)
}
