package api

import "neophora/lib/type/uintval"

// GetBlockHashByBlockHeightInHex ...
// as an example:
//
// ```
// $ curl https://example.neophora.io -d '{"jsonrpc": "2.0","id": 1,"method": "GetBlockHashByBlockHeightInHex","params":{"BlockHeight":2100000}}'
// {"id":1,"result":"c04f6eb83a783a3a3b50aca65a7f4c138c61b1f325bd4330cd75ad2193e363a1","error":null}
// ```
func (me *T) GetBlockHashByBlockHeightInHex(args struct {
	BlockHeight uintval.T
}, ret *string) error {
	return me.Data.GetArgsInHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "h256.blk",
		Index:  "uint.hgt",
		Keys:   []string{args.BlockHeight.Hex()},
	}, ret)
}
