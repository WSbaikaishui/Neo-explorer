package api

import (
	"neophora/lib/type/h256"
	"neophora/var/stderr"
)

// GetAssetByAssetHashInHex ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetAssetByAssetHashInHex(args struct {
	AssetHash h256.T
}, ret *string) error {
	if args.AssetHash.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	return me.Data.GetLatestUint64ValInHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.ast",
		Index:  "h256.ast-uint.hgt",
		Keys:   []string{args.AssetHash.Val()},
	}, ret)
}
