package api

import (
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetBlockByHeightInHex ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetBlockByHeightInHex(args struct {
	Height uintval.T
}, ret *string) error {
	if args.Height.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	return me.Data.GetArgsInHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.blk",
		Index:  "uint.hgt",
		Keys:   []string{args.Height.Hex()},
	}, ret)
}
