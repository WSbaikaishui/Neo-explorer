package api

import (
	"encoding/json"
	"neophora/lib/trans"
	"neophora/var/stderr"
)

// GetHeaderByHashInJSON ...
func (me *T) GetHeaderByHashInJSON(args struct {
	Hash string
}, ret *json.RawMessage) error {
	var result []byte
	if err := me.Data.GetArgs(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.hdr",
		Index:  "h256.blk",
		Keys:   []string{args.Hash},
	}, &result); err != nil {
		return err
	}
	if result == nil {
		return stderr.ErrNotFound
	}
	tr := &trans.T{V: result}
	if err := tr.BytesToJSONViaHeader(); err != nil {
		return err
	}
	*ret = tr.V.(json.RawMessage)
	return nil
}
