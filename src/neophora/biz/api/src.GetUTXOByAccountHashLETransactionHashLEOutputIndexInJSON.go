package api

import (
	"encoding/json"
	"neophora/lib/type/bins"
	"neophora/lib/type/h160"
	"neophora/lib/type/h256"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetUTXOByAccountHashLETransactionHashLEOutputIndexInJSON ...
// as an example:
//
// ```
// $ curl https://example.neophora.io -d '{"jsonrpc": "2.0","id": 1,"method": "GetUTXOByAccountHashLETransactionHashLEOutputIndexInJSON","params":{"AccountHashLE":"01bafeeafe62e651efc3a530fde170cf2f7b09bd","TransactionHashLE":"2323c44a37212d78ef302e64155217081f32553c1e57f0989f2b3ad737b46a4f","OutputIndex":0}}'
// {"id":1,"result":{"address":"AZ1QiX5nqgm8dsUY7iRyafKwmKnGP9bUhN","asset":"0xc56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b","value":"24"},"error":null}
// ```
func (me *T) GetUTXOByAccountHashLETransactionHashLEOutputIndexInJSON(args struct {
	AccountHashLE     h160.T
	TransactionHashLE h256.T
	OutputIndex       uintval.T
}, ret *json.RawMessage) error {
	if args.AccountHashLE.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.TransactionHashLE.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.OutputIndex.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	var result bins.T
	if err := me.Data.GetArgsInBins(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.utx",
		Index:  "h160.act-h256.trx-uint.num",
		Keys:   []string{args.AccountHashLE.RevVal(), args.TransactionHashLE.RevVal(), args.OutputIndex.Hex()},
	}, &result); err != nil {
		return err
	}
	if result.Valid() == false {
		return stderr.ErrNotFound
	}
	js, err := result.JSONViaUTXO()
	if err != nil {
		return stderr.ErrNotFound
	}
	*ret = js
	return nil
}
