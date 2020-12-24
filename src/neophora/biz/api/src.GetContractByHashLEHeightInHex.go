package api

import (
	"neophora/lib/type/h160"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetContractByHashLEHeightInHex ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetContractByHashLEHeightInHex(args struct {
	Hash   h160.T
	Height uintval.T
}, ret *string) error {
	if args.Hash.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.Height.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	return me.Data.GetLastValInHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.ctr",
		Index:  "h160.ctr-uint.hgt",
		Keys:   []string{args.Hash.RevVal(), args.Height.Hex()},
	}, ret)
}
