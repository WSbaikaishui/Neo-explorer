package api

import (
	"neophora/lib/type/h256"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetUTXOByTransactionHashOutputIndexInHex ...
// as an example:
//
// ```
// $ curl https://example.neophora.io -d '{"id":1,"jsonrpc":"2.0","method":"GetUTXOByTransactionHashOutputIndexInHex","params":{"TransactionHash": "0f2c2397cbdf33794489dffae1f9f38f332e58810922ea7f56b0578b99365296", "OutputIndex": 0}}'
// {"id":1,"result":"e72d286979ee6cb1b7e65dfddfb2e384100b8d148e7758de42e4168b71792c6001000000000000007335f929546270b8f811a0f9427b5712457107e7","error":null}
// ```
func (me *T) GetUTXOByTransactionHashOutputIndexInHex(args struct {
	TransactionHash h256.T
	OutputIndex     uintval.T
}, ret *string) error {
	if args.TransactionHash.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.OutputIndex.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	return me.Data.GetArgsInHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.utx",
		Index:  "h256.trx-uint.num",
		Keys:   []string{args.TransactionHash.Val(), args.OutputIndex.Hex()},
	}, ret)
}
