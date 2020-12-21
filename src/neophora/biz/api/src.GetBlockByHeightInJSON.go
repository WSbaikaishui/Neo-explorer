package api

import (
	"encoding/json"
	"fmt"
	"neophora/lib/trans"
)

// GetBlockByHeightInJSON ...
func (me *T) GetBlockByHeightInJSON(args struct {
	Height uint64
}, ret *json.RawMessage) error {
	var result []byte
	if err := me.Data.GetArgs(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.blk",
		Index:  "uint.hgt",
		Keys:   []string{fmt.Sprintf("%016x", args.Height)},
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
