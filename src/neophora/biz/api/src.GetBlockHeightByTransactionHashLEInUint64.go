package api

import (
	"neophora/lib/type/bins"
	"neophora/lib/type/h256"
	"neophora/var/stderr"
)

// GetBlockHeightByTransactionHashLEInUint64 ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetBlockHeightByTransactionHashLEInUint64(args struct {
	TransactionHashLE h256.T
}, ret *uint64) error {
	if args.TransactionHashLE.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	var result bins.T
	if err := me.Data.GetArgsInBins(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "uint.hgt",
		Index:  "h256.trx",
		Keys:   []string{args.TransactionHashLE.RevVal()},
	}, &result); err != nil {
		return err
	}
	if result.Valid() == false {
		return stderr.ErrNotFound
	}
	*ret = result.Uint64()
	return nil
}
