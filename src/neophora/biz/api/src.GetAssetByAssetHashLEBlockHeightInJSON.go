package api

import (
	"encoding/json"
	"neophora/lib/type/bins"
	"neophora/lib/type/h256"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetAssetByAssetHashLEBlockHeightInJSON ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetAssetByAssetHashLEBlockHeightInJSON(args struct {
	AssetHashLE h256.T
	BlockHeight uintval.T
}, ret *json.RawMessage) error {
	if args.AssetHashLE.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.BlockHeight.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	var result bins.T
	if err := me.Data.GetLastValInBins(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.ast",
		Index:  "h256.ast-uint.hgt",
		Keys:   []string{args.AssetHashLE.RevVal(), args.BlockHeight.Hex()},
	}, &result); err != nil {
		return err
	}
	if result.Valid() == false {
		return stderr.ErrNotFound
	}
	js, err := result.JSONViaAsset()
	if err != nil {
		return stderr.ErrNotFound
	}
	*ret = js
	return nil
}
