package api

import (
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetBlockHashDataByBlockHeightInHex ...
// as an example:
//
// ```
// $ curl https://example.neophora.io -d '{"jsonrpc": "2.0","id": 1,"method": "GetBlockHashDataByBlockHeightInHex","params":{"BlockHeight":2400000}}'
// {"id":1,"result":"000000006f9ea8d3d05cbbc6078e0ec78842847bcbd612e94b624daeeea1356d1b328af27e2bb12ddc21dffbacb1af393f5ee99a0e9cd698ad42327196124291a1709745c6bf255b009f2400847cd0ea5390b83459e75d652b5d3827bf04c165bbe9ef95cca4bf55","error":null}
// ```
func (me *T) GetBlockHashDataByBlockHeightInHex(args struct {
	BlockHeight uintval.T
}, ret *string) error {
	if args.BlockHeight.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	return me.Data.GetLastValInHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.hdt",
		Index:  "uint.hgt",
		Keys:   []string{args.BlockHeight.Hex()},
	}, ret)
}
