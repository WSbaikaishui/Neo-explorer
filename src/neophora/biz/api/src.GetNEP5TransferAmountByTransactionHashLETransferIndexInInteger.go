package api

import (
	"encoding/json"
	"neophora/lib/type/bins"
	"neophora/lib/type/h256"
	"neophora/lib/type/uintval"
	"neophora/var/pad"
	"neophora/var/stderr"
)

// GetNEP5TransferAmountByTransactionHashLETransferIndexInInteger ...
// as an example:
//
// ```
// $ curl http://127.0.0.1:8888 -d '{"jsonrpc": "2.0","id": 1,"method": "GetNEP5TransferAmountByTransactionHashLETransferIndexInInteger","params":{"TransactionHashLE":"9edfff14e7ac6d3c0bda2976de78fa774a119124cec7343dd66dad5cf4ffeddf", "TransferIndex":0}}'
// {"id":1,"result":1650000000000,"error":null}
// ```
func (me *T) GetNEP5TransferAmountByTransactionHashLETransferIndexInInteger(args struct {
	TransactionHashLE h256.T
	TransferIndex     uintval.T
}, ret *json.RawMessage) error {
	if args.TransactionHashLE.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.TransferIndex.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	var result bins.T
	if err := me.Data.GetLatestValInBins(struct {
		Target string
		Index  string
		Keys   []string
		C      uint
	}{
		Target: "bigu.tsf",
		Index:  "h256.trx-uint.num-h160.src-h160.dst-h160.ctr",
		Keys:   []string{args.TransactionHashLE.RevVal(), args.TransferIndex.Hex(), pad.MAXH160, pad.MAXH160, pad.MAXH160},
		C:      2,
	}, &result); err != nil {
		return err
	}
	if result.Valid() == false {
		return stderr.ErrNotFound
	}
	*ret = json.RawMessage(result.BigString())
	return nil
}
