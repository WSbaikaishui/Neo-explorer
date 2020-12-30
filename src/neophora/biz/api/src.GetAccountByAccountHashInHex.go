package api

import (
	"neophora/lib/type/h160"
	"neophora/var/stderr"
)

// GetAccountByAccountHashInHex ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetAccountByAccountHashInHex(args struct {
	AccountHash h160.T
}, ret *string) error {
	if args.AccountHash.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	return me.Data.GetLatestUint64ValInHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.act",
		Index:  "h160.act-uint.hgt",
		Keys:   []string{args.AccountHash.Val()},
	}, ret)
}
