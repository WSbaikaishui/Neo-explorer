package api

import (
	"neophora/lib/type/h256"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetStorageByContractHashLEKeyHashBlockHeightInHex ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetStorageByContractHashLEKeyHashBlockHeightInHex(args struct {
	ContractHashLE h256.T
	KeyHash        h256.T
	BlockHeight    uintval.T
}, ret *string) error {
	if args.ContractHashLE.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.KeyHash.Valid() == false {
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
		Target: "bins.str",
		Index:  "h160.ctr-h256.key-uint.hgt",
		Keys:   []string{args.ContractHashLE.RevVal(), args.KeyHash.Val(), args.BlockHeight.Hex()},
	}, ret)
}
