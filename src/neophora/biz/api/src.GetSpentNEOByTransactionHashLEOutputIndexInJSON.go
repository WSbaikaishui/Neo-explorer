package api

import (
	"encoding/json"
	"neophora/lib/type/bins"
	"neophora/lib/type/h256"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetSpentNEOByTransactionHashLEOutputIndexInJSON ...
// as an example:
//
// ```
// $ curl https://example.neophora.io -d '{"jsonrpc": "2.0","id": 1,"method": "GetSpentNEOByTransactionHashLEOutputIndexInJSON","params":{"TransactionHashLE":"22e1a1617f0d53d6f8e504d6d663d5c33f07698b4e3cb0006019710761dd0e16","OutputIndex":0}}'
// {"id":1,"result":{"burn":2475663,"mint":2400029},"error":null}
// ```
func (me *T) GetSpentNEOByTransactionHashLEOutputIndexInJSON(args struct {
	TransactionHashLE h256.T
	OutputIndex       uintval.T
}, ret *json.RawMessage) error {
	if args.TransactionHashLE.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.OutputIndex.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	var result bins.T
	if err := me.Data.GetLatestUint64ValInBins(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "u128.spt",
		Index:  "h256.trx-uint.num-uint.hgt",
		Keys:   []string{args.TransactionHashLE.RevVal(), args.OutputIndex.Hex()},
	}, &result); err != nil {
		return err
	}
	if result.Valid() == false {
		return stderr.ErrNotFound
	}
	js, err := result.JSONViaSpentNEO()
	if err != nil {
		return stderr.ErrNotFound
	}
	*ret = js
	return nil
}
