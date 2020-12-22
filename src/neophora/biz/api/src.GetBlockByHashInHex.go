package api

import (
	"neophora/lib/type/h256"
	"neophora/var/stderr"
)

// GetBlockByHashInHex ...
func (me *T) GetBlockByHashInHex(args struct {
	Hash h256.T
}, ret *string) error {
	if args.Hash.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	return me.Data.GetArgsInHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.blk",
		Index:  "h256.blk",
		Keys:   []string{args.Hash.Val()},
	}, ret)
}
