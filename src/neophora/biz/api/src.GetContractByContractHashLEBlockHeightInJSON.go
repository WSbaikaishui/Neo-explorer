package api

import (
	"encoding/json"
	"neophora/lib/type/bins"
	"neophora/lib/type/h256"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetContractByContractHashLEBlockHeightInJSON ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetContractByContractHashLEBlockHeightInJSON(args struct {
	ContractHashLE h256.T
	BlockHeight    uintval.T
}, ret *json.RawMessage) error {
	if args.ContractHashLE.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.BlockHeight.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	var result bins.T
	if err := me.Data.GetLastValInBins(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.ctr",
		Index:  "h160.ctr-uint.hgt",
		Keys:   []string{args.ContractHashLE.RevVal(), args.BlockHeight.Hex()},
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
