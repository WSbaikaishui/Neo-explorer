package api

import (
	"neophora/lib/type/h256"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetUTXOByAssetHashLETransactionHashOutputIndexInHex ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetUTXOByAssetHashLETransactionHashOutputIndexInHex(args struct {
	AssetHashLE     h256.T
	TransactionHash h256.T
	OutputIndex     uintval.T
}, ret *string) error {
	if args.AssetHashLE.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.TransactionHash.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.OutputIndex.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	return me.Data.GetArgsInHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.utx",
		Index:  "h256.ast-h256.trx-uint.num",
		Keys:   []string{args.AssetHashLE.RevVal(), args.TransactionHash.Val(), args.OutputIndex.Hex()},
	}, ret)
}
