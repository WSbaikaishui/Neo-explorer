package api

import (
	"encoding/json"
	"fmt"
	"neophora/lib/transex"
)

// GetHeaderByHeightInJSON ...
func (me *T) GetHeaderByHeightInJSON(args struct {
	Height uint64
}, ret *json.RawMessage) error {
	var result []byte
	var tr transex.T
	if err := me.Data.GetArgs(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "header",
		Index:  "height",
		Keys:   []string{fmt.Sprintf("%016x", args.Height)},
	}, &result); err != nil {
		return err
	}
	tr.V = result
	if err := tr.BytesToJSONViaHeader(); err != nil {
		return err
	}
	*ret = tr.V.(json.RawMessage)
	return nil
}
