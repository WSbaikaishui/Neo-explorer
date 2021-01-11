package api

import (
	"neophora/lib/type/h256"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetCoinStateByTransactionHashLEOutputIndexBlockHeightInHex ...
// as an example:
//
// ```
// $ curl https://example.neophora.io -d '{"jsonrpc": "2.0","id": 1,"method": "GetCoinStateByTransactionHashLEOutputIndexBlockHeightInHex","params":{"TransactionHashLE":"e2e209f861f04dab0c93704f0686d52fd35887616caac0204e2284019f2e25db","OutputIndex":0, "BlockHeight":3599999}}'
// {"id":1,"result":"0000000000000001","error":null}
// ```
func (me *T) GetCoinStateByTransactionHashLEOutputIndexBlockHeightInHex(args struct {
	TransactionHashLE h256.T
	OutputIndex       uintval.T
	BlockHeight       uintval.T
}, ret *string) error {
	if args.TransactionHashLE.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.OutputIndex.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.BlockHeight.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	return me.Data.GetLastValInHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "uint.cst",
		Index:  "h256.trx-uint.num-uint.hgt",
		Keys:   []string{args.TransactionHashLE.RevVal(), args.OutputIndex.Hex(), args.BlockHeight.Hex()},
	}, ret)
}
