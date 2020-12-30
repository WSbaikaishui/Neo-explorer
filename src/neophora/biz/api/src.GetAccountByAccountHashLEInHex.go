package api

import (
	"neophora/lib/type/h160"
	"neophora/var/stderr"
)

// GetAccountByAccountHashLEInHex ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetAccountByAccountHashLEInHex(args struct {
	AccountHashLE h160.T
}, ret *string) error {
	if args.AccountHashLE.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	return me.Data.GetLatestUint64ValInHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.act",
		Index:  "h160.act-uint.hgt",
		Keys:   []string{args.AccountHashLE.RevVal()},
	}, ret)
}
