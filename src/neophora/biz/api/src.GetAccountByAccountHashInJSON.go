package api

import (
	"encoding/json"
	"neophora/lib/type/bins"
	"neophora/lib/type/h160"
	"neophora/var/stderr"
)

// GetAccountByAccountHashInJSON ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetAccountByAccountHashInJSON(args struct {
	AccountHash h160.T
}, ret *json.RawMessage) error {
	if args.AccountHash.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	var result bins.T
	if err := me.Data.GetLastValInBins(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.act",
		Index:  "h160.act-uint.hgt",
		Keys:   []string{args.AccountHash.Val()},
	}, &result); err != nil {
		return err
	}
	if result.Valid() == false {
		return stderr.ErrNotFound
	}
	js, err := result.JSONViaAccount()
	if err != nil {
		return stderr.ErrNotFound
	}
	*ret = js
	return nil
}
