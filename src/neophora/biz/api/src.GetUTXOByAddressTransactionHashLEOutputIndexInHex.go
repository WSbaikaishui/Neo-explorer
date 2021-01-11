package api

import (
	"neophora/lib/type/addr"
	"neophora/lib/type/h256"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetUTXOByAddressTransactionHashLEOutputIndexInHex ...
// as an example:
//
// ```
// $ curl https://example.neophora.io -d '{"jsonrpc": "2.0","id": 1,"method": "GetUTXOByAddressTransactionHashLEOutputIndexInHex","params":{"Address":"AZ1QiX5nqgm8dsUY7iRyafKwmKnGP9bUhN" ,"TransactionHashLE":"2323c44a37212d78ef302e64155217081f32553c1e57f0989f2b3ad737b46a4f","OutputIndex":0}}'
// {"id":1,"result":"9b7cffdaa674beae0f930ebe6085af9093e5fe56b34a5c220ccdcf6efc336fc500180d8f00000000bd097b2fcf70e1fd30a5c3ef51e662feeafeba01","error":null}
// ```
func (me *T) GetUTXOByAddressTransactionHashLEOutputIndexInHex(args struct {
	Address           addr.T
	TransactionHashLE h256.T
	OutputIndex       uintval.T
}, ret *string) error {
	if args.Address.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.TransactionHashLE.Valid() == false {
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
		Keys:   []string{args.Address.H160(), args.TransactionHashLE.RevVal(), args.OutputIndex.Hex()},
	}, ret)
}
