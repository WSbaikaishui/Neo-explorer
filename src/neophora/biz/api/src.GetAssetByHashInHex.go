package api

import (
	"neophora/lib/type/h256"
	"neophora/var/stderr"
)

// GetAssetByHashInHex ...
func (me *T) GetAssetByHashInHex(args struct {
	Hash h256.T
}, ret *string) error {
	if args.Hash.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	return me.Data.GetLastestUint64ValInHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.ast",
		Index:  "h256.ast",
		Keys:   []string{args.Hash.Val()},
	}, ret)
}
