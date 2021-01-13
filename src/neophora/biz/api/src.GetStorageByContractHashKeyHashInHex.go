package api

import (
	"neophora/lib/type/h160"
	"neophora/lib/type/h256"
	"neophora/var/stderr"
)

// GetStorageByContractHashKeyHashInHex ...
// as an example:
//
// ```
// $ curl https://example.neophora.io -d '{"jsonrpc": "2.0","id": 1,"method": "GetStorageByContractHashKeyHashInHex","params":{"ContractHash":"e345419e7377286ee5b0a39b56e30f6213ab9e4d","KeyHash":"21b103ee94e0feac3577d378d093b0f4be10b3f6989bbfe6be047e37b8236997"}}'
// {"id":1,"result":"00071e09cc575eb43100","error":null}
// ```
func (me *T) GetStorageByContractHashKeyHashInHex(args struct {
	ContractHash h160.T
	KeyHash      h256.T
}, ret *string) error {
	if args.ContractHash.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.KeyHash.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	return me.Data.GetLatestUint64ValInHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.str",
		Index:  "h160.ctr-h256.key-uint.hgt",
		Keys:   []string{args.ContractHash.Val(), args.KeyHash.Val()},
	}, ret)
}
