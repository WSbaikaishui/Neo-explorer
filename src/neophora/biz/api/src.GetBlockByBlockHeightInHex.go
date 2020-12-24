package api

import (
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetBlockByBlockHeightInHex ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetBlockByBlockHeightInHex(args struct {
	BlockHeight uintval.T
}, ret *string) error {
	if args.BlockHeight.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	return me.Data.GetArgsInHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.blk",
		Index:  "uint.hgt",
		Keys:   []string{args.BlockHeight.Hex()},
	}, ret)
}
