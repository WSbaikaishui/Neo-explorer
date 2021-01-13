package api

import (
	"neophora/lib/type/h160"
	"neophora/lib/type/h256"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetStorageByContractHashLEKeyHashLEBlockHeightInHex ...
// as an example:
//
// ```
// $ curl https://example.neophora.io -d '{"jsonrpc": "2.0","id": 1,"method": "GetStorageByContractHashLEKeyHashLEBlockHeightInHex","params":{"ContractHashLE":"4d9eab13620fe3569ba3b0e56e2877739e4145e3","KeyHashLE":"976923b8377e04bee6bf9b98f6b310bef4b093d078d37735acfee094ee03b121","BlockHeight":99999999}}'
// {"id":1,"result":"0007f5ff4e435db43100","error":null}
// ```
func (me *T) GetStorageByContractHashLEKeyHashLEBlockHeightInHex(args struct {
	ContractHashLE h160.T
	KeyHashLE      h256.T
	BlockHeight    uintval.T
}, ret *string) error {
	if args.ContractHashLE.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.KeyHashLE.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.BlockHeight.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	return me.Data.GetLastValInHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.str",
		Index:  "h160.ctr-h256.key-uint.hgt",
		Keys:   []string{args.ContractHashLE.RevVal(), args.KeyHashLE.RevVal(), args.BlockHeight.Hex()},
	}, ret)
}
