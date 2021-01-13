package api

import (
	"neophora/lib/type/h160"
	"neophora/lib/type/h256"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetStorageByContractHashLEKeyHashBlockHeightInHex ...
// as an example:
//
// ```
// $ curl https://example.neophora.io -d '{"jsonrpc": "2.0","id": 1,"method": "GetStorageByContractHashLEKeyHashBlockHeightInHex","params":{"ContractHashLE":"4d9eab13620fe3569ba3b0e56e2877739e4145e3","HexKey":"21b103ee94e0feac3577d378d093b0f4be10b3f6989bbfe6be047e37b8236997","BlockHeight":9999999}}'
// {"id":1,"result":"0007324861c968b43100","error":null}
// ```
func (me *T) GetStorageByContractHashLEKeyHashBlockHeightInHex(args struct {
	ContractHashLE h160.T
	KeyHash        h256.T
	BlockHeight    uintval.T
}, ret *string) error {
	if args.ContractHashLE.Valid() == false {
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
		Keys:   []string{args.ContractHashLE.RevVal(), args.KeyHash.Val(), args.BlockHeight.Hex()},
	}, ret)
}
