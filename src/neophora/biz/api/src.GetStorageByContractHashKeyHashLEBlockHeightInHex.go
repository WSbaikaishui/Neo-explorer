package api

import (
	"neophora/lib/type/h160"
	"neophora/lib/type/h256"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetStorageByContractHashKeyHashLEBlockHeightInHex ...
// as an example:
//
// ```
// $ curl https://example.neophora.io -d '{"jsonrpc": "2.0","id": 1,"method": "GetStorageByContractHashKeyHashLEBlockHeightInHex","params":{"ContractHash":"e345419e7377286ee5b0a39b56e30f6213ab9e4d","KeyHashLE":"976923b8377e04bee6bf9b98f6b310bef4b093d078d37735acfee094ee03b121","BlockHeight":99999999}}'
// {"id":1,"result":"0007324861c968b43100","error":null}
// ```
func (me *T) GetStorageByContractHashKeyHashLEBlockHeightInHex(args struct {
	ContractHash h160.T
	KeyHashLE    h256.T
	BlockHeight  uintval.T
}, ret *string) error {
	if args.ContractHash.Valid() == false {
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
		Keys:   []string{args.ContractHash.Val(), args.KeyHashLE.RevVal(), args.BlockHeight.Hex()},
	}, ret)
}
