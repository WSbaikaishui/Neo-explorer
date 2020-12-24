package api

import (
	"encoding/json"
	"neophora/lib/type/bins"
	"neophora/lib/type/h256"
	"neophora/var/stderr"
)

// GetTransactionByTransactionHashLEInJSON ...
func (me *T) GetTransactionByTransactionHashLEInJSON(args struct {
	TransactionHashLE h256.T
}, ret *json.RawMessage) error {
	var result bins.T
	if err := me.Data.GetArgsInBins(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.trx",
		Index:  "h256.trx",
		Keys:   []string{args.TransactionHashLE.RevVal()},
	}, &result); err != nil {
		return err
	}
	if result.Valid() == false {
		return stderr.ErrNotFound
	}
	js, err := result.JSONViaTransaction()
	if err != nil {
		return stderr.ErrNotFound
	}
	*ret = js
	return nil
}
