package api

import (
	"neophora/lib/type/h160"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetContractByContractHashLEBlockHeightInHex ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetContractByContractHashLEBlockHeightInHex(args struct {
	ContractHashLE h160.T
	BlockHeight    uintval.T
}, ret *string) error {
	if args.ContractHashLE.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.BlockHeight.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	return me.Data.GetLastValInHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.ctr",
		Index:  "h160.ctr-uint.hgt",
		Keys:   []string{args.ContractHashLE.RevVal(), args.BlockHeight.Hex()},
	}, ret)
}
