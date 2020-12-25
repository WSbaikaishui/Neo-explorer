package api

import (
	"neophora/lib/type/h256"
	"neophora/var/stderr"
)

// GetStorageByContractHashKeyHashBlockInHex ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetStorageByContractHashKeyHashBlockInHex(args struct {
	ContractHash h256.T
	KeyHash      h256.T
}, ret *string) error {
	if args.ContractHash.Valid() == false {
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
		Keys:   []string{args.ContractHash.Val(), args.KeyHash.Val()},
	}, ret)
}
