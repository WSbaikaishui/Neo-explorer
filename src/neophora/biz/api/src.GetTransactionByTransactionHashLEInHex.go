package api

import (
	"neophora/lib/type/h256"
)

// GetTransactionByTransactionHashLEInHex ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetTransactionByTransactionHashLEInHex(args struct {
	TransactionHashLE h256.T
}, ret *string) error {
	return me.Data.GetArgsInHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.trx",
		Index:  "h256.trx",
		Keys:   []string{args.TransactionHashLE.RevVal()},
	}, ret)
}
