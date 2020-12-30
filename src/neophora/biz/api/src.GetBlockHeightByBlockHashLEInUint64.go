package api

import (
	"neophora/lib/type/bins"
	"neophora/lib/type/h256"
	"neophora/var/stderr"
)

// GetBlockHeightByBlockHashLEInUint64 ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetBlockHeightByBlockHashLEInUint64(args struct {
	BlockHashLE h256.T
}, ret *uint64) error {
	if args.BlockHashLE.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	var result bins.T
	if err := me.Data.GetArgsInBins(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "uint.hgt",
		Index:  "h256.blk",
		Keys:   []string{args.BlockHashLE.RevVal()},
	}, &result); err != nil {
		return err
	}
	if result.Valid() == false {
		return stderr.ErrNotFound
	}
	*ret = result.Uint64()
	return nil
}
