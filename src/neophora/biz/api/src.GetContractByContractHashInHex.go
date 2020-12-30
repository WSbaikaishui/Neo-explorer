package api

import (
	"neophora/lib/type/h160"
	"neophora/var/stderr"
)

// GetContractByContractHashInHex ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetContractByContractHashInHex(args struct {
	ContractHash h160.T
}, ret *string) error {
	if args.ContractHash.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	return me.Data.GetLatestUint64ValInHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.ctr",
		Index:  "h160.ctr-uint.hgt",
		Keys:   []string{args.ContractHash.Val()},
	}, ret)
}
