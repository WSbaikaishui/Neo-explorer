package api

import (
	"neophora/lib/type/bins"
	"neophora/lib/type/h256"
	"neophora/var/stderr"
)

// GetBlockHeightByBlockHashInUint64 ...
// as an example:
//
// ```
// $ curl https://example.neophora.io -d '{"jsonrpc": "2.0","id": 1,"method": "GetBlockHeightByBlockHashInUint64","params":{"BlockHash":"405b0dcc1704505feebd6f7a2121e3c979b949827bb1fa1eaf29a205086c156b"}}'
// {"id":1,"result":3600000,"error":null}
// ```
func (me *T) GetBlockHeightByBlockHashInUint64(args struct {
	BlockHash h256.T
}, ret *uint64) error {
	if args.BlockHash.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	var result bins.T
	if err := me.Data.GetArgsInBins(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "uint.hgt",
		Index:  "h256.blk",
		Keys:   []string{args.BlockHash.Val()},
	}, &result); err != nil {
		return err
	}
	if result.Valid() == false {
		return stderr.ErrNotFound
	}
	*ret = result.Uint64()
	return nil
}
