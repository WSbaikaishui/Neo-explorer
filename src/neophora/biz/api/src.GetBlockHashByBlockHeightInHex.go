package api

import "neophora/lib/type/uintval"

// GetBlockHashByBlockHeightInHex ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetBlockHashByBlockHeightInHex(args struct {
	BlockHeight uintval.T
}, ret *string) error {
	return me.Data.GetArgsInHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "h256.blk",
		Index:  "uint.hgt",
		Keys:   []string{args.BlockHeight.Hex()},
	}, ret)
}
