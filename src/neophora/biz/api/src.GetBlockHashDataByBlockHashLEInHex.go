package api

import (
	"neophora/lib/type/h256"
	"neophora/var/stderr"
)

// GetBlockHashDataByBlockHashLEInHex ...
// as an example:
//
// ```
// $ curl https://example.neophora.io -d '{"jsonrpc": "2.0","id": 1,"method": "GetBlockHashDataByBlockHashLEInHex","params":{"BlockHashLE":"6b156c0805a229af1efab17b8249b979c9e321217a6fbdee5f500417cc0d5b40"}}'
// {"id":1,"result":"0000000054594b55854bdc8b722cf7945e773ca39f969ab707e7c3616974eb5cf379964283301274b99eec5dc4e9d74027665511b9c55c4230b475ac8c09f41548d280e427beac5c80ee36007278e2194b834dfc4e4e04879cfe60ba3b296b1ff08f112f6071756f","error":null}
// ```
func (me *T) GetBlockHashDataByBlockHashLEInHex(args struct {
	BlockHashLE h256.T
}, ret *string) error {
	if args.BlockHashLE.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	return me.Data.GetArgsInHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.hdt",
		Index:  "h256.blk",
		Keys:   []string{args.BlockHashLE.RevVal()},
	}, ret)
}
