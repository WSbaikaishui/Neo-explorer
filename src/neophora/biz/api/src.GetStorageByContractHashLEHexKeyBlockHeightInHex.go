package api

import (
	"neophora/lib/type/h256"
	"neophora/lib/type/hexs"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetStorageByContractHashLEHexKeyBlockHeightInHex ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetStorageByContractHashLEHexKeyBlockHeightInHex(args struct {
	ContractHashLE h256.T
	HexKey         hexs.T
	BlockHeight    uintval.T
}, ret *string) error {
	if args.ContractHashLE.Valid() == false {
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
		Keys:   []string{args.ContractHashLE.RevVal(), args.HexKey.H256(), args.BlockHeight.Hex()},
	}, ret)
}
