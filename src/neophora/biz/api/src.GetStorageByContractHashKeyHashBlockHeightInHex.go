package api

import (
	"neophora/lib/type/h160"
	"neophora/lib/type/h256"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetStorageByContractHashKeyHashBlockHeightInHex ...
// as an example:
//
// ```
// $ curl https://example.neophora.io -d '{"jsonrpc": "2.0","id": 1,"method": "GetStorageByContractHashKeyHashBlockHeightInHex","params":{"ContractHash":"e345419e7377286ee5b0a39b56e30f6213ab9e4d","KeyHash":"21b103ee94e0feac3577d378d093b0f4be10b3f6989bbfe6be047e37b8236997","BlockHeight":99999999}}'
// {"id":1,"result":"0007f5ff4e435db43100","error":null}
// ```
func (me *T) GetStorageByContractHashKeyHashBlockHeightInHex(args struct {
	ContractHash h160.T
	KeyHash      h256.T
	BlockHeight  uintval.T
}, ret *string) error {
	if args.ContractHash.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.KeyHash.Valid() == false {
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
		Keys:   []string{args.ContractHash.Val(), args.KeyHash.Val(), args.BlockHeight.Hex()},
	}, ret)
}
