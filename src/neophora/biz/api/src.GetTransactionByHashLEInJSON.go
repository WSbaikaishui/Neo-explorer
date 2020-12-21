package api

import (
	"encoding/json"
	"neophora/lib/trans"
)

// GetTransactionByHashLEInJSON ...
func (me *T) GetTransactionByHashLEInJSON(args struct {
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
		Target: "bins.trx",
		Index:  "h256.trx",
		Keys:   []string{tr.V.(string)},
	}, &result); err != nil {
		return err
	}
	tr.V = result
	if err := tr.BytesToJSONViaTX(); err != nil {
		return err
	}
	*ret = tr.V.(json.RawMessage)
	return nil
}
