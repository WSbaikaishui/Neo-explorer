package api

import (
	"neophora/lib/type/h256"
	"neophora/var/stderr"
)

// GetBlockHashDataByBlockHashInHex ...
// as an example:
//
// ```
// $ curl https://example.neophora.io -d '{"jsonrpc": "2.0","id": 1,"method": "GetBlockHashDataByBlockHashInHex","params":{"BlockHash":"405b0dcc1704505feebd6f7a2121e3c979b949827bb1fa1eaf29a205086c156b"}}'
// {"id":1,"result":"0000000054594b55854bdc8b722cf7945e773ca39f969ab707e7c3616974eb5cf379964283301274b99eec5dc4e9d74027665511b9c55c4230b475ac8c09f41548d280e427beac5c80ee36007278e2194b834dfc4e4e04879cfe60ba3b296b1ff08f112f6071756f","error":null}
// ```
func (me *T) GetBlockHashDataByBlockHashInHex(args struct {
	BlockHash h256.T
}, ret *string) error {
	if args.BlockHash.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	return me.Data.GetArgsInHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.hdt",
		Index:  "h256.blk",
		Keys:   []string{args.BlockHash.Val()},
	}, ret)
}
