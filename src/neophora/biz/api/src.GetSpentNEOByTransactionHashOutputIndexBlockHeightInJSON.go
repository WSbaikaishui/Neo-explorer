package api

import (
	"encoding/json"
	"neophora/lib/type/bins"
	"neophora/lib/type/h256"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetSpentNEOByTransactionHashOutputIndexBlockHeightInJSON ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetSpentNEOByTransactionHashOutputIndexBlockHeightInJSON(args struct {
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
		Target: "u128.spt",
		Index:  "h256.trx-uint.num-uint.hgt",
		Keys:   []string{args.TransactionHash.Val(), args.OutputIndex.Hex(), args.BlockHeight.Hex()},
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
