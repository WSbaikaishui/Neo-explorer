package api

import (
	"neophora/lib/type/h160"
	"neophora/var/stderr"
)

// GetNEP5TokenSymbolByContractHashLEInString ...
// as an example:
//
// ```
// $  curl https://example.neophora.io -d '{"jsonrpc": "2.0","id": 1,"method": "GetNEP5TokenSymbolByContractHashLEInString","params":{"ContractHashLE":"e8f98440ad0d7a6e76d84fb1c3d3f8a16e162e97"}}'
// {"id":1,"result":"EXT","error":null}
// ```
func (me *T) GetNEP5TokenSymbolByContractHashLEInString(args struct {
	ContractHashLE h160.T
}, ret *string) error {
	if args.ContractHashLE.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	return me.Data.GetArgsInString(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "strv.tsb",
		Index:  "h160.ctr",
		Keys:   []string{args.ContractHashLE.RevVal()},
	}, ret)
}
