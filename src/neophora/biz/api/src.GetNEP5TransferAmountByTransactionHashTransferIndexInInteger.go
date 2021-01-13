package api

import (
	"encoding/json"
	"neophora/lib/type/bins"
	"neophora/lib/type/h256"
	"neophora/lib/type/uintval"
	"neophora/var/pad"
	"neophora/var/stderr"
)

// GetNEP5TransferAmountByTransactionHashTransferIndexInInteger ...
// as an example:
//
// ```
// $ curl http://127.0.0.1:8888 -d '{"jsonrpc": "2.0","id": 1,"method": "GetNEP5TransferAmountByTransactionHashTransferIndexInInteger","params":{"TransactionHash":"dfedfff45cad6dd63d34c7ce2491114a77fa78de7629da0b3c6dace714ffdf9e", "TransferIndex":0}}'
// {"id":1,"result":1650000000000,"error":null}
// ```
func (me *T) GetNEP5TransferAmountByTransactionHashTransferIndexInInteger(args struct {
	TransactionHash h256.T
	TransferIndex   uintval.T
}, ret *json.RawMessage) error {
	if args.TransactionHash.Valid() == false {
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
		Keys:   []string{args.TransactionHash.Val(), args.TransferIndex.Hex(), pad.MAXH160, pad.MAXH160, pad.MAXH160},
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
