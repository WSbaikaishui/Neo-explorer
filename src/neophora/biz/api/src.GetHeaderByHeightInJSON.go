package api

import (
	"encoding/json"
	"fmt"
	"neophora/lib/trans"
	"neophora/var/stderr"
)

// GetHeaderByHeightInJSON ...
func (me *T) GetHeaderByHeightInJSON(args struct {
	Height uint64
}, ret *json.RawMessage) error {
	var result []byte
	if err := me.Data.GetArgs(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.hdr",
		Index:  "uint.hgt",
		Keys:   []string{fmt.Sprintf("%016x", args.Height)},
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
