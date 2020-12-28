package api

import (
	"encoding/json"
	"neophora/lib/type/bins"
	"neophora/lib/type/h256"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetUTXOByTransactionHashOutputIndexInJSON ...
// as an example:
//
// ```
// $ curl https://example.neophora.io -d '{"id":1,"jsonrpc":"2.0","method":"GetUTXOByTransactionHashOutputIndexInJSON","params":{"TransactionHash": "0f2c2397cbdf33794489dffae1f9f38f332e58810922ea7f56b0578b99365296", "OutputIndex": 0}}'
// {"id":1,"result":{"address":"ASH41gtWftHvhuYhZz1jj7ee7z9vp9D9wk","asset":"0x602c79718b16e442de58778e148d0b1084e3b2dffd5de6b7b16cee7969282de7","value":"0.00000001"},"error":null}
// ```
func (me *T) GetUTXOByTransactionHashOutputIndexInJSON(args struct {
	TransactionHash h256.T
	OutputIndex     uintval.T
}, ret *json.RawMessage) error {
	if args.TransactionHash.Valid() == false {
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
		Index:  "h256.trx-uint.num",
		Keys:   []string{args.TransactionHash.Val(), args.OutputIndex.Hex()},
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
