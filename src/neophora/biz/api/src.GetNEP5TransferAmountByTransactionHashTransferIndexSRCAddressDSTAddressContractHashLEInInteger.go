package api

import (
	"encoding/json"
	"neophora/lib/type/addr"
	"neophora/lib/type/bins"
	"neophora/lib/type/h160"
	"neophora/lib/type/h256"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetNEP5TransferAmountByTransactionHashTransferIndexSRCAddressDSTAddressContractHashLEInInteger ...
// as an example:
//
// ```
// $ curl https://example.neophora.io -d '{"jsonrpc": "2.0","id": 1,"method": "GetNEP5TransferAmountByTransactionHashTransferIndexSRCAddressDSTAddressContractHashLEInInteger","params":{"TransactionHash":"dfedfff45cad6dd63d34c7ce2491114a77fa78de7629da0b3c6dace714ffdf9e","TransferIndex":0,"SRCAddress":"Aa1hdLWLw441vQ94hzpwDj8FY7QEh1zsr6","DSTAddress":"AZcuN1PmBJmSdt8ubY6JqDBGCrHVDpo4kR","ContractHashLE":"e8f98440ad0d7a6e76d84fb1c3d3f8a16e162e97"}}'
// {"id":1,"result":1650000000000,"error":null}
// ```
func (me *T) GetNEP5TransferAmountByTransactionHashTransferIndexSRCAddressDSTAddressContractHashLEInInteger(args struct {
	TransactionHash h256.T
	TransferIndex   uintval.T
	SRCAddress      addr.T
	DSTAddress      addr.T
	ContractHashLE  h160.T
}, ret *json.RawMessage) error {
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
		Index:  "h256.trx-uint.num-h160.src-h160.dst-h160.ctr",
		Keys:   []string{args.TransactionHash.Val(), args.TransferIndex.Hex(), args.SRCAddress.H160(), args.DSTAddress.H160(), args.ContractHashLE.RevVal()},
	}, &result); err != nil {
		return err
	}
	if result.Valid() == false {
		return stderr.ErrNotFound
	}
	*ret = json.RawMessage(result.BigString())
	return nil
}
