package api

import (
	"encoding/json"
	"neophora/lib/type/bins"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetBlockByHeightInJSON ...
func (me *T) GetBlockByHeightInJSON(args struct {
	Height uintval.T
}, ret *json.RawMessage) error {
	if args.Height.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	var result bins.T
	if err := me.Data.GetArgsInBins(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.blk",
		Index:  "uint.hgt",
		Keys:   []string{args.Height.Hex()},
	}, &result); err != nil {
		return err
	}
	if result.Valid() == false {
		return stderr.ErrNotFound
	}
	js, err := result.JSONViaBlock()
	if err != nil {
		return stderr.ErrNotFound
	}
	*ret = js
	return nil
}
