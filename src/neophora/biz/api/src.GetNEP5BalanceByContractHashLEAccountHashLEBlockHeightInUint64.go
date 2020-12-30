package api

import (
	"neophora/lib/type/bins"
	"neophora/lib/type/h160"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetNEP5BalanceByContractHashLEAccountHashLEBlockHeightInUint64 ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetNEP5BalanceByContractHashLEAccountHashLEBlockHeightInUint64(args struct {
	ContractHashLE h160.T
	AccountHashLE  h160.T
	BlockHeight    uintval.T
}, ret *uint64) error {
	var result bins.T
	if args.ContractHashLE.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.AccountHashLE.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.BlockHeight.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if err := me.Data.GetLastValInBins(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bigu.bal",
		Index:  "h160.ctr-h160.act-uint.hgt",
		Keys:   []string{args.ContractHashLE.RevVal(), args.AccountHashLE.RevVal(), args.BlockHeight.Hex()},
	}, &result); err != nil {
		return nil
	}
	if result.Valid() == false {
		return stderr.ErrNotFound
	}
	*ret = result.Uint64()
	return nil
}
