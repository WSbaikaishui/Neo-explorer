package api

import (
	"neophora/lib/type/h256"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetUTXOByHashLEIndexInHex ...
func (me *T) GetUTXOByHashLEIndexInHex(args struct {
	TransactionHashLE h256.T
	OutputIndex       uintval.T
}, ret *string) error {
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
		Index:  "h256.trx-uint.num",
		Keys:   []string{args.TransactionHashLE.RevVal(), args.OutputIndex.Hex()},
	}, ret)
}
