package api

import (
	"neophora/lib/type/h160"
	"neophora/var/stderr"
)

// GetNEP5TokenNameByContractHashInString ...
// as an example:
//
// ```
// TODO
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
