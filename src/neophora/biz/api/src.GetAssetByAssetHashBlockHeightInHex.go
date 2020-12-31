package api

import (
	"neophora/lib/type/h256"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetAssetByAssetHashBlockHeightInHex ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetAssetByAssetHashBlockHeightInHex(args struct {
	AssetHash   h256.T
	BlockHeight uintval.T
}, ret *string) error {
	if args.AssetHash.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.BlockHeight.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	return me.Data.GetLastValInHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.ast",
		Index:  "h256.ast-uint.hgt",
		Keys:   []string{args.AssetHash.Val(), args.BlockHeight.Hex()},
	}, ret)
}
