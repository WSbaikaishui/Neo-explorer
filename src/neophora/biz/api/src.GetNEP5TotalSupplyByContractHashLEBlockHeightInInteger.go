package api

import (
	"encoding/json"
	"neophora/lib/type/bins"
	"neophora/lib/type/h160"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetNEP5TotalSupplyByContractHashLEBlockHeightInInteger ...
// as an example:
//
// ```
// $ curl https://example.neophora.io -d '{"jsonrpc": "2.0","id": 1,"method": "GetNEP5TotalSupplyByContractHashLEBlockHeightInInteger","params":{"ContractHashLE":"e8f98440ad0d7a6e76d84fb1c3d3f8a16e162e97","BlockHeight":2400003}}'
// {"id":1,"result":1000000000000000000,"error":null}
// ```
func (me *T) GetNEP5TotalSupplyByContractHashLEBlockHeightInInteger(args struct {
	ContractHashLE h160.T
	BlockHeight    uintval.T
}, ret *json.RawMessage) error {
	var result bins.T
	if args.ContractHashLE.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.BlockHeight.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if err := me.Data.GetArgsInBins(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bigu.tts",
		Index:  "h160.ctr-uint.hgt",
		Keys:   []string{args.ContractHashLE.RevVal(), args.BlockHeight.Hex()},
	}, &result); err != nil {
		return nil
	}
	if result.Valid() == false {
		return stderr.ErrNotFound
	}
	*ret = json.RawMessage(result.BigString())
	return nil
}
