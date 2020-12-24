package api

import (
	"neophora/lib/type/uintval"
)

// GetHeaderByBlockHeightInHex ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetHeaderByBlockHeightInHex(args struct {
	BlockHeight uintval.T
}, ret *string) error {
	return me.Data.GetArgsInHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.hdr",
		Index:  "uint.hgt",
		Keys:   []string{args.BlockHeight.Hex()},
	}, ret)
}
