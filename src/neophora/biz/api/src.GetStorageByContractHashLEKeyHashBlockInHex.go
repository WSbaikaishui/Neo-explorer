package api

import (
	"neophora/lib/type/h256"
	"neophora/var/stderr"
)

// GetStorageByContractHashLEKeyHashBlockInHex ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetStorageByContractHashLEKeyHashBlockInHex(args struct {
	ContractHashLE h256.T
	KeyHash        h256.T
}, ret *string) error {
	if args.ContractHashLE.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.KeyHash.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	return me.Data.GetLatestUint64ValInHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.str",
		Index:  "h160.ctr-h256.key-uint.hgt",
		Keys:   []string{args.ContractHashLE.RevVal(), args.KeyHash.Val()},
	}, ret)
}
