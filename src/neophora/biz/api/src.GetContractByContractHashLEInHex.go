package api

import (
	"neophora/lib/type/h160"
	"neophora/var/stderr"
)

// GetContractByContractHashLEInHex ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetContractByContractHashLEInHex(args struct {
	ContractHashLE h160.T
}, ret *string) error {
	if args.ContractHashLE.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	return me.Data.GetLatestUint64ValInHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.ctr",
		Index:  "h160.ctr-uint.hgt",
		Keys:   []string{args.ContractHashLE.RevVal()},
	}, ret)
}
