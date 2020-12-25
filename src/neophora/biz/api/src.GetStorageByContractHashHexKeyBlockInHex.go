package api

import (
	"neophora/lib/type/h256"
	"neophora/lib/type/hexs"
	"neophora/var/stderr"
)

// GetStorageByContractHashHexKeyBlockInHex ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetStorageByContractHashHexKeyBlockInHex(args struct {
	ContractHash h256.T
	HexKey       hexs.T
}, ret *string) error {
	if args.ContractHash.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.HexKey.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	return me.Data.GetLatestUint64ValInHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.str",
		Index:  "h160.ctr-h256.key-uint.hgt",
		Keys:   []string{args.ContractHash.Val(), args.HexKey.H256()},
	}, ret)
}
