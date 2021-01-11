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
// $ curl https://example.neophora.io -d '{"jsonrpc": "2.0","id": 1,"method": "GetAssetByAssetHashLEInJSON","params":{"AssetHashLE":"c56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b"}}'
// {"id":1,"result":{"admin":"0x4b5acd30ba7ec77199561afa0bbd49b5e94517da","amount":"100000000","available":"100000000","expiration":4000000,"frozen":false,"id":"0xc56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b","issuer":"0x4b5acd30ba7ec77199561afa0bbd49b5e94517da","name":"[{\"lang\":\"zh-CN\",\"name\":\"小蚁股\"},{\"lang\":\"en\",\"name\":\"AntShare\"}]","owner":"00","precision":0,"type":"GoverningToken"},"error":null}
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
		Index:  "h256.ast-uint.hgt",
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
