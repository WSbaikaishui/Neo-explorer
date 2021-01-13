package api

import (
	"encoding/json"
	"neophora/lib/type/bins"
	"neophora/lib/type/h160"
	"neophora/var/stderr"
)

// GetNEP5InfoByContractHashInJSON ...
// as an example:
//
// ```
// $ curl https://example.neophora.io -d '{"jsonrpc": "2.0","id": 1,"method": "GetNEP5InfoByContractHashInJSON","params":{"ContractHash":"972e166ea1f8d3c3b14fd8766e7a0dad4084f9e8"}}'
// {"id":1,"result":{"decimals":8,"name":"Experience Token","symbol":"EXT","totalsupply":1000000000000000000},"error":null}
// ```
func (me *T) GetNEP5InfoByContractHashInJSON(args struct {
	ContractHash h160.T
}, ret *json.RawMessage) error {
	if args.ContractHash.Valid() == false {
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
		Keys:   []string{args.ContractHash.Val()},
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
		Keys:   []string{args.ContractHash.Val()},
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
		Keys:   []string{args.ContractHash.Val()},
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
		Keys:   []string{args.ContractHash.Val()},
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
