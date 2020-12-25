package api

import (
	"neophora/lib/type/h256"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetStorageByContractHashKeyHashLEBlockHeightInHex ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetStorageByContractHashKeyHashLEBlockHeightInHex(args struct {
	ContractHash h256.T
	KeyHashLE    h256.T
	BlockHeight  uintval.T
}, ret *string) error {
	if args.ContractHash.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.KeyHashLE.Valid() == false {
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
		Keys:   []string{args.ContractHash.Val(), args.KeyHashLE.RevVal(), args.BlockHeight.Hex()},
	}, ret)
}
