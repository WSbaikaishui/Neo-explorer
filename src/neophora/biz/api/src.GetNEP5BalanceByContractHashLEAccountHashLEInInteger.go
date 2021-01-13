package api

import (
	"encoding/json"
	"neophora/lib/type/bins"
	"neophora/lib/type/h160"
	"neophora/var/stderr"
)

// GetNEP5BalanceByContractHashLEAccountHashLEInInteger ...
// as an example:
//
// ```
// $ curl https://example.neophora.io -d '{"jsonrpc": "2.0","id": 1,"method": "GetNEP5BalanceByContractHashLEAccountHashLEInInteger","params":{"ContractHashLE":"e8f98440ad0d7a6e76d84fb1c3d3f8a16e162e97","AccountHashLE":"c80fb928f11b365655d0cf5600633ea393c9c630"}}'
// {"id":1,"result":0,"error":null}
// ```
func (me *T) GetNEP5BalanceByContractHashLEAccountHashLEInInteger(args struct {
	ContractHashLE h160.T
	AccountHashLE  h160.T
}, ret *json.RawMessage) error {
	var result bins.T
	if args.ContractHashLE.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.AccountHashLE.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if err := me.Data.GetLatestUint64ValInBins(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bigu.bal",
		Index:  "h160.ctr-h160.act-uint.hgt",
		Keys:   []string{args.ContractHashLE.RevVal(), args.AccountHashLE.RevVal()},
	}, &result); err != nil {
		return nil
	}
	if result.Valid() == false {
		return stderr.ErrNotFound
	}
	*ret = json.RawMessage(result.BigString())
	return nil
}
