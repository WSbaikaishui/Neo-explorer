package api

import (
	"encoding/json"
	"neophora/lib/type/bins"
	"neophora/lib/type/h256"
	"neophora/var/stderr"
)

// GetAssetByAssetHashLEInJSON ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetAssetByAssetHashLEInJSON(args struct {
	AssetHashLE h256.T
}, ret *json.RawMessage) error {
	if args.AssetHashLE.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	var result bins.T
	if err := me.Data.GetLatestUint64ValInBins(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.ast",
		Index:  "h256.ast",
		Keys:   []string{args.AssetHashLE.RevVal()},
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
