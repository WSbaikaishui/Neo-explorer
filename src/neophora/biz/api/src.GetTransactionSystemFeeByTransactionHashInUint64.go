package api

import (
	"neophora/lib/type/bins"
	"neophora/lib/type/h256"
	"neophora/var/stderr"
)

// GetTransactionSystemFeeByTransactionHashInUint64 ...
// as an example:
//
// ```
// $ curl https://example.neophora.io -d '{"id":1,"jsonrpc":"2.0","method":"GetTransactionSystemFeeByTransactionHashInUint64","params":{"TransactionHash": "f87eca76ac0eea8e307b958cf91401c61d6ec5971d3c86cfe902bcfb3c0d5eda"}}'
// {"id":1,"result":0,"error":null}
// ```
func (me *T) GetTransactionSystemFeeByTransactionHashInUint64(args struct {
	TransactionHash h256.T
}, ret *uint64) error {
	var result bins.T
	if err := me.Data.GetArgsInBins(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "uint.fos",
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
