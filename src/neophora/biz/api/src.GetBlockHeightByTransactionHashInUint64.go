package api

import (
	"neophora/lib/type/bins"
	"neophora/lib/type/h256"
	"neophora/var/stderr"
)

// GetBlockHeightByTransactionHashInUint64 ...
// as an example:
//
// ```
// $ curl https://example.neophora.io -d '{"jsonrpc": "2.0","id": 1,"method": "GetBlockHeightByTransactionHashInUint64","params":{"TransactionHash":"db252e9f0184224e20c0aa6c618758d32fd586064f70930cab4df061f809e2e2"}}'
// {"id":1,"result":3599999,"error":null}
// ```
func (me *T) GetBlockHeightByTransactionHashInUint64(args struct {
	TransactionHash h256.T
}, ret *uint64) error {
	if args.TransactionHash.Valid() == false {
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
