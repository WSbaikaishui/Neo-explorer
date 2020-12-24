package api

import (
	"encoding/json"
	"neophora/lib/type/bins"
	"neophora/lib/type/h256"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetSpentNEOByHashIndexHeightInJSON ...
func (me *T) GetSpentNEOByHashIndexHeightInJSON(args struct {
	Hash   h256.T
	Index  uintval.T
	Height uintval.T
}, ret *json.RawMessage) error {
	if args.Hash.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.Index.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.Height.Valid() == false {
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
		Keys:   []string{args.Hash.Val(), args.Index.Hex(), args.Height.Hex()},
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
