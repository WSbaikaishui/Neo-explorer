package api

import (
	"neophora/lib/type/addr"
	"neophora/lib/type/bins"
	"neophora/lib/type/h160"
	"neophora/lib/type/h256"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetNEP5TransferAmountByContractHashLEDSTAddressTransactionHashTransferIndexSRCAddressInUint64 ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetNEP5TransferAmountByContractHashLEDSTAddressTransactionHashTransferIndexSRCAddressInUint64(args struct {
	ContractHashLE  h160.T
	DSTAddress      addr.T
	TransactionHash h256.T
	TransferIndex   uintval.T
	SRCAddress      addr.T
}, ret *uint64) error {
	if args.TransactionHash.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.TransferIndex.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.SRCAddress.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.DSTAddress.Valid() == false {
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
		Keys:   []string{args.ContractHashLE.RevVal(), args.DSTAddress.H160(), args.TransactionHash.Val(), args.TransferIndex.Hex(), args.SRCAddress.H160()},
	}, &result); err != nil {
		return err
	}
	if result.Valid() == false {
		return stderr.ErrNotFound
	}
	*ret = result.Uint64()
	return nil
}