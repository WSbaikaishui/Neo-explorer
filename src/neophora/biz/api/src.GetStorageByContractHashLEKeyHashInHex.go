package api

import (
	"neophora/lib/type/h160"
	"neophora/lib/type/h256"
	"neophora/var/stderr"
)

// GetStorageByContractHashLEKeyHashInHex ...
// as an example:
//
// ```
// $ curl https://example.neophora.io -d '{"jsonrpc": "2.0","id": 1,"method": "GetStorageByContractHashLEKeyHashInHex","params":{"ContractHashLE":"4d9eab13620fe3569ba3b0e56e2877739e4145e3","KeyHash":"21b103ee94e0feac3577d378d093b0f4be10b3f6989bbfe6be047e37b8236997"}}'
// {"id":1,"result":"0007f5ff4e435db43100","error":null}
// ```
func (me *T) GetStorageByContractHashLEKeyHashInHex(args struct {
	ContractHashLE h160.T
	KeyHash        h256.T
}, ret *string) error {
	if args.ContractHashLE.Valid() == false {
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
		Keys:   []string{args.ContractHashLE.RevVal(), args.KeyHash.Val()},
	}, ret)
}
