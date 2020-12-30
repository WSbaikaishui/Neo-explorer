package api

import (
	"neophora/lib/type/bins"
	"neophora/lib/type/h160"
	"neophora/lib/type/h256"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetNEP5TransferAmountByTransactionHashOutputIndexSRCHashLEDSTHashLEContractHashInUint64 ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetNEP5TransferAmountByTransactionHashOutputIndexSRCHashLEDSTHashLEContractHashInUint64(args struct {
	TransactionHash h256.T
	OutputIndex     uintval.T
	SRCHashLE       h160.T
	DSTHashLE       h160.T
	ContractHash    h160.T
}, ret *uint64) error {
	if args.TransactionHash.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.OutputIndex.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.SRCHashLE.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.DSTHashLE.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.ContractHash.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	var result bins.T
	if err := me.Data.GetArgsInBins(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bigu.tsf",
		Index:  "h256.trx-uint.num-h160.src-h160.dst-h160.ctr",
		Keys:   []string{args.TransactionHash.Val(), args.OutputIndex.Hex(), args.SRCHashLE.RevVal(), args.DSTHashLE.RevVal(), args.ContractHash.Val()},
	}, &result); err != nil {
		return err
	}
	if result.Valid() == false {
		return stderr.ErrNotFound
	}
	*ret = result.Uint64()
	return nil
}
