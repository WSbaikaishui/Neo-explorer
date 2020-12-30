package api

import (
	"neophora/lib/type/bins"
	"neophora/lib/type/h160"
	"neophora/lib/type/h256"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetNEP5TransferAmountByContractHashLESRCHashDSTHashTransactionHashOutputIndexInUint64 ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetNEP5TransferAmountByContractHashLESRCHashDSTHashTransactionHashOutputIndexInUint64(args struct {
	ContractHashLE  h160.T
	SRCHash         h160.T
	DSTHash         h160.T
	TransactionHash h256.T
	OutputIndex     uintval.T
}, ret *uint64) error {
	if args.TransactionHash.Valid() == false {
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
		Index:  "h160.ctr-h160.src-h160.dst-h256.trx-uint.num",
		Keys:   []string{args.ContractHashLE.RevVal(), args.SRCHash.Val(), args.DSTHash.Val(), args.TransactionHash.Val(), args.OutputIndex.Hex()},
	}, &result); err != nil {
		return err
	}
	if result.Valid() == false {
		return stderr.ErrNotFound
	}
	*ret = result.Uint64()
	return nil
}
