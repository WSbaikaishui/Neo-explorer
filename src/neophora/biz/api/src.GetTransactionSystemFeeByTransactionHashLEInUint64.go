package api

import (
	"neophora/lib/type/bins"
	"neophora/lib/type/h256"
	"neophora/var/stderr"
)

// GetTransactionSystemFeeByTransactionHashLEInUint64 ...
// as an example:
//
// ```
// $ curl https://example.neophora.io -d '{"id":1,"jsonrpc":"2.0","method":"GetTransactionSystemFeeByTransactionHashLEInUint64","params":{"TransactionHashLE": "da5e0d3cfbbc02e9cf863c1d97c56e1dc60114f98c957b308eea0eac76ca7ef8"}}'
// {"id":1,"result":0,"error":null}
// ```
func (me *T) GetTransactionSystemFeeByTransactionHashLEInUint64(args struct {
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
		Target: "uint.fos",
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
