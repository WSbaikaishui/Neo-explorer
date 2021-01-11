package api

import (
	"neophora/lib/type/h160"
	"neophora/var/stderr"
)

// GetNEP5TokenNameByContractHashInString ...
// as an example:
//
// ```
// $ curl https://example.neophora.io -d '{"jsonrpc": "2.0","id": 1,"method": "GetNEP5TokenNameByContractHashInString","params":{"ContractHash":"972e166ea1f8d3c3b14fd8766e7a0dad4084f9e8"}}'
// {"id":1,"result":"Experience Token","error":null}
// ```
func (me *T) GetNEP5TokenNameByContractHashInString(args struct {
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
		Target: "strv.tnm",
		Index:  "h160.ctr",
		Keys:   []string{args.ContractHash.Val()},
	}, ret)
}
