package api

import (
	"encoding/json"
	"neophora/lib/trans"
)

// GetBlockByHashInJSON ...
func (me *T) GetBlockByHashInJSON(args struct {
	Hash string
}, ret *json.RawMessage) error {
	var result []byte
	if err := me.Data.GetArgs(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.blk",
		Index:  "h256.blk",
		Keys:   []string{args.Hash},
	}, &result); err != nil {
		return err
	}
	tr := &trans.T{V: result}
	if err := tr.BytesToJSONViaBlock(); err != nil {
		return err
	}
	*ret = tr.V.(json.RawMessage)
	return nil
}
