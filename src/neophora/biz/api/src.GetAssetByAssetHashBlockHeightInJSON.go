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
// $ curl https://example.neophora.io -d '{"jsonrpc": "2.0","id": 1,"method": "GetAssetByAssetHashBlockHeightInJSON","params":{"AssetHash":"9b7cffdaa674beae0f930ebe6085af9093e5fe56b34a5c220ccdcf6efc336fc5","BlockHeight":2600000}}'
// {"id":1,"result":{"admin":"0x4b5acd30ba7ec77199561afa0bbd49b5e94517da","amount":"100000000","available":"100000000","expiration":4000000,"frozen":false,"id":"0xc56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b","issuer":"0x4b5acd30ba7ec77199561afa0bbd49b5e94517da","name":"[{\"lang\":\"zh-CN\",\"name\":\"小蚁股\"},{\"lang\":\"en\",\"name\":\"AntShare\"}]","owner":"00","precision":0,"type":"GoverningToken"},"error":null}
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
