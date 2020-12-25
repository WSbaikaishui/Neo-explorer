package api

import (
	"neophora/lib/type/h256"
	"neophora/lib/type/hexs"
	"neophora/var/stderr"
)

// GetStorageByContractHashLEHexKeyBlockInHex ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetStorageByContractHashLEHexKeyBlockInHex(args struct {
	ContractHashLE h256.T
	HexKey         hexs.T
}, ret *string) error {
	if args.ContractHashLE.Valid() == false {
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
		Keys:   []string{args.ContractHashLE.Val(), args.HexKey.H256()},
	}, ret)
}
