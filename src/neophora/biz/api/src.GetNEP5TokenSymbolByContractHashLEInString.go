package api

import (
	"neophora/lib/type/h160"
	"neophora/var/stderr"
)

// GetNEP5TokenSymbolByContractHashLEInString ...
// as an example:
//
// ```
// TODO
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
