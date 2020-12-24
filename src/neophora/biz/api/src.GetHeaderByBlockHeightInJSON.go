package api

import (
	"encoding/json"
	"neophora/lib/type/bins"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetHeaderByBlockHeightInJSON ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetHeaderByBlockHeightInJSON(args struct {
	BlockHeight uintval.T
}, ret *json.RawMessage) error {
	var result bins.T
	if err := me.Data.GetArgsInBins(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.hdr",
		Index:  "uint.hgt",
		Keys:   []string{args.BlockHeight.Hex()},
	}, &result); err != nil {
		return err
	}
	if result.Valid() == false {
		return stderr.ErrNotFound
	}
	js, err := result.JSONViaHeader()
	if err != nil {
		return stderr.ErrNotFound
	}
	*ret = js
	return nil
}
