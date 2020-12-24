package api

import "neophora/lib/type/h256"

// GetTransactionByTransactionHashInHex ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetTransactionByTransactionHashInHex(args struct {
	TransactionHash h256.T
}, ret *string) error {
	return me.Data.GetArgsInHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.trx",
		Index:  "h256.trx",
		Keys:   []string{args.TransactionHash.Val()},
	}, ret)
}
