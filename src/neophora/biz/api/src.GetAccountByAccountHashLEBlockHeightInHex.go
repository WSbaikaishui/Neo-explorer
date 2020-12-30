package api

import (
	"neophora/lib/type/h160"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetAccountByAccountHashLEBlockHeightInHex ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetAccountByAccountHashLEBlockHeightInHex(args struct {
	AccountHashLE h160.T
	BlockHeight   uintval.T
}, ret *string) error {
	if args.AccountHashLE.Valid() == false {
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
		Target: "bins.act",
		Index:  "h160.act-uint.hgt",
		Keys:   []string{args.AccountHashLE.RevVal(), args.BlockHeight.Hex()},
	}, ret)
}
