package api

import (
	"encoding/json"
	"neophora/lib/type/bins"
	"neophora/lib/type/h160"
	"neophora/var/stderr"
)

// GetNEP5InfoByContractHashLEInJSON ...
// as an example:
//
// ```
// $ curl https://example.neophora.io -d '{"jsonrpc": "2.0","id": 1,"method": "GetNEP5InfoByContractHashLEInJSON","params":{"ContractHashLE":"e8f98440ad0d7a6e76d84fb1c3d3f8a16e162e97"}}'
// {"id":1,"result":{"decimals":8,"name":"Experience Token","symbol":"EXT","totalsupply":1000000000000000000},"error":null}
// ```
func (me *T) GetNEP5InfoByContractHashLEInJSON(args struct {
	ContractHashLE h160.T
}, ret *json.RawMessage) error {
	if args.ContractHashLE.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	var name string
	var symbol string
	var decimals bins.T
	var totalsupply bins.T
	if err := me.Data.GetArgsInString(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "strv.tnm",
		Index:  "h160.ctr",
		Keys:   []string{args.ContractHashLE.RevVal()},
	}, &name); err != nil {
		return err
	}
	if err := me.Data.GetArgsInString(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "strv.tsb",
		Index:  "h160.ctr",
		Keys:   []string{args.ContractHashLE.RevVal()},
	}, &symbol); err != nil {
		return err
	}
	if err := me.Data.GetArgsInBins(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bigu.tde",
		Index:  "h160.ctr",
		Keys:   []string{args.ContractHashLE.RevVal()},
	}, &decimals); err != nil {
		return err
	}
	if decimals.Valid() == false {
		return stderr.ErrNotFound
	}
	if err := me.Data.GetLatestUint64ValInBins(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bigu.tts",
		Index:  "h160.ctr-uint.hgt",
		Keys:   []string{args.ContractHashLE.RevVal()},
	}, &totalsupply); err != nil {
		return err
	}
	if totalsupply.Valid() == false {
		return stderr.ErrNotFound
	}
	marshalled, err := json.Marshal(map[string]interface{}{
		"name":        name,
		"totalsupply": json.RawMessage(totalsupply.BigString()),
		"symbol":      symbol,
		"decimals":    json.RawMessage(decimals.BigString()),
	})
	if err != nil {
		return err
	}
	*ret = json.RawMessage(marshalled)
	return nil
}
