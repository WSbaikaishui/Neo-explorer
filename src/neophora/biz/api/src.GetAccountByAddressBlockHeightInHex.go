package api

import (
	"neophora/lib/type/addr"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetAccountByAddressBlockHeightInHex ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetAccountByAddressBlockHeightInHex(args struct {
	Address     addr.T
	BlockHeight uintval.T
}, ret *string) error {
	if args.Address.Valid() == false {
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
		Keys:   []string{args.Address.H160(), args.BlockHeight.Hex()},
	}, ret)
}
