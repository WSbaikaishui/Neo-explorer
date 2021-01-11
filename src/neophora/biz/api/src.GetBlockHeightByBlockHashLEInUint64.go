package api

import (
	"neophora/lib/type/bins"
	"neophora/lib/type/h256"
	"neophora/var/stderr"
)

// GetBlockHeightByBlockHashLEInUint64 ...
// as an example:
//
// ```
// $ curl https://example.neophora.io -d '{"jsonrpc": "2.0","id": 1,"method": "GetBlockHeightByBlockHashLEInUint64","params":{"BlockHashLE":"6b156c0805a229af1efab17b8249b979c9e321217a6fbdee5f500417cc0d5b40"}}'
// {"id":1,"result":3600000,"error":null}
// ```
func (me *T) GetBlockHeightByBlockHashLEInUint64(args struct {
	BlockHashLE h256.T
}, ret *uint64) error {
	if args.BlockHashLE.Valid() == false {
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
		Keys:   []string{args.BlockHashLE.RevVal()},
	}, &result); err != nil {
		return err
	}
	if result.Valid() == false {
		return stderr.ErrNotFound
	}
	*ret = result.Uint64()
	return nil
}
