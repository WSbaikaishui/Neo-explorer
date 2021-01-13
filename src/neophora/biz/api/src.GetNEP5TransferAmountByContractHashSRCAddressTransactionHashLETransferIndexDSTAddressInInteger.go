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

// GetNEP5TransferAmountByContractHashSRCAddressTransactionHashLETransferIndexDSTAddressInInteger ...
// as an example:
//
// ```
// $ curl https://example.neophora.io -d '{"jsonrpc": "2.0","id": 1,"method": "GetNEP5TransferAmountByContractHashSRCAddressTransactionHashLETransferIndexDSTAddressInInteger","params":{"ContractHash":"972e166ea1f8d3c3b14fd8766e7a0dad4084f9e8","SRCAddress":"Aa1hdLWLw441vQ94hzpwDj8FY7QEh1zsr6","DSTAddress":"AZcuN1PmBJmSdt8ubY6JqDBGCrHVDpo4kR","TransactionHashLE":"9edfff14e7ac6d3c0bda2976de78fa774a119124cec7343dd66dad5cf4ffeddf","TransferIndex":0}}'
// {"id":1,"result":1650000000000,"error":null}
// ```
func (me *T) GetNEP5TransferAmountByContractHashSRCAddressTransactionHashLETransferIndexDSTAddressInInteger(args struct {
	ContractHash      h160.T
	SRCAddress        addr.T
	TransactionHashLE h256.T
	TransferIndex     uintval.T
	DSTAddress        addr.T
}, ret *json.RawMessage) error {
	if args.TransactionHashLE.Valid() == false {
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
		Index:  "h160.ctr-h160.src-h256.trx-uint.num-h160.dst",
		Keys:   []string{args.ContractHash.Val(), args.SRCAddress.H160(), args.TransactionHashLE.RevVal(), args.TransferIndex.Hex(), args.DSTAddress.H160()},
	}, &result); err != nil {
		return err
	}
	if result.Valid() == false {
		return stderr.ErrNotFound
	}
	*ret = json.RawMessage(result.BigString())
	return nil
}
