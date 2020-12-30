package api

import (
	"neophora/lib/type/bins"
	"neophora/lib/type/h160"
	"neophora/lib/type/h256"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetNEP5TransferAmountByContractHashLEDSTHashTransactionHashLEOutputIndexSRCHashInUint64 ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetNEP5TransferAmountByContractHashLEDSTHashTransactionHashLEOutputIndexSRCHashInUint64(args struct {
	ContractHashLE    h160.T
	DSTHash           h160.T
	TransactionHashLE h256.T
	OutputIndex       uintval.T
	SRCHash           h160.T
}, ret *uint64) error {
	if args.TransactionHashLE.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.OutputIndex.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.SRCHash.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.DSTHash.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.ContractHashLE.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	var result bins.T
	if err := me.Data.GetArgsInBins(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bigu.tsf",
		Index:  "h160.ctr-h160.dst-h256.trx-uint.num-h160.src",
		Keys:   []string{args.ContractHashLE.RevVal(), args.DSTHash.Val(), args.TransactionHashLE.RevVal(), args.OutputIndex.Hex(), args.SRCHash.Val()},
	}, &result); err != nil {
		return err
	}
	if result.Valid() == false {
		return stderr.ErrNotFound
	}
	*ret = result.Uint64()
	return nil
}
