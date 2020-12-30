package api

import (
	"encoding/json"
	"neophora/lib/type/bins"
	"neophora/lib/type/h160"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetAccountByAccountHashLEBlockHeightInJSON ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetAccountByAccountHashLEBlockHeightInJSON(args struct {
	AccountHashLE h160.T
	BlockHeight   uintval.T
}, ret *json.RawMessage) error {
	if args.AccountHashLE.Valid() == false {
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
		Target: "bins.act",
		Index:  "h160.act-uint.hgt",
		Keys:   []string{args.AccountHashLE.RevVal(), args.BlockHeight.Hex()},
	}, &result); err != nil {
		return err
	}
	if result.Valid() == false {
		return stderr.ErrNotFound
	}
	js, err := result.JSONViaAccount()
	if err != nil {
		return stderr.ErrNotFound
	}
	*ret = js
	return nil
}
