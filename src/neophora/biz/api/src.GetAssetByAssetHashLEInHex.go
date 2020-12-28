package api

import (
	"neophora/lib/type/h256"
	"neophora/var/stderr"
)

// GetAssetByAssetHashLEInHex ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetAssetByAssetHashLEInHex(args struct {
	AssetHashLE h256.T
}, ret *string) error {
	if args.AssetHashLE.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	return me.Data.GetLatestUint64ValInHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.ast",
		Index:  "h256.ast-uint.hgt",
		Keys:   []string{args.AssetHashLE.RevVal()},
	}, ret)
}
