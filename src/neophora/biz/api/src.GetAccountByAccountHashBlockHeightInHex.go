package api

import (
	"neophora/lib/type/h160"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetAccountByAccountHashBlockHeightInHex ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetAccountByAccountHashBlockHeightInHex(args struct {
	AccountHash h160.T
	BlockHeight uintval.T
}, ret *string) error {
	if args.AccountHash.Valid() == false {
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
		Keys:   []string{args.AccountHash.Val(), args.BlockHeight.Hex()},
	}, ret)
}
