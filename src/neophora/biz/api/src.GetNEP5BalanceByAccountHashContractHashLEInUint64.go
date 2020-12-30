package api

import (
	"neophora/lib/type/bins"
	"neophora/lib/type/h160"
	"neophora/var/stderr"
)

// GetNEP5BalanceByAccountHashContractHashLEInUint64 ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetNEP5BalanceByAccountHashContractHashLEInUint64(args struct {
	AccountHash    h160.T
	ContractHashLE h160.T
}, ret *uint64) error {
	var result bins.T
	if args.ContractHashLE.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.AccountHash.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if err := me.Data.GetLatestUint64ValInBins(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bigu.bal",
		Index:  "h160.act-h160.ctr-uint.hgt",
		Keys:   []string{args.AccountHash.Val(), args.ContractHashLE.RevVal()},
	}, &result); err != nil {
		return nil
	}
	if result.Valid() == false {
		return stderr.ErrNotFound
	}
	*ret = result.Uint64()
	return nil
}
