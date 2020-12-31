package api

import (
	"neophora/lib/type/h160"
	"neophora/lib/type/h256"
	"neophora/var/stderr"
)

// GetStorageByContractHashKeyHashLEBlockInHex ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetStorageByContractHashKeyHashLEBlockInHex(args struct {
	ContractHash h160.T
	KeyHashLE    h256.T
}, ret *string) error {
	if args.ContractHash.Valid() == false {
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
		Keys:   []string{args.ContractHash.Val(), args.KeyHashLE.RevVal()},
	}, ret)
}
