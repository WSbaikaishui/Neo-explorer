package api

import (
	"neophora/lib/type/bins"
	"neophora/lib/type/h160"
	"neophora/var/stderr"
)

// GetNEP5DecimalsByContractHashInUint64 ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetNEP5DecimalsByContractHashInUint64(args struct {
	ContractHash h160.T
}, ret *uint64) error {
	var result bins.T
	if args.ContractHash.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if err := me.Data.GetArgsInBins(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bigu.tde",
		Index:  "h160.ctr",
		Keys:   []string{args.ContractHash.Val()},
	}, &result); err != nil {
		return nil
	}
	if result.Valid() == false {
		return stderr.ErrNotFound
	}
	*ret = result.Uint64()
	return nil
}
