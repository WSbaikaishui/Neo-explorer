package api

import (
	"encoding/json"
	"neophora/lib/type/addr"
	"neophora/lib/type/bins"
	"neophora/var/stderr"
)

// GetAccountByAddressInJSON ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetAccountByAddressInJSON(args struct {
	Address addr.T
}, ret *json.RawMessage) error {
	if args.Address.Valid() == false {
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
		Keys:   []string{args.Address.H160()},
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
