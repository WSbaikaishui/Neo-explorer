package api

import (
	"neophora/lib/type/h256"
	"neophora/var/stderr"
)

// GetBlockHashDataByBlockHashLEInHex ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetBlockHashDataByBlockHashLEInHex(args struct {
	BlockHashLE h256.T
}, ret *string) error {
	if args.BlockHashLE.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	return me.Data.GetArgsInHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.hdt",
		Index:  "h256.blk",
		Keys:   []string{args.BlockHashLE.RevVal()},
	}, ret)
}
