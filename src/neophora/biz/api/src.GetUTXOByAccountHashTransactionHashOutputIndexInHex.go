package api

import (
	"neophora/lib/type/h160"
	"neophora/lib/type/h256"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetUTXOByAccountHashTransactionHashOutputIndexInHex ...
// as an example:
//
// ```
// $ curl https://example.neophora.io -d '{"jsonrpc": "2.0","id": 1,"method": "GetUTXOByAccountHashTransactionHashOutputIndexInHex","params":{"AccountHash":"bd097b2fcf70e1fd30a5c3ef51e662feeafeba01","TransactionHash":"4f6ab437d73a2b9f98f0571e3c55321f08175215642e30ef782d21374ac42323","OutputIndex":0}}'
// {"id":1,"result":"9b7cffdaa674beae0f930ebe6085af9093e5fe56b34a5c220ccdcf6efc336fc500180d8f00000000bd097b2fcf70e1fd30a5c3ef51e662feeafeba01","error":null}
// ```
func (me *T) GetUTXOByAccountHashTransactionHashOutputIndexInHex(args struct {
	AccountHash     h160.T
	TransactionHash h256.T
	OutputIndex     uintval.T
}, ret *string) error {
	if args.AccountHash.Valid() == false {
		return stderr.ErrInvalidArgs
	}
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
		Index:  "h160.act-h256.trx-uint.num",
		Keys:   []string{args.AccountHash.Val(), args.TransactionHash.Val(), args.OutputIndex.Hex()},
	}, ret)
}
