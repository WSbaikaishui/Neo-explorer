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

// GetNEP5TransferAmountByContractHashDSTAddressTransactionHashTransferIndexSRCAddressInInteger ...
// as an example:
//
// ```
// $ curl https://example.neophora.io -d '{"jsonrpc": "2.0","id": 1,"method": "GetNEP5TransferAmountByContractHashDSTAddressTransactionHashTransferIndexSRCAddressInInteger","params":{"ContractHash":"972e166ea1f8d3c3b14fd8766e7a0dad4084f9e8","DSTAddress":"AZcuN1PmBJmSdt8ubY6JqDBGCrHVDpo4kR","SRCAddress":"Aa1hdLWLw441vQ94hzpwDj8FY7QEh1zsr6","TransactionHash":"dfedfff45cad6dd63d34c7ce2491114a77fa78de7629da0b3c6dace714ffdf9e","TransferIndex":0}}'
// {"id":1,"result":1650000000000,"error":null}
// ```
func (me *T) GetNEP5TransferAmountByContractHashDSTAddressTransactionHashTransferIndexSRCAddressInInteger(args struct {
	ContractHash    h160.T
	DSTAddress      addr.T
	TransactionHash h256.T
	TransferIndex   uintval.T
	SRCAddress      addr.T
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
		Index:  "h160.ctr-h160.dst-h256.trx-uint.num-h160.src",
		Keys:   []string{args.ContractHash.Val(), args.DSTAddress.H160(), args.TransactionHash.Val(), args.TransferIndex.Hex(), args.SRCAddress.H160()},
	}, &result); err != nil {
		return err
	}
	if result.Valid() == false {
		return stderr.ErrNotFound
	}
	*ret = json.RawMessage(result.BigString())
	return nil
}
