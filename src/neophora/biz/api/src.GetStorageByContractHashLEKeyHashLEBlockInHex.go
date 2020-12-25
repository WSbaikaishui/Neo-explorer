package api

import (
	"neophora/lib/type/h256"
	"neophora/var/stderr"
)

// GetStorageByContractHashLEKeyHashLEBlockInHex ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetStorageByContractHashLEKeyHashLEBlockInHex(args struct {
	ContractHashLE h256.T
	KeyHashLE      h256.T
}, ret *string) error {
	if args.ContractHashLE.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.KeyHashLE.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	return me.Data.GetLatestUint64ValInHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.str",
		Index:  "h160.ctr-h256.key-uint.hgt",
		Keys:   []string{args.ContractHashLE.RevVal(), args.KeyHashLE.RevVal()},
	}, ret)
}
