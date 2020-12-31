package api

import (
	"neophora/lib/type/h160"
	"neophora/lib/type/hexs"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetStorageByContractHashHexKeyBlockHeightInHex ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetStorageByContractHashHexKeyBlockHeightInHex(args struct {
	ContractHash h160.T
	HexKey       hexs.T
	BlockHeight  uintval.T
}, ret *string) error {
	if args.ContractHash.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.HexKey.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.BlockHeight.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	return me.Data.GetLastValInHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.str",
		Index:  "h160.ctr-h256.key-uint.hgt",
		Keys:   []string{args.ContractHash.Val(), args.HexKey.H256(), args.BlockHeight.Hex()},
	}, ret)
}
