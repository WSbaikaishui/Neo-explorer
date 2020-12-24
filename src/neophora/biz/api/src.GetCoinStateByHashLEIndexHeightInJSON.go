package api

import (
	"encoding/json"
	"neophora/lib/type/bins"
	"neophora/lib/type/h256"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetCoinStateByHashLEIndexHeightInJSON ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetCoinStateByHashLEIndexHeightInJSON(args struct {
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
		Target: "uint.cst",
		Index:  "h256.trx-uint.num-uint.hgt",
		Keys:   []string{args.Hash.RevVal(), args.Index.Hex(), args.Height.Hex()},
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
