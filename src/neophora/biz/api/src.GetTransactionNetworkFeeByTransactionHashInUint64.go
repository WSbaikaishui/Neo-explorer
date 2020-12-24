package api

import (
	"neophora/lib/type/bins"
	"neophora/lib/type/h256"
	"neophora/var/stderr"
)

// GetTransactionNetworkFeeByTransactionHashInUint64 ...
func (me *T) GetTransactionNetworkFeeByTransactionHashInUint64(args struct {
	TransactionHash h256.T
}, ret *uint64) error {
	var result bins.T
	if err := me.Data.GetArgsInBins(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "uint.fon",
		Index:  "h256.trx",
		Keys:   []string{args.TransactionHash.Val()},
	}, &result); err != nil {
		return err
	}
	if result.Valid() == false {
		return stderr.ErrNotFound
	}
	*ret = result.Uint64()
	return nil
}
