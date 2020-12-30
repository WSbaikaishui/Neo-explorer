package api

import (
	"encoding/json"
	"neophora/lib/type/bins"
	"neophora/lib/type/h160"
	"neophora/lib/type/h256"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetUTXOByAccountHashLETransactionHashLEOutputIndexInJSON ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetUTXOByAccountHashLETransactionHashLEOutputIndexInJSON(args struct {
	AccountHashLE     h160.T
	TransactionHashLE h256.T
	OutputIndex       uintval.T
}, ret *json.RawMessage) error {
	if args.AccountHashLE.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.TransactionHashLE.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.OutputIndex.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	var result bins.T
	if err := me.Data.GetArgsInBins(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.utx",
		Index:  "h160.act-h256.trx-uint.num",
		Keys:   []string{args.AccountHashLE.RevVal(), args.TransactionHashLE.RevVal(), args.OutputIndex.Hex()},
	}, &result); err != nil {
		return err
	}
	if result.Valid() == false {
		return stderr.ErrNotFound
	}
	js, err := result.JSONViaUTXO()
	if err != nil {
		return stderr.ErrNotFound
	}
	*ret = js
	return nil
}
