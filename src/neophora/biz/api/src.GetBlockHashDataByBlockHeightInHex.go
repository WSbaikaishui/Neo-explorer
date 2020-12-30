package api

import (
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetBlockHashDataByBlockHeightInHex ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetBlockHashDataByBlockHeightInHex(args struct {
	BlockHeight uintval.T
}, ret *string) error {
	if args.BlockHeight.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	return me.Data.GetLastValInHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.hdt",
		Index:  "uint.hgt",
		Keys:   []string{args.BlockHeight.Hex()},
	}, ret)
}
