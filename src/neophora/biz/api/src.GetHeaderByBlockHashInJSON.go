package api

import (
	"encoding/json"
	"neophora/lib/type/bins"
	"neophora/lib/type/h256"
	"neophora/var/stderr"
)

// GetHeaderByBlockHashInJSON ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetHeaderByBlockHashInJSON(args struct {
	BlockHash h256.T
}, ret *json.RawMessage) error {
	var result bins.T
	if err := me.Data.GetArgsInBins(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.hdr",
		Index:  "h256.blk",
		Keys:   []string{args.BlockHash.Val()},
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
