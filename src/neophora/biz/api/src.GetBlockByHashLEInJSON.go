package api

import (
	"encoding/json"
	"neophora/lib/trans"
)

// GetBlockByHashLEInJSON ...
func (me *T) GetBlockByHashLEInJSON(args struct {
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
		Target: "bins.blk",
		Index:  "h256.blk",
		Keys:   []string{tr.V.(string)},
	}, &result); err != nil {
		return err
	}
	tr.V = result
	if err := tr.BytesToJSONViaBlock(); err != nil {
		return err
	}
	*ret = tr.V.(json.RawMessage)
	return nil
}
