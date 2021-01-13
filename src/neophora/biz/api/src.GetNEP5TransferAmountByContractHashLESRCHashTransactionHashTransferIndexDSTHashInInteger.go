package api

import (
	"encoding/json"
	"neophora/lib/type/bins"
	"neophora/lib/type/h160"
	"neophora/lib/type/h256"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetNEP5TransferAmountByContractHashLESRCHashTransactionHashTransferIndexDSTHashInInteger ...
// as an example:
//
// ```
// $ curl https://example.neophora.io -d '{"jsonrpc": "2.0","id": 1,"method": "GetNEP5TransferAmountByContractHashLESRCHashTransactionHashTransferIndexDSTHashInInteger","params":{"ContractHashLE":"e8f98440ad0d7a6e76d84fb1c3d3f8a16e162e97","SRCHash":"30c6c993a33e630056cfd05556361bf128b90fc8","DSTHash":"c7cb40c82c7920d2191f8924a6b73659dff3bfc3","TransactionHash":"dfedfff45cad6dd63d34c7ce2491114a77fa78de7629da0b3c6dace714ffdf9e","TransferIndex":0}}'
// {"id":1,"result":1650000000000,"error":null}
// ```
func (me *T) GetNEP5TransferAmountByContractHashLESRCHashTransactionHashTransferIndexDSTHashInInteger(args struct {
	ContractHashLE  h160.T
	SRCHash         h160.T
	TransactionHash h256.T
	TransferIndex   uintval.T
	DSTHash         h160.T
}, ret *json.RawMessage) error {
	if args.TransactionHash.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.TransferIndex.Valid() == false {
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
		Index:  "h160.ctr-h160.src-h256.trx-uint.num-h160.dst",
		Keys:   []string{args.ContractHashLE.RevVal(), args.SRCHash.Val(), args.TransactionHash.Val(), args.TransferIndex.Hex(), args.DSTHash.Val()},
	}, &result); err != nil {
		return err
	}
	if result.Valid() == false {
		return stderr.ErrNotFound
	}
	*ret = json.RawMessage(result.BigString())
	return nil
}
