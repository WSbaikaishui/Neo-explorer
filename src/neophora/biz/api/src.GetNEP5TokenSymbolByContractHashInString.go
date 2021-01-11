package api

import (
	"neophora/lib/type/h160"
	"neophora/var/stderr"
)

// GetNEP5TokenSymbolByContractHashInString ...
// as an example:
//
// ```
// $ curl https://example.neophora.io -d '{"jsonrpc": "2.0","id": 1,"method": "GetNEP5TokenSymbolByContractHashInString","params":{"ContractHash":"972e166ea1f8d3c3b14fd8766e7a0dad4084f9e8"}}'
// {"id":1,"result":"EXT","error":null}
// ```
func (me *T) GetNEP5TokenSymbolByContractHashInString(args struct {
	ContractHash h160.T
}, ret *string) error {
	if args.ContractHash.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	return me.Data.GetArgsInString(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "strv.tsb",
		Index:  "h160.ctr",
		Keys:   []string{args.ContractHash.Val()},
	}, ret)
}
