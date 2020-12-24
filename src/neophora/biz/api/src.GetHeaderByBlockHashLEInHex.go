package api

import (
	"neophora/lib/type/h256"
)

// GetHeaderByBlockHashLEInHex ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetHeaderByBlockHashLEInHex(args struct {
	BlockHashLE h256.T
}, ret *string) error {
	return me.Data.GetArgsInHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.hdr",
		Index:  "h256.blk",
		Keys:   []string{args.BlockHashLE.RevVal()},
	}, ret)
}
