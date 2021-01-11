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
// $ curl https://example.neophora.io -d '{"jsonrpc": "2.0","id": 1,"method": "GetAssetByAssetHashBlockHeightInHex","params":{"AssetHash":"9b7cffdaa674beae0f930ebe6085af9093e5fe56b34a5c220ccdcf6efc336fc5","BlockHeight":2600000}}'
// {"id":1,"result":"009b7cffdaa674beae0f930ebe6085af9093e5fe56b34a5c220ccdcf6efc336fc500455b7b226c616e67223a227a682d434e222c226e616d65223a22e5b08fe89a81e882a1227d2c7b226c616e67223a22656e222c226e616d65223a22416e745368617265227d5d0000c16ff28623000000c16ff286230000000000000000000000000000000000000000000000000000000000000000da1745e9b549bd0bfa1a569971c77eba30cd5a4bda1745e9b549bd0bfa1a569971c77eba30cd5a4b00093d0000","error":null}
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
