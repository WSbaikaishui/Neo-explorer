package api

import (
	"encoding/json"
	"neophora/lib/trans"
	"neophora/var/stderr"
)

// GetHeaderByHashLEInJSON ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetHeaderByHashLEInJSON(args struct {
	Hash string
}, ret *json.RawMessage) error {
	var result []byte
	tr := &trans.T{V: args.Hash}
	if err := tr.HexReverse(); err != nil {
		return err
	}
	if err := me.Data.GetArgs(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.hdr",
		Index:  "h256.blk",
		Keys:   []string{tr.V.(string)},
	}, &result); err != nil {
		return err
	}
	if result == nil {
		return stderr.ErrNotFound
	}
	tr.V = result
	if err := tr.BytesToJSONViaHeader(); err != nil {
		return err
	}
	*ret = tr.V.(json.RawMessage)
	return nil
}
