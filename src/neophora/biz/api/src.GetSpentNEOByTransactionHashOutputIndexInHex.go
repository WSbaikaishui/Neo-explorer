package api

import (
	"neophora/lib/type/h256"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetSpentNEOByTransactionHashOutputIndexInHex ...
// as an example:
//
// ```
// $ curl https://example.neophora.io -d '{"jsonrpc": "2.0","id": 1,"method": "GetSpentNEOByTransactionHashOutputIndexInHex","params":{"TransactionHash":"160edd610771196000b03c4e8b69073fc3d563d6d604e5f8d6530d7f61a1e122","OutputIndex":0}}'
// {"id":1,"result":"0000000000249f1d000000000025c68f","error":null}
// ```
func (me *T) GetSpentNEOByTransactionHashOutputIndexInHex(args struct {
	TransactionHash h256.T
	OutputIndex     uintval.T
}, ret *string) error {
	if args.TransactionHash.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.OutputIndex.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	return me.Data.GetLatestUint64ValInHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "u128.spt",
		Index:  "h256.trx-uint.num-uint.hgt",
		Keys:   []string{args.TransactionHash.Val(), args.OutputIndex.Hex()},
	}, ret)
}
