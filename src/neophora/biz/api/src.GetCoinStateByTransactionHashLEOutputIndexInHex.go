package api

import (
	"neophora/lib/type/h256"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetCoinStateByTransactionHashLEOutputIndexInHex ...
// as an example:
//
// ```
// $ curl https://example.neophora.io -d '{"jsonrpc": "2.0","id": 1,"method": "GetCoinStateByTransactionHashLEOutputIndexInHex","params":{"TransactionHashLE":"e2e209f861f04dab0c93704f0686d52fd35887616caac0204e2284019f2e25db","OutputIndex":0}}'
// {"id":1,"result":"0000000000000003","error":null}
// ```
func (me *T) GetCoinStateByTransactionHashLEOutputIndexInHex(args struct {
	TransactionHashLE h256.T
	OutputIndex       uintval.T
}, ret *string) error {
	if args.TransactionHashLE.Valid() == false {
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
		Target: "uint.cst",
		Index:  "h256.trx-uint.num-uint.hgt",
		Keys:   []string{args.TransactionHashLE.RevVal(), args.OutputIndex.Hex()},
	}, ret)
}
