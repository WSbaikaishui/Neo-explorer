package api

import (
	"encoding/json"
	"neophora/lib/type/bins"
	"neophora/lib/type/h256"
	"neophora/var/stderr"
)

// GetContractByContractLEHashInJSON ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetContractByContractLEHashInJSON(args struct {
	ContractHashLE h256.T
}, ret *json.RawMessage) error {
	if args.ContractHashLE.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	var result bins.T
	if err := me.Data.GetLatestUint64ValInBins(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.ctr",
		Index:  "h160.ctr-uint.hgt",
		Keys:   []string{args.ContractHashLE.RevVal()},
	}, &result); err != nil {
		return err
	}
	if result.Valid() == false {
		return stderr.ErrNotFound
	}
	js, err := result.JSONViaContract()
	if err != nil {
		return stderr.ErrNotFound
	}
	*ret = js
	return nil
}
