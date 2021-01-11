package api

import (
	"encoding/json"
	"neophora/lib/type/bins"
	"neophora/lib/type/h256"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetCoinStateByTransactionHashOutputIndexBlockHeightInJSON ...
// as an example:
//
// ```
// $ curl https://example.neophora.io -d '{"jsonrpc": "2.0","id": 1,"method": "GetCoinStateByTransactionHashOutputIndexBlockHeightInJSON","params":{"TransactionHash":"db252e9f0184224e20c0aa6c618758d32fd586064f70930cab4df061f809e2e2","OutputIndex":0, "BlockHeight":3599999}}'
// {"id":1,"result":"Untracted State: 0000000000000001","error":null}
// ```
func (me *T) GetCoinStateByTransactionHashOutputIndexBlockHeightInJSON(args struct {
	TransactionHash h256.T
	OutputIndex     uintval.T
	BlockHeight     uintval.T
}, ret *json.RawMessage) error {
	if args.TransactionHash.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.OutputIndex.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.BlockHeight.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	var result bins.T
	if err := me.Data.GetLastValInBins(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "uint.cst",
		Index:  "h256.trx-uint.num-uint.hgt",
		Keys:   []string{args.TransactionHash.Val(), args.OutputIndex.Hex(), args.BlockHeight.Hex()},
	}, &result); err != nil {
		return err
	}
	if result.Valid() == false {
		return stderr.ErrNotFound
	}
	js, err := result.JSONViaCoinState()
	if err != nil {
		return stderr.ErrNotFound
	}
	*ret = js
	return nil
}
