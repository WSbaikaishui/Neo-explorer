package api

import (
	"encoding/json"
	"neophora/lib/type/bins"
	"neophora/lib/type/h256"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetAssetByAssetHashBlockHeightInJSON ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetAssetByAssetHashBlockHeightInJSON(args struct {
	AssetHash   h256.T
	BlockHeight uintval.T
}, ret *json.RawMessage) error {
	if args.AssetHash.Valid() == false {
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
		Keys:   []string{args.AssetHash.Val(), args.BlockHeight.Hex()},
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
