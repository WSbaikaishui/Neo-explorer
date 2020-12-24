package api

import (
	"encoding/json"
	"neophora/lib/trans"
	"neophora/lib/type/h256"
	"neophora/var/stderr"
)

// GetTransactionByTransactionHashInJSON ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetTransactionByTransactionHashInJSON(args struct {
	TransactionHash h256.T
}, ret *json.RawMessage) error {
	var result []byte
	if err := me.Data.GetArgs(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.trx",
		Index:  "h256.trx",
		Keys:   []string{args.TransactionHash.Val()},
	}, &result); err != nil {
		return err
	}
	if result == nil {
		return stderr.ErrNotFound
	}
	tr := &trans.T{V: result}
	if err := tr.BytesToJSONViaTX(); err != nil {
		return err
	}
	*ret = tr.V.(json.RawMessage)
	return nil
}
